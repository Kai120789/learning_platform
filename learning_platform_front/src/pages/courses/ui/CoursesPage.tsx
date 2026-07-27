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
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("courses.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("courses.subtitle")}
                </Label>
            </div>

            <div className="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
                {mockCourses.map((course) => (
                    <Card key={course.id}>
                        <CardHeader className="space-y-3">
                            <div className="flex items-start justify-between gap-3">
                                <CardTitle className="text-lg">{course.title}</CardTitle>
                                <Badge variant="outline">{course.subjectTitle}</Badge>
                            </div>
                            <CardDescription>{course.description}</CardDescription>
                        </CardHeader>
                        <CardContent className="space-y-3">
                            <div className="text-sm text-muted-foreground">
                                {course.tutorName}
                            </div>
                            <div className="text-sm">
                                {t("courses.modules", { count: course.modulesCount })}
                            </div>
                            <div className="text-sm">
                                {t("courses.enrolled", { count: course.enrolledCount })}
                            </div>
                            <div className="space-y-2">
                                <div className="flex justify-between text-sm">
                                    <span>{t("courses.progress")}</span>
                                    <span>{course.progress}%</span>
                                </div>
                                <div className="h-2 rounded-full bg-muted overflow-hidden">
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
