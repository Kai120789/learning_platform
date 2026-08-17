import { useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { createFolder } from "@/entities/material"
import { getUserFullData } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type CreateFolderDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    parentFolderId: number | null
}

export function CreateFolderDialog({
    open,
    onOpenChange,
    parentFolderId,
}: CreateFolderDialogProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const userData = useAppSelector(getUserFullData)
    const [title, setTitle] = useState("")

    const onSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault()
        const tutorId = userData?.user.userID
        if (!tutorId || !title.trim()) return

        const response = await dispatch(createFolder({
            title: title.trim(),
            parent_folder_id: parentFolderId,
            tutor_id: tutorId,
        }))

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("materials.createFolderSuccess"),
                type: "success",
            }))
            setTitle("")
            onOpenChange(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("materials.createFolderError"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left">
                        {t("materials.createFolderTitle")}
                    </DialogTitle>
                </DialogHeader>
                <Separator className="my-1" />
                <form className="w-full" onSubmit={onSubmit}>
                    <FieldGroup className="w-full gap-3">
                        <Field className="w-full">
                            <FieldLabel>{t("materials.folderName")}</FieldLabel>
                            <Input
                                className="w-full"
                                required
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                autoFocus
                            />
                        </Field>
                        <Button type="submit" className="w-full">
                            {t("materials.createFolder")}
                        </Button>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
