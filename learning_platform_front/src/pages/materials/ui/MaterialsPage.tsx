import { useTranslation } from "react-i18next"
import { MaterialsBrowser } from "@/widgets/materialsBrowser"
import { Label } from "@/shared/ui/Label"

export default function MaterialsPage() {
    const { t } = useTranslation()

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("materials.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("materials.subtitle")}
                </Label>
            </div>

            <MaterialsBrowser />
        </div>
    )
}
