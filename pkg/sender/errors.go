package sender

import "fmt"

type DownloadErrorType int

const (
	InvalidFileIdError DownloadErrorType = iota
	FailedToDownloadFileError
)

type DownloadError struct {
	Type DownloadErrorType
	Err  error
}

func (e *DownloadError) Error() string {
	return fmt.Sprintf("download error: [%d] %v", e.Type, e.Err)
}
