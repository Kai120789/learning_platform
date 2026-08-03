import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    bindLessonToScheduleSlot,
    getSlotTimeLabel,
    type ScheduleSlotData,
} from "@/entities/schedule"
import {
    getLessonLabel,
    getLessons,
    getLessonsByTutorId,
} from "@/entities/lesson"
import { getUserFullData } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
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
    const userData = useAppSelector(getUserFullData)
    const lessons = useAppSelector(getLessons)

    const [slotId, setSlotId] = useState<number | "">("")
    const [lessonId, setLessonId] = useState<number | "">("")
    const [isSubmitting, setIsSubmitting] = useState(false)

    const sortedSlots = useMemo(
        () => [...freeSlots].sort((a, b) => a.startTime.localeCompare(b.startTime)),
        [freeSlots],
    )

    const bindableLessons = useMemo(
        () => [...(lessons ?? [])]
            .filter((lesson) => lesson.status === "SCHEDULED" || lesson.status === "IN_PROCESS")
            .sort((a, b) => a.startTime.localeCompare(b.startTime)),
        [lessons],
    )

    useEffect(() => {
        if (!isOpen || !userData?.user.userID) return
        dispatch(getLessonsByTutorId(userData.user.userID))
    }, [dispatch, isOpen, userData?.user.userID])

    const resolvedLessonId = lessonId !== ""
        ? lessonId
        : (bindableLessons[0]?.id ?? "")

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (open) {
            const preferredSlot = preselectedSlotId
                && sortedSlots.some((slot) => slot.id === preselectedSlotId)
                ? preselectedSlotId
                : sortedSlots[0]?.id
            setSlotId(preferredSlot ?? "")
            setLessonId("")
        }
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (!slotId || !resolvedLessonId) {
            dispatch(notificationActions.addNotification({
                message: t("bindLesson.invalid"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        const response = await dispatch(bindLessonToScheduleSlot({
            scheduleSlotID: Number(slotId),
            lessonID: Number(resolvedLessonId),
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
                            <FieldLabel>{t("bindLesson.lesson")}</FieldLabel>
                            <NativeSelect
                                required
                                value={resolvedLessonId}
                                disabled={bindableLessons.length === 0}
                                onChange={(e) => setLessonId(Number(e.target.value))}
                            >
                                {bindableLessons.length === 0 ? (
                                    <option value="">{t("bindLesson.noLessons")}</option>
                                ) : (
                                    bindableLessons.map((lesson) => (
                                        <option key={lesson.id} value={lesson.id}>
                                            {getLessonLabel(lesson)}
                                        </option>
                                    ))
                                )}
                            </NativeSelect>
                        </Field>
                        <Field className="pt-3">
                            <Button
                                type="submit"
                                className="w-full"
                                disabled={
                                    isSubmitting
                                    || sortedSlots.length === 0
                                    || bindableLessons.length === 0
                                }
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
