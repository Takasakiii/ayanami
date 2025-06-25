package database

import (
	"gorm.io/gorm"
)

type GormDatabase struct {
	connection *gorm.DB
}

func NewGormDatabase() *GormDatabase {
	return &GormDatabase{}
}
