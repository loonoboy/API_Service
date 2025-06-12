package redisDB

import (
	"API_Service/internal/dto"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type ArticleRedis struct {
	rdb *redis.Client
}

func NewArticleRedis(rdb *redis.Client) *ArticleRedis {
	return &ArticleRedis{rdb: rdb}
}

const keyRedis = "articles:recent"

func (r ArticleRedis) GetArticles() ([]dto.Article, error) {
	const op = "repository.redisDB.article_redis.GetArticles"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	val, err := r.rdb.Get(ctx, keyRedis).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
	}

	var articles []dto.Article
	if err := json.Unmarshal([]byte(val), &articles); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return articles, nil
}

func (r ArticleRedis) SetRecentArticles(articles []dto.Article) error {
	const op = "repository.redisDB.article_redis.SetRecentArticles"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	data, err := json.Marshal(articles)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = r.rdb.Set(ctx, keyRedis, data, 5*time.Minute).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("%s: completed with timeout", op)
		} else {
			return fmt.Errorf("%s: %w", op, err)
		}
	}
	return nil
}
