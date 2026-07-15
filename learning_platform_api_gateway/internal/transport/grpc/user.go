package grpc

import (
	"context"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/authDto"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/userDto"
	"time"
)

type UserClient struct {
	client userGRPC.UserClient
}

/*
GetAllUsersWithData
ChangePassword
ChangeEmail
UpdateUserInfo
UpdateUserSettings
*/

func NewUserGrpcConnection(userGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		userGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewUserClient(connection *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: userGRPC.NewUserClient(connection),
	}
}

func (u *UserClient) GetUserByEmail(email string) (*userDto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserByEmail(ctx, &userGRPC.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	return &userDto.GetUser{
		UserID:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}

func (u *UserClient) GetUserById(id int64) (*userDto.GetUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserById(ctx, &userGRPC.GetUserByIdRequest{UserId: id})
	if err != nil {
		return nil, err
	}

	return &userDto.GetUser{
		UserID:       res.GetUserId(),
		Email:        res.GetEmail(),
		PasswordHash: res.GetPasswordHash(),
	}, nil
}

func (u *UserClient) GetUserData(id int64) (*userDto.UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.GetUserData(ctx, &userGRPC.GetUserDataRequest{UserId: id})
	if err != nil {
		return nil, err
	}

	return &userDto.UserData{
		UserID: res.GetUserId(),
		Email:  res.GetEmail(),
		Role:   protoToEnumRole(res.GetRole()),
		Status: protoToEnumUserStatus(res.GetStatus()),
		UserInfo: userDto.UserInfo{
			Name:       res.GetUserInfo().GetName(),
			Surname:    res.GetUserInfo().GetSurname(),
			Patronymic: res.GetUserInfo().Patronymic,
			City:       res.GetUserInfo().City,
			About:      res.GetUserInfo().About,
			Avatar:     res.GetUserInfo().Avatar,
			Gender:     protoToEnumGender(res.GetUserInfo().GetGender()),
			BirthDate:  mapDate(res.GetUserInfo().BirthDate),
		},
		UserSettings: userDto.UserSettings{
			Is2FaEnabled:           res.GetUserSettings().GetIs_2FaEnabled(),
			IsNotificationsEnabled: res.GetUserSettings().GetIsNotificationsEnabled(),
			Language:               protoToEnumLanguage(res.GetUserSettings().GetLanguage()),
			Theme:                  protoToEnumTheme(res.GetUserSettings().GetTheme()),
		},
	}, nil
}

func (u *UserClient) CreateUser(newUser authDto.RegisterRequest) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var birthDate *userGRPC.Date
	if newUser.BirthDate != nil {
		birthDate = &userGRPC.Date{
			Day:   int32(newUser.BirthDate.Day()),
			Month: int32(newUser.BirthDate.Month()),
			Year:  int32(newUser.BirthDate.Year()),
		}
	}

	res, err := u.client.CreateUser(ctx, &userGRPC.CreateUserRequest{
		Email:        newUser.Email,
		Name:         newUser.Name,
		Surname:      newUser.Surname,
		Patronymic:   newUser.Patronymic,
		Role:         enumToProtoRole(newUser.Role),
		Gender:       enumToProtoGender(newUser.Gender),
		Language:     enumToProtoLanguage(newUser.Language),
		PasswordHash: newUser.Password,
		BirthDate:    birthDate,
	})
	if err != nil {
		return nil, err
	}

	resUserId := res.GetUserId()

	return &resUserId, nil
}

func protoToEnumRole(role userGRPC.UserRole) enum.UserRole {
	switch role {
	case userGRPC.UserRole_TUTOR:
		return enum.RoleTutor
	case userGRPC.UserRole_STUDENT:
		return enum.RoleStudent
	case userGRPC.UserRole_ADMIN:
		return enum.RoleAdmin
	default:
		return ""
	}
}

func protoToEnumUserStatus(status userGRPC.UserStatus) enum.UserStatus {
	switch status {
	case userGRPC.UserStatus_ACTIVE:
		return enum.StatusActive
	case userGRPC.UserStatus_INACTIVE:
		return enum.StatusInactive
	case userGRPC.UserStatus_BANNED:
		return enum.StatusBanned
	default:
		return ""
	}
}

func protoToEnumGender(gender userGRPC.UserGender) enum.UserGender {
	switch gender {
	case userGRPC.UserGender_MALE:
		return enum.GenderMale
	case userGRPC.UserGender_FEMALE:
		return enum.GenderFemale
	default:
		return enum.GenderUnknown
	}
}

func protoToEnumLanguage(language userGRPC.UserLanguage) enum.UserLanguage {
	switch language {
	case userGRPC.UserLanguage_RU:
		return enum.LanguageRU
	case userGRPC.UserLanguage_EN:
		return enum.LanguageEN
	default:
		return ""
	}
}

func protoToEnumTheme(theme userGRPC.UserTheme) enum.UserTheme {
	switch theme {
	case userGRPC.UserTheme_LIGHT:
		return enum.ThemeLight
	case userGRPC.UserTheme_DARK:
		return enum.ThemeDark
	default:
		return ""
	}
}

func enumToProtoRole(role enum.UserRole) userGRPC.UserRole {
	switch role {
	case enum.RoleTutor:
		return userGRPC.UserRole_TUTOR
	case enum.RoleStudent:
		return userGRPC.UserRole_STUDENT
	case enum.RoleAdmin:
		return userGRPC.UserRole_ADMIN
	default:
		return userGRPC.UserRole_USER_ROLE_UNSPECIFIED
	}
}

func enumToProtoStatus(status enum.UserStatus) userGRPC.UserStatus {
	switch status {
	case enum.StatusActive:
		return userGRPC.UserStatus_ACTIVE
	case enum.StatusInactive:
		return userGRPC.UserStatus_INACTIVE
	case enum.StatusBanned:
		return userGRPC.UserStatus_BANNED
	default:
		return userGRPC.UserStatus_STATUS_UNSPECIFIED
	}
}

func enumToProtoGender(gender enum.UserGender) userGRPC.UserGender {
	switch gender {
	case enum.GenderMale:
		return userGRPC.UserGender_MALE
	case enum.GenderFemale:
		return userGRPC.UserGender_FEMALE
	default:
		return userGRPC.UserGender_ENUM_GENDER_UNSPECIFIED
	}
}

func enumToProtoTheme(theme enum.UserTheme) userGRPC.UserTheme {
	switch theme {
	case enum.ThemeLight:
		return userGRPC.UserTheme_LIGHT
	case enum.ThemeDark:
		return userGRPC.UserTheme_DARK
	default:
		return userGRPC.UserTheme_ENUM_THEME_UNSPECIFIED
	}
}

func enumToProtoLanguage(language enum.UserLanguage) userGRPC.UserLanguage {
	switch language {
	case enum.LanguageRU:
		return userGRPC.UserLanguage_RU
	case enum.LanguageEN:
		return userGRPC.UserLanguage_EN
	default:
		return userGRPC.UserLanguage_ENUM_LANGUAGE_UNSPECIFIED
	}
}

func mapDate(bd *userGRPC.Date) *time.Time {
	if bd == nil {
		return nil
	}

	birthDate := time.Date(
		int(bd.GetYear()),
		time.Month(bd.GetMonth()),
		int(bd.GetDay()),
		0, 0, 0, 0,
		time.UTC,
	)

	return &birthDate
}
