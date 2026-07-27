import { useTranslation } from "react-i18next"
import { mockWeekSlots } from "@/shared/mocks"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { DayColumn } from "./DayColumn"

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
    const weekDays = getWeekDates()
    const todayKey = new Date().toDateString()

    return (
        <Card>
            <CardHeader className="pb-2">
                <CardTitle>{t("main.weekSchedule")}</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="overflow-x-auto">
                    <div className="grid min-w-[720px] grid-cols-7 overflow-hidden rounded-xl border h-72">
                        {weekDays.map((day) => {
                            const slots = mockWeekSlots
                                .filter((slot) => slot.weekday === day.weekday)
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
