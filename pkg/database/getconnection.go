package database

import "gorm.io/gorm"

func (d *GormDatabase) GetConnection() *gorm.DB {
	if d.connection == nil {
		panic("database connection is nil")
	}

	return d.connection
}
