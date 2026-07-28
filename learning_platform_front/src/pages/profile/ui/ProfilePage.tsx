import { useTranslation } from "react-i18next"
import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getUserFullData } from "@/entities/user"
import { mockLessons } from "@/shared/mocks"
import { Label } from "@/shared/ui/Label"
import { ProfileActivityCard } from "./ProfileActivityCard"
import { ProfileHeaderCard } from "./ProfileHeaderCard"
import { ProfileStatsGrid } from "./ProfileStatsGrid"
import { ProfileUpcomingCard } from "./ProfileUpcomingCard"

export default function ProfilePage() {
    const { t } = useTranslation()
    const userData = useAppSelector(getUserFullData)

    const name = userData?.userInfo.name ?? "Иван"
    const surname = userData?.userInfo.surname ?? "Петров"
    const email = userData?.user.email ?? "tutor@example.com"
    const role = userData?.user.role ?? "TUTOR"
    const city = userData?.userInfo.city ?? "Москва"
    const about = userData?.userInfo.about ?? "Преподаватель математики"
    const avatar = userData?.userInfo.avatar

    const subjects = [...new Set(mockLessons.map((lesson) => lesson.subjectTitle))]

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("profile.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("profile.subtitle")}
                </Label>
            </div>

            <ProfileHeaderCard
                name={name}
                surname={surname}
                email={email}
                role={role}
                city={city}
                about={about}
                avatar={avatar}
                memberSince={new Date("2025-09-01")}
                subjects={subjects}
            />

            <ProfileStatsGrid />

            <div className="grid gap-5 lg:grid-cols-[2fr_1fr]">
                <ProfileActivityCard />
                <ProfileUpcomingCard />
            </div>
        </div>
    )
}
