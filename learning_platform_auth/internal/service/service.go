package service

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/dto"
	"learning-platform/auth/internal/utils"
	"time"
)

type AuthService struct {
	config *config.Config
	logger *zap.Logger
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
	logger *zap.Logger,
	redis RedisStorage,
) *AuthService {
	return &AuthService{
		config: config,
		logger: logger,
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
	}, s.logger)
	if err != nil {
		s.logger.Error("create jwt tokens error", zap.Error(err))
		return nil, err
	}

	err = s.redis.SetSession(loginData.UserId, *tokenBundle, time.Duration(s.config.RefreshTokenLifeTime)*time.Hour*24)
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
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
	}, s.logger)
	if err != nil {
		s.logger.Error("create jwt tokens error", zap.Error(err))
		return nil, err
	}

	err = s.redis.SetSession(registerData.UserId, *tokenBundle, time.Duration(s.config.RefreshTokenLifeTime)*time.Hour*24)
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
	}

	return &dto.RegisterResponse{
		UserId:    registerData.UserId,
		SessionId: tokenBundle.SessionId,
	}, nil
}

func (s *AuthService) RefreshTokens(refreshToken string) (*string, error) {
	accessClaims, err := utils.GetTokenClaims(refreshToken, s.config.SignedKey)
	if err != nil {
		s.logger.Error("get refresh token claims error", zap.Error(err))
		return nil, err
	}

	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      accessClaims.UserId,
		UserEmail:   accessClaims.UserEmail,
		SignedKey:   s.config.SignedKey,
		Issuer:      s.config.Issuer,
		AccessTime:  s.config.AccessTokenLifeTime,
		RefreshTime: s.config.RefreshTokenLifeTime,
		SessionId:   &accessClaims.SessionId,
	}, s.logger)
	if err != nil {
		s.logger.Error("create jwt tokens error", zap.Error(err))
		return nil, err
	}

	err = s.redis.SetTokens(*tokenBundle, time.Duration(s.config.RefreshTokenLifeTime)*time.Hour*24)
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
	}

	return &tokenBundle.AccessToken, nil
}

func (s *AuthService) CheckPassword(password string, passwordHash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	if err != nil {
		s.logger.Error("attempt login with incorrect password", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (s *AuthService) GeneratePasswordHash(password string) (*string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), int(s.config.Salt))
	if err != nil {
		s.logger.Error("generate hash for password error", zap.Error(err))
		return nil, err
	}

	resPasswordHash := string(passwordHash)

	return &resPasswordHash, nil
}

func (s *AuthService) Logout(accessToken string) error {
	accessClaims, err := utils.GetTokenClaims(accessToken, s.config.SignedKey)
	if err != nil {
		s.logger.Error("get access token claims error", zap.Error(err))
		return err
	}

	err = s.redis.DeleteTokens(accessClaims.SessionId, accessClaims.UserId)
	if err != nil {
		s.logger.Error("delete token error", zap.Error(err))
		return err
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
