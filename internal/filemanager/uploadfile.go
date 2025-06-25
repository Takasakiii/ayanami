package filemanager

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/crypt"
	"github.com/Takasakiii/ayanami/internal/database"
	"github.com/Takasakiii/ayanami/internal/file"
	"mime/multipart"
)

func (f *FileManager) UploadFile(
	fileRaw *multipart.FileHeader,
	password string,
	originIp string,
	userAgent string) (string, error) {

	if password == "" {
		password = f.config.Secret
	}

	absFile, err := file.ToFile(fileRaw)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile tofile: %v", err)
	}

	fileWithMeta, err := f.addFileInfo(absFile)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile addfileinfo: %v", err)
	}

	fileWithMeta.Content, err = crypt.Encrypt(password, fileWithMeta.Content)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile encrypt: %v", err)
	}
	fileWithMeta.Size = int64(len(fileWithMeta.Content))

	contentUrl, err := f.sender.Send(fileWithMeta)
	if err != nil {
		return "", fmt.Errorf("filemanager uploadfile send: %v", err)
	}

	model := database.File{
		Ip:        originIp,
		FileName:  contentUrl,
		UserAgent: userAgent,
		MimeType:  absFile.MimeType,
	}

	f.database.Create(&model)
	return contentUrl, nil
}
