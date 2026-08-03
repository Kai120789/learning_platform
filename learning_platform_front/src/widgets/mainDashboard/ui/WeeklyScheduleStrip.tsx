import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { format } from "date-fns"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { DayColumn } from "./DayColumn"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    getLessons,
    getLessonsByStudentId,
    getLessonsByTutorId,
    mapLessonsToWeekItems,
    updateLessonStatus,
    type LessonData,
} from "@/entities/lesson"
import { getUserFullData, useCanEdit } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { EditLessonModal } from "@/widgets/editLessonModal"
import { StudentLessonModal } from "@/widgets/studentLessonModal"

const WEEKDAY_KEYS = ["mon", "tue", "wed", "thu", "fri", "sat", "sun"] as const

function getWeekDates(referenceDate = new Date()) {
    const day = referenceDate.getDay()
    const mondayOffset = day === 0 ? -6 : 1 - day
    const monday = new Date(referenceDate)
    monday.setHours(0, 0, 0, 0)
    monday.setDate(referenceDate.getDate() + mondayOffset)

    return WEEKDAY_KEYS.map((key, index) => {
        const date = new Date(monday)
        date.setDate(monday.getDate() + index)
        return { key, weekday: index as 0 | 1 | 2 | 3 | 4 | 5 | 6, date }
    })
}

export function WeeklyScheduleStrip() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const canEdit = useCanEdit()
    const userData = useAppSelector(getUserFullData)
    const lessons = useAppSelector(getLessons)
    const weekDays = getWeekDates()
    const todayKey = new Date().toDateString()

    const [openedLesson, setOpenedLesson] = useState<LessonData | null>(null)
    const [editingLesson, setEditingLesson] = useState<LessonData | null>(null)

    useEffect(() => {
        if (!userData?.user.userID) return
        if (canEdit) {
            dispatch(getLessonsByTutorId(userData.user.userID))
        } else {
            dispatch(getLessonsByStudentId(userData.user.userID))
        }
    }, [canEdit, dispatch, userData?.user.userID])

    const weekItems = useMemo(
        () => mapLessonsToWeekItems(lessons),
        [lessons],
    )

    const lessonsById = useMemo(() => {
        const map = new Map<number, LessonData>()
        ;(lessons ?? []).forEach((lesson) => map.set(lesson.id, lesson))
        return map
    }, [lessons])

    const activeLesson = openedLesson
        ? (lessonsById.get(openedLesson.id) ?? openedLesson)
        : null

    const onSelectLesson = (lessonId: number) => {
        const lesson = lessonsById.get(lessonId)
        if (lesson) setOpenedLesson(lesson)
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
        if (response.meta.requestStatus === "fulfilled") {
            setOpenedLesson(null)
        }
    }

    return (
        <>
            <Card>
                <CardHeader className="pb-2">
                    <CardTitle>{t("main.weekSchedule")}</CardTitle>
                </CardHeader>
                <CardContent>
                    <div className="overflow-x-auto">
                        <div className="grid min-w-[640px] grid-cols-7 overflow-hidden rounded-xl border h-72 lg:h-80">
                            {weekDays.map((day) => {
                                const dateKey = format(day.date, "yyyy-MM-dd")
                                const items = weekItems
                                    .filter((item) =>
                                        item.weekday === day.weekday && item.date === dateKey
                                    )
                                    .sort((a, b) => a.start.localeCompare(b.start))

                                return (
                                    <DayColumn
                                        key={day.key}
                                        dayKey={day.key}
                                        dateLabel={day.date.toLocaleDateString(undefined, {
                                            day: "numeric",
                                            month: "short",
                                        })}
                                        isToday={day.date.toDateString() === todayKey}
                                        items={items}
                                        onSelectLesson={onSelectLesson}
                                    />
                                )
                            })}
                        </div>
                    </div>
                </CardContent>
            </Card>

            <StudentLessonModal
                isOpen={activeLesson !== null}
                setIsOpen={(open) => {
                    if (!open) setOpenedLesson(null)
                }}
                lesson={activeLesson}
                canEdit={canEdit}
                onEdit={(lesson) => {
                    setOpenedLesson(null)
                    setEditingLesson(lesson)
                }}
                onStart={onStart}
                onComplete={onComplete}
                onCancel={onCancel}
            />

            {canEdit && (
                <EditLessonModal
                    isOpen={editingLesson !== null}
                    setIsOpen={(open) => {
                        if (!open) setEditingLesson(null)
                    }}
                    lesson={editingLesson}
                />
            )}
        </>
    )
}
