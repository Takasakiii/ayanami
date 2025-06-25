package service

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/crypt"
)

func (s *FileService) DownloadFile(fileId string, password string) (*file.AbstractFile, error) {
	downloadedFile, downloadErr := s.downloader.Download(fileId)
	if downloadErr != nil {
		return nil, downloadErr
	}

	if password == "" {
		password = s.config.Secret
	}

	decryptFile, err := crypt.Decrypt(password, downloadedFile.Content)
	if err != nil {
		return nil, fmt.Errorf("filemanager downloadfile decrypt: %v", err)
	}
	downloadedFile.Content = decryptFile
	downloadedFile.Size = int64(len(decryptFile))

	originalFile, err := s.restoreOriginalFile(downloadedFile)
	if err != nil {
		return nil, fmt.Errorf("filemanager downloadfile restoreoriginalfile: %v", err)
	}

	return originalFile, nil
}
