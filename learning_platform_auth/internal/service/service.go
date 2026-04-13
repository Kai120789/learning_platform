package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/dto"
	"learning-platform/auth/internal/utils"
	"time"
)

type AuthService struct {
	config *config.Config
	redis  RedisStorage
}

type RedisStorage interface {
	SetSession(userId int64, tokenBundle dto.TokenBundle, ttl time.Duration) error
	SetTokens(tokenBundle dto.TokenBundle, ttl time.Duration) error
	DeleteTokens(sessionId string, userId int64) error
	DeleteAllUserSessions(userId int64) error
}

func New(
	config *config.Config,
	redis RedisStorage,
) *AuthService {
	return &AuthService{
		config: config,
		redis:  redis,
	}
}

func (s *AuthService) Login(loginData dto.LoginRequest) (*dto.LoginResponse, error) {
	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      loginData.UserId,
		UserEmail:   loginData.Email,
		SignedKey:   s.config.SignedKey,
		Issuer:      s.config.Issuer,
		AccessTime:  s.config.AccessTokenLifeTime,
		RefreshTime: s.config.RefreshTokenLifeTime,
	})
	if err != nil {
		return nil, fmt.Errorf("create jwt tokens error: %w", err)
	}

	err = s.redis.SetSession(loginData.UserId, *tokenBundle, time.Duration(s.config.RefreshTokenLifeTime)*time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("set tokens error: %w", err)
	}

	return &dto.LoginResponse{
		SessionId: tokenBundle.SessionId,
		UserId:    loginData.UserId,
	}, nil
}

func (s *AuthService) Register(registerData dto.RegisterRequest) (*dto.RegisterResponse, error) {
	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      registerData.UserId,
		UserEmail:   registerData.Email,
		SignedKey:   s.config.SignedKey,
		Issuer:      s.config.Issuer,
		AccessTime:  s.config.AccessTokenLifeTime,
		RefreshTime: s.config.RefreshTokenLifeTime,
	})
	if err != nil {
		return nil, fmt.Errorf("create jwt tokens error: %w", err)
	}

	err = s.redis.SetSession(registerData.UserId, *tokenBundle, time.Duration(s.config.RefreshTokenLifeTime)*time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("set tokens error: %w", err)
	}

	return &dto.RegisterResponse{
		UserId:    registerData.UserId,
		SessionId: tokenBundle.SessionId,
	}, nil
}

func (s *AuthService) RefreshTokens(refreshToken string) (*string, error) {
	accessClaims, err := utils.GetTokenClaims(refreshToken, s.config.SignedKey)
	if err != nil {
		return nil, fmt.Errorf("get refresh token claims error: %w", err)
	}

	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      accessClaims.UserId,
		UserEmail:   accessClaims.UserEmail,
		SignedKey:   s.config.SignedKey,
		Issuer:      s.config.Issuer,
		AccessTime:  s.config.AccessTokenLifeTime,
		RefreshTime: s.config.RefreshTokenLifeTime,
		SessionId:   &accessClaims.SessionId,
	})
	if err != nil {
		return nil, fmt.Errorf("create jwt tokens error: %w", err)
	}

	err = s.redis.SetTokens(*tokenBundle, time.Duration(s.config.RefreshTokenLifeTime)*time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("set tokens error: %w", err)
	}

	return &tokenBundle.AccessToken, nil
}

func (s *AuthService) CheckPassword(password string, passwordHash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	if err != nil {
		return false, fmt.Errorf("attempt login with incorrect password: %w", err)
	}

	return true, nil
}

func (s *AuthService) GeneratePasswordHash(password string) (*string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), int(s.config.Salt))
	if err != nil {
		return nil, fmt.Errorf("generate hash for password error: %w", err)
	}

	resPasswordHash := string(passwordHash)

	return &resPasswordHash, nil
}

func (s *AuthService) Logout(accessToken string) error {
	accessClaims, err := utils.GetTokenClaims(accessToken, s.config.SignedKey)
	if err != nil {
		return fmt.Errorf("get access token claims error: %w", err)
	}

	err = s.redis.DeleteTokens(accessClaims.SessionId, accessClaims.UserId)
	if err != nil {
		return fmt.Errorf("delete token error: %w", err)
	}

	return nil
}

func (s *AuthService) LogoutAll(userId int64) error {
	return s.redis.DeleteAllUserSessions(userId)
}

func (s *AuthService) ChangePassword() {}

func (s *AuthService) ForceChangePassword() {}

func (s *AuthService) ChangeEmail() {}

func (s *AuthService) ForceChangeEmail() {}
