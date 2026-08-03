import { useEffect, useMemo } from "react"
import { useTranslation } from "react-i18next"
import { format } from "date-fns"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { DayColumn } from "./DayColumn"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    getSchedules,
    getSchedulesByTutorId,
    mapSchedulesToWeekSlots,
} from "@/entities/schedule"
import { getUserFullData, useCanEdit } from "@/entities/user"

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
    const schedules = useAppSelector(getSchedules)
    const weekDays = getWeekDates()
    const todayKey = new Date().toDateString()

    useEffect(() => {
        if (!canEdit || !userData?.user.userID) return
        dispatch(getSchedulesByTutorId(userData.user.userID))
    }, [canEdit, dispatch, userData?.user.userID])

    const weekSlots = useMemo(
        () => mapSchedulesToWeekSlots(schedules, {
            freeSlot: t("schedule.freeSlot"),
            bookedLesson: (lessonId) => t("schedule.bookedLesson", { id: lessonId }),
        }),
        [schedules, t],
    )

    return (
        <Card>
            <CardHeader className="pb-2">
                <CardTitle>{t("main.weekSchedule")}</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="overflow-x-auto">
                    <div className="grid min-w-[640px] grid-cols-7 overflow-hidden rounded-xl border h-72 lg:h-80">
                        {weekDays.map((day) => {
                            const dateKey = format(day.date, "yyyy-MM-dd")
                            const slots = weekSlots
                                .filter((slot) =>
                                    slot.weekday === day.weekday && slot.date === dateKey
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
                                    slots={slots}
                                />
                            )
                        })}
                    </div>
                </div>
            </CardContent>
        </Card>
    )
}
