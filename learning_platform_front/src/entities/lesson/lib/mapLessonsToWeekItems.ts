import { addMinutes, format, getDay } from "date-fns"
import type { LessonData, LessonStatus } from "../types/types"

export type LessonWeekItem = {
    id: number
    weekday: 0 | 1 | 2 | 3 | 4 | 5 | 6
    date: string
    start: string
    end: string
    title: string
    lessonId: number
    status: LessonStatus
}

function getWeekdayIndex(date: Date): 0 | 1 | 2 | 3 | 4 | 5 | 6 {
    const day = getDay(date)
    return (day === 0 ? 6 : day - 1) as 0 | 1 | 2 | 3 | 4 | 5 | 6
}

export function mapLessonToWeekItem(lesson: LessonData): LessonWeekItem {
    const start = new Date(lesson.startTime)
    const duration = lesson.duration && lesson.duration > 0 ? lesson.duration : 60
    const end = addMinutes(start, duration)

    return {
        id: lesson.id,
        weekday: getWeekdayIndex(start),
        date: format(start, "yyyy-MM-dd"),
        start: format(start, "HH:mm"),
        end: format(end, "HH:mm"),
        title: `#${lesson.id}`,
        lessonId: lesson.id,
        status: lesson.status,
    }
}

export function mapLessonsToWeekItems(
    lessons: LessonData[] | null | undefined,
): LessonWeekItem[] {
    return (lessons ?? [])
        .filter((lesson) => lesson.status !== "CANCELLED")
        .map(mapLessonToWeekItem)
        .sort((a, b) => a.date.localeCompare(b.date) || a.start.localeCompare(b.start))
}
