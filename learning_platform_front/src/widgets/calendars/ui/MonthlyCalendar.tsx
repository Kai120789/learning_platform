import { useState } from "react"
import {
    addMonths,
    eachDayOfInterval,
    endOfMonth,
    format,
    isSameMonth,
    isToday,
    startOfMonth,
    startOfWeek,
    endOfWeek,
} from "date-fns"
import { ru } from "date-fns/locale"
import { FaArrowLeft, FaArrowRight } from "react-icons/fa";


type CalendarEvent = {
    id: number
    title: string
    group?: string
    date: string
    start: number
    end: number
}

type MonthlyCalendarProps = {
    events: CalendarEvent[]
}


export default function MonthlyCalendar({
    events,
}: MonthlyCalendarProps) {
    const [month, setMonth] = useState(new Date())

    const days = eachDayOfInterval({
        start: startOfWeek(startOfMonth(month), {
            weekStartsOn: 1,
        }),
        end: endOfWeek(endOfMonth(month), {
            weekStartsOn: 1,
        }),
    })

    return (
        <div className="px-6">
            <div className="rounded-lg border overflow-hidden">
                <div className="flex items-center justify-between p-4">
                    <FaArrowLeft
                        size={25}
                        className="hover:text-primary/50"
                        onClick={() =>
                            setMonth(addMonths(month, -1))
                        }
                    />

                    <h2 className="text-xl font-semibold capitalize">
                        {format(month, "LLLL yyyy", {
                            locale: ru,
                        })}
                    </h2>

                    <FaArrowRight
                        size={25}
                        className="hover:text-primary/50"
                        onClick={() =>
                            setMonth(addMonths(month, 1))
                        }
                    />
                </div>

                <div className="grid grid-cols-7 overflow-hidden rounded-lg">
                    {["Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"].map(day => (
                        <div
                            key={day}
                            className="border-b p-2 text-center text-sm text-muted-foreground"
                        >
                            {day}
                        </div>

                    ))}
                    {days.map(day => {
                        const dayEvents = events.filter(
                            event =>
                                event.date ===
                                format(day, "yyyy-MM-dd")
                        )

                        return (
                            <div
                                key={day.toString()}
                                className={`min-h-32 border p-2
                                ${!isSameMonth(day, month)
                                        ? "bg-muted text-muted-foreground"
                                        : ""
                                    }
                                ${isToday(day)
                                        ? "border-8 border-primary font-semibold"
                                        : ""
                                    }
                            `}
                            >
                                <div
                                    className={`flex h-7 w-7 items-center justify-center rounded-full text-sm`}
                                >
                                    {format(day, "d")}
                                </div>
                                <div className="mt-2 space-y-1">
                                    {dayEvents.map(event => (
                                        <div
                                            key={event.id}
                                            className="rounded-md border bg-primary/10 p-2 text-xs"
                                        >
                                            <div className="font-medium">
                                                {event.title}
                                            </div>

                                            {event.group && (
                                                <div className="text-muted-foreground">
                                                    {event.group}
                                                </div>
                                            )}
                                        </div>
                                    ))}
                                </div>
                            </div>
                        )
                    })}
                </div>
            </div>
        </div>
    )
}