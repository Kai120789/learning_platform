package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto"
)

type RedisStorage struct {
	client *redis.Client
	logger *zap.Logger
}

func New(client *redis.Client, logger *zap.Logger) *RedisStorage {
	return &RedisStorage{
		client: client,
		logger: logger,
	}
}

func Connection(redisUrl string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func (r *RedisStorage) GetTokens(sessionId string) (*dto.RedisTokens, error) {
	tokenBundle, err := r.client.HGetAll(
		context.Background(),
		fmt.Sprintf("tokenBundle:%s", sessionId),
	).Result()
	if err != nil {
		r.logger.Error("error get tokens to redis", zap.Error(err))
		return nil, err
	}

	data, err := json.Marshal(tokenBundle)
	if err != nil {
		r.logger.Error("marshal map to json error", zap.Error(err))
		return nil, err
	}

	var redisTokens dto.RedisTokens
	err = json.Unmarshal(data, &redisTokens)

	return &redisTokens, nil
}
