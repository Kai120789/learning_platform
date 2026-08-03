import { useTranslation } from "react-i18next"
import { Link } from "react-router-dom"
import { format } from "date-fns"
import type { Locale } from "date-fns"
import { FaArrowLeft, FaPen, FaTrash } from "react-icons/fa"
import { getRouteSchedule } from "@/app/router/routePaths"
import type { ScheduleData } from "@/entities/schedule"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { cn } from "@/shared/lib/utils"

export type ScheduleTab = "week" | "month"

type ScheduleDetailsHeaderProps = {
    scheduleId: number
    schedule: ScheduleData | null | undefined
    tab: ScheduleTab
    onTabChange: (tab: ScheduleTab) => void
    dateLocale: Locale
    onEdit: () => void
    onDeleteSchedule: () => void
}

export function ScheduleDetailsHeader({
    scheduleId,
    schedule,
    tab,
    onTabChange,
    dateLocale,
    onEdit,
    onDeleteSchedule,
}: ScheduleDetailsHeaderProps) {
    const { t } = useTranslation()

    const tabs: { id: ScheduleTab; label: string }[] = [
        { id: "week", label: t("schedule.week") },
        { id: "month", label: t("schedule.month") },
    ]

    return (
        <div className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
            <div className="space-y-2">
                <Link
                    to={getRouteSchedule()}
                    className="inline-flex items-center gap-2 text-sm text-muted-foreground hover:text-foreground"
                >
                    <FaArrowLeft className="size-3" />
                    {t("schedule.backToList")}
                </Link>
                <div className="space-y-1">
                    <Label className="text-xl lg:text-2xl">
                        {schedule?.title?.trim()
                            || t("schedule.scheduleItem", { id: scheduleId })}
                    </Label>
                    {schedule && (
                        <Label className="text-sm lg:text-base font-normal text-primary/50">
                            {format(new Date(schedule.startTime), "d MMM yyyy", { locale: dateLocale })}
                            {" – "}
                            {format(new Date(schedule.endTime), "d MMM yyyy", { locale: dateLocale })}
                        </Label>
                    )}
                </div>
            </div>
            <div className="flex flex-wrap items-center gap-2 sm:justify-end">
                <div className="flex rounded-lg border border-border p-0.5 bg-secondary/60">
                    {tabs.map((item) => (
                        <button
                            key={item.id}
                            type="button"
                            onClick={() => onTabChange(item.id)}
                            className={cn(
                                "rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors cursor-pointer",
                                item.id === tab
                                    ? "bg-primary text-primary-foreground shadow-sm"
                                    : "text-secondary-foreground hover:bg-muted"
                            )}
                        >
                            {item.label}
                        </button>
                    ))}
                </div>
                {schedule && (
                    <Button
                        type="button"
                        size="sm"
                        variant="outline"
                        onClick={onEdit}
                    >
                        <FaPen className="size-3" />
                        {t("schedule.editPeriod")}
                    </Button>
                )}
                <Button
                    type="button"
                    size="sm"
                    variant="outline"
                    onClick={onDeleteSchedule}
                >
                    <FaTrash className="size-3" />
                    {t("schedule.deleteSchedule")}
                </Button>
            </div>
        </div>
    )
}
