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
		return fmt.Errorf("error set tokens to redis: %w", err)
	}

	redisSessions, err := r.client.HGet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
	).Result()
	if errors.Is(err, redis.Nil) {
		err := r.setUserSessions(userId, []string{tokenBundle.SessionId})
		if err != nil {
			return fmt.Errorf("error set first session to user-session to redis: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error get user sessions from redis: %w", err)
	} else {
		var sessions []string
		err = json.Unmarshal([]byte(redisSessions), &sessions)
		if err != nil {
			return fmt.Errorf("error unmarshal sessions from redis: %w", err)
		}

		sessions = append(sessions, tokenBundle.SessionId)
		err := r.setUserSessions(userId, sessions)
		if err != nil {
			return fmt.Errorf("error set new session in user-sessions to redis: %w", err)
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
		return fmt.Errorf("error set tokens to redis: %w", err)
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
		return fmt.Errorf("error delete tokens from redis by session %s: %w", sessionId, err)
	}

	redisSessions, err := r.client.HGet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
	).Result()
	if err != nil {
		return fmt.Errorf("error get user sessions from redis: %w", err)
	}

	var sessions []string
	err = json.Unmarshal([]byte(redisSessions), &sessions)
	if err != nil {
		return fmt.Errorf("error unmarshal tokens from redis: %w", err)
	}

	for i := 0; i < len(sessions); i++ {
		if sessions[i] == sessionId {
			sessions = append(sessions[:i], sessions[i+1:]...)
		}
	}

	err = r.setUserSessions(userId, sessions)
	if err != nil {
		return fmt.Errorf("error set new sessions in user-sessions to redis: %w", err)
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
		return fmt.Errorf("error get user sessions from redis: %w", err)
	}

	var sessions []string
	err = json.Unmarshal([]byte(redisSessions), &sessions)

	for _, session := range sessions {
		err := r.DeleteTokens(session, userId)
		if err != nil {
			return fmt.Errorf("error delete tokens: %w", err)
		}
	}

	_, err = r.client.Del(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
	).Result()
	if err != nil {
		return fmt.Errorf("failed delete sessions for user %d: %w", userId, err)
	}

	return nil
}

func (r *RedisStorage) setUserSessions(userId int64, sessions []string) error {
	jsonSessions, err := json.Marshal(sessions)
	if err != nil {
		return fmt.Errorf("marshal body with tokens error: %w", err)
	}

	_, err = r.client.HSet(
		context.Background(),
		fmt.Sprintf("userSessions:%d", userId),
		"session_id",
		jsonSessions,
	).Result()
	if err != nil {
		return fmt.Errorf("error set tokens in user-sessions to redis: %w", err)
	}
	return nil
}
