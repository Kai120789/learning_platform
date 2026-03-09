package service

import (
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto"
)

type AuthService struct {
	client      AuthClient
	logger      *zap.Logger
	userService UserAuthService
}

type AuthClient interface {
	Login(req dto.LoginRequest, userId int64) (*dto.LoginResponse, error)
	Register(req dto.RegisterRequest, userId int64) (*dto.RegisterResponse, error)
	RefreshTokens(accessToken string) (*string, error)
	CheckPassword(password string, passwordHash string) (bool, error)
	GeneratePasswordHash(password string) (*string, error)
	Logout(accessToken string) error
}

type UserAuthService interface {
	GetUserByEmail(email string) (*dto.GetUser, error)
	CreateUser(newUser dto.RegisterRequest) (*int64, error)
}

func NewAuthService(
	client AuthClient,
	logger *zap.Logger,
	userService UserAuthService,
) *AuthService {
	return &AuthService{
		client:      client,
		logger:      logger,
		userService: userService,
	}
}

func (a *AuthService) Login(loginReq dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userService.GetUserByEmail(loginReq.Email)
	if err != nil {
		a.logger.Error("failed get user by email", zap.Error(err))
		return nil, err
	}

	if user == nil {
		a.logger.Error("user does not exists", zap.Error(err))
		return nil, err
	}

	isValid, err := a.client.CheckPassword(loginReq.Password, user.PasswordHash)
	if err != nil || !isValid {
		a.logger.Error("attempt login with incorrect password")
		return nil, err
	}

	res, err := a.client.Login(loginReq, user.UserId)
	if err != nil {
		a.logger.Error("failed login", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (a *AuthService) Register(registerReq dto.RegisterRequest) (*dto.RegisterResponse, error) {
	passwordHash, err := a.client.GeneratePasswordHash(registerReq.Password)
	if err != nil {
		a.logger.Error("failed generate hash from password", zap.Error(err))
		return nil, err
	}

	registerReq.Password = *passwordHash

	userId, err := a.userService.CreateUser(registerReq)
	if err != nil {
		a.logger.Error("failed create user", zap.Error(err))
		return nil, err
	}

	res, err := a.client.Register(registerReq, *userId)
	if err != nil {
		a.logger.Error("failed register", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (a *AuthService) RefreshTokens(accessToken string) (*string, error) {
	res, err := a.client.RefreshTokens(accessToken)
	if err != nil {
		a.logger.Error("failed refresh tokens", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (a *AuthService) Logout(accessToken string) error {
	err := a.client.Logout(accessToken)
	if err != nil {
		a.logger.Error("failed logout", zap.Error(err))
		return err
	}

	return nil
}
