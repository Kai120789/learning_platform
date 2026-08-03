import type { GroupUser, ShortUserInfo } from "@/entities/group"

export function displayName(user: ShortUserInfo | GroupUser) {
    return `${user.name} ${user.surname}`.trim()
}

export function toShortUser(user: GroupUser | ShortUserInfo): ShortUserInfo {
    return {
        id: user.id,
        name: user.name,
        surname: user.surname,
        patronymic: user.patronymic,
        tg_username: "tgUsername" in user ? user.tgUsername : (user as ShortUserInfo).tg_username,
    }
}

export function placeholderUser(id: number): ShortUserInfo {
    return {
        id,
        name: `#${id}`,
        surname: "",
    }
}
