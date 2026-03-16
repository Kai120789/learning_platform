package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"learning-platform/api-gateway/internal/dto"
	"learning-platform/api-gateway/internal/utils"
	"net/http"
	"time"
)

type CustomJwtClaims struct {
	UserId    int64  `json:"user_id"`
	UserEmail string `json:"user_email"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

type AuthService interface {
	RefreshTokens(refreshToken string) (*string, error)
	GetTokens(sessionId string) (*dto.RedisTokens, error)
}

func JWT(secretKey []byte, refreshTokenTTL int64, authService AuthService) func(http.Handler) http.Handler {
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

			redisTokens, err := authService.GetTokens(cookie.Value)
			if err != nil {
				http.Error(w, "tokens not found from redis", http.StatusNotFound)
				http.SetCookie(w, utils.DeleteCookie("session_id"))
				return
			}

			accessClaims := &CustomJwtClaims{}

			_, err = jwt.ParseWithClaims(redisTokens.AccessToken, accessClaims, jwtKey(secretKey))

			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					refreshClaims := &jwt.RegisteredClaims{}

					refreshToken, err := jwt.ParseWithClaims(redisTokens.RefreshToken, refreshClaims, jwtKey(secretKey))

					if err != nil || !refreshToken.Valid {
						http.Error(w, "invalid refresh token", http.StatusUnauthorized)
						http.SetCookie(w, utils.DeleteCookie("session_id"))
						return
					}

					newAccessToken, err := authService.RefreshTokens(redisTokens.RefreshToken)
					if err != nil {
						http.Error(w, "refresh tokens error", http.StatusInternalServerError)
						http.SetCookie(w, utils.DeleteCookie("session_id"))
						return
					}

					_, err = jwt.ParseWithClaims(*newAccessToken, accessClaims, jwtKey(secretKey))
					if err != nil {
						http.Error(w, "invalid new access token", http.StatusUnauthorized)
						http.SetCookie(w, utils.DeleteCookie("session_id"))
						return
					}

					cookie := utils.CreateCookie(
						"session_id",
						accessClaims.SessionId,
						time.Now().Add(time.Duration(refreshTokenTTL)*time.Hour*24),
					)
					http.SetCookie(w, cookie)
				} else {
					http.Error(w, "invalid access token", http.StatusUnauthorized)
					http.SetCookie(w, utils.DeleteCookie("session_id"))
					return
				}
			}

			ctx := context.WithValue(r.Context(), "user_id", accessClaims.UserId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func jwtKey(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	}
}
