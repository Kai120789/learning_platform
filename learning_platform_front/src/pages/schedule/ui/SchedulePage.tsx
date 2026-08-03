import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { CalendarDays, LayoutList, Plus } from "lucide-react"
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
import { cn } from "@/shared/lib/utils"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { CreateLessonModal } from "@/widgets/createLessonModal"
import { EditLessonModal } from "@/widgets/editLessonModal"
import { LessonCard } from "./LessonCard"
import { LessonsCalendarView } from "./LessonsCalendarView"

type ScheduleLayout = "calendar" | "list"

export default function SchedulePage() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const canEdit = useCanEdit()
    const userData = useAppSelector(getUserFullData)
    const lessons = useAppSelector(getLessons)
    const isLoading = useAppSelector(getLessonsIsLoading)

    const [layout, setLayout] = useState<ScheduleLayout>("calendar")
    const [isCreateOpen, setIsCreateOpen] = useState(false)
    const [createAt, setCreateAt] = useState<Date | null>(null)
    const [createKey, setCreateKey] = useState(0)
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

    const openCreate = (date?: Date) => {
        setCreateAt(date ?? null)
        setCreateKey((key) => key + 1)
        setIsCreateOpen(true)
    }

    const onCreateOpenChange = (open: boolean) => {
        setIsCreateOpen(open)
        if (!open) setCreateAt(null)
    }

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

    const layouts: {
        id: ScheduleLayout
        label: string
        icon: typeof CalendarDays
    }[] = [
        { id: "calendar", label: t("schedule.layoutCalendar"), icon: CalendarDays },
        { id: "list", label: t("schedule.layoutList"), icon: LayoutList },
    ]

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
                <div className="space-y-1">
                    <Label className="text-xl lg:text-2xl">
                        {t("schedule.title")}
                    </Label>
                    <Label className="text-sm lg:text-base font-normal text-primary/50">
                        {canEdit ? t("schedule.tutorSubtitle") : t("schedule.studentSubtitle")}
                    </Label>
                </div>

                <div className="flex flex-wrap items-center gap-2 sm:justify-end">
                    <div
                        className="inline-flex rounded-lg border border-border bg-background p-0.5"
                        role="tablist"
                        aria-label={t("schedule.layoutMode")}
                    >
                        {layouts.map(({ id, label, icon: Icon }) => {
                            const active = layout === id
                            return (
                                <button
                                    key={id}
                                    type="button"
                                    role="tab"
                                    aria-selected={active}
                                    onClick={() => setLayout(id)}
                                    className={cn(
                                        "inline-flex cursor-pointer items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                                        active
                                            ? "bg-primary text-primary-foreground shadow-sm"
                                            : "text-muted-foreground hover:bg-muted hover:text-foreground",
                                    )}
                                >
                                    <Icon className="size-3.5 shrink-0" />
                                    {label}
                                </button>
                            )
                        })}
                    </div>

                    {canEdit && (
                        <Button
                            size="sm"
                            className="rounded-full shrink-0 md:h-8 md:gap-1.5 md:px-2.5 md:text-sm"
                            onClick={() => openCreate()}
                        >
                            <Plus className="size-3 md:size-3.5" />
                            {t("lessons.create")}
                        </Button>
                    )}
                </div>
            </div>

            {layout === "list" ? (
                <>
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
                </>
            ) : (
                <LessonsCalendarView
                    lessons={lessons}
                    isLoading={isLoading}
                    canEdit={canEdit}
                    onCreateAt={(date) => openCreate(date)}
                    onEdit={setEditingLesson}
                    onStart={onStart}
                    onComplete={onComplete}
                    onCancel={onCancel}
                />
            )}

            {canEdit && (
                <>
                    <CreateLessonModal
                        key={createKey}
                        isOpen={isCreateOpen}
                        setIsOpen={onCreateOpenChange}
                        defaultStartTime={createAt}
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
