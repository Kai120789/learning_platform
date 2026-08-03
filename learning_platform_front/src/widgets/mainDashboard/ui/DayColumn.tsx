import { useTranslation } from "react-i18next"
import type { WeekDaySlotView } from "@/entities/schedule"

type DayColumnProps = {
    dayKey: string
    dateLabel: string
    isToday: boolean
    slots: WeekDaySlotView[]
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
            <div className="shrink-0 border-b px-2 py-1.5 text-center">
                <div className={`text-sm font-medium ${isToday ? "text-primary" : ""}`}>
                    {t(`common.weekdays.${dayKey}`)}
                </div>
                <div className="text-xs text-muted-foreground">{dateLabel}</div>
            </div>

            <div className="min-h-0 flex-1 space-y-1.5 overflow-y-auto p-1.5">
                {slots.length === 0 ? (
                    <div className="px-1 py-6 text-center text-xs text-muted-foreground">
                        —
                    </div>
                ) : (
                    slots.map((slot) => (
                        <div
                            key={slot.id}
                            className="rounded-md border bg-background px-2 py-1.5 space-y-0.5"
                        >
                            <div className="text-[11px] font-medium text-muted-foreground">
                                {slot.start} – {slot.end}
                            </div>
                            <div className="line-clamp-1 text-xs font-medium">
                                {slot.title}
                            </div>
                            {slot.subtitle && (
                                <div className="truncate text-[11px] text-muted-foreground">
                                    {t(`scheduleSlotStatus.${slot.status}`, {
                                        defaultValue: slot.status,
                                    })}
                                </div>
                            )}
                        </div>
                    ))
                )}
            </div>
        </div>
    )
}
