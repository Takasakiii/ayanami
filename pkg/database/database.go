package database

import (
	"github.com/Takasakiii/ayanami/pkg/config"
	"gorm.io/gorm"
)

type GormDatabase struct {
	connection *gorm.DB
	config     *config.Database
}

func NewGormDatabase(conf *config.Config) *GormDatabase {
	return &GormDatabase{
		config: &conf.Database,
	}
}
