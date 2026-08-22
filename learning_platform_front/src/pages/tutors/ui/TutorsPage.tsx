import { useTranslation } from "react-i18next"
import { Label } from "@/shared/ui/Label"
import { TutorsCatalog } from "@/widgets/tutorsCatalog"

export default function TutorsPage() {
    const { t } = useTranslation()

    return (
        <div className="space-y-6 px-6 py-8 lg:px-20 lg:py-10">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("tutors.title")}
                </Label>
                <Label className="text-sm font-normal text-primary/50 lg:text-base">
                    {t("tutors.subtitle")}
                </Label>
            </div>
            <TutorsCatalog />
        </div>
    )
}
