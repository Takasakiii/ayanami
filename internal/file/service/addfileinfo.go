package service

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
)

func (s *FileService) addFileInfo(fileInput *file.AbstractFile) (*file.AbstractFile, error) {
	finalFile := *fileInput

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(finalFile); err != nil {
		return nil, fmt.Errorf("filemodifier addFileInfo: encoder:  %v", err)
	}

	b := buf.Bytes()

	return &file.AbstractFile{
		FileName: s.cuid.Generate(),
		Size:     int64(len(b)),
		MimeType: "application/x-tar",
		Content:  b,
	}, nil
}
