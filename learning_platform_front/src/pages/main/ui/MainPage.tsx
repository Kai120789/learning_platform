import { useTranslation } from "react-i18next"
import {
    GroupsRowList,
    HomeworkList,
    WeeklyScheduleStrip,
} from "@/widgets/mainDashboard"
import { Label } from "@/shared/ui/Label"

export default function MainPage() {
    const { t } = useTranslation()

    return (
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("main.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("main.subtitle")}
                </Label>
            </div>

            <WeeklyScheduleStrip />

            <div className="grid gap-5 lg:grid-cols-2">
                <HomeworkList />
                <GroupsRowList />
            </div>
        </div>
    )
}
