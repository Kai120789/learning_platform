export type LoginRequestDTO = {
    email: string
    password: string
}

export type LoginResponseDTO = {
    session_id: string
}