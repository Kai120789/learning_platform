import { useState } from "react"
import { useTranslation } from "react-i18next"
import { FaPlus } from "react-icons/fa"
import { MonthlyCalendar, WeeklyCalendar, type CalendarEvent } from "@/widgets/calendars"
import { CreateLessonModal } from "@/widgets/createLessonModal"
import { Card, CardContent } from "@/shared/ui/Card"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"

type ScheduleTab = "week" | "month"

export default function SchedulePage() {
    const { t } = useTranslation()
    const [tab, setTab] = useState<ScheduleTab>("month")
    const [isOpen, setIsOpen] = useState(false)

    const events: CalendarEvent[] = [
        {
            id: 1,
            title: t("schedule.mockEvents.math"),
            group: t("schedule.mockEvents.groupA"),
            date: "2026-07-24",
            start: 10,
            end: 11,
        },
        {
            id: 2,
            title: t("schedule.mockEvents.physics"),
            group: t("schedule.mockEvents.groupB"),
            date: "2026-07-25",
            start: 14,
            end: 15,
        },
    ]

    const tabs: { id: ScheduleTab; label: string }[] = [
        { id: "week", label: t("schedule.week") },
        { id: "month", label: t("schedule.month") },
    ]

    return (
        <div className="flex flex-col py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <div className="flex justify-between items-center gap-4">
                    <Label className="text-2xl lg:text-4xl">
                        {t("schedule.title")}
                    </Label>
                    <Button onClick={() => setIsOpen(true)} size="lg" className="rounded-full">
                        <FaPlus className="size-3" />
                        {t("schedule.createLesson")}
                    </Button>
                </div>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("schedule.subtitle")}
                </Label>
            </div>
            <Card>
                <CardContent className="p-4 overflow-auto space-y-10">
                    <div className="flex flex-row justify-end gap-2 px-6">
                        {tabs.map((item) => (
                            <Button
                                key={item.id}
                                size="lg"
                                variant={item.id === tab ? "default" : "outline"}
                                onClick={() => setTab(item.id)}
                            >
                                {item.label}
                            </Button>
                        ))}
                    </div>

                    {tab === "month" ? (
                        <MonthlyCalendar events={events} />
                    ) : (
                        <WeeklyCalendar events={events} />
                    )}
                </CardContent>
            </Card>
            <CreateLessonModal
                isOpen={isOpen}
                setIsOpen={setIsOpen}
            />
        </div>
    )
}
