import MonthCalendar from "@/widgets/calendars/ui/MonthlyCalendar";
import { useState } from "react";
import WeeklyCalendar from "@/widgets/calendars/ui/WeeklyCalendar.tsx";
import { Card, CardContent } from "@/shared/ui/Card.tsx";
import { Button } from "@/shared/ui/Button.tsx";
import { Label } from "@/shared/ui/Label.tsx";

const events = [
    {
        id: 1,
        title: "Математика",
        group: "Группа А",
        date: "2026-07-10",
        start: 10,
        end: 11,
    },
    {
        id: 2,
        title: "Физика",
        group: "Группа Б",
        date: "2026-07-15",
        start: 10,
        end: 11,
    },
]

const tabs = ["Неделя", "Месяц"]


export default function SchedulePage() {
    const [tab, setTab] = useState("Месяц")

    return (
        <div className="flex flex-col py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    Расписания
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    Отслеживайте запланированные уроки и домашние задания
                </Label>
            </div>
            <Card>
                <CardContent className="p-4 overflow-auto space-y-10">
                    <div className="flex flex-row justify-end gap-2 px-6">
                        {tabs.map(t => {
                            return (
                                <Button
                                    size="lg"
                                    variant={t == tab ? "outline" : "default"}
                                    onClick={() => setTab(t)}
                                >
                                    {t}
                                </Button>
                            )
                        })}
                    </div>

                    {tab == "Месяц"
                        ? (
                            <MonthCalendar
                                events={events}
                            />
                        )
                        : (
                            <WeeklyCalendar
                                events={events}
                            />
                        )
                    }
                </CardContent>
            </Card>
        </div>
    )
}
