import { useTranslation } from "react-i18next"
import { useNavigate } from "react-router-dom"
import { format } from "date-fns"
import type { Locale } from "date-fns"
import { FaTrash } from "react-icons/fa"
import { getRouteScheduleDetails } from "@/app/router/routePaths"
import {
    scheduleSlotStatusClass,
    type ScheduleData,
} from "@/entities/schedule"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Card, CardContent } from "@/shared/ui/Card"
import { cn } from "@/shared/lib/utils"

type TutorScheduleListProps = {
    schedules: ScheduleData[]
    dateLocale: Locale
    onDeleteSchedule: (e: React.MouseEvent, schedule: ScheduleData) => void
}

export function TutorScheduleList({
    schedules,
    dateLocale,
    onDeleteSchedule,
}: TutorScheduleListProps) {
    const { t } = useTranslation()
    const navigate = useNavigate()

    return (
        <div className="flex flex-col gap-2">
            {schedules.map((schedule) => {
                const freeCount = schedule.slots.filter((slot) => slot.status === "FREE").length
                const bookedCount = schedule.slots.filter((slot) => slot.status === "BOOKED").length

                return (
                    <Card
                        key={schedule.id}
                        size="sm"
                        className="cursor-pointer transition-all hover:shadow-md"
                        onClick={() => navigate(getRouteScheduleDetails(schedule.id))}
                    >
                        <CardContent className="flex items-center justify-between gap-4 py-3">
                            <div className="min-w-0 space-y-1">
                                <div className="text-sm font-medium truncate">
                                    {schedule.title?.trim()
                                        || t("schedule.scheduleItem", { id: schedule.id })}
                                </div>
                                <div className="text-xs text-muted-foreground">
                                    {format(new Date(schedule.startTime), "d MMM yyyy", { locale: dateLocale })}
                                    {" – "}
                                    {format(new Date(schedule.endTime), "d MMM yyyy", { locale: dateLocale })}
                                </div>
                            </div>

                            <div className="flex items-center gap-2">
                                <span className="text-xs text-muted-foreground">
                                    {t("schedule.slotsCount", { count: schedule.slots.length })}
                                </span>
                                <Badge
                                    variant="outline"
                                    className={cn("text-[10px]", scheduleSlotStatusClass("FREE"))}
                                >
                                    {t("schedule.freeCount", { count: freeCount })}
                                </Badge>
                                <Badge
                                    variant="outline"
                                    className={cn("text-[10px]", scheduleSlotStatusClass("BOOKED"))}
                                >
                                    {t("schedule.bookedCount", { count: bookedCount })}
                                </Badge>
                                <Button
                                    type="button"
                                    size="icon-sm"
                                    variant="outline"
                                    onClick={(e) => onDeleteSchedule(e, schedule)}
                                    aria-label={t("schedule.deleteSchedule")}
                                >
                                    <FaTrash className="size-3.5" />
                                </Button>
                            </div>
                        </CardContent>
                    </Card>
                )
            })}
        </div>
    )
}
