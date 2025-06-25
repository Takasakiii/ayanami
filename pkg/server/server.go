package server

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.Server
	File   file.Service
}

func (s *Server) StartWebServer() {
	g := gin.Default()
	s.router(g)

	err := g.Run(fmt.Sprintf("%s:%d", s.Config.BindHost, s.Config.BindPort))
	if err != nil {
		panic(err)
	}
}
