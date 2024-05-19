package minio

import (
	"github.com/gofiber/storage/minio"
	"os"
)

func NewClient() *minio.Storage {
	return minio.New(minio.Config{
		Bucket:   os.Getenv("MINIO_BUCKET"),
		Endpoint: os.Getenv("MINIO_ENDPOINT"),
		Credentials: minio.Credentials{
			AccessKeyID:     os.Getenv("MINIO_USER"),
			SecretAccessKey: os.Getenv("MINIO_PASSWORD"),
		},
	})
}
