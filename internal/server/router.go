package server

import "github.com/gin-gonic/gin"

func (s Server) router(engine *gin.Engine) {
	file := engine.Group("/files")
	file.POST("/", s.uploadFile)
	file.GET("/:fileId", s.downloadFile)

	engine.GET("/", s.indexPage)
}
