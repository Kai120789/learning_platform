import { useTranslation } from "react-i18next"
import { mockPractices } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"

export function HomeworkList() {
    const { t } = useTranslation()
    const homeworks = mockPractices.filter(
        (practice) => practice.type === "homework" || practice.status === "opened" || practice.status === "pending"
    )

    return (
        <Card className="h-full">
            <CardHeader>
                <CardTitle>{t("main.homework")}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
                {homeworks.map((practice) => (
                    <div
                        key={practice.id}
                        className="flex items-start justify-between gap-3 border-b border-border pb-3 last:border-0 last:pb-0"
                    >
                        <div className="min-w-0 space-y-1">
                            <div className="truncate font-medium">{practice.title}</div>
                            <div className="truncate text-sm text-muted-foreground">
                                {practice.groupTitle} · {practice.subjectTitle}
                            </div>
                            <div className="text-xs text-muted-foreground">
                                {t("practices.deadline")}: {new Date(practice.endTime).toLocaleDateString()}
                            </div>
                        </div>
                        <Badge variant="secondary" className="shrink-0">
                            {t(`practiceStatus.${practice.status}`)}
                        </Badge>
                    </div>
                ))}
            </CardContent>
        </Card>
    )
}
