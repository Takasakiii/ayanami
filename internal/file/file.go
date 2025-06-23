package file

import (
	"fmt"
	"io"
	"mime/multipart"
)

type File struct {
	FileName string
	Size     int64
	MimeType string
	Content  []byte
}

func ToFile(formFile *multipart.FileHeader) (File, error) {
	openedFile, err := formFile.Open()
	defer func(openedFile multipart.File) {
		_ = openedFile.Close()
	}(openedFile)

	if err != nil {
		return File{}, fmt.Errorf("file tofile open: %v", err)
	}

	data, err := io.ReadAll(openedFile)
	if err != nil {
		return File{}, fmt.Errorf("file tofile read: %v", err)
	}

	file := File{
		FileName: formFile.Filename,
		Size:     formFile.Size,
		MimeType: formFile.Header.Get("Content-Type"),
		Content:  data,
	}

	return file, nil
}
