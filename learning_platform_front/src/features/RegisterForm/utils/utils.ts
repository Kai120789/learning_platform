import { UserGenderEnum, UserLanguageEnum } from "@/shared/enums/user"

export const languageToTranslationKey = (language: UserLanguageEnum): string => {
    return `languages.${language}`
}

export const genderToTranslationKey = (gender: UserGenderEnum): string => {
    return `genders.${gender}`
}
