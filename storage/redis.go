package storage

import (
	"github.com/redis/go-redis/v9"
	"movieWishlistAPI/cfg"
)

type RedisStorage struct {
	RedisClient *redis.Client
}

func NewRedisClient(config *cfg.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})
}
func (r *RedisStorage) Close() {
	r.RedisClient.Close()
}
