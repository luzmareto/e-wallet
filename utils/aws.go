package utils

import (
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSS3Client interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (uploadedFile string, errUpload error)
}

type awsS3Client struct {
	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
	env        *Config
}

func NewAWSS3Client(env Config) AWSS3Client {
	cfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				env.AwsAccessKeyID,
				env.AwsSecretKeyID,
				"",
			),
		))

	cfg.Region = env.AwsRegion
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	downloader := manager.NewDownloader(client)

	return &awsS3Client{
		s3Client:   client,
		uploader:   uploader,
		downloader: downloader,
		env:        &env,
	}
}

// Upload implements AWSS3Client.
func (awsClient *awsS3Client) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (uploadedFile string, errUpload error) {
	// Generate a random file name
	filename := RandomFileName(file)

	mime, err := GetMIMEType(file)
	if err != nil {
		errUpload = &CustomError{
			Err: sql.ErrConnDone,
			Msg: fmt.Sprintf("invalid file mimetype %s", mime),
		}
		return
	}

	if !IsAllowedFileType(mime) {
		errUpload = &CustomError{
			Err: sql.ErrConnDone,
			Msg: fmt.Sprintf("unsuported file mimetype %s, only jpg, jpeg or png ", mime),
		}
		return
	}

	f, errUpload := file.Open()
	if errUpload != nil {
		return
	}

	defer f.Close()

	objectKey := fmt.Sprintf("%s/%s", folder, filename)
	_, errUpload = awsClient.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(awsClient.env.AwsS3Bucket),
		Key:    aws.String(objectKey),
		Body:   f,
	})
	uploadedFile = filename
	return
}
