import { useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { format, setHours, setMinutes, setSeconds, setMilliseconds } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { Plus } from "lucide-react"
import {
    mapLessonsToCalendarEvents,
    type LessonData,
} from "@/entities/lesson"
import { lessonStatusClass } from "@/shared/lib/statusStyles"
import { cn } from "@/shared/lib/utils"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { MonthlyCalendar, WeeklyCalendar } from "@/widgets/calendars"
import { StudentLessonModal } from "@/widgets/studentLessonModal"

type CalendarSpan = "week" | "month"

type LessonsCalendarViewProps = {
    lessons: LessonData[] | null
    isLoading: boolean
    canEdit: boolean
    onCreateAt?: (date: Date) => void
    onEdit: (lesson: LessonData) => void
    onStart: (lesson: LessonData) => void
    onComplete: (lesson: LessonData) => void
    onCancel: (lesson: LessonData) => void
}

function atDayHour(day: Date, hour: number) {
    return setMilliseconds(setSeconds(setMinutes(setHours(day, hour), 0), 0), 0)
}

export function LessonsCalendarView({
    lessons,
    isLoading,
    canEdit,
    onCreateAt,
    onEdit,
    onStart,
    onComplete,
    onCancel,
}: LessonsCalendarViewProps) {
    const { t, i18n } = useTranslation()
    const [span, setSpan] = useState<CalendarSpan>("month")
    const [selectedDate, setSelectedDate] = useState(new Date())
    const [openedLesson, setOpenedLesson] = useState<LessonData | null>(null)
    const dateLocale = i18n.language === "ru" ? ru : enUS

    const events = useMemo(
        () => mapLessonsToCalendarEvents(lessons),
        [lessons],
    )

    const selectedKey = format(selectedDate, "yyyy-MM-dd")
    const selectedEvents = events
        .filter((event) => event.date === selectedKey)
        .sort((a, b) => a.start - b.start)

    const lessonsById = useMemo(() => {
        const map = new Map<number, LessonData>()
        ;(lessons ?? []).forEach((lesson) => map.set(lesson.id, lesson))
        return map
    }, [lessons])

    const activeLesson = openedLesson
        ? (lessonsById.get(openedLesson.id) ?? openedLesson)
        : null

    const spans: { id: CalendarSpan; label: string }[] = [
        { id: "week", label: t("schedule.week") },
        { id: "month", label: t("schedule.month") },
    ]

    if (lessons === null && isLoading) {
        return <p className="text-sm text-muted-foreground">{t("common.loading")}</p>
    }

    return (
        <>
            <div className="grid gap-4 lg:grid-cols-[minmax(0,1fr)_300px]">
                <Card className="overflow-hidden">
                    <CardHeader className="flex flex-row items-center justify-end space-y-0 px-3 pb-0 pt-3 lg:px-4">
                        <div
                            className="inline-flex rounded-lg border border-border bg-background p-0.5"
                            role="group"
                            aria-label={t("schedule.calendarSpan")}
                        >
                            {spans.map((item) => {
                                const active = span === item.id
                                return (
                                    <button
                                        key={item.id}
                                        type="button"
                                        aria-pressed={active}
                                        onClick={() => setSpan(item.id)}
                                        className={cn(
                                            "cursor-pointer rounded-md px-2.5 py-1 text-xs font-medium transition-colors",
                                            active
                                                ? "bg-primary text-primary-foreground shadow-sm"
                                                : "text-muted-foreground hover:bg-muted hover:text-foreground",
                                        )}
                                    >
                                        {item.label}
                                    </button>
                                )
                            })}
                        </div>
                    </CardHeader>
                    <CardContent className="p-3 pt-2 lg:p-4 lg:pt-3">
                        {span === "month" ? (
                            <MonthlyCalendar
                                events={events}
                                selectedDate={selectedDate}
                                onSelectDate={setSelectedDate}
                                onCreateAt={canEdit ? (date) => onCreateAt?.(atDayHour(date, 10)) : undefined}
                            />
                        ) : (
                            <WeeklyCalendar
                                events={events}
                                selectedDate={selectedDate}
                                onSelectDate={setSelectedDate}
                                onCreateAt={canEdit
                                    ? (date, hour) => onCreateAt?.(atDayHour(date, hour))
                                    : undefined}
                            />
                        )}
                    </CardContent>
                </Card>

                <Card>
                    <CardHeader className="px-4 pb-2">
                        <div className="flex items-start justify-between gap-2">
                            <div className="min-w-0 space-y-1">
                                <CardTitle className="text-sm">
                                    {t("schedule.dayDetails")}
                                </CardTitle>
                                <p className="text-xs text-muted-foreground capitalize">
                                    {format(selectedDate, "EEEE, d MMMM", { locale: dateLocale })}
                                </p>
                            </div>
                            {canEdit && onCreateAt && (
                                <Button
                                    type="button"
                                    size="xs"
                                    variant="outline"
                                    className="md:h-7 md:gap-1 md:px-2.5 md:text-[0.8rem] md:[&_svg:not([class*='size-'])]:size-3.5"
                                    onClick={() => onCreateAt(atDayHour(selectedDate, 10))}
                                >
                                    <Plus className="size-3" />
                                    {t("lessons.create")}
                                </Button>
                            )}
                        </div>
                    </CardHeader>
                    <CardContent className="space-y-2 px-4 pb-4">
                        {selectedEvents.length === 0 ? (
                            <p className="text-sm text-muted-foreground">
                                {t("schedule.noEvents")}
                            </p>
                        ) : (
                            selectedEvents.map((event) => {
                                const lesson = lessonsById.get(event.lessonId ?? event.id)
                                return (
                                    <button
                                        type="button"
                                        key={event.id}
                                        className="flex w-full cursor-pointer items-start justify-between gap-2 rounded-md px-2 py-2 text-left transition-colors hover:bg-muted/50"
                                        onClick={() => {
                                            if (lesson) setOpenedLesson(lesson)
                                        }}
                                    >
                                        <div className="min-w-0 space-y-1">
                                            <div className="truncate text-sm font-medium">
                                                {t("schedule.bookedLesson", { id: event.lessonId ?? event.id })}
                                            </div>
                                            <div className="text-xs text-muted-foreground">
                                                {event.timeLabel}
                                            </div>
                                        </div>
                                        {lesson && (
                                            <Badge
                                                variant="outline"
                                                className={`shrink-0 text-[10px] ${lessonStatusClass(lesson.status)}`}
                                            >
                                                {t(`lessonStatus.${lesson.status}`)}
                                            </Badge>
                                        )}
                                    </button>
                                )
                            })
                        )}
                    </CardContent>
                </Card>
            </div>

            <StudentLessonModal
                isOpen={activeLesson !== null}
                setIsOpen={(open) => {
                    if (!open) setOpenedLesson(null)
                }}
                lesson={activeLesson}
                canEdit={canEdit}
                onEdit={(lesson) => {
                    setOpenedLesson(null)
                    onEdit(lesson)
                }}
                onStart={onStart}
                onComplete={onComplete}
                onCancel={onCancel}
            />
        </>
    )
}
