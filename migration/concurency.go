package migration

import (
	"seagate-hackathon/db"
	"time"
)

const NumberOfWorker = 5

var requestChan = make(chan int)
var execChan = make(chan *ObjectMigrationController)
var onDone = make(chan uint) // channel receiving the id of done object
var onErr = make(chan uint)  // channel receiving the id of failed object
var largeObjectList = make(map[uint]*ObjectMigrationController)
var partTracking = make(map[uint]int64)

func onCompleteObject(oId uint) {
	delete(largeObjectList, oId)
	delete(partTracking, oId)
	UpdateObjectStatus(oId, db.Done)
	CheckMigrationAndSet(oId)
}

func onErrorObject(oId uint) {
	delete(largeObjectList, oId)
	delete(partTracking, oId)
	UpdateObjectStatus(oId, db.Failed)
	CheckMigrationAndSet(oId)
}

func onRequest() bool {
	for u, o := range largeObjectList {
		if o.getNumberOfParts() > partTracking[u] {
			partTracking[u]++
			execChan <- o
			return true
		}
	}

	for {
		object := GetNotStartedAndSet()
		if object != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		sourceClient, destClient := initMigrationClients(object.MigrationID)
		migrationController := NewObjectMigrationController(sourceClient, destClient, object.Key, object.ID)
		migrationController.prepareObject()

		if migrationController.isSmall() {
			execChan <- migrationController
			return true
		}

		largeObjectList[object.ID] = migrationController
		partTracking[object.ID] = 1
		execChan <- migrationController
		return true
	}
}

func worker() {
	requestChan <- 0
	for o := range execChan {
		if o == nil { // exit signal
			return
		}

		if !o.isSmall() {
			o.migratePart()
		} else {
			o.migrateSmallObject()
		}

		requestChan <- 0
	}
}

func master() {
	for i := 0; i < NumberOfWorker; i++ {
		go worker()
	}

	for {
		select {
		case <-requestChan:
			if !onRequest() {
				return
			}
		case oId := <-onDone:
			onCompleteObject(oId)
		case oId := <-onErr:
			onErrorObject(oId)
		}
	}
}

func Init() {
	UpdateInProgressObjectsStatus()
	go master()
}
