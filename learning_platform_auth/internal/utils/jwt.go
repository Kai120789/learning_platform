package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"learning-platform/auth/internal/dto"
	"time"
)

type CustomJwtClaims struct {
	UserID    int64  `json:"user_id"`
	UserEmail string `json:"user_email"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

func CreateJWT(createJwtDto dto.CreateJWT) (*dto.TokenBundle, error) {
	var sessionId string
	if createJwtDto.SessionID != nil {
		sessionId = *createJwtDto.SessionID
	} else {
		sessionId = uuid.New().String()
	}
	accessClaims := CustomJwtClaims{
		createJwtDto.UserID,
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
		return nil, fmt.Errorf("signed access token for user %d: %w", createJwtDto.UserID, err)
	}
	refreshClaims := CustomJwtClaims{
		createJwtDto.UserID,
		createJwtDto.UserEmail,
		sessionId,
		jwt.RegisteredClaims{
			Issuer:    createJwtDto.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(createJwtDto.RefreshTime) * time.Hour * 24)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(createJwtDto.SignedKey))
	if err != nil {
		return nil, fmt.Errorf("signed refresh token for user %d: %w", createJwtDto.UserID, err)
	}

	return &dto.TokenBundle{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
		SessionID:    sessionId,
	}, nil
}

func GetTokenClaims(JWTToken string, signedKey string) (*CustomJwtClaims, error) {
	token, err := jwt.ParseWithClaims(JWTToken, &CustomJwtClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomJwtClaims)
	if !ok {
		return nil, fmt.Errorf("get token claims for user %d: %w", claims.UserID, err)
	}

	return claims, nil
}
