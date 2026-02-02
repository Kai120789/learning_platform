package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"learning-platform/auth/internal/dto"
	"time"
)

type CustomJwtClaims struct {
	UserId    int64  `json:"user_id"`
	UserEmail string `json:"user_email"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

func CreateJWT(createJwtDto dto.CreateJWT, log *zap.Logger) (*dto.TokenBundle, error) {
	sessionId := uuid.New().String()
	accessClaims := CustomJwtClaims{
		createJwtDto.UserId,
		createJwtDto.UserEmail,
		sessionId,
		jwt.RegisteredClaims{
			Issuer:    createJwtDto.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(createJwtDto.AccessTime) * time.Minute)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(createJwtDto.SignedKey))
	if err != nil {
		log.Error("sign access token error", zap.Error(err))
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    createJwtDto.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(createJwtDto.RefreshTime) * time.Hour * 24)),
		},
	)
	signedRefreshToken, err := refreshToken.SignedString([]byte(createJwtDto.SignedKey))
	if err != nil {
		log.Error("sign refresh token error", zap.Error(err))
		return nil, err
	}

	return &dto.TokenBundle{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
		SessionId:    sessionId,
	}, nil
}

func GetAccessTokenClaims(accessToken string, signedKey string) (*CustomJwtClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomJwtClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomJwtClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
