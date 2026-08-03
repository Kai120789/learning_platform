import { addMinutes, format } from "date-fns"
import type { LessonData } from "../types/types"
import { getLessonLabel } from "./getLessonLabel"

export type LessonCalendarEvent = {
    id: number
    title: string
    date: string
    start: number
    end: number
    timeLabel: string
    lessonId: number
}

export function mapLessonToCalendarEvent(lesson: LessonData): LessonCalendarEvent {
    const start = new Date(lesson.startTime)
    const duration = lesson.duration && lesson.duration > 0 ? lesson.duration : 60
    const end = addMinutes(start, duration)

    return {
        id: lesson.id,
        title: getLessonLabel(lesson),
        date: format(start, "yyyy-MM-dd"),
        start: start.getHours(),
        end: Math.max(end.getHours(), start.getHours() + 1),
        timeLabel: `${format(start, "HH:mm")} – ${format(end, "HH:mm")}`,
        lessonId: lesson.id,
    }
}

export function mapLessonsToCalendarEvents(
    lessons: LessonData[] | null | undefined,
): LessonCalendarEvent[] {
    return (lessons ?? [])
        .filter((lesson) => lesson.status !== "CANCELLED")
        .map(mapLessonToCalendarEvent)
        .sort((a, b) => a.date.localeCompare(b.date) || a.start - b.start)
}
