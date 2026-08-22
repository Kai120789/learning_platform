import { useMemo, useState } from "react"
import {
    addMonths,
    eachDayOfInterval,
    endOfDay,
    endOfMonth,
    format,
    isSameDay,
    isSameMonth,
    isToday,
    isWithinInterval,
    startOfDay,
    startOfMonth,
    startOfWeek,
    endOfWeek,
} from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { ArrowLeft, ArrowRight } from "lucide-react"
import { useTranslation } from "react-i18next"
import { cn } from "@/shared/lib/utils"

export type CalendarEvent = {
    id: number
    title: string
    group?: string
    date: string
    start: number
    end: number
    startMinutes?: number
    endMinutes?: number
    timeLabel?: string
    status?: "FREE" | "BOOKED"
    lessonStatus?: "SCHEDULED" | "IN_PROCESS" | "COMPLETED" | "CANCELLED"
    lessonId?: number | null
}

type MonthlyCalendarProps = {
    events: CalendarEvent[]
    selectedDate: Date
    onSelectDate: (date: Date) => void
    onCreateAt?: (date: Date) => void
    periodStart?: string | Date
    periodEnd?: string | Date
}

export default function MonthlyCalendar({
    events,
    selectedDate,
    onSelectDate,
    onCreateAt,
    periodStart,
    periodEnd,
}: MonthlyCalendarProps) {
    const { t, i18n } = useTranslation()
    const [monthOverride, setMonthOverride] = useState<Date | null>(null)
    const dateLocale = i18n.language === "ru" ? ru : enUS
    const month = monthOverride ?? startOfMonth(selectedDate)

    const period = useMemo(() => {
        if (!periodStart || !periodEnd) return null
        const start = startOfDay(new Date(periodStart))
        const end = endOfDay(new Date(periodEnd))
        if (Number.isNaN(start.getTime()) || Number.isNaN(end.getTime())) return null
        if (end < start) return null
        return { start, end }
    }, [periodStart, periodEnd])

    const weekdays = [
        t("common.weekdays.mon"),
        t("common.weekdays.tue"),
        t("common.weekdays.wed"),
        t("common.weekdays.thu"),
        t("common.weekdays.fri"),
        t("common.weekdays.sat"),
        t("common.weekdays.sun"),
    ]

    const days = eachDayOfInterval({
        start: startOfWeek(startOfMonth(month), {
            weekStartsOn: 1,
        }),
        end: endOfWeek(endOfMonth(month), {
            weekStartsOn: 1,
        }),
    })

    return (
        <div className="rounded-lg border overflow-hidden">
            <div className="flex items-center justify-between px-3 py-2 border-b">
                <ArrowLeft
                    className="size-4 cursor-pointer hover:text-primary/50"
                    onClick={() => setMonthOverride(addMonths(month, -1))}
                />

                <h2 className="text-sm font-semibold capitalize">
                    {format(month, "LLLL yyyy", { locale: dateLocale })}
                </h2>

                <ArrowRight
                    className="size-4 cursor-pointer hover:text-primary/50"
                    onClick={() => setMonthOverride(addMonths(month, 1))}
                />
            </div>

            <div className="grid grid-cols-7">
                {weekdays.map((day) => (
                    <div
                        key={day}
                        className="border-b p-1.5 text-center text-[11px] text-muted-foreground"
                    >
                        {day}
                    </div>
                ))}
                {days.map((day) => {
                    const dayKey = format(day, "yyyy-MM-dd")
                    const dayEvents = events.filter((event) => event.date === dayKey)
                    const selected = isSameDay(day, selectedDate)
                    const inPeriod = period
                        ? isWithinInterval(day, { start: period.start, end: period.end })
                        : false
                    const isPeriodStart = period ? isSameDay(day, period.start) : false
                    const isPeriodEnd = period ? isSameDay(day, period.end) : false

                    return (
                        <button
                            type="button"
                            key={day.toString()}
                            onClick={() => onSelectDate(day)}
                            onDoubleClick={() => onCreateAt?.(day)}
                            className={cn(
                                "min-h-[5.5rem] cursor-pointer border p-1.5 text-left transition-colors hover:bg-muted/60",
                                !isSameMonth(day, month) && "bg-muted/40 text-muted-foreground",
                                inPeriod && !isPeriodStart && !isPeriodEnd && "bg-primary/10",
                                (isPeriodStart || isPeriodEnd) && "bg-primary/25 border-2 border-primary/60",
                                isToday(day) && !inPeriod && "bg-primary/5",
                                selected && "ring-2 ring-inset ring-primary",
                            )}
                        >
                            <div
                                className={cn(
                                    "mb-1 flex size-6 items-center justify-center rounded-full text-xs",
                                    isToday(day) && "bg-primary text-primary-foreground",
                                )}
                            >
                                {format(day, "d")}
                            </div>
                            <div className="space-y-0.5">
                                {dayEvents.slice(0, 2).map((event) => (
                                    <div
                                        key={event.id}
                                        className="truncate rounded border bg-primary/10 px-1 py-0.5 text-[10px] leading-tight"
                                    >
                                        {event.title}
                                    </div>
                                ))}
                                {dayEvents.length > 2 && (
                                    <div className="text-[10px] text-muted-foreground">
                                        +{dayEvents.length - 2}
                                    </div>
                                )}
                            </div>
                        </button>
                    )
                })}
            </div>
        </div>
    )
}
