export type RegisterRequestDTO = {
    name: string
    surname: string
    last_name?: string
    role: RegisterRoleEnum
    email: string
    password: string
}

export enum RegisterRoleEnum {
    USER = "USER",
    TUTOR = "TUTOR"
}