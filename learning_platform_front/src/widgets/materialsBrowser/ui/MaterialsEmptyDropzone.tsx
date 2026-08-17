import { useTranslation } from "react-i18next"
import { cn } from "@/shared/lib/utils"

type MaterialsEmptyDropzoneProps = {
    canUpload: boolean
    className?: string
}

export function MaterialsEmptyDropzone({
    canUpload,
    className,
}: MaterialsEmptyDropzoneProps) {
    const { t } = useTranslation()

    return (
        <div
            className={cn(
                "flex min-h-56 flex-col items-center justify-center rounded-xl border border-dashed border-border px-6 py-16 text-center",
                className
            )}
        >
            <p className="text-muted-foreground">{t("materials.emptyFolder")}</p>
            {canUpload && (
                <p className="mt-1 text-sm text-muted-foreground/70">
                    {t("materials.dropHint")}
                </p>
            )}
        </div>
    )
}
