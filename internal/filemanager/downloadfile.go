package filemanager

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/crypt"
	"github.com/Takasakiii/ayanami/internal/file"
)

func (f *FileManager) DownloadFile(fileId string, password string) (*file.File, error) {
	downloadedFile, downloadErr := f.downloader.Download(fileId)
	if downloadErr != nil {
		return nil, downloadErr
	}

	if password == "" {
		password = f.config.Secret
	}

	decryptFile, err := crypt.Decrypt(password, downloadedFile.Content)
	if err != nil {
		return nil, fmt.Errorf("filemanager downloadfile decrypt: %v", err)
	}
	downloadedFile.Content = decryptFile
	downloadedFile.Size = int64(len(decryptFile))

	originalFile, err := f.restoreOriginalFile(downloadedFile)
	if err != nil {
		return nil, fmt.Errorf("filemanager downloadfile restoreoriginalfile: %v", err)
	}

	return originalFile, nil
}
