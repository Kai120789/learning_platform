export type LoginRequestDTO = {
    email: string
    password: string
}

export type LoginResponseDTO = {
    user_id: number
    session_id: string
}