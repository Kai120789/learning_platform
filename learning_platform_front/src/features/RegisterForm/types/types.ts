import type { UserGenderEnum, UserLanguageEnum } from "@/shared/enums/user"

export type RegisterRequestDTO = {
    email: string
    name: string
    surname: string
    patronymic?: string
    role: RegisterRoleEnum
    gender: UserGenderEnum
    language: UserLanguageEnum
    birth_date?: Date
    password: string
}

export type RegisterResponseDTO = {
    session_id: string
}

export enum RegisterRoleEnum {
    STUDENT = "STUDENT",
    TUTOR = "TUTOR"
}
