package filemanager

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/config"
	"github.com/Takasakiii/ayanami/internal/sender"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type FileManager struct {
	config     *config.File
	sender     sender.Sender
	cuid       func() string
	downloader sender.Downloader
	database   *gorm.DB
}

func NewFileManager(
	config *config.File,
	sender sender.Sender,
	downloader sender.Downloader,
	db *gorm.DB) (FileManager, error) {

	cuid, err := cuid2.Init()

	if err != nil {
		return FileManager{}, fmt.Errorf("filemanager init cuid: %v", err)
	}

	return FileManager{
		config:     config,
		sender:     sender,
		downloader: downloader,
		cuid:       cuid,
		database:   db,
	}, nil
}
