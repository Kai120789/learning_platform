import { useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import {
    bindLessonToScheduleSlot,
    getSlotTimeLabel,
    type ScheduleSlotData,
} from "@/entities/schedule"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { NativeSelect } from "@/shared/ui/NativeSelect"
import { Separator } from "@/shared/ui/Separator"

type BindLessonModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    freeSlots: ScheduleSlotData[]
    preselectedSlotId?: number | null
}

export function BindLessonModal({
    isOpen,
    setIsOpen,
    freeSlots,
    preselectedSlotId = null,
}: BindLessonModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const [slotId, setSlotId] = useState<number | "">("")
    const [lessonId, setLessonId] = useState("")
    const [isSubmitting, setIsSubmitting] = useState(false)

    const sortedSlots = useMemo(
        () => [...freeSlots].sort((a, b) => a.startTime.localeCompare(b.startTime)),
        [freeSlots],
    )

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (open) {
            const preferred = preselectedSlotId
                && sortedSlots.some((slot) => slot.id === preselectedSlotId)
                ? preselectedSlotId
                : sortedSlots[0]?.id
            setSlotId(preferred ?? "")
            setLessonId("")
        }
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const parsedLessonId = Number(lessonId)
        if (!slotId || !Number.isFinite(parsedLessonId) || parsedLessonId <= 0) {
            dispatch(notificationActions.addNotification({
                message: t("bindLesson.invalid"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        const response = await dispatch(bindLessonToScheduleSlot({
            scheduleSlotID: Number(slotId),
            lessonID: parsedLessonId,
        }))
        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("bindLesson.success"),
                type: "success",
            }))
            setIsOpen(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("bindLesson.error"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("bindLesson.title")}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />
                <form className="w-full" onSubmit={onSubmit}>
                    <FieldGroup className="w-full gap-3">
                        <Field className="w-full">
                            <FieldLabel>{t("bindLesson.slot")}</FieldLabel>
                            <NativeSelect
                                required
                                value={slotId}
                                disabled={sortedSlots.length === 0}
                                onChange={(e) => setSlotId(Number(e.target.value))}
                            >
                                {sortedSlots.length === 0 ? (
                                    <option value="">{t("bindLesson.noFreeSlots")}</option>
                                ) : (
                                    sortedSlots.map((slot) => (
                                        <option key={slot.id} value={slot.id}>
                                            {getSlotTimeLabel(slot)}
                                        </option>
                                    ))
                                )}
                            </NativeSelect>
                        </Field>
                        <Field className="w-full">
                            <FieldLabel>{t("bindLesson.lessonId")}</FieldLabel>
                            <Input
                                className="w-full"
                                required
                                type="number"
                                min={1}
                                value={lessonId}
                                onChange={(e) => setLessonId(e.target.value)}
                                placeholder="1"
                            />
                        </Field>
                        <Field className="pt-3">
                            <Button
                                type="submit"
                                className="w-full"
                                disabled={isSubmitting || sortedSlots.length === 0}
                            >
                                {isSubmitting ? t("common.loading") : t("bindLesson.submit")}
                            </Button>
                        </Field>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
