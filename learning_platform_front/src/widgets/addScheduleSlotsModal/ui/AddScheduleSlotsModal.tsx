import { useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { FaPlus, FaTrash } from "react-icons/fa"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { createScheduleSlot, type ScheduleData } from "@/entities/schedule"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"

type SlotDraft = {
    key: string
    time: string
    duration: number
}

type AddScheduleSlotsModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    schedule: ScheduleData
    defaultDate?: Date
}

function toDateInputValue(date: Date) {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, "0")
    const day = String(date.getDate()).padStart(2, "0")
    return `${year}-${month}-${day}`
}

function createEmptySlot(time = "09:00"): SlotDraft {
    return {
        key: `${Date.now()}-${Math.random().toString(36).slice(2, 7)}`,
        time,
        duration: 60,
    }
}

function toIsoFromDateAndTime(dateValue: string, time: string) {
    return new Date(`${dateValue}T${time}`).toISOString()
}

function nextTimeAfter(time: string, durationMinutes: number) {
    const [hours, minutes] = time.split(":").map(Number)
    const total = hours * 60 + minutes + Math.max(15, durationMinutes)
    const nextHours = Math.floor(total / 60) % 24
    const nextMinutes = total % 60
    return `${String(nextHours).padStart(2, "0")}:${String(nextMinutes).padStart(2, "0")}`
}

export function AddScheduleSlotsModal({
    isOpen,
    setIsOpen,
    schedule,
    defaultDate,
}: AddScheduleSlotsModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const lockedDate = useMemo(
        () => toDateInputValue(defaultDate ?? new Date(schedule.startTime)),
        [defaultDate, schedule.startTime],
    )
    const [slots, setSlots] = useState<SlotDraft[]>([createEmptySlot()])
    const [isSubmitting, setIsSubmitting] = useState(false)

    const resetForm = () => {
        setSlots([createEmptySlot()])
    }

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (open) resetForm()
    }

    const updateSlot = (key: string, patch: Partial<SlotDraft>) => {
        setSlots((prev) => prev.map((slot) => (
            slot.key === key ? { ...slot, ...patch } : slot
        )))
    }

    const removeSlot = (key: string) => {
        setSlots((prev) => (prev.length <= 1 ? prev : prev.filter((slot) => slot.key !== key)))
    }

    const addSlot = () => {
        const last = slots[slots.length - 1]
        const nextTime = last
            ? nextTimeAfter(last.time, last.duration)
            : "09:00"
        setSlots((prev) => [...prev, createEmptySlot(nextTime)])
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (slots.length === 0 || slots.some((slot) => !slot.time || slot.duration < 15)) {
            dispatch(notificationActions.addNotification({
                message: t("createSchedule.invalidSlots"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        let failed = false
        for (const slot of slots) {
            const response = await dispatch(createScheduleSlot({
                schedule_id: schedule.id,
                start_time: toIsoFromDateAndTime(lockedDate, slot.time),
                duration: slot.duration,
            }))
            if (response.meta.requestStatus !== "fulfilled") {
                failed = true
                break
            }
        }
        setIsSubmitting(false)

        if (failed) {
            dispatch(notificationActions.addNotification({
                message: t("addScheduleSlots.error"),
                type: "error",
            }))
            return
        }

        dispatch(notificationActions.addNotification({
            message: t("addScheduleSlots.success"),
            type: "success",
        }))
        setIsOpen(false)
        resetForm()
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-lg p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("addScheduleSlots.title")}
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
                                value={lockedDate}
                                disabled
                                readOnly
                            />
                        </Field>

                        <div className="space-y-2">
                            <div className="flex items-center justify-between gap-2">
                                <FieldLabel>{t("createSchedule.slots")}</FieldLabel>
                                <Button
                                    type="button"
                                    size="xs"
                                    variant="outline"
                                    onClick={addSlot}
                                >
                                    <FaPlus className="size-3" />
                                    {t("createSchedule.addSlot")}
                                </Button>
                            </div>

                            <div className="space-y-2 max-h-64 overflow-y-auto">
                                {slots.map((slot, index) => (
                                    <div
                                        key={slot.key}
                                        className="grid grid-cols-[1fr_88px_auto] gap-2 items-end rounded-lg border p-2.5"
                                    >
                                        <Field className="w-full">
                                            <FieldLabel className="text-xs text-muted-foreground">
                                                {t("createSchedule.slotStart", { index: index + 1 })}
                                            </FieldLabel>
                                            <Input
                                                className="w-full"
                                                required
                                                type="time"
                                                value={slot.time}
                                                onChange={(e) => updateSlot(slot.key, {
                                                    time: e.target.value,
                                                })}
                                            />
                                        </Field>
                                        <Field className="w-full">
                                            <FieldLabel className="text-xs text-muted-foreground">
                                                {t("createSchedule.duration")}
                                            </FieldLabel>
                                            <Input
                                                className="w-full"
                                                required
                                                type="number"
                                                min={15}
                                                step={15}
                                                value={slot.duration}
                                                onChange={(e) => updateSlot(slot.key, {
                                                    duration: Number(e.target.value),
                                                })}
                                            />
                                        </Field>
                                        <Button
                                            type="button"
                                            size="icon-sm"
                                            variant="outline"
                                            disabled={slots.length <= 1}
                                            onClick={() => removeSlot(slot.key)}
                                            aria-label={t("createSchedule.removeSlot")}
                                        >
                                            <FaTrash className="size-3" />
                                        </Button>
                                    </div>
                                ))}
                            </div>
                        </div>

                        <Field className="pt-3">
                            <Button type="submit" className="w-full" disabled={isSubmitting}>
                                {isSubmitting ? t("common.loading") : t("common.add")}
                            </Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
