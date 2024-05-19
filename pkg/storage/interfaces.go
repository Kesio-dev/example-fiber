package storage

import "mime/multipart"

// RedisClient интерфейс для операций с Redis
type RedisClient interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

// MinioClient интерфейс для операций с MinIO
type MinioClient interface {
	UploadFile(fileData multipart.File, filename string) error
	GetFile(key string) ([]byte, error)
}
