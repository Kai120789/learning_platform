import { useTranslation } from "react-i18next"
import { format } from "date-fns"
import type { Locale } from "date-fns"
import { FaLink, FaPen, FaPlus, FaTrash, FaUnlink } from "react-icons/fa"
import {
    scheduleSlotStatusClass,
    type ScheduleCalendarEvent,
    type ScheduleData,
    type ScheduleSlotData,
} from "@/entities/schedule"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { cn } from "@/shared/lib/utils"

type ScheduleDayDetailsProps = {
    selectedDate: Date
    dateLocale: Locale
    selectedEvents: ScheduleCalendarEvent[]
    schedule: ScheduleData | null | undefined
    onAddSlots: () => void
    onEditSlot: (slot: ScheduleSlotData) => void
    onDeleteSlot: (slot: ScheduleSlotData) => void
    onUnbindLesson: (slot: ScheduleSlotData) => void
    onBindSlot: (slotId: number) => void
}

export function ScheduleDayDetails({
    selectedDate,
    dateLocale,
    selectedEvents,
    schedule,
    onAddSlots,
    onEditSlot,
    onDeleteSlot,
    onUnbindLesson,
    onBindSlot,
}: ScheduleDayDetailsProps) {
    const { t } = useTranslation()

    return (
        <div className="space-y-4">
            <Card>
                <CardHeader className="px-4 pb-2">
                    <div className="flex items-start justify-between gap-2">
                        <div className="min-w-0 space-y-1">
                            <CardTitle className="text-sm">
                                {t("schedule.dayDetails")}
                            </CardTitle>
                            <p className="text-xs text-muted-foreground capitalize">
                                {format(selectedDate, "EEEE, d MMMM", { locale: dateLocale })}
                            </p>
                        </div>
                        {schedule && (
                            <Button
                                type="button"
                                size="xs"
                                variant="outline"
                                onClick={onAddSlots}
                            >
                                <FaPlus className="size-3" />
                                {t("schedule.addSlot")}
                            </Button>
                        )}
                    </div>
                </CardHeader>
                <CardContent className="space-y-2 px-4 pb-4">
                    {selectedEvents.length === 0 ? (
                        <p className="text-sm text-muted-foreground">
                            {t("schedule.noEvents")}
                        </p>
                    ) : (
                        selectedEvents.map((event) => {
                            const slot = schedule?.slots.find((item) => item.id === event.id)
                            const isFree = slot?.status === "FREE"

                            return (
                                <div
                                    key={event.id}
                                    role={isFree ? "button" : undefined}
                                    tabIndex={isFree ? 0 : undefined}
                                    onClick={() => {
                                        if (slot && isFree) onBindSlot(slot.id)
                                    }}
                                    onKeyDown={(e) => {
                                        if (!slot || !isFree) return
                                        if (e.key === "Enter" || e.key === " ") {
                                            e.preventDefault()
                                            onBindSlot(slot.id)
                                        }
                                    }}
                                    className={cn(
                                        "rounded-lg border p-2.5 space-y-2 transition-colors",
                                        isFree && "border-dashed border-primary/40 bg-primary/5 cursor-pointer hover:bg-primary/10 hover:border-primary/60",
                                    )}
                                >
                                    <div className="flex items-start justify-between gap-2">
                                        <div className="min-w-0 space-y-1">
                                            <div className="text-sm font-medium">{event.title}</div>
                                            <Badge
                                                variant="outline"
                                                className={cn(
                                                    "text-[10px]",
                                                    event.status && scheduleSlotStatusClass(event.status),
                                                )}
                                            >
                                                {t(`scheduleSlotStatus.${event.status}`)}
                                            </Badge>
                                        </div>
                                        {slot && (
                                            <div
                                                className="flex items-center gap-1 shrink-0"
                                                onClick={(e) => e.stopPropagation()}
                                            >
                                                <Button
                                                    type="button"
                                                    size="icon-xs"
                                                    variant="outline"
                                                    onClick={() => onEditSlot(slot)}
                                                    aria-label={t("schedule.editSlot")}
                                                >
                                                    <FaPen className="size-3" />
                                                </Button>
                                                {isFree ? (
                                                    <Button
                                                        type="button"
                                                        size="icon-xs"
                                                        variant="outline"
                                                        onClick={() => onDeleteSlot(slot)}
                                                        aria-label={t("schedule.deleteSlot")}
                                                    >
                                                        <FaTrash className="size-3" />
                                                    </Button>
                                                ) : (
                                                    <Button
                                                        type="button"
                                                        size="icon-xs"
                                                        variant="outline"
                                                        onClick={() => onUnbindLesson(slot)}
                                                        aria-label={t("schedule.unbindLesson")}
                                                    >
                                                        <FaUnlink className="size-3" />
                                                    </Button>
                                                )}
                                            </div>
                                        )}
                                    </div>
                                    <div className="flex items-end justify-between gap-2">
                                        <span className="text-xs text-muted-foreground">
                                            {event.timeLabel}
                                        </span>
                                        {slot && isFree && (
                                            <Button
                                                type="button"
                                                size="icon-xs"
                                                variant="ghost"
                                                className="-mb-1 -mr-1 text-primary hover:bg-primary/10 hover:text-primary"
                                                onClick={(e) => {
                                                    e.stopPropagation()
                                                    onBindSlot(slot.id)
                                                }}
                                                aria-label={t("schedule.bindLesson")}
                                            >
                                                <FaLink className="size-3" />
                                            </Button>
                                        )}
                                    </div>
                                </div>
                            )
                        })
                    )}
                </CardContent>
            </Card>
        </div>
    )
}
