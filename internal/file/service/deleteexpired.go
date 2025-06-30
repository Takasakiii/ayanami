package service

import (
	"fmt"
	"time"
)

const daysToDelete = 30

func (s *FileService) DeleteExpired() error {
	expirationDate := time.
		Now().
		UTC().
		AddDate(0, 0, -daysToDelete)
	err := s.repository.DeleteExpired(expirationDate)
	if err != nil {
		return fmt.Errorf("fileservice deleteexpired: %v", err)
	}

	return nil
}
