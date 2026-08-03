import { useMemo } from "react"
import {
    addDays,
    endOfDay,
    format,
    isSameDay,
    isWithinInterval,
    startOfDay,
    startOfWeek,
} from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { FaArrowLeft, FaArrowRight } from "react-icons/fa"
import { useTranslation } from "react-i18next"
import { lessonStatusClass } from "@/shared/lib/statusStyles"
import { cn } from "@/shared/lib/utils"
import { Badge } from "@/shared/ui/Badge"
import type { CalendarEvent } from "./MonthlyCalendar"

type WeeklyScheduleProps = {
    events: CalendarEvent[]
    selectedDate: Date
    onSelectDate: (date: Date) => void
    onCreateAt?: (date: Date, hour: number) => void
    periodStart?: string | Date
    periodEnd?: string | Date
}

const HOUR_START = 6
const HOUR_COUNT = 19
const hours = Array.from({ length: HOUR_COUNT }, (_, i) => HOUR_START + i)
const ROW_HEIGHT_PX = 48
const EVENT_INSET_X_PX = 6
const EVENT_INSET_Y_PX = 3
const DAY_START_MINUTES = HOUR_START * 60

function eventStartMinutes(event: CalendarEvent) {
    if (typeof event.startMinutes === "number") return event.startMinutes
    return event.start * 60
}

function eventEndMinutes(event: CalendarEvent) {
    if (typeof event.endMinutes === "number") return event.endMinutes
    return Math.max(event.end, event.start + 1) * 60
}

export default function WeeklySchedule({
    events,
    selectedDate,
    onSelectDate,
    onCreateAt,
    periodStart,
    periodEnd,
}: WeeklyScheduleProps) {
    const { t, i18n } = useTranslation()
    const dateLocale = i18n.language === "ru" ? ru : enUS

    const weekStart = startOfWeek(selectedDate, { weekStartsOn: 1 })
    const weekDays = Array.from({ length: 7 }, (_, i) => addDays(weekStart, i))

    const period = useMemo(() => {
        if (!periodStart || !periodEnd) return null
        const start = startOfDay(new Date(periodStart))
        const end = endOfDay(new Date(periodEnd))
        if (Number.isNaN(start.getTime()) || Number.isNaN(end.getTime())) return null
        if (end < start) return null
        return { start, end }
    }, [periodStart, periodEnd])

    const dayEvents = events
        .filter((event) => event.date === format(selectedDate, "yyyy-MM-dd"))
        .sort((a, b) => eventStartMinutes(a) - eventStartMinutes(b))

    const gridHeight = HOUR_COUNT * ROW_HEIGHT_PX

    const shiftWeek = (delta: number) => {
        onSelectDate(addDays(selectedDate, delta * 7))
    }

    return (
        <div className="rounded-lg border overflow-hidden">
            <div className="flex items-center justify-between border-b px-3 py-2">
                <FaArrowLeft
                    size={16}
                    className="cursor-pointer hover:text-primary/50"
                    onClick={() => shiftWeek(-1)}
                />

                <div className="text-sm font-semibold capitalize">
                    {format(selectedDate, "LLLL yyyy", { locale: dateLocale })}
                </div>

                <FaArrowRight
                    size={16}
                    className="cursor-pointer hover:text-primary/50"
                    onClick={() => shiftWeek(1)}
                />
            </div>

            <div className="grid grid-cols-7 border-b">
                {weekDays.map((date) => {
                    const inPeriod = period
                        ? isWithinInterval(date, { start: period.start, end: period.end })
                        : false
                    const isPeriodStart = period ? isSameDay(date, period.start) : false
                    const isPeriodEnd = period ? isSameDay(date, period.end) : false

                    return (
                        <button
                            type="button"
                            key={date.toString()}
                            onClick={() => onSelectDate(date)}
                            className={cn(
                                "cursor-pointer border-r p-2 text-xs last:border-r-0 hover:bg-muted",
                                inPeriod && !isPeriodStart && !isPeriodEnd && "bg-primary/10",
                                (isPeriodStart || isPeriodEnd) && "border-2 border-primary/60 bg-primary/25 last:border-2",
                                isSameDay(date, selectedDate) && "ring-2 ring-inset ring-primary",
                            )}
                        >
                            <div className="text-muted-foreground">
                                {format(date, "EEE", { locale: dateLocale })}
                            </div>
                            <div className="font-semibold text-sm">
                                {format(date, "d")}
                            </div>
                        </button>
                    )
                })}
            </div>

            <div className="max-h-[420px] overflow-y-auto">
                <div className="flex" style={{ height: gridHeight }}>
                    <div className="w-14 shrink-0 border-r">
                        {hours.map((hour) => (
                            <div
                                key={hour}
                                className="border-b px-2 py-1.5 text-xs text-muted-foreground last:border-b-0"
                                style={{ height: ROW_HEIGHT_PX }}
                            >
                                {hour}:00
                            </div>
                        ))}
                    </div>

                    <div className="relative min-w-0 flex-1">
                        {hours.map((hour) => (
                            <div
                                key={hour}
                                className={cn(
                                    "absolute inset-x-0 border-b border-border/80 last:border-b-0",
                                    onCreateAt && "cursor-pointer hover:bg-muted/40",
                                )}
                                style={{
                                    top: (hour - HOUR_START) * ROW_HEIGHT_PX,
                                    height: ROW_HEIGHT_PX,
                                }}
                                onClick={() => onCreateAt?.(selectedDate, hour)}
                                onKeyDown={(e) => {
                                    if (!onCreateAt) return
                                    if (e.key === "Enter" || e.key === " ") {
                                        e.preventDefault()
                                        onCreateAt(selectedDate, hour)
                                    }
                                }}
                                role={onCreateAt ? "button" : undefined}
                                tabIndex={onCreateAt ? 0 : undefined}
                            />
                        ))}

                        {dayEvents.map((event) => {
                            const startMin = Math.max(eventStartMinutes(event), DAY_START_MINUTES)
                            const endMin = Math.min(
                                eventEndMinutes(event),
                                DAY_START_MINUTES + HOUR_COUNT * 60,
                            )
                            if (endMin <= startMin) return null

                            const top =
                                ((startMin - DAY_START_MINUTES) / 60) * ROW_HEIGHT_PX
                                + EVENT_INSET_Y_PX
                            const height = Math.max(
                                ((endMin - startMin) / 60) * ROW_HEIGHT_PX
                                - EVENT_INSET_Y_PX * 2,
                                22,
                            )

                            return (
                                <div
                                    key={event.id}
                                    className="pointer-events-none absolute z-10 overflow-hidden rounded-md border border-primary/25 bg-primary/15 px-2 py-1 shadow-sm"
                                    style={{
                                        top,
                                        height,
                                        left: EVENT_INSET_X_PX,
                                        right: EVENT_INSET_X_PX,
                                    }}
                                    title={event.timeLabel ?? event.title}
                                >
                                    <div className="flex min-w-0 items-center justify-between gap-2">
                                        <div className="min-w-0 truncate text-xs font-medium leading-tight text-foreground">
                                            {event.title}
                                        </div>
                                        {event.lessonStatus && (
                                            <Badge
                                                variant="outline"
                                                className={cn(
                                                    "shrink-0 text-[10px]",
                                                    lessonStatusClass(event.lessonStatus),
                                                )}
                                            >
                                                {t(`lessonStatus.${event.lessonStatus}`)}
                                            </Badge>
                                        )}
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                </div>
            </div>
        </div>
    )
}
