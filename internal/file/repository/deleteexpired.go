package repository

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"time"
)

func (f *FileRepository) DeleteExpired(maxTime time.Time) error {
	maxTimeUnixTime := maxTime.Unix()
	res := f.database.
		GetConnection().
		Where("created_at <= ?", maxTimeUnixTime).
		Where("permanent = ?", false).
		Delete(&file.File{})

	if res.Error != nil {
		return fmt.Errorf("filerepository deleteexpired: %v", res.Error)
	}
	return nil
}
