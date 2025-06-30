package repository

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
)

func (f *FileRepository) AddFile(data *file.File) error {
	res := f.database.GetConnection().Create(data)
	if res.Error != nil {
		return fmt.Errorf("filerepository addfile: %v", res.Error)
	}
	return nil
}
