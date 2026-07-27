import { useTranslation } from "react-i18next"
import { MaterialsBrowser } from "@/widgets/materialsBrowser"
import { Label } from "@/shared/ui/Label"

export default function MaterialsPage() {
    const { t } = useTranslation()

    return (
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("materials.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("materials.subtitle")}
                </Label>
            </div>

            <MaterialsBrowser />
        </div>
    )
}
