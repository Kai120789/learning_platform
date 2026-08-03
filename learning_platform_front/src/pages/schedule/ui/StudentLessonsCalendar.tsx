import { useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { format } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import {
    mapLessonsToCalendarEvents,
    type LessonData,
} from "@/entities/lesson"
import { lessonStatusClass } from "@/shared/lib/statusStyles"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { MonthlyCalendar, WeeklyCalendar } from "@/widgets/calendars"
import { StudentLessonModal } from "@/widgets/studentLessonModal"

export type ScheduleTab = "week" | "month"

type StudentLessonsCalendarProps = {
    lessons: LessonData[] | null
    isLoading: boolean
    tab: ScheduleTab
}

export function StudentLessonsCalendar({
    lessons,
    isLoading,
    tab,
}: StudentLessonsCalendarProps) {
    const { t, i18n } = useTranslation()
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

    if (lessons === null && isLoading) {
        return <p className="text-sm text-muted-foreground">{t("common.loading")}</p>
    }

    if (lessons !== null && lessons.length === 0) {
        return <p className="text-sm text-muted-foreground">{t("schedule.noStudentLessons")}</p>
    }

    return (
        <>
            <div className="grid gap-4 lg:grid-cols-[minmax(0,1fr)_300px]">
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

                <Card>
                    <CardHeader className="px-4 pb-2">
                        <CardTitle className="text-sm">
                            {t("schedule.dayDetails")}
                        </CardTitle>
                        <p className="text-xs text-muted-foreground capitalize">
                            {format(selectedDate, "EEEE, d MMMM", { locale: dateLocale })}
                        </p>
                    </CardHeader>
                    <CardContent className="space-y-2 px-4 pb-4">
                        {selectedEvents.length === 0 ? (
                            <p className="text-sm text-muted-foreground">
                                {t("schedule.noEvents")}
                            </p>
                        ) : (
                            selectedEvents.map((event) => {
                                const lesson = lessonsById.get(event.lessonId)
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
                                                {t("schedule.bookedLesson", { id: event.lessonId })}
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
                isOpen={openedLesson !== null}
                setIsOpen={(open) => {
                    if (!open) setOpenedLesson(null)
                }}
                lesson={openedLesson}
            />
        </>
    )
}
