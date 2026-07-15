package service

import (
	"fmt"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/models/enum"
	"learning-platform/users/internal/utils"
)

type UserBaseService struct {
	storage             UserBaseStorage
	userInfoService     UserInfo
	userSettingsService UserSettings
}

type UserBaseStorage interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUserById(userID int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ChangePassword(userID int64, newPasswordHash string) error
	ChangeEmail(userID int64, newEmail string) error
}

type UserInfo interface {
	CreateUserInfo(userID int64, userDto dto.CreateUser) error
	GetUserInfo(userID int64) (*models.UserInfo, error)
}

type UserSettings interface {
	CreateUserSettings(userID int64, language enum.UserLanguage) error
	GetUserSettings(userID int64) (*models.UserSettings, error)
}

func NewUserBaseService(
	storage UserBaseStorage,
	userInfoService UserInfo,
	userSettingsService UserSettings,
) *UserBaseService {
	return &UserBaseService{
		storage:             storage,
		userInfoService:     userInfoService,
		userSettingsService: userSettingsService,
	}
}

func (u *UserBaseService) CreateUser(userDto dto.CreateUser) (*int64, error) {
	userID, err := u.storage.CreateUser(userDto)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	err = u.userInfoService.CreateUserInfo(*userID, userDto)
	if err != nil {
		return nil, fmt.Errorf("create user (info): %w", err)
	}

	err = u.userSettingsService.CreateUserSettings(*userID, userDto.Language)
	if err != nil {
		return nil, fmt.Errorf("create user (settings): %w", err)
	}

	return userID, nil
}

func (u *UserBaseService) GetUserData(userID int64) (*dto.UserData, error) {
	user, err := u.storage.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("get user data: %w", err)
	}

	userInfo, err := u.userInfoService.GetUserInfo(userID)
	if err != nil {
		return nil, fmt.Errorf("get user data (info): %w", err)
	}

	userSettings, err := u.userSettingsService.GetUserSettings(userID)
	if err != nil {
		return nil, fmt.Errorf("get user data (settings): %w", err)
	}

	return formUserDto(user, userInfo, userSettings), nil
}

func (u *UserBaseService) GetUserById(userID int64) (*models.User, error) {
	user, err := u.storage.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}

func (u *UserBaseService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.storage.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return user, nil
}

func (u *UserBaseService) ChangePassword(userID int64, newPasswordHash string) error {
	err := u.storage.ChangePassword(userID, newPasswordHash)
	if err != nil {
		return fmt.Errorf("change password: %w", err)
	}
	return nil
}
func (u *UserBaseService) ChangeEmail(userID int64, newEmail string) error {
	err := u.storage.ChangeEmail(userID, newEmail)
	if err != nil {
		return fmt.Errorf("change email: %w", err)
	}
	return nil
}

func formUserDto(
	user *models.User,
	userInfo *models.UserInfo,
	userSettings *models.UserSettings,
) *dto.UserData {
	return &dto.UserData{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Status: user.Status,
		UserInfo: dto.UserInfoResponse{
			Name:       userInfo.Name,
			Surname:    userInfo.Surname,
			Patronymic: utils.DBStringToOptional(userInfo.Patronymic),
			City:       utils.DBStringToOptional(userInfo.City),
			About:      utils.DBStringToOptional(userInfo.About),
			Avatar:     utils.DBStringToOptional(userInfo.Avatar),
			Gender:     userInfo.Gender,
			BirthDate:  &userInfo.BirthDate.Time,
		},
		UserSettings: dto.UserSettingsResponse{
			Is2FaEnabled:           userSettings.Is2FaEnabled,
			IsNotificationsEnabled: userSettings.IsNotificationsEnabled,
			Language:               userSettings.Language,
			Theme:                  userSettings.Theme,
		},
	}
}
