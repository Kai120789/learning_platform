import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { FaPlus } from "react-icons/fa"
import { enUS, ru } from "date-fns/locale"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    deleteSchedule,
    getSchedules,
    getSchedulesIsLoading,
    getSchedulesByTutorId,
    type ScheduleData,
} from "@/entities/schedule"
import {
    getLessons,
    getLessonsByStudentId,
    getLessonsIsLoading,
} from "@/entities/lesson"
import { getUserFullData, useCanEdit } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { cn } from "@/shared/lib/utils"
import { CreateScheduleModal } from "@/widgets/createScheduleModal"
import { StudentLessonsCalendar, type ScheduleTab } from "./StudentLessonsCalendar"
import { TutorScheduleList } from "./TutorScheduleList"

export default function SchedulePage() {
    const { t, i18n } = useTranslation()
    const dispatch = useAppDispatch()
    const canEdit = useCanEdit()
    const userData = useAppSelector(getUserFullData)
    const schedules = useAppSelector(getSchedules)
    const isLoading = useAppSelector(getSchedulesIsLoading)
    const lessons = useAppSelector(getLessons)
    const lessonsLoading = useAppSelector(getLessonsIsLoading)
    const [isOpen, setIsOpen] = useState(false)
    const [studentTab, setStudentTab] = useState<ScheduleTab>("month")
    const dateLocale = i18n.language === "ru" ? ru : enUS

    const studentTabs: { id: ScheduleTab; label: string }[] = [
        { id: "week", label: t("schedule.week") },
        { id: "month", label: t("schedule.month") },
    ]

    useEffect(() => {
        if (!userData?.user.userID) return
        if (canEdit) {
            dispatch(getSchedulesByTutorId(userData.user.userID))
        } else {
            dispatch(getLessonsByStudentId(userData.user.userID))
        }
    }, [canEdit, dispatch, userData?.user.userID])

    const sortedSchedules = useMemo(
        () => [...(schedules ?? [])].sort((a, b) => a.startTime.localeCompare(b.startTime)),
        [schedules],
    )

    const onDeleteSchedule = async (
        e: React.MouseEvent,
        schedule: ScheduleData,
    ) => {
        e.stopPropagation()
        const response = await dispatch(deleteSchedule(schedule.id))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteError"),
                type: "error",
            }))
        }
    }

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
                <div className="space-y-1">
                    <Label className="text-xl lg:text-2xl">
                        {t("schedule.title")}
                    </Label>
                    <Label className="text-sm lg:text-base font-normal text-primary/50">
                        {canEdit ? t("schedule.subtitle") : t("schedule.studentSubtitle")}
                    </Label>
                </div>

                {canEdit ? (
                    <Button onClick={() => setIsOpen(true)} className="rounded-full shrink-0">
                        <FaPlus className="size-3" />
                        {t("schedule.createSchedule")}
                    </Button>
                ) : (
                    <div className="flex shrink-0 self-end sm:self-start rounded-lg border border-border p-0.5 bg-secondary/60">
                        {studentTabs.map((item) => (
                            <button
                                key={item.id}
                                type="button"
                                onClick={() => setStudentTab(item.id)}
                                className={cn(
                                    "rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors cursor-pointer",
                                    item.id === studentTab
                                        ? "bg-primary text-primary-foreground shadow-sm"
                                        : "text-secondary-foreground hover:bg-muted",
                                )}
                            >
                                {item.label}
                            </button>
                        ))}
                    </div>
                )}
            </div>

            {!canEdit && (
                <StudentLessonsCalendar
                    lessons={lessons}
                    isLoading={lessonsLoading}
                    tab={studentTab}
                />
            )}

            {canEdit && schedules === null && isLoading && (
                <p className="text-sm text-muted-foreground">{t("common.loading")}</p>
            )}

            {canEdit && schedules !== null && schedules.length === 0 && (
                <p className="text-sm text-muted-foreground">{t("schedule.noSchedules")}</p>
            )}

            {canEdit && sortedSchedules.length > 0 && (
                <TutorScheduleList
                    schedules={sortedSchedules}
                    dateLocale={dateLocale}
                    onDeleteSchedule={onDeleteSchedule}
                />
            )}

            {canEdit && (
                <CreateScheduleModal
                    isOpen={isOpen}
                    setIsOpen={setIsOpen}
                />
            )}
        </div>
    )
}
