package sender

import (
	"github.com/Takasakiii/ayanami/internal/file"
)

type Sender interface {
	Send(file *file.AbstractFile) (string, error)
}

type Downloader interface {
	Download(fileId string) (*file.AbstractFile, *DownloadError)
}
