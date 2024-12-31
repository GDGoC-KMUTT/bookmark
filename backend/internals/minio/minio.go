package minio

import (
	"backend/internals/config"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"log"
	"net/url"
	"strings"
)

var MinioClient *minio.Client

func SetUpMinio() {
	// Parse the URL
	parsedURL, err := url.Parse(*config.Env.MinioS3Endpoint)
	if err != nil {
		fmt.Println("Error parsing minio URL:", err)
		return
	}

	// Get the hostname (e.g., test.com)
	endpoint := parsedURL.Host

	// Strip any port if present (e.g., test.com:443 -> test.com)
	endpoint = strings.Split(endpoint, ":")[0]
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(*config.Env.MinioS3AccessKey, *config.Env.MinioS3SecretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	logrus.Infof("[MINIO] Minio client started successfully.")

	MinioClient = minioClient
}
