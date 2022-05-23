package models

import "seagate-hackathon/db"

type ObjectsStatusCount struct {
	Status db.StatusEnum
	Count  int
}

func GetMigrationObjects(migration_pk string, migration *db.Migration, objects *[]db.Object) (err error) {
	// retrieve migration
	if err := db.DbSession.Where("ID = ?", migration_pk).First(&migration).Error; err != nil {
		return err
	}

	// look up objects under migration
	if err := db.DbSession.Where(migration).Find(&objects).Error; err != nil {
		return err
	}
	return nil
}

func GroupByStatus(migration_pk uint) ([]ObjectsStatusCount, error) {
	var status []ObjectsStatusCount
	if err := db.DbSession.Table("object").
		Select("status, count(id) as count").
		Where(&db.Object{MigrationID: migration_pk}).
		Group("status").
		Scan(&status).Error; err != nil {
		return []ObjectsStatusCount{}, err
	}

	return status, nil
}
