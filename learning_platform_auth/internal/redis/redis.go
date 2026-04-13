package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/dto"
	"time"
)

type RedisStorage struct {
	client *redis.Client
	config *config.Config
}

func New(
	client *redis.Client,
	config *config.Config,
) *RedisStorage {
	return &RedisStorage{
		client: client,
		config: config,
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

func (r *RedisStorage) SetSession(userId int64, tokenBundle dto.TokenBundle, ttl time.Duration) error {
	err := r.SetTokens(tokenBundle, ttl)
	if err != nil {
		return fmt.Errorf("set tokens to redis for user %d: %w", userId, err)
	}

	redisSessions, err := r.client.HGet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
	).Result()
	if errors.Is(err, redis.Nil) {
		err := r.setUserSessions(userId, []string{tokenBundle.SessionId})
		if err != nil {
			return fmt.Errorf("set first session to redis for user %d: %w", userId, err)
		}
	} else if err != nil {
		return fmt.Errorf("get user %d sessions from redis: %w", userId, err)
	} else {
		var sessions []string
		err = json.Unmarshal([]byte(redisSessions), &sessions)
		if err != nil {
			return fmt.Errorf("unmarshal sessions from redis for user %d: %w", userId, err)
		}

		sessions = append(sessions, tokenBundle.SessionId)
		err := r.setUserSessions(userId, sessions)
		if err != nil {
			return fmt.Errorf("set new session to redis for user %d: %w", userId, err)
		}
	}
	return nil
}

func (r *RedisStorage) SetTokens(tokenBundle dto.TokenBundle, ttl time.Duration) error {
	_, err := r.client.HSet(
		context.Background(),
		fmt.Sprintf("tokenBundle:%s", tokenBundle.SessionId),
		"access_token", tokenBundle.AccessToken,
		"refresh_token", tokenBundle.RefreshToken,
	).Result()
	if err != nil {
		return fmt.Errorf("set tokens to redis: %w", err)
	}

	r.client.Expire(context.Background(), fmt.Sprintf("tokenBundle:%s", tokenBundle.SessionId), ttl)

	return nil
}

func (r *RedisStorage) DeleteTokens(sessionId string, userId int64) error {
	_, err := r.client.Del(
		context.Background(),
		fmt.Sprintf("tokenBundle:%s", sessionId),
	).Result()
	if err != nil {
		return fmt.Errorf("delete tokens from redis by session %s for user %d: %w", sessionId, userId, err)
	}

	redisSessions, err := r.client.HGet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
	).Result()
	if err != nil {
		return fmt.Errorf("get user %d sessions from redis: %w", userId, err)
	}

	var sessions []string
	err = json.Unmarshal([]byte(redisSessions), &sessions)
	if err != nil {
		return fmt.Errorf("unmarshal tokens from redis for user %d: %w", userId, err)
	}

	for i := 0; i < len(sessions); i++ {
		if sessions[i] == sessionId {
			sessions = append(sessions[:i], sessions[i+1:]...)
		}
	}

	err = r.setUserSessions(userId, sessions)
	if err != nil {
		return fmt.Errorf("set new sessions to redis for user %d: %w", userId, err)
	}

	return nil
}

func (r *RedisStorage) DeleteAllUserSessions(userId int64) error {
	redisSessions, err := r.client.HGet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
	).Result()
	if err != nil {
		return fmt.Errorf("get user %d sessions from redis: %w", userId, err)
	}

	var sessions []string
	err = json.Unmarshal([]byte(redisSessions), &sessions)
	if err != nil {
		return fmt.Errorf("unmarshal tokens from redis for user %d: %w", userId, err)
	}

	for _, session := range sessions {
		err := r.DeleteTokens(session, userId)
		if err != nil {
			return fmt.Errorf("delete tokens for user %d: %w", userId, err)
		}
	}

	_, err = r.client.Del(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
	).Result()
	if err != nil {
		return fmt.Errorf("delete sessions for user %d: %w", userId, err)
	}

	return nil
}

func (r *RedisStorage) setUserSessions(userId int64, sessions []string) error {
	jsonSessions, err := json.Marshal(sessions)
	if err != nil {
		return fmt.Errorf("marshal body with tokens for user %d: %w", userId, err)
	}

	_, err = r.client.HSet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
		jsonSessions,
	).Result()
	if err != nil {
		return fmt.Errorf("set tokens to redis for user %d: %w", userId, err)
	}
	return nil
}
