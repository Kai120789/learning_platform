import { GraduationCap, Settings, Shield, User } from "lucide-react"
import type { LucideIcon } from "lucide-react"

export type Section = {
    ID: SectionTabs,
    Title: string
    Icon: LucideIcon
}

export enum SectionTabs {
    PROFILE = "PROFILE",
    ACCOUNT = "ACCOUNT",
    SETTINGS = "SETTINGS",
    TEACHING = "TEACHING",
}

export const Sections: Section[] = [
    {
        ID: SectionTabs.PROFILE,
        Title: "Личные данные",
        Icon: User,
    },
    {
        ID: SectionTabs.ACCOUNT,
        Title: "Аккаунт",
        Icon: Shield,
    },
    {
        ID: SectionTabs.TEACHING,
        Title: "Преподавание",
        Icon: GraduationCap,
    },
    {
        ID: SectionTabs.SETTINGS,
        Title: "Настройки",
        Icon: Settings,
    },
]
