import { useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { updateSchedule, type ScheduleData } from "@/entities/schedule"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type EditScheduleModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    schedule: ScheduleData
}

function toDateInputValue(iso: string) {
    const date = new Date(iso)
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, "0")
    const day = String(date.getDate()).padStart(2, "0")
    return `${year}-${month}-${day}`
}

function periodStartIso(dateValue: string) {
    return new Date(`${dateValue}T00:00:00`).toISOString()
}

function periodEndIso(dateValue: string) {
    return new Date(`${dateValue}T23:59:59`).toISOString()
}

export function EditScheduleModal({
    isOpen,
    setIsOpen,
    schedule,
}: EditScheduleModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const [title, setTitle] = useState(schedule.title ?? "")
    const [startDate, setStartDate] = useState(toDateInputValue(schedule.startTime))
    const [endDate, setEndDate] = useState(toDateInputValue(schedule.endTime))
    const [isSubmitting, setIsSubmitting] = useState(false)

    const syncFromSchedule = () => {
        setTitle(schedule.title ?? "")
        setStartDate(toDateInputValue(schedule.startTime))
        setEndDate(toDateInputValue(schedule.endTime))
    }

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (open) syncFromSchedule()
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (endDate < startDate) {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.invalidPeriod"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        const response = await dispatch(updateSchedule({
            scheduleID: schedule.id,
            request: {
                id: schedule.id,
                title: title.trim(),
                start_time: periodStartIso(startDate),
                end_time: periodEndIso(endDate),
            },
        }))
        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("editSchedule.success"),
                type: "success",
            }))
            setIsOpen(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("editSchedule.error"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("editSchedule.title")}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />
                <form className="w-full" onSubmit={onSubmit}>
                    <FieldGroup className="w-full gap-3">
                        <Field className="w-full">
                            <FieldLabel>{t("createSchedule.name")}</FieldLabel>
                            <Input
                                className="w-full"
                                required
                                type="text"
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                            />
                        </Field>

                        <div className="grid gap-3 sm:grid-cols-2">
                            <Field className="w-full">
                                <FieldLabel>{t("createSchedule.startDate")}</FieldLabel>
                                <Input
                                    className="w-full"
                                    required
                                    type="date"
                                    value={startDate}
                                    onChange={(e) => {
                                        const nextDate = e.target.value
                                        setStartDate(nextDate)
                                        if (endDate < nextDate) setEndDate(nextDate)
                                    }}
                                />
                            </Field>
                            <Field className="w-full">
                                <FieldLabel>{t("createSchedule.endDate")}</FieldLabel>
                                <Input
                                    className="w-full"
                                    required
                                    type="date"
                                    min={startDate}
                                    value={endDate}
                                    onChange={(e) => setEndDate(e.target.value)}
                                />
                            </Field>
                        </div>
                        <Field className="pt-3">
                            <Button type="submit" className="w-full" disabled={isSubmitting}>
                                {isSubmitting ? t("common.loading") : t("common.save")}
                            </Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
