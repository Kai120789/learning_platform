import { format } from "date-fns"
import type { LessonData } from "../types/types"

export function getLessonLabel(lesson: LessonData) {
    const start = new Date(lesson.startTime)
    const time = Number.isNaN(start.getTime())
        ? "—"
        : format(start, "d MMM, HH:mm")

    return `#${lesson.id} · ${time} · ${lesson.duration} мин`
}
