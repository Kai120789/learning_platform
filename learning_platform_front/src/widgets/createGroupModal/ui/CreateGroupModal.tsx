import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { createGroup } from "@/entities/group"
import { getSubjects } from "@/entities/subject"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"
import { Textarea } from "@/shared/ui/Textarea"
import { useState } from "react"
import { useTranslation } from "react-i18next"

type CreateGroupModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
}

export function CreateGroupModal({
    isOpen,
    setIsOpen
}: CreateGroupModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()

    const subjects = useAppSelector(getSubjects)

    const [title, setTitle] = useState<string>("")
    const [description, setDescription] = useState<string>("")
    const [subjectId, setSubjectId] = useState<number>(1)

    const onSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault();
        const response = await dispatch(createGroup({
            title: title,
            description: description,
            subject_id: subjectId,
        }))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("groups.createSuccess"),
                type: "success"
            }))
            setIsOpen(false)
            setTitle("")
            setDescription("")
            setSubjectId(1)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("groups.createError"),
                type: "error"
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle className="text-xl text-left line-clamp-2 pr-10">
                        {t("groups.createTitle")}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />
                <form className="p-6 md:p-8" onSubmit={onSubmit}>
                    <FieldGroup>
                        <Field>
                            <FieldLabel>{t("groups.name")}</FieldLabel>
                            <Input
                                required
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                            />
                        </Field>
                        <Field>
                            <FieldLabel>{t("groups.description")}</FieldLabel>
                            <Textarea
                                required
                                value={description}
                                onChange={(e) => setDescription(e.target.value)}
                                className="w-full break-words min-h-50"
                            />
                        </Field>
                        <Field>
                            <FieldLabel>{t("groups.subject")}</FieldLabel>
                            <select
                                className="border border-input rounded-lg p-2"
                                value={subjectId}
                                onChange={(e) => setSubjectId(Number(e.target.value))}
                            >
                                {subjects && subjects.map((subject) => (
                                    <option key={subject.id} value={subject.id}>
                                        {subject.title + " " + subject.type}
                                    </option>
                                ))}
                            </select>
                        </Field>
                        <Field>
                            <Button size="lg" type="submit">{t("common.create")}</Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
