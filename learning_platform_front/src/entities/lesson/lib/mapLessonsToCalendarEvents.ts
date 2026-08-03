import { addMinutes, format } from "date-fns"
import type { LessonData, LessonStatus } from "../types/types"
import { getLessonLabel } from "./getLessonLabel"

export type LessonCalendarEvent = {
    id: number
    title: string
    date: string
    start: number
    end: number
    startMinutes: number
    endMinutes: number
    timeLabel: string
    lessonId: number
    lessonStatus: LessonStatus
}

export function mapLessonToCalendarEvent(lesson: LessonData): LessonCalendarEvent {
    const start = new Date(lesson.startTime)
    const duration = lesson.duration && lesson.duration > 0 ? lesson.duration : 60
    const end = addMinutes(start, duration)
    const startMinutes = start.getHours() * 60 + start.getMinutes()
    const endMinutes = end.getHours() * 60 + end.getMinutes()
        + (end.getDate() !== start.getDate() ? 24 * 60 : 0)

    return {
        id: lesson.id,
        title: getLessonLabel(lesson),
        date: format(start, "yyyy-MM-dd"),
        start: start.getHours(),
        end: Math.max(end.getHours(), start.getHours() + 1),
        startMinutes,
        endMinutes: Math.max(endMinutes, startMinutes + 15),
        timeLabel: `${format(start, "HH:mm")} – ${format(end, "HH:mm")}`,
        lessonId: lesson.id,
        lessonStatus: lesson.status,
    }
}

export function mapLessonsToCalendarEvents(
    lessons: LessonData[] | null | undefined,
): LessonCalendarEvent[] {
    return (lessons ?? [])
        .filter((lesson) => lesson.status !== "CANCELLED")
        .map(mapLessonToCalendarEvent)
        .sort((a, b) => a.date.localeCompare(b.date) || a.startMinutes - b.startMinutes)
}
