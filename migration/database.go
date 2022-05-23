package migration

import (
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

func GetNotStartedAndSet() *db.Object {
	var object = &db.Object{}
	if err := db.DbSession.Where(&db.Object{Status: db.NotStarted}).First(object).Error; err != nil {
		return nil
	}

	object.Status = db.InProgress
	db.DbSession.Save(object)
	return object
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
