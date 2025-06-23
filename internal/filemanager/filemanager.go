package filemanager

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/config"
	"github.com/Takasakiii/ayanami/internal/sender"
	"github.com/nrednav/cuid2"
)

type FileManager struct {
	config *config.File
	sender sender.Sender
	cuid   func() string
}

func NewFileManager(config *config.File, sender sender.Sender) (FileManager, error) {
	cuid, err := cuid2.Init()

	if err != nil {
		return FileManager{}, fmt.Errorf("filemanager init cuid: %v", err)
	}

	return FileManager{
		config: config,
		sender: sender,
		cuid:   cuid,
	}, nil
}
