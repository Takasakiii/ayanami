package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database *gorm.DB

func GetDatabase() Database {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate()
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
