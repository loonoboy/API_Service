package redisDB

import (
	"API_Service/internal/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisDB(cfg config.RDB) (*redis.Client, error) {
	const op = "repository.redisDB.NewRedisDB"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return client, nil
}
