package service

import (
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/dto"
	"learning-platform/auth/internal/redis"
	"learning-platform/auth/internal/transport/grpc"
	"learning-platform/auth/internal/utils"
)

type AuthService struct {
	config *config.Config
	logger *zap.Logger
	redis  *redis.RedisStorage
	api    *grpc.UserApi
}

type AuthStorage interface {
}

func New(
	config *config.Config,
	logger *zap.Logger,
	redis *redis.RedisStorage,
	api *grpc.UserApi,
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

	_ = tokenBundle

	err = s.redis.SetTokens("testAccessToken", "testRefreshToken", "testSid")
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
	}

	return &dto.LoginResponse{}, nil
}

func (s *AuthService) Register() {}

func (s *AuthService) RefreshTokens() {}

func (s *AuthService) Logout() {}

func (s *AuthService) LogoutAll() {}

func (s *AuthService) ChangePassword() {}

func (s *AuthService) ForceChangePassword() {}

func (s *AuthService) ChangeEmail() {}

func (s *AuthService) ForceChangeEmail() {}
