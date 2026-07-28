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
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("main.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
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
