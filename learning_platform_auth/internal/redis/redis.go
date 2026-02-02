package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"learning-platform/auth/internal/dto"
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

func (r *RedisStorage) SetTokens(userId int64, tokenBundle dto.TokenBundle) error {
	_, err := r.client.HSet(
		context.Background(),
		fmt.Sprintf("tokenBundle:%s", tokenBundle.AccessToken),
		"refresh_token", tokenBundle.RefreshToken,
		"session_id", tokenBundle.SessionId,
	).Result()
	if err != nil {
		r.logger.Error("error set tokens to redis", zap.Error(err))
		return err
	}

	redisTokens, err := r.client.HGet(
		context.Background(),
		fmt.Sprintf("userAccess:%d", userId),
		"access_tokens",
	).Result()
	if errors.Is(err, redis.Nil) {
		err := r.setUserAccess(userId, []string{tokenBundle.AccessToken})
		if err != nil {
			r.logger.Error("error set first token to user-access to redis", zap.Error(err))
			return err
		}
	} else if err != nil {
		r.logger.Error("error get user access tokens from redis", zap.Error(err))
		return err
	} else {
		var tokens []string
		err = json.Unmarshal([]byte(redisTokens), &tokens)
		if err != nil {
			r.logger.Error("error unmarshal tokens from redis", zap.Error(err))
			return err
		}

		tokens = append(tokens, tokenBundle.AccessToken)
		err := r.setUserAccess(userId, tokens)
		if err != nil {
			r.logger.Error("error set new tokens in user-access to redis", zap.Error(err))
			return err
		}
	}
	return nil
}

func (r *RedisStorage) setUserAccess(userId int64, accessTokens []string) error {
	jsonTokens, err := json.Marshal(accessTokens)
	if err != nil {
		r.logger.Error("marshal body with tokens error", zap.Error(err))
		return err
	}

	_, err = r.client.HSet(
		context.Background(),
		fmt.Sprintf("userAccess:%d", userId),
		"access_tokens",
		jsonTokens,
	).Result()
	if err != nil {
		r.logger.Error("error set tokens in user-access to redis", zap.Error(err))
		return err
	}
	return nil
}
