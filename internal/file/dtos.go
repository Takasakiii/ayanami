package file

import (
	"mime/multipart"
)

type AbstractFile struct {
	FileName string
	Size     int64
	MimeType string
	Content  []byte
}

type UploadFileData struct {
	File       *multipart.FileHeader
	Password   string
	OriginalIp string
	UserAgent  string
}
