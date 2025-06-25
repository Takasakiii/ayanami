package internal

import (
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/config"
)

type Server interface {
	GetConfig() *config.Server
	GetFile() file.Service
}
