import { useTranslation } from "react-i18next"
import { Check, ExternalLink, Layout, Pen, Play, Users, Video, X } from "lucide-react"
import type { LessonData } from "@/entities/lesson"
import { lessonStatusClass } from "@/shared/lib/statusStyles"
import { useCountdown } from "@/shared/lib/useCountdown"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Separator } from "@/shared/ui/Separator"

type StudentLessonModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    lesson: LessonData | null
    canEdit?: boolean
    onEdit?: (lesson: LessonData) => void
    onStart?: (lesson: LessonData) => void
    onComplete?: (lesson: LessonData) => void
    onCancel?: (lesson: LessonData) => void
}

function LessonTiming({
    lesson,
    canEdit,
}: {
    lesson: LessonData
    canEdit: boolean
}) {
    const { t } = useTranslation()
    const { isPast, label } = useCountdown(
        lesson.status === "SCHEDULED" ? lesson.startTime : null,
    )

    if (lesson.status === "CANCELLED" || lesson.status === "COMPLETED") {
        return null
    }

    if (lesson.status === "IN_PROCESS") {
        return (
            <p className="text-amber-700/80 dark:text-amber-200/45">
                {t("lessons.inProcessHint")}
            </p>
        )
    }

    if (!isPast) {
        return (
            <p>
                {t("lessons.startsIn", { time: label })}
            </p>
        )
    }

    if (canEdit) {
        return (
            <p className="text-amber-700/80 dark:text-amber-200/45">
                {t("lessons.waitingStartTutor")}
            </p>
        )
    }

    return (
        <p className="text-amber-700/80 dark:text-amber-200/45">
            {t("lessons.waitingStartStudent")}
        </p>
    )
}

export function StudentLessonModal({
    isOpen,
    setIsOpen,
    lesson,
    canEdit = false,
    onEdit,
    onStart,
    onComplete,
    onCancel,
}: StudentLessonModalProps) {
    const { t } = useTranslation()
    const showJoinActions = Boolean(
        lesson
        && lesson.status === "IN_PROCESS"
        && (lesson.meetLink || lesson.boardId),
    )
    const showTutorScheduleActions = Boolean(
        canEdit && lesson && lesson.status === "SCHEDULED",
    )
    const showTutorComplete = Boolean(
        canEdit && lesson && lesson.status === "IN_PROCESS",
    )

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="flex flex-wrap items-center gap-2.5 pr-8 text-left text-base">
                        <span className="min-w-0">
                            {lesson
                                ? t("lessons.itemTitle", { id: lesson.id })
                                : t("lessons.title")}
                        </span>
                        {lesson && (
                            <Badge
                                variant="outline"
                                className={`h-7 px-2.5 text-sm ${lessonStatusClass(lesson.status)}`}
                            >
                                {t(`lessonStatus.${lesson.status}`)}
                            </Badge>
                        )}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />

                {lesson ? (
                    <div className="space-y-4">
                        <div className="space-y-1.5 text-sm text-muted-foreground">
                            <p>
                                {new Date(lesson.startTime).toLocaleString()}
                            </p>
                            <p>
                                {t("common.duration")}: {t("common.minutes", { count: lesson.duration })}
                            </p>
                            <p className="flex items-center gap-1.5">
                                <Users className="size-3.5" />
                                {t("lessons.participants")}: {lesson.userIds.length}
                            </p>
                            <p className="flex items-center gap-1.5">
                                <Video className="size-3.5" />
                                {t("lessons.media")}: {lesson.mediaItems.length}
                            </p>
                            <LessonTiming lesson={lesson} canEdit={canEdit} />
                        </div>

                        <div className="flex flex-col gap-2">
                            {showTutorScheduleActions && (
                                <>
                                    <Button
                                        size="sm"
                                        variant="outline"
                                        className="w-full"
                                        onClick={() => onEdit?.(lesson)}
                                    >
                                        <Pen className="size-3.5" />
                                        {t("common.edit")}
                                    </Button>
                                    <div className="grid grid-cols-2 gap-2">
                                        <Button size="sm" onClick={() => onStart?.(lesson)}>
                                            <Play className="size-3.5" />
                                            {t("lessons.start")}
                                        </Button>
                                        <Button
                                            size="sm"
                                            variant="outline"
                                            onClick={() => onCancel?.(lesson)}
                                        >
                                            <X className="size-3.5" />
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
                                            size="sm"
                                            onClick={() => window.open(lesson.meetLink!, "_blank", "noreferrer")}
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
                                </div>
                            )}

                            {showTutorComplete && (
                                <Button
                                    size="sm"
                                    className="w-full"
                                    onClick={() => onComplete?.(lesson)}
                                >
                                    <Check className="size-3.5" />
                                    {t("lessons.complete")}
                                </Button>
                            )}
                        </div>
                    </div>
                ) : (
                    <p className="text-sm text-muted-foreground">{t("common.loading")}</p>
                )}
            </DialogContent>
        </Dialog>
    )
}
