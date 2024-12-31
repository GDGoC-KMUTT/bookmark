package utilServices

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"mime/multipart"
)

type minioService struct {
	minioClient *minio.Client
}

func NewMinioService(minioClient *minio.Client) MinioService {
	return &minioService{
		minioClient: minioClient,
	}
}

func (r *minioService) PutObject(ctx context.Context, bucketName string, objectName string, reader io.Reader, fileHeader *multipart.FileHeader) error {
	_, err := r.minioClient.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")},
	)
	if err != nil {
		return err
	}
	return nil
}
