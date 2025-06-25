package database

import "gorm.io/gorm"

type Database interface {
	GetConnection() *gorm.DB
	Migrate(tables ...interface{}) error
	ConnectDatabase() error
}
