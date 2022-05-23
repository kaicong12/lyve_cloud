package main

import (
	"os"
	"seagate-hackathon/migration"

	"seagate-hackathon/db"
	"seagate-hackathon/routes"
	"seagate-hackathon/utils"
)

func CreateMockMigration() {
	migration.CreateMigrationDB(db.Migration{
		AwsAccessKey:   os.Getenv("AWS_ACCESS_KEY"),
		AwsSecretKey:   os.Getenv("AWS_SECRET_KEY"),
		AwsRegionName:  "ap-southeast-1",
		AwsBucket:      "linh-testing-nhan",
		LyveAccessKey:  os.Getenv("LYVE_ACCESS_KEY"),
		LyveSecretKey:  os.Getenv("LYVE_SECRET_KEY"),
		LyveRegionName: "ap-southeast-1",
		LyveBucket:     "active-learning-linh",
		Status:         0,
	})

	migration.CreateObjectDB(db.Object{
		Key:         "hackathon/sample (1).zip",
		MigrationID: 1,
	})
}

func main() {
	db.InitDB()

	utils.InitializeLogger()
	migration.Init()
	r := routes.SetUpRouter()
	r.Run()
}
