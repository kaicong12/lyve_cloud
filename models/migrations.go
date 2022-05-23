package models

import (
	"fmt"
	"seagate-hackathon/db"
)

func GetAllMigrations(migrations *[]db.Migration, order_by string, search string, page int, page_size int) (err error) {
	base_query := db.DbSession.Order(order_by).
		Offset(page - 1*page_size).
		Limit(page_size).
		Find(migrations)

	if search != "" {
		base_query = base_query.Where("name like = ", fmt.Sprintf("%%%s%%", search))
	}

	if err := base_query.Error; err != nil {
		return err
	}
	return nil
}

func GetAMigration(migration *db.Migration, pk string) (err error) {
	if err := db.DbSession.First(migration, pk).Error; err != nil {
		return err
	}
	return nil
}

func CreateAMigration(newMigration *db.Migration) (err error) {
	if err := db.DbSession.Create(newMigration).Error; err != nil {
		return err
	}

	// create objects and object parts

	return nil
}
