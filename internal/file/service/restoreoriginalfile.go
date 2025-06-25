package service

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
)

func (s *FileService) restoreOriginalFile(fileInput *file.AbstractFile) (*file.AbstractFile, error) {
	var finalFile file.AbstractFile

	buf := bytes.NewReader(fileInput.Content)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&finalFile); err != nil {
		return nil, fmt.Errorf("filemodifier restoreOriginalFile: decoder:  %v", err)
	}

	return &finalFile, nil
}
