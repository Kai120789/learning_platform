import type { UserGenderEnum, UserLanguageEnum, UserRoleEnum, UserStatusEnum, UserThemeEnum } from "@/shared/enums/user";

export interface UserSchema {
    data: {
        user: {
            userID: number;
            email: string;
            role: UserRoleEnum;
            status: UserStatusEnum;
        };
        userInfo: {
            name: string;
            surname: string;
            lastname?: string;
            city?: string;
            about?: string;
        };
        userSettings: {
            isNotificationsEnabled: boolean;
            is2FaEnabled: boolean;
        };
    } | null;

    isLoading: boolean;
    error?: string;
}

export type UserData = {
    user_id: number
    email: string
    role: UserRoleEnum
    status: UserStatusEnum
    user_info: UserInfo
    user_settings: UserSettings
}

export type UserInfo = {
    user_id: number
    name: string
    surname: string
    patronymic?: string
    city?: string
    about?: string
    avatar?: string
    gender: UserGenderEnum
    birth_date?: string
}

export type UserSettings = {
    user_id: number
    is_notifications_enabled: boolean
    is_2_fa_enabled: boolean
    language: UserLanguageEnum
    theme: UserThemeEnum
}
