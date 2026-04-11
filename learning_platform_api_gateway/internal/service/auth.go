package service

import (
	"fmt"
	"learning-platform/api-gateway/internal/dto"
	"learning-platform/api-gateway/internal/redis"
)

type AuthService struct {
	client      AuthClient
	userService UserAuthService
	redis       RedisStorage
}

type AuthClient interface {
	Login(req dto.LoginRequest, userId int64) (*dto.LoginResponse, error)
	Register(req dto.RegisterRequest, userId int64) (*dto.RegisterResponse, error)
	RefreshTokens(accessToken string) (*string, error)
	CheckPassword(password string, passwordHash string) (bool, error)
	GeneratePasswordHash(password string) (*string, error)
	Logout(accessToken string) error
	LogoutAll(userId int64) error
}

type UserAuthService interface {
	GetUserByEmail(email string) (*dto.GetUser, error)
	CreateUser(newUser dto.RegisterRequest) (*int64, error)
	GetUserById(id int64) (*dto.GetUser, error)
}

type RedisStorage interface {
	GetTokens(sessionId string) (*dto.RedisTokens, error)
}

func NewAuthService(
	client AuthClient,
	userService UserAuthService,
	redis *redis.RedisStorage,
) *AuthService {
	return &AuthService{
		client:      client,
		userService: userService,
		redis:       redis,
	}
}

func (a *AuthService) Login(loginReq dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userService.GetUserByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user does not exists, %w", err)
	}

	isValid, err := a.client.CheckPassword(loginReq.Password, user.PasswordHash)
	if err != nil || !isValid {
		return nil, err
	}

	res, err := a.client.Login(loginReq, user.UserId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AuthService) Register(registerReq dto.RegisterRequest) (*dto.RegisterResponse, error) {
	passwordHash, err := a.client.GeneratePasswordHash(registerReq.Password)
	if err != nil {
		return nil, err
	}

	registerReq.Password = *passwordHash

	userId, err := a.userService.CreateUser(registerReq)
	if err != nil {
		return nil, err
	}

	res, err := a.client.Register(registerReq, *userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AuthService) RefreshTokens(refreshToken string) (*string, error) {
	res, err := a.client.RefreshTokens(refreshToken)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AuthService) Logout(sessionId string) error {
	tokens, err := a.redis.GetTokens(sessionId)
	if err != nil {
		return err
	}

	err = a.client.Logout(tokens.AccessToken)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) LogoutAll(userId int64) error {
	user, err := a.userService.GetUserById(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user does not exists, %w", err)
	}

	err = a.client.LogoutAll(userId)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) GetTokens(sessionId string) (*dto.RedisTokens, error) {
	return a.redis.GetTokens(sessionId)
}
