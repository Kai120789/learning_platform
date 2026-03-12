package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type CustomJwtClaims struct {
	UserId    int64  `json:"user_id"`
	UserEmail string `json:"user_email"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

func JWT(secretKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")
			if err != nil {
				http.Error(w, "cookie not found", http.StatusNotFound)
				return
			}

			if cookie.Value == "" {
				http.Error(w, "incorrect cookie value", http.StatusBadRequest)
				return
			}

			claims := &CustomJwtClaims{}

			token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (any, error) {
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
