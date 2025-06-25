package sender

import (
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"io"
	"net/http"
)

func processFileFromRequest(res *http.Response, fileName string) (*file.AbstractFile, *DownloadError) {
	body, err := io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		var bodyResponse string
		if err == nil {
			bodyResponse = string(body)
		}

		return nil, &DownloadError{
			Type: FailedToDownloadFileError,
			Err:  fmt.Errorf("filebin download: http response: [%d] %s", res.StatusCode, bodyResponse),
		}
	}

	return &file.AbstractFile{
		FileName: fileName,
		Size:     int64(len(body)),
		MimeType: res.Header.Get("content-type"),
		Content:  body,
	}, nil
}
