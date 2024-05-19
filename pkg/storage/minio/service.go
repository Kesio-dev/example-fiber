package minio

import (
	"airdrop/pkg/storage"
	"bytes"
	"github.com/gofiber/storage/minio"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"strings"
)

type Service struct {
	client *minio.Storage
}

// NewService создает новый сервис MinIO
func NewService(client *minio.Storage) storage.MinioClient {
	return &Service{client: client}
}

func (s *Service) UploadFile(fileData multipart.File, filename string) error {
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, fileData); err != nil {
		return err
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	memType := strings.Split(filename, ".")
	filename = newUUID.String() + "." + memType[len(memType)-1]

	return s.client.Set(filename, buffer.Bytes(), 0)
}

func (s *Service) GetFile(key string) ([]byte, error) {
	return s.client.Get(key)
}
