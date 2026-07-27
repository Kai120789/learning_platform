import { useTranslation } from "react-i18next"
import { CheckCircle2, FileUp, Video } from "lucide-react"
import { mockLessons, mockMaterials, mockPractices } from "@/shared/mocks"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"

type ActivityItem = {
    id: string
    titleKey: string
    subtitle: string
    date: Date
    icon: typeof Video
}

const buildActivity = (): ActivityItem[] => {
    const lessonItems: ActivityItem[] = mockLessons
        .filter((lesson) => lesson.status === "COMPLETED")
        .map((lesson) => ({
            id: `lesson-${lesson.id}`,
            titleKey: "profile.activityLessonCompleted",
            subtitle: `${lesson.subjectTitle} · ${lesson.groupTitle}`,
            date: new Date(lesson.startTime),
            icon: Video,
        }))

    const practiceItems: ActivityItem[] = mockPractices
        .filter((practice) => practice.status === "completed")
        .map((practice) => ({
            id: `practice-${practice.id}`,
            titleKey: "profile.activityPracticeChecked",
            subtitle: `${practice.title} · ${practice.groupTitle}`,
            date: new Date(practice.endTime),
            icon: CheckCircle2,
        }))

    const materialItems: ActivityItem[] = mockMaterials
        .slice(0, 3)
        .map((material) => ({
            id: `material-${material.id}`,
            titleKey: "profile.activityMaterialAdded",
            subtitle: material.title,
            date: new Date(material.updatedAt),
            icon: FileUp,
        }))

    return [...lessonItems, ...practiceItems, ...materialItems]
        .sort((a, b) => b.date.getTime() - a.date.getTime())
        .slice(0, 6)
}

export function ProfileActivityCard() {
    const { t, i18n } = useTranslation()
    const activity = buildActivity()

    return (
        <Card>
            <CardHeader>
                <CardTitle>{t("profile.activity")}</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="space-y-1">
                    {activity.map((item, index) => (
                        <div key={item.id} className="flex gap-3">
                            <div className="flex flex-col items-center">
                                <div className="flex size-8 shrink-0 items-center justify-center rounded-full bg-muted">
                                    <item.icon className="size-4 text-muted-foreground" />
                                </div>
                                {index < activity.length - 1 && (
                                    <div className="my-1 w-px flex-1 bg-border" />
                                )}
                            </div>
                            <div className="min-w-0 flex-1 pb-4">
                                <p className="text-sm font-medium">
                                    {t(item.titleKey)}
                                </p>
                                <p className="truncate text-sm text-muted-foreground">
                                    {item.subtitle}
                                </p>
                                <p className="text-xs text-muted-foreground">
                                    {item.date.toLocaleDateString(i18n.language, {
                                        day: "numeric",
                                        month: "long",
                                    })}
                                </p>
                            </div>
                        </div>
                    ))}
                </div>
            </CardContent>
        </Card>
    )
}
