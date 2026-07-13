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
    lastname?: string
    city?: string
    about?: string
}

export type UserSettings = {
    user_id: number
    is_notifications_enabled: boolean
    is_2_fa_enabled: boolean
}

export enum UserRoleEnum {
    ADMIN = "ADMIN",
    STUDENT = "STUDENT",
    TUTOR = "TUTOR"
}

export enum UserStatusEnum {
    ACTIVE = "ACTIVE",
    INACTIVE = "INACTIVE",
    BANNED = "BANNED"
}