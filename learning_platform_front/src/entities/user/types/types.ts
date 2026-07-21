import type { UserGenderEnum, UserLanguageEnum, UserRoleEnum, UserStatusEnum, UserThemeEnum } from "@/shared/enums/user";

export interface UserSchema {
    data: UserFullData | null
    isLoading: boolean
    error?: string
}

export type UserDataResponse = {
    user_id: number
    email: string
    role: UserRoleEnum
    status: UserStatusEnum
    user_info: UserInfoResponse
    user_settings: UserSettingsResponse
}

export type UserInfoRequest = {
    name: string
    surname: string
    patronymic?: string
    city?: string
    about?: string
    gender: UserGenderEnum
    birth_date?: Date
}

export type UserInfoResponse = {
    name: string
    surname: string
    patronymic?: string
    tg_link?: string
    city?: string
    about?: string
    avatar?: string
    gender: UserGenderEnum
    birth_date?: Date
}

export type UserSettingsRequest = {
    is_notifications_enabled: boolean
    is_2_fa_enabled: boolean
    language: UserLanguageEnum
}

export type UserSettingsResponse = {
    is_notifications_enabled: boolean
    is_2_fa_enabled: boolean
    language: UserLanguageEnum
    theme: UserThemeEnum
}

export type UserFullData = {
    user: {
        userID: number
        email: string
        role: UserRoleEnum
        status: UserStatusEnum
    };
    userInfo: {
        name: string
        surname: string
        patronymic?: string
        tgLink?: string
        city?: string
        about?: string
        avatar?: string
        gender: UserGenderEnum
        birthDate?: Date
    };
    userSettings: {
        isNotificationsEnabled: boolean
        is2FaEnabled: boolean
        language: UserLanguageEnum
        theme: UserThemeEnum
    };
}