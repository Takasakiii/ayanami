package file

import "time"

type Service interface {
	DownloadFile(fileId string, password string) (*AbstractFile, error)
	UploadFile(data UploadFileData) (string, error)
	DeleteExpired() error
}

type Repository interface {
	AddFile(data *File) error
	DeleteExpired(maxTime time.Time) error
	GetExpired(maxTime time.Time) ([]string, error)
}
