import { useTranslation } from "react-i18next"
import { ExternalLink, Layout, Pen, Play, Check, Users, Video, X } from "lucide-react"
import type { LessonData } from "@/entities/lesson"
import { lessonStatusClass } from "@/shared/lib/statusStyles"
import { cn } from "@/shared/lib/utils"
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
import { LessonTiming } from "./LessonTiming"

type LessonCardProps = {
    lesson: LessonData
    canEdit: boolean
    onEdit: (lesson: LessonData) => void
    onStart: (lesson: LessonData) => void
    onComplete: (lesson: LessonData) => void
    onCancel: (lesson: LessonData) => void
}

export function LessonCard({
    lesson,
    canEdit,
    onEdit,
    onStart,
    onComplete,
    onCancel,
}: LessonCardProps) {
    const { t } = useTranslation()
    const isCancelled = lesson.status === "CANCELLED"
    const showJoinActions = Boolean(lesson.meetLink || lesson.boardId)
        && lesson.status === "IN_PROCESS"
    const showTutorScheduleActions = canEdit && lesson.status === "SCHEDULED"
    const showTutorComplete = canEdit && lesson.status === "IN_PROCESS"
    const hasFooter = showJoinActions || showTutorScheduleActions || showTutorComplete

    return (
        <Card
            size="sm"
            aria-disabled={isCancelled}
            className={cn(
                "h-full",
                isCancelled && "cursor-not-allowed opacity-55 text-muted-foreground ring-border/5 grayscale",
            )}
        >
            <CardHeader className="space-y-2">
                <div className="flex items-start justify-between gap-2">
                    <CardTitle
                        className={cn(
                            "text-sm font-medium line-clamp-2 min-h-10",
                            isCancelled && "text-muted-foreground",
                        )}
                    >
                        {t("lessons.itemTitle", { id: lesson.id })}
                    </CardTitle>
                    <Badge
                        variant="outline"
                        className={`shrink-0 text-[10px] ${lessonStatusClass(lesson.status)}`}
                    >
                        {t(`lessonStatus.${lesson.status}`)}
                    </Badge>
                </div>
                <CardDescription className="text-xs line-clamp-1">
                    {new Date(lesson.startTime).toLocaleString()}
                </CardDescription>
            </CardHeader>
            <CardContent
                className={cn(
                    "flex-1 space-y-1.5 text-xs text-muted-foreground",
                    isCancelled && "opacity-90",
                )}
            >                <div>
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
                <LessonTiming lesson={lesson} canEdit={canEdit} />
            </CardContent>
            {hasFooter && (
                <CardFooter className="flex-col items-stretch gap-2 min-h-10">
                    {showTutorScheduleActions && (
                        <>
                            <Button
                                size="xs"
                                variant="outline"
                                className="w-full"
                                onClick={() => onEdit(lesson)}
                            >
                                <Pen className="size-3" />
                                {t("common.edit")}
                            </Button>
                            <div className="grid grid-cols-2 gap-2">
                                <Button size="xs" onClick={() => onStart(lesson)}>
                                    <Play className="size-3" />
                                    {t("lessons.start")}
                                </Button>
                                <Button size="xs" variant="outline" onClick={() => onCancel(lesson)}>
                                    <X className="size-3" />
                                    {t("lessons.cancel")}
                                </Button>
                            </div>
                        </>
                    )}

                    {showJoinActions && (
                        <div className={
                            lesson.meetLink && lesson.boardId
                                ? "grid grid-cols-2 gap-2"
                                : "grid grid-cols-1 gap-2"
                        }>
                            {lesson.meetLink && (
                                <Button
                                    size="xs"
                                    onClick={() => window.open(lesson.meetLink!, "_blank", "noreferrer")}
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
                        </div>
                    )}

                    {showTutorComplete && (
                        <Button size="xs" className="w-full" onClick={() => onComplete(lesson)}>
                            <Check className="size-3" />
                            {t("lessons.complete")}
                        </Button>
                    )}
                </CardFooter>
            )}
        </Card>
    )
}
