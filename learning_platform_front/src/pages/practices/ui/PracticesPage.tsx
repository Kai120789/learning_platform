import { useTranslation } from "react-i18next"
import { mockPractices } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/shared/ui/Card"
import { Label } from "@/shared/ui/Label"

export default function PracticesPage() {
    const { t } = useTranslation()

    return (
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("practices.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("practices.subtitle")}
                </Label>
            </div>

            <div className="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
                {mockPractices.map((practice) => {
                    const score = practice.correctAnswersCount != null
                        ? Math.round((practice.correctAnswersCount / practice.exercisesCount) * 100)
                        : null

                    return (
                        <Card key={practice.id}>
                            <CardHeader className="space-y-3">
                                <div className="flex items-start justify-between gap-3">
                                    <CardTitle className="text-lg">{practice.title}</CardTitle>
                                    <Badge variant="secondary">
                                        {t(`practiceStatus.${practice.status}`)}
                                    </Badge>
                                </div>
                                <CardDescription>
                                    {practice.subjectTitle} · {practice.groupTitle}
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="space-y-2 text-sm">
                                <div>
                                    {t("practices.deadline")}: {new Date(practice.endTime).toLocaleDateString()}
                                </div>
                                <div>
                                    {t("practices.exercises", { count: practice.exercisesCount })}
                                </div>
                                {score != null && (
                                    <div>{t("practices.score", { score })}</div>
                                )}
                            </CardContent>
                        </Card>
                    )
                })}
            </div>
        </div>
    )
}
