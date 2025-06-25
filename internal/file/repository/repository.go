package repository

import (
	"github.com/Takasakiii/ayanami/pkg/database"
)

type FileRepository struct {
	database database.Database
}

func NewFileRepository(database database.Database) *FileRepository {
	return &FileRepository{
		database: database,
	}
}
