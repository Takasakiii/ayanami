package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (d *GormDatabase) ConnectDatabase() error {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("database connection failed: %v", err)
	}
	db = db.Debug()

	d.connection = db
	return nil
}
