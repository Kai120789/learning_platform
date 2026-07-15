import { UserGenderEnum, UserLanguageEnum } from "@/shared/enums/user"

export const enumToStringLanguage = (language: UserLanguageEnum): string => {
    switch (language) {
        case UserLanguageEnum.RU:
            return "Русский"
        case UserLanguageEnum.EN:
            return "Английский"
        default:
            return "Русский"
    }
}

export const enumToStringGender = (gender: UserGenderEnum): string => {
    switch (gender) {
        case UserGenderEnum.MALE:
            return "Мужской"
        case UserGenderEnum.FEMALE:
            return "Женский"
        case UserGenderEnum.UNKNOWN:
            return "Не выбрано"
        default:
            return "Не выбрано"
    }
}