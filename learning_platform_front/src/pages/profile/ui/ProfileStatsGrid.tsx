import { useTranslation } from "react-i18next"
import { FolderOpen, TrendingUp, Users, UsersRound, Video } from "lucide-react"
import { mockHomeGroups, mockLessons, mockMaterials } from "@/shared/mocks"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"

type StatItem = {
    key: string
    value: number
    change: number
    icon: typeof Video
}

export function ProfileStatsGrid() {
    const { t } = useTranslation()

    const studentsCount = new Set(
        mockLessons.flatMap((lesson) => lesson.userIds)
    ).size

    const stats: StatItem[] = [
        {
            key: "statLessons",
            value: mockLessons.length,
            change: 3,
            icon: Video,
        },
        {
            key: "statStudents",
            value: studentsCount,
            change: 2,
            icon: Users,
        },
        {
            key: "statGroups",
            value: mockHomeGroups.length,
            change: 1,
            icon: UsersRound,
        },
        {
            key: "statMaterials",
            value: mockMaterials.length,
            change: 4,
            icon: FolderOpen,
        },
    ]

    return (
        <div className="grid auto-rows-fr gap-4 sm:grid-cols-2 xl:grid-cols-4">
            {stats.map((stat) => (
                <Card key={stat.key} className="h-full">
                    <CardHeader className="flex flex-row items-center justify-between">
                        <CardTitle className="text-sm font-medium text-muted-foreground">
                            {t(`profile.${stat.key}`)}
                        </CardTitle>
                        <stat.icon className="size-4 text-muted-foreground" />
                    </CardHeader>
                    <CardContent className="mt-auto space-y-1">
                        <div className="text-xl font-bold">{stat.value}</div>
                        <div className="flex items-center gap-1 text-xs text-emerald-600/80 dark:text-emerald-300/65">
                            <TrendingUp className="size-3" />
                            {t("profile.statChange", { count: stat.change })}
                        </div>
                    </CardContent>
                </Card>
            ))}
        </div>
    )
}
