package sender

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/config"
	"github.com/nrednav/cuid2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FileBin struct {
	config     *config.FileBin
	cuid       func() string
	httpClient *http.Client
}

func NewFileBin(config *config.FileBin) (*FileBin, error) {
	cuid, err := cuid2.Init()
	if err != nil {
		return nil, err
	}

	return &FileBin{
		config:     config,
		cuid:       cuid,
		httpClient: &http.Client{},
	}, nil
}

func (f *FileBin) Send(file *file.AbstractFile) (string, error) {
	bin := f.cuid()
	cid := f.cuid()
	fileName := url.QueryEscape(file.FileName)

	urlFileBin := fmt.Sprintf("%s/%s/%s", f.config.BaseUrl, bin, fileName)

	reader := bytes.NewReader(file.Content)
	req, err := http.NewRequest(http.MethodPost, urlFileBin, reader)
	if err != nil {
		return "", fmt.Errorf("filebin send: new request: %v", err)
	}

	req.ContentLength = file.Size
	req.TransferEncoding = nil
	req.Header.Add("content-type", "application/octet-stream")
	req.Header.Add("cid", fmt.Sprintf("Ayanami-%s", cid))

	res, err := f.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("filebin send: http request: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(res.Body)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		var bodyResponse string
		if err == nil {
			bodyResponse = string(body)
		}

		return "", fmt.Errorf("filebin send: http response: [%d] %s", res.StatusCode, bodyResponse)
	}

	return fmt.Sprintf("%s_%s", bin, fileName), nil
}

func (f *FileBin) Download(fileId string) (*file.AbstractFile, *DownloadError) {
	binFileName := strings.Split(fileId, "_")
	if len(binFileName) != 2 {
		return nil, &DownloadError{
			Type: InvalidFileIdError,
			Err:  errors.New("filebin download: invalid file id"),
		}
	}

	fileUrl := fmt.Sprintf("%s/%s/%s", f.config.BaseUrl, binFileName[0], binFileName[1])

	req, err := http.NewRequest(http.MethodGet, fileUrl, nil)
	if err != nil {
		return nil, &DownloadError{
			Type: FailedToDownloadFileError,
			Err:  fmt.Errorf("filebin download: new request: %v", err),
		}
	}

	req.AddCookie(&http.Cookie{
		Name:  "verified",
		Value: "2024-05-24",
	})

	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, &DownloadError{
			Type: FailedToDownloadFileError,
			Err:  fmt.Errorf("filebin download: http response: %v", err),
		}
	}

	var finalFileName string
	if finalFileName, err = url.QueryUnescape(binFileName[1]); err != nil {
		finalFileName = binFileName[1]
	}

	return processFileFromRequest(res, finalFileName)
}
