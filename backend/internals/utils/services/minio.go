package utilServices

import (
	"context"
	"io"
	"mime/multipart"
)

type MinioService interface {
	PutObject(ctx context.Context, bucketName string, objectName string, reader io.Reader, fileHeader *multipart.FileHeader) error
}
