import { useTranslation } from "react-i18next"
import { formatFileSize, mimeToMediaKind } from "@/entities/material"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Separator } from "@/shared/ui/Separator"
import type { MaterialsTarget } from "../model/types"

type MaterialsInfoDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    target: MaterialsTarget | null
}

export function MaterialsInfoDialog({
    open,
    onOpenChange,
    target,
}: MaterialsInfoDialogProps) {
    const { t } = useTranslation()

    if (!target) return null

    const rows = target.kind === "folder"
        ? [
            { label: t("materials.name"), value: target.folder.title },
            { label: t("materials.type"), value: t("materials.folders") },
            { label: "ID", value: String(target.folder.id) },
        ]
        : [
            { label: t("materials.name"), value: target.file.title },
            {
                label: t("materials.type"),
                value: t(`mediaType.${mimeToMediaKind(target.file.mimeType)}`),
            },
            { label: t("materials.mime"), value: target.file.mimeType || "—" },
            { label: t("materials.size"), value: formatFileSize(target.file.size) },
            { label: "ID", value: String(target.file.id) },
        ]

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left">
                        {t("materials.infoTitle")}
                    </DialogTitle>
                </DialogHeader>
                <Separator className="my-1" />
                <dl className="space-y-3">
                    {rows.map((row) => (
                        <div key={row.label} className="grid grid-cols-[100px_1fr] gap-3 text-sm">
                            <dt className="text-muted-foreground">{row.label}</dt>
                            <dd className="min-w-0 break-all font-medium">{row.value}</dd>
                        </div>
                    ))}
                </dl>
            </DialogContent>
        </Dialog>
    )
}
