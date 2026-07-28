import { useState } from "react"
import { useTranslation } from "react-i18next"
import { format } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { FaPlus } from "react-icons/fa"
import { MonthlyCalendar, WeeklyCalendar, type CalendarEvent } from "@/widgets/calendars"
import { CreateLessonModal } from "@/widgets/createLessonModal"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { cn } from "@/shared/lib/utils"

type ScheduleTab = "week" | "month"

export default function SchedulePage() {
    const { t, i18n } = useTranslation()
    const [tab, setTab] = useState<ScheduleTab>("month")
    const [isOpen, setIsOpen] = useState(false)
    const [selectedDate, setSelectedDate] = useState(new Date())
    const dateLocale = i18n.language === "ru" ? ru : enUS
    const todayKey = format(new Date(), "yyyy-MM-dd")

    const events: CalendarEvent[] = [
        {
            id: 1,
            title: t("schedule.mockEvents.math"),
            group: t("schedule.mockEvents.groupA"),
            date: "2026-07-24",
            start: 10,
            end: 11,
        },
        {
            id: 2,
            title: t("schedule.mockEvents.physics"),
            group: t("schedule.mockEvents.groupB"),
            date: "2026-07-25",
            start: 14,
            end: 15,
        },
        {
            id: 3,
            title: t("schedule.mockEvents.math"),
            group: t("schedule.mockEvents.groupA"),
            date: todayKey,
            start: 12,
            end: 13,
        },
    ]

    const tabs: { id: ScheduleTab; label: string }[] = [
        { id: "week", label: t("schedule.week") },
        { id: "month", label: t("schedule.month") },
    ]

    const selectedKey = format(selectedDate, "yyyy-MM-dd")
    const selectedEvents = events
        .filter((event) => event.date === selectedKey)
        .sort((a, b) => a.start - b.start)

    const upcomingEvents = events
        .filter((event) => event.date >= todayKey)
        .sort((a, b) => a.date.localeCompare(b.date) || a.start - b.start)
        .slice(0, 5)

    return (
        <div className="flex flex-col py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
                <div className="space-y-1">
                    <Label className="text-xl lg:text-2xl">
                        {t("schedule.title")}
                    </Label>
                    <Label className="text-sm lg:text-base font-normal text-primary/50">
                        {t("schedule.subtitle")}
                    </Label>
                </div>
                <div className="flex flex-wrap items-center gap-2 sm:justify-end">
                    <div className="flex rounded-lg border border-border p-0.5 bg-secondary/60">
                        {tabs.map((item) => (
                            <button
                                key={item.id}
                                type="button"
                                onClick={() => setTab(item.id)}
                                className={cn(
                                    "rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                                    item.id === tab
                                        ? "bg-primary text-primary-foreground shadow-sm"
                                        : "text-secondary-foreground hover:bg-muted"
                                )}
                            >
                                {item.label}
                            </button>
                        ))}
                    </div>
                    <Button onClick={() => setIsOpen(true)} size="sm" className="rounded-full">
                        <FaPlus className="size-3" />
                        {t("schedule.createLesson")}
                    </Button>
                </div>
            </div>

            <div className="grid gap-4 lg:grid-cols-[minmax(0,1fr)_280px]">
                <Card className="overflow-hidden">
                    <CardContent className="p-3 lg:p-4">
                        {tab === "month" ? (
                            <MonthlyCalendar
                                events={events}
                                selectedDate={selectedDate}
                                onSelectDate={setSelectedDate}
                            />
                        ) : (
                            <WeeklyCalendar
                                events={events}
                                selectedDate={selectedDate}
                                onSelectDate={setSelectedDate}
                            />
                        )}
                    </CardContent>
                </Card>

                <div className="space-y-4">
                    <Card>
                        <CardHeader className="pb-2">
                            <CardTitle className="text-sm">
                                {t("schedule.dayDetails")}
                            </CardTitle>
                            <p className="text-xs text-muted-foreground capitalize">
                                {format(selectedDate, "EEEE, d MMMM", { locale: dateLocale })}
                            </p>
                        </CardHeader>
                        <CardContent className="space-y-2">
                            {selectedEvents.length === 0 ? (
                                <p className="text-sm text-muted-foreground">
                                    {t("schedule.noEvents")}
                                </p>
                            ) : (
                                selectedEvents.map((event) => (
                                    <div
                                        key={event.id}
                                        className="rounded-lg border p-2.5 space-y-1"
                                    >
                                        <div className="text-sm font-medium">{event.title}</div>
                                        {event.group && (
                                            <div className="text-xs text-muted-foreground">
                                                {event.group}
                                            </div>
                                        )}
                                        <Badge variant="secondary" className="text-[10px]">
                                            {event.start}:00 – {event.end}:00
                                        </Badge>
                                    </div>
                                ))
                            )}
                        </CardContent>
                    </Card>

                    <Card>
                        <CardHeader className="pb-2">
                            <CardTitle className="text-sm">
                                {t("schedule.upcoming")}
                            </CardTitle>
                        </CardHeader>
                        <CardContent className="space-y-2">
                            {upcomingEvents.map((event) => (
                                <button
                                    type="button"
                                    key={`upcoming-${event.id}`}
                                    className="w-full rounded-lg border p-2.5 text-left transition-colors hover:bg-muted/50"
                                    onClick={() => setSelectedDate(new Date(`${event.date}T12:00:00`))}
                                >
                                    <div className="text-sm font-medium">{event.title}</div>
                                    <div className="text-xs text-muted-foreground">
                                        {format(new Date(`${event.date}T12:00:00`), "d MMM", { locale: dateLocale })}
                                        {" · "}
                                        {event.start}:00
                                    </div>
                                </button>
                            ))}
                        </CardContent>
                    </Card>
                </div>
            </div>

            <CreateLessonModal
                isOpen={isOpen}
                setIsOpen={setIsOpen}
            />
        </div>
    )
}
