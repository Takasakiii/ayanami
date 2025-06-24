package server

import (
	"context"
	"github.com/Takasakiii/ayanami/prisma/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (s Server) logMiddleware(c *gin.Context) {
	ctx := context.Background()
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("[WARN] invalid file in request")
	}

	fileName := file.Filename
	mimeType := file.Header.Get("Content-Type")
	ip := c.Request.Header.Get("CF-Connecting-Ip")
	if ip == "" {
		ip = c.RemoteIP()
	}

	userAgent := c.Request.Header.Get("User-Agent")
	localNow := time.Now()

	res, err := s.Database.UploadedFiles.CreateOne(
		db.UploadedFiles.CreatedAt.Set(localNow.UTC()),
		db.UploadedFiles.UserAgent.Set(userAgent),
		db.UploadedFiles.IP.Set(ip),
		db.UploadedFiles.FileName.Set(fileName),
		db.UploadedFiles.MimeType.Set(mimeType),
	).Exec(ctx)
	if err != nil {
		log.Printf("[WARN] failed to create file log, connection aborted: %v\n", err)
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: "failed to upload file",
		})
		c.Abort()
		return
	}

	log.Printf("[INFO] saved log, log id: %d", res.ID)
}
