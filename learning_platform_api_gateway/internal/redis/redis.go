package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"learning-platform/api-gateway/internal/dto/authDto"
)

type RedisStorage struct {
	client *redis.Client
}

func New(client *redis.Client) *RedisStorage {
	return &RedisStorage{
		client: client,
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

func (r *RedisStorage) GetTokens(sessionId string) (*authDto.RedisTokens, error) {
	tokenBundle, err := r.client.HGetAll(
		context.Background(),
		fmt.Sprintf("tokenBundle:%s", sessionId),
	).Result()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(tokenBundle)
	if err != nil {
		return nil, err
	}

	var redisTokens authDto.RedisTokens
	err = json.Unmarshal(data, &redisTokens)

	return &redisTokens, nil
}
