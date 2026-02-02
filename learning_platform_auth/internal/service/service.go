package service

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
	DeleteTokens(accessToken string, userId int64) error
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

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(loginData.Password),
	)
	if err != nil {
		s.logger.Error("attempt login with incorrect password", zap.Error(err))
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), int(s.config.Salt))
	if err != nil {
		return nil, err
	}
	registerData.Password = string(passwordHash)

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

func (s *AuthService) RefreshTokens(accessToken string) (*string, error) {
	accessClaims, err := utils.GetAccessTokenClaims(accessToken, s.config.SignedKey)
	if err != nil {
		s.logger.Error("get access token claims error", zap.Error(err))
		return nil, err
	}

	err = s.redis.DeleteTokens(accessToken, accessClaims.UserId)
	if err != nil {
		s.logger.Error("delete token error", zap.Error(err))
		return nil, err
	}

	tokenBundle, err := utils.CreateJWT(dto.CreateJWT{
		UserId:      accessClaims.UserId,
		UserEmail:   accessClaims.UserEmail,
		SignedKey:   s.config.SignedKey,
		Issuer:      s.config.Issuer,
		AccessTime:  s.config.AccessTokenLifeTime,
		RefreshTime: s.config.RefreshTokenLifeTime,
	}, s.logger)
	if err != nil {
		s.logger.Error("create jwt tokens error", zap.Error(err))
		return nil, err
	}

	err = s.redis.SetTokens(accessClaims.UserId, *tokenBundle)
	if err != nil {
		s.logger.Error("set tokens error")
		return nil, err
	}

	return &tokenBundle.AccessToken, nil
}

func (s *AuthService) Logout(accessToken string) error {
	accessClaims, err := utils.GetAccessTokenClaims(accessToken, s.config.SignedKey)
	if err != nil {
		s.logger.Error("get access token claims error", zap.Error(err))
		return err
	}

	err = s.redis.DeleteTokens(accessToken, accessClaims.UserId)
	if err != nil {
		s.logger.Error("delete token error", zap.Error(err))
		return err
	}

	return nil
}

func (s *AuthService) LogoutAll() {}

func (s *AuthService) ChangePassword() {}

func (s *AuthService) ForceChangePassword() {}

func (s *AuthService) ChangeEmail() {}

func (s *AuthService) ForceChangeEmail() {}
