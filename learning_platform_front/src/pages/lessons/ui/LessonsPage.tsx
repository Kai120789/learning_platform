import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    getLessons,
    getLessonsByStudentId,
    getLessonsByTutorId,
    getLessonsIsLoading,
    updateLessonStatus,
    type LessonData,
} from "@/entities/lesson"
import { getUserFullData, useCanEdit } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { CreateLessonModal } from "@/widgets/createLessonModal"
import { EditLessonModal } from "@/widgets/editLessonModal"
import { Plus } from "lucide-react"
import { LessonCard } from "./LessonCard"

export default function LessonsPage() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const canEdit = useCanEdit()
    const userData = useAppSelector(getUserFullData)
    const lessons = useAppSelector(getLessons)
    const isLoading = useAppSelector(getLessonsIsLoading)

    const [isCreateOpen, setIsCreateOpen] = useState(false)
    const [editingLesson, setEditingLesson] = useState<LessonData | null>(null)

    useEffect(() => {
        if (!userData?.user.userID) return
        if (canEdit) {
            dispatch(getLessonsByTutorId(userData.user.userID))
        } else {
            dispatch(getLessonsByStudentId(userData.user.userID))
        }
    }, [canEdit, dispatch, userData?.user.userID])

    const sortedLessons = useMemo(
        () => [...(lessons ?? [])].sort((a, b) => a.startTime.localeCompare(b.startTime)),
        [lessons],
    )

    const onStart = async (lesson: LessonData) => {
        const response = await dispatch(updateLessonStatus({
            lessonID: lesson.id,
            status: "IN_PROCESS",
        }))
        dispatch(notificationActions.addNotification({
            message: response.meta.requestStatus === "fulfilled"
                ? t("lessons.startSuccess")
                : t("lessons.startError"),
            type: response.meta.requestStatus === "fulfilled" ? "success" : "error",
        }))
    }

    const onComplete = async (lesson: LessonData) => {
        const response = await dispatch(updateLessonStatus({
            lessonID: lesson.id,
            status: "COMPLETED",
        }))
        dispatch(notificationActions.addNotification({
            message: response.meta.requestStatus === "fulfilled"
                ? t("lessons.completeSuccess")
                : t("lessons.completeError"),
            type: response.meta.requestStatus === "fulfilled" ? "success" : "error",
        }))
    }

    const onCancel = async (lesson: LessonData) => {
        const response = await dispatch(updateLessonStatus({
            lessonID: lesson.id,
            status: "CANCELLED",
        }))
        dispatch(notificationActions.addNotification({
            message: response.meta.requestStatus === "fulfilled"
                ? t("lessons.cancelSuccess")
                : t("lessons.cancelError"),
            type: response.meta.requestStatus === "fulfilled" ? "success" : "error",
        }))
    }

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <div className="flex items-center justify-between gap-4">
                    <Label className="text-xl lg:text-2xl">
                        {t("lessons.title")}
                    </Label>
                    {canEdit && (
                        <Button className="rounded-full" onClick={() => setIsCreateOpen(true)}>
                            <Plus className="size-3" />
                            {t("lessons.create")}
                        </Button>
                    )}
                </div>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("lessons.subtitle")}
                </Label>
            </div>

            {lessons === null && isLoading && (
                <p className="text-sm text-muted-foreground">{t("common.loading")}</p>
            )}

            {lessons !== null && sortedLessons.length === 0 && (
                <p className="text-sm text-muted-foreground">{t("lessons.empty")}</p>
            )}

            {sortedLessons.length > 0 && (
                <div className="grid auto-rows-fr gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                    {sortedLessons.map((lesson) => (
                        <LessonCard
                            key={lesson.id}
                            lesson={lesson}
                            canEdit={canEdit}
                            onEdit={setEditingLesson}
                            onStart={onStart}
                            onComplete={onComplete}
                            onCancel={onCancel}
                        />
                    ))}
                </div>
            )}

            {canEdit && (
                <>
                    <CreateLessonModal
                        isOpen={isCreateOpen}
                        setIsOpen={setIsCreateOpen}
                    />
                    <EditLessonModal
                        isOpen={Boolean(editingLesson)}
                        setIsOpen={(open) => {
                            if (!open) setEditingLesson(null)
                        }}
                        lesson={editingLesson}
                    />
                </>
            )}
        </div>
    )
}
