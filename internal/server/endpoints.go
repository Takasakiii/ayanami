package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/Takasakiii/ayanami/internal/sender"
	"github.com/Takasakiii/ayanami/internal/server/internal/templates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s Server) uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	if file.Size == 0 {
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: "file is empty",
		})
		return
	}

	if file.Size > 1024*1024 {
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: "file is too big",
		})
		return
	}

	password := c.PostForm("password")

	ip := c.Request.Header.Get("CF-Connecting-Ip")
	if ip == "" {
		ip = c.RemoteIP()
	}
	userAgent := c.Request.Header.Get("User-Agent")

	fileId, err := s.FileManager.UploadFile(file, password, ip, userAgent)
	if err != nil {
		log.Printf("[ERROR] upload file error: %v\n", err)
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: "upload file error",
		})
		return
	}

	fileUrl := fmt.Sprintf("%s/files/%s", s.Config.ServerUrl, fileId)
	c.JSON(http.StatusOK, uploadFileResponse{
		Url: fileUrl,
	})
}

func (s Server) downloadFile(c *gin.Context) {
	fileId, found := c.Params.Get("fileId")
	if !found {
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: "fileId is missing",
		})
		return
	}

	password := c.Query("password")

	downloadedFile, downErr := s.FileManager.DownloadFile(fileId, password)
	if downErr != nil {
		var castedDownloadError *sender.DownloadError
		if errors.As(downErr, &castedDownloadError) && castedDownloadError.Type == sender.InvalidFileIdError {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: "fileId is invalid",
			})
			return
		}

		log.Printf("[ERROR] Download file error: %v\n", downErr)
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: "download file error",
		})

		return
	}

	c.Header("Content-Type", downloadedFile.MimeType)
	c.Header("Accept-Length", fmt.Sprintf("%d", downloadedFile.Size))

	_, err := c.Writer.Write(downloadedFile.Content)
	if err != nil {
		log.Printf("[ERROR] Download file error: %v\n", err)
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: "download file error",
		})
		return
	}
}

func (s Server) indexPage(c *gin.Context) {
	background := context.Background()
	page := templates.IndexPage()
	_ = page.Render(background, c.Writer)
}
