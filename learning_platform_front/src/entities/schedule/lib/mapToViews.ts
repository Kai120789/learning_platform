import { addMinutes, format, getDay } from "date-fns"
import type { ScheduleData, ScheduleSlotData, ScheduleSlotStatus, WeekDaySlotView } from "../types/types"

export type ScheduleCalendarEvent = {
    id: number
    title: string
    date: string
    start: number
    end: number
    timeLabel: string
    status: ScheduleSlotStatus
    lessonId?: number | null
}

function toDate(value: string) {
    return new Date(value)
}

function getSlotEnd(slot: ScheduleSlotData) {
    const start = toDate(slot.startTime)
    const duration = slot.duration && slot.duration > 0 ? slot.duration : 60
    return addMinutes(start, duration)
}

function getWeekdayIndex(date: Date): 0 | 1 | 2 | 3 | 4 | 5 | 6 {
    const day = getDay(date)
    return (day === 0 ? 6 : day - 1) as 0 | 1 | 2 | 3 | 4 | 5 | 6
}

export function getSlotTimeLabel(slot: ScheduleSlotData) {
    const start = toDate(slot.startTime)
    const end = getSlotEnd(slot)
    return `${format(start, "HH:mm")} – ${format(end, "HH:mm")}`
}

export function mapSlotToCalendarEvent(
    slot: ScheduleSlotData,
    labels: { freeSlot: string; bookedLesson: (lessonId: number) => string },
): ScheduleCalendarEvent {
    const start = toDate(slot.startTime)
    const end = getSlotEnd(slot)

    return {
        id: slot.id,
        title: slot.lessonId
            ? labels.bookedLesson(slot.lessonId)
            : labels.freeSlot,
        date: format(start, "yyyy-MM-dd"),
        start: start.getHours(),
        end: Math.max(end.getHours(), start.getHours() + 1),
        timeLabel: getSlotTimeLabel(slot),
        status: slot.status,
        lessonId: slot.lessonId ?? null,
    }
}

export function mapSchedulesToCalendarEvents(
    schedules: ScheduleData[] | null | undefined,
    labels: { freeSlot: string; bookedLesson: (lessonId: number) => string },
): ScheduleCalendarEvent[] {
    return (schedules ?? [])
        .flatMap((schedule) => schedule.slots)
        .map((slot) => mapSlotToCalendarEvent(slot, labels))
        .sort((a, b) => a.date.localeCompare(b.date) || a.start - b.start)
}

export function mapSchedulesToWeekSlots(
    schedules: ScheduleData[] | null | undefined,
    labels: { freeSlot: string; bookedLesson: (lessonId: number) => string },
): WeekDaySlotView[] {
    return (schedules ?? [])
        .flatMap((schedule) => schedule.slots)
        .map((slot) => {
            const start = toDate(slot.startTime)
            const end = getSlotEnd(slot)

            return {
                id: slot.id,
                weekday: getWeekdayIndex(start),
                date: format(start, "yyyy-MM-dd"),
                start: format(start, "HH:mm"),
                end: format(end, "HH:mm"),
                title: slot.lessonId
                    ? labels.bookedLesson(slot.lessonId)
                    : labels.freeSlot,
                subtitle: slot.status,
                status: slot.status,
            }
        })
}
