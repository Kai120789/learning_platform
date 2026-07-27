import { useTranslation } from "react-i18next"
import type { ScheduleSlotMock } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"

type DayColumnProps = {
    dayKey: string
    dateLabel: string
    isToday: boolean
    slots: ScheduleSlotMock[]
}

export function DayColumn({
    dayKey,
    dateLabel,
    isToday,
    slots,
}: DayColumnProps) {
    const { t } = useTranslation()

    return (
        <div
            className={`flex min-h-0 min-w-0 flex-col border-r last:border-r-0 ${
                isToday ? "bg-primary/5" : ""
            }`}
        >
            <div className="shrink-0 border-b px-2 py-2 text-center">
                <div className={`text-sm font-medium ${isToday ? "text-primary" : ""}`}>
                    {t(`common.weekdays.${dayKey}`)}
                </div>
                <div className="text-xs text-muted-foreground">{dateLabel}</div>
            </div>

            <div className="min-h-0 flex-1 space-y-2 overflow-y-auto p-2">
                {slots.length === 0 ? (
                    <div className="px-1 py-6 text-center text-xs text-muted-foreground">
                        —
                    </div>
                ) : (
                    slots.map((slot) => (
                        <div
                            key={slot.id}
                            className="rounded-lg border bg-background p-2 space-y-1"
                        >
                            <div className="text-xs font-medium text-muted-foreground">
                                {slot.start} – {slot.end}
                            </div>
                            <div className="line-clamp-2 text-sm font-medium">
                                {slot.subjectTitle}
                            </div>
                            <div className="truncate text-xs text-muted-foreground">
                                {slot.groupTitle}
                            </div>
                            <Badge variant="secondary" className="text-[10px]">
                                {t(`lessonStatus.${slot.status}`)}
                            </Badge>
                        </div>
                    ))
                )}
            </div>
        </div>
    )
}
