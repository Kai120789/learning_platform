import type { ShortUserInfo } from "@/entities/group"
import type { UserRoleEnum } from "@/shared/enums/user"

export type GetUsersWithPaginationRequest = {
    search: string,
    page: number,
    limit: number,
    role: UserRoleEnum
}

export type GetUsersWithPaginationResponse = {
    users: ShortUserInfo[]
    count: number
}