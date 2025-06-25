package service

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"io"
	"mime/multipart"
)

func toFile(formFile *multipart.FileHeader) (*file.AbstractFile, error) {
	openedFile, err := formFile.Open()
	defer func(openedFile multipart.File) {
		_ = openedFile.Close()
	}(openedFile)

	if err != nil {
		return nil, fmt.Errorf("file tofile open: %v", err)
	}

	data, err := io.ReadAll(openedFile)
	if err != nil {
		return nil, fmt.Errorf("file tofile read: %v", err)
	}

	abstractFile := file.AbstractFile{
		FileName: formFile.Filename,
		Size:     formFile.Size,
		MimeType: formFile.Header.Get("Content-Type"),
		Content:  data,
	}

	return &abstractFile, nil
}
