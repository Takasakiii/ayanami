package handlers

import (
	"errors"
	"fmt"
	filePkg "github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/sender"
	"github.com/Takasakiii/ayanami/pkg/server/internal"
	"github.com/Takasakiii/ayanami/pkg/server/internal/responses"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadFile(s internal.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		if file.Size == 0 {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Error: "file is empty",
			})
			return
		}

		if file.Size > 100*1024*1024 {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
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
		fileId, err := s.GetFile().UploadFile(filePkg.UploadFileData{
			File:       file,
			Password:   password,
			OriginalIp: ip,
			UserAgent:  userAgent,
		})
		if err != nil {
			log.Printf("[ERROR] upload file error: %v\n", err)
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: "upload file error",
			})
			return
		}

		fileUrl := fmt.Sprintf("%s/files/%s", s.GetConfig().ServerUrl, fileId)
		c.JSON(http.StatusOK, responses.UploadFileResponse{
			Url: fileUrl,
		})
	}
}

func DownloadFile(s internal.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileId, found := c.Params.Get("fileId")
		if !found {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Error: "fileId is missing",
			})
			return
		}

		password := c.Query("password")

		downloadedFile, downErr := s.GetFile().DownloadFile(fileId, password)
		if downErr != nil {
			var castedDownloadError *sender.DownloadError
			if errors.As(downErr, &castedDownloadError) && castedDownloadError.Type == sender.InvalidFileIdError {
				c.JSON(http.StatusBadRequest, responses.ErrorResponse{
					Error: "fileId is invalid",
				})
				return
			}

			log.Printf("[ERROR] Download file error: %v\n", downErr)
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: "download file error",
			})

			return
		}

		c.Header("Content-Type", downloadedFile.MimeType)
		c.Header("Accept-Length", fmt.Sprintf("%d", downloadedFile.Size))
		c.Header("Cache-Control", "public, max-age=31536000, s-maxage=31536000, immutable, stale-while-revalidate=86400, stale-if-error=2592000")
		c.Header("Vary", "Accept-Encoding")

		_, err := c.Writer.Write(downloadedFile.Content)
		if err != nil {
			log.Printf("[ERROR] Download file error: %v\n", err)
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: "download file error",
			})
			return
		}
	}
}
