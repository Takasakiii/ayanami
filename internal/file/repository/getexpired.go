package repository

import (
	"fmt"
	"time"
)

func (f *FileRepository) GetExpired(maxTime time.Time) ([]string, error) {
	maxTimeUnix := maxTime.Unix()
	var filesExpired []string
	res := f.database.
		GetConnection().
		Where("created_at <= ?", maxTimeUnix).
		Where("permanent = ?", false).
		Select("file_name").
		Find(&filesExpired)

	if res.Error != nil {
		return nil, fmt.Errorf("filerepository getexpired: %v", res.Error)
	}

	return filesExpired, nil
}
