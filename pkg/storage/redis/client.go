package redis

import (
	"github.com/redis/go-redis/v9"
	"os"
)

// NewClient создает новый клиент Redis
func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // используйте номер DB по умолчанию
	})
}
