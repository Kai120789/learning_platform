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
import { lessonStatusClass } from "@/shared/lib/statusStyles"

export default function LessonsPage() {
    const { t } = useTranslation()

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("lessons.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("lessons.subtitle")}
                </Label>
            </div>

            <div className="grid auto-rows-fr gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {mockLessons.map((lesson) => (
                    <Card key={lesson.id} size="sm" className="h-full">
                        <CardHeader className="space-y-2">
                            <div className="flex items-start justify-between gap-2">
                                <CardTitle className="text-sm font-medium line-clamp-2 min-h-10">
                                    {lesson.subjectTitle}
                                </CardTitle>
                                <Badge
                                    variant="outline"
                                    className={`shrink-0 text-[10px] ${lessonStatusClass(lesson.status)}`}
                                >
                                    {t(`lessonStatus.${lesson.status}`)}
                                </Badge>
                            </div>
                            <CardDescription className="text-xs line-clamp-1">
                                {lesson.groupTitle} · {lesson.tutorName}
                            </CardDescription>
                        </CardHeader>
                        <CardContent className="flex-1 space-y-1.5 text-xs text-muted-foreground">
                            <div>{new Date(lesson.startTime).toLocaleString()}</div>
                            <div>
                                {t("common.duration")}: {t("common.minutes", { count: lesson.duration })}
                            </div>
                            <div className="flex items-center gap-1.5">
                                <Users className="size-3.5" />
                                {t("lessons.participants")}: {lesson.userIds.length}
                            </div>
                            <div className="flex items-center gap-1.5">
                                <Video className="size-3.5" />
                                {t("lessons.media")}: {lesson.mediaItems.length}
                            </div>
                        </CardContent>
                        <CardFooter className="gap-2 flex-wrap min-h-10">
                            {lesson.meetLink && (
                                <Button
                                    size="xs"
                                    onClick={() => window.open(lesson.meetLink, "_blank", "noreferrer")}
                                >
                                    <ExternalLink className="size-3" />
                                    {t("lessons.meet")}
                                </Button>
                            )}
                            {lesson.boardId && (
                                <Button size="xs" variant="outline">
                                    <Layout className="size-3" />
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
