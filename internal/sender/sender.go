package sender

import (
	"github.com/Takasakiii/ayanami/internal/file"
)

type Sender interface {
	Send(file *file.File) (string, error)
	Download(fileId string) (*file.File, *DownloadError)
}
