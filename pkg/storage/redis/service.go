package redis

import (
	"airdrop/pkg/storage"
	"context"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	client *redis.Client
}

// NewService создает новый сервис Redis
func NewService(client *redis.Client) storage.RedisClient {
	return &Service{client: client}
}

func (s *Service) Set(key string, value interface{}) error {
	return s.client.Set(context.Background(), key, value, 0).Err()
}

func (s *Service) Get(key string) (interface{}, error) {
	return s.client.Get(context.Background(), key).Result()
}
