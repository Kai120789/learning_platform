import { useParams } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { Label } from "@/shared/ui/Label"
import { TutorProfile } from "@/widgets/tutorProfile"

export default function TutorPage() {
    const { t } = useTranslation()
    const { tutorId } = useParams()
    const parsedId = Number(tutorId)

    return (
        <div className="space-y-6 px-6 py-8 lg:px-20 lg:py-10">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("tutors.profileTitle")}
                </Label>
                <Label className="text-sm font-normal text-primary/50 lg:text-base">
                    {t("tutors.profileSubtitle")}
                </Label>
            </div>
            {Number.isFinite(parsedId) && parsedId > 0 ? (
                <TutorProfile tutorId={parsedId} />
            ) : (
                <div className="rounded-2xl border border-dashed border-border px-4 py-12 text-center text-sm text-muted-foreground">
                    {t("tutors.notFound")}
                </div>
            )}
        </div>
    )
}
