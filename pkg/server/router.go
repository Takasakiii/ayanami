package server

import (
	"github.com/Takasakiii/ayanami/pkg/server/internal"
	"github.com/Takasakiii/ayanami/pkg/server/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) router(engine *gin.Engine) {
	file := engine.Group("/files")
	file.POST("/", handlers.UploadFile(s))
	file.GET("/:fileId", handlers.DownloadFile(s))

	engine.GET("/", handlers.IndexPage())
	engine.StaticFS("/static/", http.FS(internal.GetStaticFiles()))
}
