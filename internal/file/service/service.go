package service

import (
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/config"
	"github.com/Takasakiii/ayanami/pkg/cuid"
	"github.com/Takasakiii/ayanami/pkg/sender"
)

type FileService struct {
	config     *config.File
	sender     sender.Sender
	cuid       cuid.Generator
	downloader sender.Downloader
	repository file.Repository
}

func NewService(
	conf *config.Config,
	sender sender.Sender,
	cuid cuid.Generator,
	downloader sender.Downloader,
	repository file.Repository) *FileService {
	return &FileService{
		config:     &conf.File,
		sender:     sender,
		cuid:       cuid,
		downloader: downloader,
		repository: repository,
	}
}
