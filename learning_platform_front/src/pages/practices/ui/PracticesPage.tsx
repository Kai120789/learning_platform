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
import { practiceStatusClass } from "@/shared/lib/statusStyles"

export default function PracticesPage() {
    const { t } = useTranslation()

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("practices.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("practices.subtitle")}
                </Label>
            </div>

            <div className="grid auto-rows-fr gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {mockPractices.map((practice) => {
                    const score = practice.correctAnswersCount != null
                        ? Math.round((practice.correctAnswersCount / practice.exercisesCount) * 100)
                        : null

                    return (
                        <Card key={practice.id} size="sm" className="h-full">
                            <CardHeader className="space-y-2">
                                <div className="flex items-start justify-between gap-2">
                                    <CardTitle className="text-sm font-medium line-clamp-2 min-h-10">
                                        {practice.title}
                                    </CardTitle>
                                    <Badge
                                        variant="outline"
                                        className={`shrink-0 text-[10px] ${practiceStatusClass(practice.status)}`}
                                    >
                                        {t(`practiceStatus.${practice.status}`)}
                                    </Badge>
                                </div>
                                <CardDescription className="text-xs line-clamp-1">
                                    {practice.subjectTitle} · {practice.groupTitle}
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="mt-auto space-y-1 text-xs">
                                <div>
                                    {t("practices.deadline")}: {new Date(practice.endTime).toLocaleDateString()}
                                </div>
                                <div>
                                    {t("practices.exercises", { count: practice.exercisesCount })}
                                </div>
                                <div className="min-h-4">
                                    {score != null ? t("practices.score", { score }) : "\u00A0"}
                                </div>
                            </CardContent>
                        </Card>
                    )
                })}
            </div>
        </div>
    )
}
