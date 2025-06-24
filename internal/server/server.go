package server

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/config"
	"github.com/Takasakiii/ayanami/internal/database"
	"github.com/Takasakiii/ayanami/internal/filemanager"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config      *config.Server
	FileManager *filemanager.FileManager
	Database    database.Database
}

func (s Server) StartWebServer() {
	g := gin.Default()
	s.router(g)

	err := g.Run(fmt.Sprintf("%s:%d", s.Config.BindHost, s.Config.BindPort))
	if err != nil {
		panic(err)
	}
}
