import { useTranslation } from "react-i18next"
import { ExternalLink, Layout, Users, Video } from "lucide-react"
import { mockLessons } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/shared/ui/Card"
import { Label } from "@/shared/ui/Label"

export default function LessonsPage() {
    const { t } = useTranslation()

    return (
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("lessons.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("lessons.subtitle")}
                </Label>
            </div>

            <div className="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
                {mockLessons.map((lesson) => (
                    <Card key={lesson.id}>
                        <CardHeader className="space-y-3">
                            <div className="flex items-start justify-between gap-3">
                                <CardTitle className="text-lg">
                                    {lesson.subjectTitle}
                                </CardTitle>
                                <Badge variant="secondary">
                                    {t(`lessonStatus.${lesson.status}`)}
                                </Badge>
                            </div>
                            <CardDescription>
                                {lesson.groupTitle} · {lesson.tutorName}
                            </CardDescription>
                        </CardHeader>
                        <CardContent className="space-y-3">
                            <div className="text-sm text-muted-foreground">
                                {new Date(lesson.startTime).toLocaleString()}
                            </div>
                            <div className="text-sm">
                                {t("common.duration")}: {t("common.minutes", { count: lesson.duration })}
                            </div>
                            <div className="flex items-center gap-2 text-sm text-muted-foreground">
                                <Users className="size-4" />
                                {t("lessons.participants")}: {lesson.userIds.length}
                            </div>
                            <div className="flex items-center gap-2 text-sm text-muted-foreground">
                                <Video className="size-4" />
                                {t("lessons.media")}: {lesson.mediaItems.length}
                            </div>
                        </CardContent>
                        <CardFooter className="gap-2 flex-wrap">
                            {lesson.meetLink && (
                                <Button
                                    size="sm"
                                    onClick={() => window.open(lesson.meetLink, "_blank", "noreferrer")}
                                >
                                    <ExternalLink className="size-3.5" />
                                    {t("lessons.meet")}
                                </Button>
                            )}
                            {lesson.boardId && (
                                <Button size="sm" variant="outline">
                                    <Layout className="size-3.5" />
                                    {t("lessons.board")}
                                </Button>
                            )}
                        </CardFooter>
                    </Card>
                ))}
            </div>
        </div>
    )
}
