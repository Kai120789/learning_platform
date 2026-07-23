import { useState } from "react"
import {
    addDays,
    format,
    startOfWeek,
} from "date-fns"
import { ru } from "date-fns/locale"
import { FaArrowLeft, FaArrowRight } from "react-icons/fa"

type CalendarEvent = {
    id: number
    title: string
    group?: string
    date: string
    start: number
    end: number
}

type WeeklyScheduleProps = {
    events: CalendarEvent[]
}


const hours = Array.from(
    { length: 19 },
    (_, i) => i + 6
)


export default function WeeklySchedule({
    events,
}: WeeklyScheduleProps) {
    const [week, setWeek] = useState(new Date())
    const [day, setDay] = useState(new Date())

    const weekDays = Array.from(
        { length: 7 },
        (_, i) =>
            addDays(
                startOfWeek(week, {
                    weekStartsOn: 1,
                }),
                i
            )
    )

    const dayEvents = events.filter(
        event =>
            event.date === format(day, "yyyy-MM-dd")
    )

    return (
        <div className="px-6">
            <div className="rounded-lg border overflow-hidden">
                <div className="flex items-center justify-between border-b p-4">
                    <FaArrowLeft
                        size={25}
                        className="hover:text-primary/50"
                        onClick={() =>
                            setWeek(
                                addDays(week, -7)
                            )
                        }
                    />

                    <div className="font-semibold text-lg">
                        {format(day, "LLLL yyyy", {
                            locale: ru,
                        })}
                    </div>

                    <FaArrowRight
                        size={25}
                        className="hover:text-primary/50"
                        onClick={() =>
                            setWeek(
                                addDays(week, 7)
                            )
                        }
                    />
                </div>

                <div className="grid grid-cols-7 border-b">
                    {weekDays.map(date => (
                        <button
                            key={date.toString()}
                            onClick={() => setDay(date)}
                            className={` border-r p-3 text-sm hover:bg-muted
                            ${format(date, "yyyy-MM-dd") ===
                                    format(day, "yyyy-MM-dd")
                                    ? "bg-primary/10"
                                    : ""
                                }
                        `}
                        >
                            <div className="text-muted-foreground">
                                {format(date, "EEE", {
                                    locale: ru,
                                })}
                            </div>
                            <div className="font-semibold text-lg">
                                {format(date, "d")}
                            </div>
                        </button>
                    ))}
                </div>

                <div className="border-t">
                    {hours.map(hour => {
                        const hourEvents =
                            dayEvents.filter(
                                event =>
                                    event.start === hour
                            )
                        return (
                            <div
                                key={hour}
                                className="flex min-h-24 border-b last:border-b-0"
                            >
                                <div className="w-20 shrink-0 border-r p-3 text-sm text-muted-foreground">
                                    {hour}:00
                                </div>

                                <div className="flex-1 p-2 space-y-2">
                                    {hourEvents.map(event => (
                                        <div
                                            key={event.id}
                                            className="rounded-lg border bg-primary/10 p-3"
                                        >
                                            <div className="font-medium">
                                                {event.title}
                                            </div>

                                            {event.group && (
                                                <div className="text-sm text-muted-foreground">
                                                    {event.group}
                                                </div>
                                            )}

                                            <div className="mt-1 text-xs text-muted-foreground">
                                                {event.start}:00 - {event.end}:00
                                            </div>
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