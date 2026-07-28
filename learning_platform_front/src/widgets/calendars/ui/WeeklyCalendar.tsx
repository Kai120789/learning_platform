import {
    addDays,
    format,
    isSameDay,
    startOfWeek,
} from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { FaArrowLeft, FaArrowRight } from "react-icons/fa"
import { useTranslation } from "react-i18next"
import { cn } from "@/shared/lib/utils"
import type { CalendarEvent } from "./MonthlyCalendar"

type WeeklyScheduleProps = {
    events: CalendarEvent[]
    selectedDate: Date
    onSelectDate: (date: Date) => void
}

const hours = Array.from({ length: 19 }, (_, i) => i + 6)

export default function WeeklySchedule({
    events,
    selectedDate,
    onSelectDate,
}: WeeklyScheduleProps) {
    const { i18n } = useTranslation()
    const dateLocale = i18n.language === "ru" ? ru : enUS

    const weekStart = startOfWeek(selectedDate, { weekStartsOn: 1 })
    const weekDays = Array.from({ length: 7 }, (_, i) => addDays(weekStart, i))

    const dayEvents = events.filter(
        (event) => event.date === format(selectedDate, "yyyy-MM-dd"),
    )

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
                {weekDays.map((date) => (
                    <button
                        type="button"
                        key={date.toString()}
                        onClick={() => onSelectDate(date)}
                        className={cn(
                            "border-r p-2 text-xs last:border-r-0 hover:bg-muted",
                            isSameDay(date, selectedDate) && "bg-primary/10",
                        )}
                    >
                        <div className="text-muted-foreground">
                            {format(date, "EEE", { locale: dateLocale })}
                        </div>
                        <div className="font-semibold text-sm">
                            {format(date, "d")}
                        </div>
                    </button>
                ))}
            </div>

            <div className="max-h-[420px] overflow-y-auto">
                {hours.map((hour) => {
                    const hourEvents = dayEvents.filter((event) => event.start === hour)

                    return (
                        <div
                            key={hour}
                            className="flex min-h-12 border-b last:border-b-0"
                        >
                            <div className="w-14 shrink-0 border-r px-2 py-1.5 text-xs text-muted-foreground">
                                {hour}:00
                            </div>

                            <div className="flex-1 space-y-1 p-1.5">
                                {hourEvents.map((event) => (
                                    <div
                                        key={event.id}
                                        className="rounded-md border bg-primary/10 px-2 py-1.5"
                                    >
                                        <div className="text-sm font-medium">
                                            {event.title}
                                        </div>
                                        {event.group && (
                                            <div className="text-xs text-muted-foreground">
                                                {event.group}
                                            </div>
                                        )}
                                        <div className="text-[11px] text-muted-foreground">
                                            {event.start}:00 – {event.end}:00
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}
