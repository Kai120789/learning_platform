import { useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { updateScheduleSlot, type ScheduleSlotData } from "@/entities/schedule"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type EditScheduleSlotModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    slot: ScheduleSlotData | null
}

type SlotForm = {
    slotId: number
    date: string
    time: string
    duration: number
}

function toSlotForm(slot: ScheduleSlotData): SlotForm {
    const date = new Date(slot.startTime)
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, "0")
    const day = String(date.getDate()).padStart(2, "0")
    const hours = String(date.getHours()).padStart(2, "0")
    const minutes = String(date.getMinutes()).padStart(2, "0")

    return {
        slotId: slot.id,
        date: `${year}-${month}-${day}`,
        time: `${hours}:${minutes}`,
        duration: slot.duration && slot.duration > 0 ? slot.duration : 60,
    }
}

function toIsoFromDateAndTime(dateValue: string, time: string) {
    return new Date(`${dateValue}T${time}`).toISOString()
}

export function EditScheduleSlotModal({
    isOpen,
    setIsOpen,
    slot,
}: EditScheduleSlotModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const [draft, setDraft] = useState<SlotForm | null>(null)
    const [isSubmitting, setIsSubmitting] = useState(false)

    const form = slot
        ? (draft?.slotId === slot.id ? draft : toSlotForm(slot))
        : null

    const patchForm = (patch: Partial<SlotForm>) => {
        if (!form) return
        setDraft({ ...form, ...patch })
    }

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (!open) setDraft(null)
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!slot || !form || !form.date || !form.time || form.duration < 15) {
            dispatch(notificationActions.addNotification({
                message: t("editScheduleSlot.error"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        const response = await dispatch(updateScheduleSlot({
            scheduleSlotID: slot.id,
            request: {
                start_time: toIsoFromDateAndTime(form.date, form.time),
                duration: form.duration,
            },
        }))
        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("editScheduleSlot.success"),
                type: "success",
            }))
            setDraft(null)
            setIsOpen(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("editScheduleSlot.error"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("editScheduleSlot.title")}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />
                <form className="w-full" onSubmit={onSubmit}>
                    <FieldGroup className="w-full gap-3">
                        <Field className="w-full">
                            <FieldLabel>{t("createSchedule.date")}</FieldLabel>
                            <Input
                                className="w-full"
                                type="date"
                                value={form?.date ?? ""}
                                disabled
                                readOnly
                            />
                        </Field>
                        <Field className="w-full">
                            <FieldLabel>{t("editScheduleSlot.startTime")}</FieldLabel>
                            <Input
                                className="w-full"
                                required
                                type="time"
                                value={form?.time ?? ""}
                                onChange={(e) => patchForm({ time: e.target.value })}
                            />
                        </Field>
                        <Field className="w-full">
                            <FieldLabel>{t("createSchedule.duration")}</FieldLabel>
                            <Input
                                className="w-full"
                                required
                                type="number"
                                min={15}
                                step={15}
                                value={form?.duration ?? 60}
                                onChange={(e) => patchForm({ duration: Number(e.target.value) })}
                            />
                        </Field>
                        <Field className="pt-3">
                            <Button type="submit" className="w-full" disabled={isSubmitting || !slot}>
                                {isSubmitting ? t("common.loading") : t("common.save")}
                            </Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
