import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type MaterialsRenameDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    initialTitle: string
    onSubmit: (title: string) => Promise<void> | void
}

export function MaterialsRenameDialog({
    open,
    onOpenChange,
    initialTitle,
    onSubmit,
}: MaterialsRenameDialogProps) {
    const { t } = useTranslation()
    const [title, setTitle] = useState(initialTitle)

    useEffect(() => {
        if (open) setTitle(initialTitle)
    }, [open, initialTitle])

    const handleSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault()
        const next = title.trim()
        if (!next) return
        await onSubmit(next)
    }

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left">
                        {t("materials.renameTitle")}
                    </DialogTitle>
                </DialogHeader>
                <Separator className="my-1" />
                <form className="w-full" onSubmit={handleSubmit}>
                    <FieldGroup className="w-full gap-3">
                        <Field className="w-full">
                            <FieldLabel>{t("materials.name")}</FieldLabel>
                            <Input
                                className="w-full"
                                required
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                autoFocus
                            />
                        </Field>
                        <Button type="submit" className="w-full">
                            {t("materials.rename")}
                        </Button>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
