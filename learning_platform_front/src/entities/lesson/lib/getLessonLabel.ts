import { addMinutes, format } from "date-fns"
import type { LessonData } from "../types/types"

export function getLessonLabel(lesson: LessonData) {
    const start = new Date(lesson.startTime)
    if (Number.isNaN(start.getTime())) return "–"

    const duration = lesson.duration && lesson.duration > 0 ? lesson.duration : 60
    const end = addMinutes(start, duration)
    return `${format(start, "HH:mm")}–${format(end, "HH:mm")}`
}
