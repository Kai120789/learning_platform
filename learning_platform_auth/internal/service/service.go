package service

import (
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/dto"
	"learning-platform/auth/internal/utils"
)

type AuthService struct {
	config *config.Config
	logger *zap.Logger
	redis  RedisStorage
	api    AuthApi
}

type AuthApi interface {
	GetUserByEmail(email string) (*dto.GetUser, error)
	CreateUser(newUser dto.RegisterRequest) (*int64, error)
}

type RedisStorage interface {
	SetTokens(userId int64, tokenBundle dto.TokenBundle) error
}

func New(
	config *config.Config,
	logger *zap.Logger,
	redis RedisStorage,
	api AuthApi,
) *AuthService {
	return &AuthService{
		config: config,
		logger: logger,
		redis:  redis,
		api:    api,
	}
}

func (s *AuthService) Login(loginData dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.api.GetUserByEmail(loginData.Email)
	if err != nil {
		s.logger.Error("get user by email error", zap.Error(err))
		return nil, err
	}

	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      user.UserId,
		UserEmail:   user.Email,
		SignedKey:   s.config.SignedKey,
		Issuer:      s.config.Issuer,
		AccessTime:  s.config.AccessTokenLifeTime,
		RefreshTime: s.config.RefreshTokenLifeTime,
	}, s.logger)
	if err != nil {
		s.logger.Error("create jwt tokens error", zap.Error(err))
		return nil, err
	}

	err = s.redis.SetTokens(user.UserId, *tokenBundle)
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: tokenBundle.AccessToken,
		UserId:      user.UserId,
	}, nil
}

func (s *AuthService) Register(registerData dto.RegisterRequest) (*dto.RegisterResponse, error) {
	userId, err := s.api.CreateUser(registerData)
	if err != nil {
		s.logger.Error("create user error", zap.Error(err))
		return nil, err
	}

	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      *userId,
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

	err = s.redis.SetTokens(*userId, *tokenBundle)
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
	}

	return &dto.RegisterResponse{
		UserId:      *userId,
		AccessToken: tokenBundle.AccessToken,
	}, nil
}

func (s *AuthService) RefreshTokens() {}

func (s *AuthService) Logout() {}

func (s *AuthService) LogoutAll() {}

func (s *AuthService) ChangePassword() {}

func (s *AuthService) ForceChangePassword() {}

func (s *AuthService) ChangeEmail() {}

func (s *AuthService) ForceChangeEmail() {}
