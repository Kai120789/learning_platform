import { useTranslation } from "react-i18next"
import {
    type ScheduleCalendarEvent,
    type ScheduleData,
} from "@/entities/schedule"
import { Card, CardContent } from "@/shared/ui/Card"
import { MonthlyCalendar, WeeklyCalendar } from "@/widgets/calendars"
import type { ScheduleTab } from "./ScheduleDetailsHeader"

type ScheduleDetailsCalendarProps = {
    isLoading: boolean
    schedule: ScheduleData | null | undefined
    tab: ScheduleTab
    events: ScheduleCalendarEvent[]
    selectedDate: Date
    onSelectDate: (date: Date) => void
}

export function ScheduleDetailsCalendar({
    isLoading,
    schedule,
    tab,
    events,
    selectedDate,
    onSelectDate,
}: ScheduleDetailsCalendarProps) {
    const { t } = useTranslation()

    return (
        <Card className="overflow-hidden">
            <CardContent className="p-3 lg:p-4">
                {isLoading && !schedule ? (
                    <p className="py-10 text-center text-sm text-muted-foreground">
                        {t("common.loading")}
                    </p>
                ) : !schedule ? (
                    <p className="py-10 text-center text-sm text-muted-foreground">
                        {t("schedule.notFound")}
                    </p>
                ) : tab === "month" ? (
                    <MonthlyCalendar
                        events={events}
                        selectedDate={selectedDate}
                        onSelectDate={onSelectDate}
                        periodStart={schedule.startTime}
                        periodEnd={schedule.endTime}
                    />
                ) : (
                    <WeeklyCalendar
                        events={events}
                        selectedDate={selectedDate}
                        onSelectDate={onSelectDate}
                        periodStart={schedule.startTime}
                        periodEnd={schedule.endTime}
                    />
                )}
            </CardContent>
        </Card>
    )
}
