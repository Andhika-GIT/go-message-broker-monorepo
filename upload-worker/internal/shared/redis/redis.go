package redispubsub

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *shared.RedisClientConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
