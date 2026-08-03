import { useTranslation } from "react-i18next"
import type { LessonWeekItem } from "@/entities/lesson"
import { lessonStatusClass } from "@/shared/lib/statusStyles"
import { Badge } from "@/shared/ui/Badge"
import { cn } from "@/shared/lib/utils"

type DayColumnProps = {
    dayKey: string
    dateLabel: string
    isToday: boolean
    items: LessonWeekItem[]
    onSelectLesson: (lessonId: number) => void
}

export function DayColumn({
    dayKey,
    dateLabel,
    isToday,
    items,
    onSelectLesson,
}: DayColumnProps) {
    const { t } = useTranslation()

    return (
        <div
            className={`flex min-h-0 min-w-0 flex-col border-r last:border-r-0 ${isToday ? "bg-primary/5" : ""
                }`}
        >
            <div className="shrink-0 border-b px-2 py-1.5 text-center">
                <div className={`text-sm font-medium ${isToday ? "text-primary" : ""}`}>
                    {t(`common.weekdays.${dayKey}`)}
                </div>
                <div className="text-xs text-muted-foreground">{dateLabel}</div>
            </div>

            <div className="min-h-0 flex-1 space-y-1.5 overflow-y-auto p-1.5">
                {items.length === 0 ? (
                    <div className="px-1 py-6 text-center text-xs text-muted-foreground">
                        –
                    </div>
                ) : (
                    items.map((item) => (
                        <button
                            type="button"
                            key={item.id}
                            className="w-full cursor-pointer rounded-md border bg-background px-2 py-1.5 text-left transition-colors hover:bg-muted/50"
                            onClick={() => onSelectLesson(item.lessonId)}
                        >
                            <div className="truncate text-xs font-medium">
                                {item.start}–{item.end}
                            </div>
                            <Badge
                                variant="outline"
                                className={cn("mt-1 text-[10px]", lessonStatusClass(item.status))}
                            >
                                {t(`lessonStatus.${item.status}`)}
                            </Badge>
                        </button>
                    ))
                )}
            </div>
        </div>
    )
}
