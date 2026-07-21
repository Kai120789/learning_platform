import { FiSettings, FiShield, FiUser } from "react-icons/fi"
import type { IconType } from "react-icons/lib"

export type Section = {
    ID: SectionTabs,
    Title: string
    Icon: IconType
}

export enum SectionTabs {
    PROFILE = "PROFILE",
    ACCOUNT = "ACCOUNT",
    SETTINGS = "SETTINGS",
}

export const Sections: Section[] = [
    {
        ID: SectionTabs.PROFILE,
        Title: "Личные данные",
        Icon: FiUser,
    },
    {
        ID: SectionTabs.ACCOUNT,
        Title: "Аккаунт",
        Icon: FiShield,
    },
    {
        ID: SectionTabs.SETTINGS,
        Title: "Настройки",
        Icon: FiSettings,
    },
]
