import { useTranslation } from "react-i18next"
import { mockCourses } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/shared/ui/Card"
import { Label } from "@/shared/ui/Label"

export default function CoursesPage() {
    const { t } = useTranslation()

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("courses.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("courses.subtitle")}
                </Label>
            </div>

            <div className="grid auto-rows-fr gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {mockCourses.map((course) => (
                    <Card key={course.id} size="sm" className="h-full">
                        <CardHeader className="space-y-2">
                            <div className="flex items-start justify-between gap-2">
                                <CardTitle className="text-sm font-medium line-clamp-2 min-h-10">
                                    {course.title}
                                </CardTitle>
                                <Badge variant="outline" className="shrink-0 text-[10px]">
                                    {course.subjectTitle}
                                </Badge>
                            </div>
                            <CardDescription className="text-xs line-clamp-2 min-h-8">
                                {course.description}
                            </CardDescription>
                        </CardHeader>
                        <CardContent className="mt-auto space-y-2 text-xs">
                            <div className="truncate text-muted-foreground">{course.tutorName}</div>
                            <div>{t("courses.modules", { count: course.modulesCount })}</div>
                            <div>{t("courses.enrolled", { count: course.enrolledCount })}</div>
                            <div className="space-y-1.5">
                                <div className="flex justify-between">
                                    <span>{t("courses.progress")}</span>
                                    <span>{course.progress}%</span>
                                </div>
                                <div className="h-1.5 rounded-full bg-muted overflow-hidden">
                                    <div
                                        className="h-full bg-primary"
                                        style={{ width: `${course.progress}%` }}
                                    />
                                </div>
                            </div>
                        </CardContent>
                    </Card>
                ))}
            </div>
        </div>
    )
}
