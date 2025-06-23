package filemanager

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
)

func (f FileManager) addFileInfo(fileInput *file.File) (*file.File, error) {
	finalFile := *fileInput

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(finalFile); err != nil {
		return nil, fmt.Errorf("filemodifier addFileInfo: encoder:  %v", err)
	}

	b := buf.Bytes()

	return &file.File{
		FileName: f.cuid(),
		Size:     int64(len(b)),
		MimeType: "application/x-tar",
		Content:  b,
	}, nil
}

func (f FileManager) restoreOriginalFile(fileInput *file.File) (*file.File, error) {
	var finalFile file.File

	buf := bytes.NewReader(fileInput.Content)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&finalFile); err != nil {
		return nil, fmt.Errorf("filemodifier restoreOriginalFile: decoder:  %v", err)
	}

	return &finalFile, nil
}
