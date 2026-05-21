package service

type Service struct {
	UserBaseService     *UserBaseService
	UserInfoService     *UserInfoService
	UserSettingsService *UserSettingsService
}

type Storage struct {
	UserBaseStorage     UserBaseStorage
	UserInfoStorage     UserInfoStorage
	UserSettingsStorage UserSettingsStorage
}

func New(
	storage *Storage,
) *Service {
	userInfoService := NewUserInfoService(storage.UserInfoStorage)
	userSettingsService := NewUserSettingsService(storage.UserSettingsStorage)
	return &Service{
		UserBaseService:     NewUserBaseService(storage.UserBaseStorage, userInfoService, userSettingsService),
		UserInfoService:     userInfoService,
		UserSettingsService: userSettingsService,
	}
}
