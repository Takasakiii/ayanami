package service

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/crypt"
)

func (s *FileService) UploadFile(data file.UploadFileData) (string, error) {

	if data.Password == "" {
		data.Password = s.config.Secret
	}

	absFile, err := toFile(data.File)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile tofile: %v", err)
	}

	fileWithMeta, err := s.addFileInfo(absFile)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile addfileinfo: %v", err)
	}

	fileWithMeta.Content, err = crypt.Encrypt(data.Password, fileWithMeta.Content)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile encrypt: %v", err)
	}
	fileWithMeta.Size = int64(len(fileWithMeta.Content))

	contentUrl, err := s.sender.Send(fileWithMeta)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile send: %v", err)
	}

	model := file.File{
		Ip:        data.OriginalIp,
		FileName:  contentUrl,
		UserAgent: data.UserAgent,
		MimeType:  absFile.MimeType,
	}

	err = s.repository.AddFile(&model)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile addfile: %v", err)
	}

	return contentUrl, nil
}
