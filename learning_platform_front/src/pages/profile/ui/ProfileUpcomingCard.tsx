import { useTranslation } from "react-i18next"
import { Clock } from "lucide-react"
import { mockLessons } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"

export function ProfileUpcomingCard() {
    const { t, i18n } = useTranslation()

    const upcoming = mockLessons
        .filter((lesson) => lesson.status === "SCHEDULED")
        .sort((a, b) => a.startTime.localeCompare(b.startTime))
        .slice(0, 4)

    return (
        <Card>
            <CardHeader>
                <CardTitle>{t("profile.upcoming")}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
                {upcoming.length === 0 ? (
                    <p className="text-sm text-muted-foreground">
                        {t("profile.noUpcoming")}
                    </p>
                ) : (
                    upcoming.map((lesson) => {
                        const start = new Date(lesson.startTime)

                        return (
                            <div
                                key={lesson.id}
                                className="rounded-lg border p-3 space-y-1.5 transition-colors hover:bg-muted/50"
                            >
                                <div className="flex items-center justify-between gap-2">
                                    <span className="truncate text-sm font-medium">
                                        {lesson.subjectTitle}
                                    </span>
                                    <Badge variant="secondary">
                                        {t(`lessonStatus.${lesson.status}`)}
                                    </Badge>
                                </div>
                                <p className="truncate text-sm text-muted-foreground">
                                    {lesson.groupTitle}
                                </p>
                                <p className="flex items-center gap-1 text-xs text-muted-foreground">
                                    <Clock className="size-3" />
                                    {start.toLocaleDateString(i18n.language, {
                                        day: "numeric",
                                        month: "short",
                                    })}
                                    {", "}
                                    {start.toLocaleTimeString(i18n.language, {
                                        hour: "2-digit",
                                        minute: "2-digit",
                                    })}
                                    {" · "}
                                    {t("common.minutes", { count: lesson.duration })}
                                </p>
                            </div>
                        )
                    })
                )}
            </CardContent>
        </Card>
    )
}
