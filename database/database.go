package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type IDatabase interface {
	Setup()
}

var DB *gorm.DB = nil

func New() *gorm.DB {
	if DB == nil {
		DB = Setup()
	}

	return DB
}

func Setup() *gorm.DB {

	filename := "database/database.sqlite"

	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err != nil {
		panic("connection failed " + err.Error())
	}

	return db
}
