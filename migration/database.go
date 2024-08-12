package migration

import (
	"fmt"
	"gorm.io/gorm/clause"
	"seagate-hackathon/db"
)

func CreateMigrationDB(migration db.Migration) {
	db.DbSession.Create(&db.Migration{
		AwsAccessKey:   migration.AwsAccessKey,
		AwsSecretKey:   migration.AwsSecretKey,
		AwsRegionName:  migration.AwsRegionName,
		AwsBucket:      migration.AwsBucket,
		LyveAccessKey:  migration.LyveAccessKey,
		LyveSecretKey:  migration.LyveSecretKey,
		LyveRegionName: migration.LyveRegionName,
		LyveBucket:     migration.LyveBucket,
		Status:         db.NotStarted,
	})
}

func CreateObjectDB(object db.Object) {
	db.DbSession.Omit(clause.Associations).Create(&db.Object{
		Key:       object.Key,
		Status:    db.NotStarted,
		Migration: object.Migration,
	})
}

func GetMigration(id uint) *db.Migration {
	var migration = &db.Migration{}
	if err := db.DbSession.First(migration, id).Error; err != nil {
		return nil
	}

	return migration
}

func GetObject(id uint) *db.Object {
	var object = &db.Object{}
	if err := db.DbSession.First(object, id).Error; err != nil {
		return nil
	}

	return object
}

func UpdateObjectStatus(oId uint, status db.StatusEnum) {
	object := GetObject(oId)
	if object == nil {
		return
	}

	object.Status = status
	db.DbSession.Save(object)
}

func GetNotStartedAndSet() (*db.Object, error) {
	var object db.Object
    tx := db.DbSession.Begin()

    if err := tx.Where("status = ?", db.NotStarted).First(&object).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there is an error
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return nil, nil // Return nil if no record is found, without an error
		// }
		fmt.Println(err, "this is error")
		return nil, err // Return the error for better handling
	}

    object.Status = db.InProgress
    if err := tx.Save(&object).Error; err != nil {
        tx.Rollback()
        return nil, err
    }

    tx.Commit()
    return &object, nil
}

func UpdateInProgressObjectsStatus() {
	db.DbSession.Model(db.Object{}).Where("status = ?", db.InProgress).Updates(db.Object{Status: db.NotStarted})
}

func CheckMigrationAndSet(oId uint) {
	object := GetObject(oId)
	var exists bool
	db.DbSession.Model(db.Object{}).Select("count(*) > 0").
		Where("(status = ? OR status = ?) AND migration_id = ?", db.NotStarted, db.InProgress, object.MigrationID).Find(&exists)
	if exists {
		return
	}

	var migration = GetMigration(object.MigrationID)
	if migration != nil {
		migration.Status = db.Done
		db.DbSession.Save(migration)
	}
}
