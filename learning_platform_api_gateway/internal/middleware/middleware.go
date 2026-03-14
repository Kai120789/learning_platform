package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"learning-platform/api-gateway/internal/redis"
	"net/http"
	"time"
)

type CustomJwtClaims struct {
	UserId    int64  `json:"user_id"`
	UserEmail string `json:"user_email"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

func JWT(secretKey []byte, redis *redis.RedisStorage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				http.Error(w, "session cookie not found", http.StatusNotFound)
				return
			}

			if cookie.Value == "" {
				http.Error(w, "incorrect cookie value", http.StatusBadRequest)
				return
			}

			redisTokens, err := redis.GetTokens(cookie.Value)
			if err != nil {
				http.Error(w, "tokens not found from redis", http.StatusNotFound)
			}

			claims := &CustomJwtClaims{}

			token, err := jwt.ParseWithClaims(redisTokens.AccessToken, claims, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
				http.Error(w, "token expired", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", claims.UserId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
