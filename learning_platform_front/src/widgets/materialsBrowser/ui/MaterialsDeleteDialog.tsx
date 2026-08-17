import { Trans, useTranslation } from "react-i18next"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Separator } from "@/shared/ui/Separator"

type MaterialsDeleteDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    names: string[]
    onConfirm: () => void
}

export function MaterialsDeleteDialog({
    open,
    onOpenChange,
    names,
    onConfirm,
}: MaterialsDeleteDialogProps) {
    const { t } = useTranslation()

    if (names.length === 0) return null

    const isBulk = names.length > 1

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left">
                        {t("materials.deleteTitle")}
                    </DialogTitle>
                </DialogHeader>
                <Separator className="my-1" />
                <p className="text-sm text-muted-foreground">
                    {isBulk ? (
                        t("materials.deleteConfirmBulk", { count: names.length })
                    ) : (
                        <Trans
                            i18nKey="materials.deleteConfirm"
                            values={{ name: names[0] }}
                            components={{
                                bold: <span className="font-semibold text-foreground" />,
                            }}
                        />
                    )}
                </p>
                <div className="mt-4 flex justify-end gap-2">
                    <Button
                        type="button"
                        variant="outline"
                        onClick={() => onOpenChange(false)}
                    >
                        {t("materials.cancelMove")}
                    </Button>
                    <Button
                        type="button"
                        variant="destructive"
                        onClick={onConfirm}
                    >
                        {t("materials.menu.delete")}
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    )
}
