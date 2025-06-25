package sender

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/pkg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nrednav/cuid2"
	"net/http"
)

type S3Sender struct {
	client *s3.Client
	config *config.S3
	cuid   func() string
}

func NewS3Sender(conf *config.S3) (S3Sender, error) {
	cfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			conf.AccessKeyId,
			conf.SecretAccessKey, ""),
		),
		awsConfig.WithRegion(conf.Region))

	if err != nil {
		return S3Sender{}, fmt.Errorf("s3 sender loaddefaultconfig: %v", err)
	}

	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.BaseEndpoint = aws.String(conf.Endpoint)
	})

	cuid, err := cuid2.Init()
	if err != nil {
		return S3Sender{}, fmt.Errorf("s3 sender init cuid: %v", err)
	}

	return S3Sender{client, conf, cuid}, nil
}

func (s *S3Sender) Send(file *file.AbstractFile) (string, error) {
	ctx := context.Background()
	fileName := fmt.Sprintf("%s_%s", s.cuid(), file.FileName)

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(file.Content),
	})

	if err != nil {
		return "", fmt.Errorf("s3 sender putobject: %v", err)
	}

	return fileName, nil
}

func (s *S3Sender) Download(fileId string) (*file.AbstractFile, *DownloadError) {
	url := fmt.Sprintf("%s/%s", s.config.BucketPublicUrl, fileId)
	response, err := http.Get(url)
	if err != nil {
		return nil, &DownloadError{
			Type: FailedToDownloadFileError,
			Err:  fmt.Errorf("s3 download: http response: %v", err),
		}
	}

	return processFileFromRequest(response, fileId)
}
