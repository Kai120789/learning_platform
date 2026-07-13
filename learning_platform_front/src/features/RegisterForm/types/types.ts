export type RegisterRequestDTO = {
    name: string
    surname: string
    lastname?: string
    role: RegisterRoleEnum
    email: string
    password: string
}

export type RegisterResponseDTO = {
    user_id: number
    session_id: string
}

export enum RegisterRoleEnum {
    STUDENT = "STUDENT",
    TUTOR = "TUTOR"
}