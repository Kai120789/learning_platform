import { useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { createSchedule } from "@/entities/schedule"
import { getUserFullData } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type CreateScheduleModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    defaultDate?: Date
}

function toDateInputValue(date: Date) {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, "0")
    const day = String(date.getDate()).padStart(2, "0")
    return `${year}-${month}-${day}`
}

function addOneMonth(date: Date) {
    const next = new Date(date)
    next.setMonth(next.getMonth() + 1)
    return next
}

function periodStartIso(dateValue: string) {
    return new Date(`${dateValue}T00:00:00`).toISOString()
}

function periodEndIso(dateValue: string) {
    return new Date(`${dateValue}T23:59:59`).toISOString()
}

export function CreateScheduleModal({
    isOpen,
    setIsOpen,
    defaultDate = new Date(),
}: CreateScheduleModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const userData = useAppSelector(getUserFullData)

    const [title, setTitle] = useState("")
    const [startDate, setStartDate] = useState(toDateInputValue(defaultDate))
    const [endDate, setEndDate] = useState(toDateInputValue(addOneMonth(defaultDate)))
    const [isSubmitting, setIsSubmitting] = useState(false)

    const resetForm = (nextDate = defaultDate) => {
        setTitle("")
        setStartDate(toDateInputValue(nextDate))
        setEndDate(toDateInputValue(addOneMonth(nextDate)))
    }

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (open) {
            resetForm(defaultDate)
        }
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (!userData?.user.userID) {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.error"),
                type: "error",
            }))
            return
        }

        if (!title.trim()) {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.error"),
                type: "error",
            }))
            return
        }

        if (endDate < startDate) {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.invalidPeriod"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        const response = await dispatch(createSchedule({
            tutor_id: userData.user.userID,
            title: title.trim(),
            start_time: periodStartIso(startDate),
            end_time: periodEndIso(endDate),
        }))
        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.success"),
                type: "success",
            }))
            setIsOpen(false)
            resetForm(defaultDate)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.error"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("createSchedule.title")}
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
                                        if (endDate < nextDate) {
                                            setEndDate(toDateInputValue(addOneMonth(new Date(`${nextDate}T12:00:00`))))
                                        }
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
                                {isSubmitting ? t("common.loading") : t("common.create")}
                            </Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
