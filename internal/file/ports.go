package file

type Service interface {
	DownloadFile(fileId string, password string) (*AbstractFile, error)
	UploadFile(data UploadFileData) (string, error)
}

type Repository interface {
	AddFile(data *File) error
}
