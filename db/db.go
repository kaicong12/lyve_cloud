package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StatusEnum int

const (
	NotStarted StatusEnum = iota
	InProgress
	Done
	Failed
)

const dbFile = "test.db"

type Migration struct {
	gorm.Model
	Name           string `binding:"required"`
	AwsAccessKey   string `binding:"required"`
	AwsSecretKey   string `binding:"required"`
	AwsRegionName  string
	AwsBucket      string `binding:"required"`
	LyveAccessKey  string `binding:"required"`
	LyveSecretKey  string `binding:"required"`
	LyveRegionName string
	LyveBucket     string `binding:"required"`
	Status         StatusEnum
}

type Object struct {
	gorm.Model
	Key           string
	ContentLength int64
	Status        StatusEnum `json:"status"`
	MigrationID   uint
	Migration     Migration `gorm:"constraint:OnDelete:CASCADE;"`
}

var DbSession *gorm.DB

func InitDB() {
	var err error
	DbSession, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = DbSession.AutoMigrate(&Migration{}, &Object{})
	if err != nil {
		panic(err)
	}
}

func ResetDB() {
	os.Remove(dbFile)
}
