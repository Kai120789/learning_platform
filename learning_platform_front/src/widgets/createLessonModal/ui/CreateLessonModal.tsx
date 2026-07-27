import { useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { notificationActions } from "@/features/notifications"
import { mockGroupsForSelect } from "@/shared/mocks"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type CreateLessonModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
}

export function CreateLessonModal({
    isOpen,
    setIsOpen,
}: CreateLessonModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()

    const [meetLink, setMeetLink] = useState("")
    const [startTime, setStartTime] = useState("")
    const [duration, setDuration] = useState(60)
    const [groupId, setGroupId] = useState(mockGroupsForSelect[0]?.id ?? 1)

    const onSubmit = (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault()

        // Mock-only: wire to lessons API later
        if (!startTime || duration <= 0) {
            dispatch(notificationActions.addNotification({
                message: t("createLesson.error"),
                type: "error",
            }))
            return
        }

        dispatch(notificationActions.addNotification({
            message: t("createLesson.success"),
            type: "success",
        }))
        setIsOpen(false)
        setMeetLink("")
        setStartTime("")
        setDuration(60)
        setGroupId(mockGroupsForSelect[0]?.id ?? 1)
    }

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle className="text-xl text-left line-clamp-2 pr-10">
                        {t("createLesson.title")}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />
                <form className="p-6 md:p-8" onSubmit={onSubmit}>
                    <FieldGroup>
                        <Field>
                            <FieldLabel>{t("createLesson.meetLink")}</FieldLabel>
                            <Input
                                type="url"
                                value={meetLink}
                                onChange={(e) => setMeetLink(e.target.value)}
                                placeholder="https://"
                            />
                        </Field>
                        <Field>
                            <FieldLabel>{t("createLesson.startTime")}</FieldLabel>
                            <Input
                                required
                                type="datetime-local"
                                value={startTime}
                                onChange={(e) => setStartTime(e.target.value)}
                            />
                        </Field>
                        <Field>
                            <FieldLabel>{t("createLesson.duration")}</FieldLabel>
                            <Input
                                required
                                type="number"
                                min={15}
                                step={15}
                                value={duration}
                                onChange={(e) => setDuration(Number(e.target.value))}
                            />
                        </Field>
                        <Field>
                            <FieldLabel>{t("createLesson.group")}</FieldLabel>
                            <select
                                className="border border-input rounded-lg p-2"
                                value={groupId}
                                onChange={(e) => setGroupId(Number(e.target.value))}
                            >
                                {mockGroupsForSelect.map((group) => (
                                    <option key={group.id} value={group.id}>
                                        {group.title}
                                    </option>
                                ))}
                            </select>
                        </Field>
                        <Field>
                            <Button size="lg" type="submit">
                                {t("common.create")}
                            </Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
