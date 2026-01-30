package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type TokenBundle struct {
	RefreshToken string `redis:"refresh_token"`
	SessionId    string `redis:"session_id"`
}

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

func (r *RedisStorage) SetTokens(
	accessToken string,
	refreshToken string,
	sessionId string,
) error {
	body, err := json.Marshal(&TokenBundle{
		RefreshToken: refreshToken,
		SessionId:    sessionId,
	})
	_, err = r.client.HSet(
		context.Background(),
		fmt.Sprintf("tokenBundle:%s", accessToken),
		body,
		15*time.Minute,
	).Result()

	if err != nil {
		r.logger.Error("error set tokens to redis", zap.Error(err))
		return err
	}

	return nil
}
