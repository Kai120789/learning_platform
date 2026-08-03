import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { useNavigate } from "react-router-dom"
import { format } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { FaPlus, FaTrash } from "react-icons/fa"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getRouteScheduleDetails } from "@/app/router/routePaths"
import {
    deleteSchedule,
    getSchedules,
    getSchedulesIsLoading,
    getSchedulesByTutorId,
    scheduleSlotStatusClass,
    type ScheduleData,
} from "@/entities/schedule"
import { getUserFullData, useCanEdit } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Card, CardContent } from "@/shared/ui/Card"
import { Label } from "@/shared/ui/Label"
import { cn } from "@/shared/lib/utils"
import { CreateScheduleModal } from "@/widgets/createScheduleModal"

export default function SchedulePage() {
    const { t, i18n } = useTranslation()
    const navigate = useNavigate()
    const dispatch = useAppDispatch()
    const canEdit = useCanEdit()
    const userData = useAppSelector(getUserFullData)
    const schedules = useAppSelector(getSchedules)
    const isLoading = useAppSelector(getSchedulesIsLoading)
    const [isOpen, setIsOpen] = useState(false)
    const dateLocale = i18n.language === "ru" ? ru : enUS

    useEffect(() => {
        if (!canEdit || !userData?.user.userID) return
        dispatch(getSchedulesByTutorId(userData.user.userID))
    }, [canEdit, dispatch, userData?.user.userID])

    const sortedSchedules = useMemo(
        () => [...(schedules ?? [])].sort((a, b) => a.startTime.localeCompare(b.startTime)),
        [schedules],
    )

    const onDeleteSchedule = async (
        e: React.MouseEvent,
        schedule: ScheduleData,
    ) => {
        e.stopPropagation()
        const response = await dispatch(deleteSchedule(schedule.id))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteError"),
                type: "error",
            }))
        }
    }

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <div className="flex justify-between items-center gap-4">
                    <Label className="text-xl lg:text-2xl">
                        {t("schedule.title")}
                    </Label>
                    {canEdit && (
                        <Button onClick={() => setIsOpen(true)} className="rounded-full">
                            <FaPlus className="size-3" />
                            {t("schedule.createSchedule")}
                        </Button>
                    )}
                </div>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("schedule.subtitle")}
                </Label>
            </div>

            {!canEdit && (
                <p className="text-sm text-muted-foreground">
                    {t("schedule.studentHint")}
                </p>
            )}

            {canEdit && schedules === null && isLoading && (
                <p className="text-sm text-muted-foreground">{t("common.loading")}</p>
            )}

            {canEdit && schedules !== null && schedules.length === 0 && (
                <p className="text-sm text-muted-foreground">{t("schedule.noSchedules")}</p>
            )}

            {canEdit && sortedSchedules.length > 0 && (
                <div className="flex flex-col gap-2">
                    {sortedSchedules.map((schedule) => {
                        const freeCount = schedule.slots.filter((slot) => slot.status === "FREE").length
                        const bookedCount = schedule.slots.filter((slot) => slot.status === "BOOKED").length

                        return (
                            <Card
                                key={schedule.id}
                                size="sm"
                                className="cursor-pointer transition-all hover:shadow-md"
                                onClick={() => navigate(getRouteScheduleDetails(schedule.id))}
                            >
                                <CardContent className="flex items-center justify-between gap-4 py-3">
                                    <div className="min-w-0 space-y-1">
                                        <div className="text-sm font-medium truncate">
                                            {schedule.title?.trim()
                                                || t("schedule.scheduleItem", { id: schedule.id })}
                                        </div>
                                        <div className="text-xs text-muted-foreground">
                                            {format(new Date(schedule.startTime), "d MMM yyyy", { locale: dateLocale })}
                                            {" – "}
                                            {format(new Date(schedule.endTime), "d MMM yyyy", { locale: dateLocale })}
                                        </div>
                                    </div>

                                    <div className="flex items-center gap-2">
                                        <span className="text-xs text-muted-foreground">
                                            {t("schedule.slotsCount", { count: schedule.slots.length })}
                                        </span>
                                        <Badge
                                            variant="outline"
                                            className={cn("text-[10px]", scheduleSlotStatusClass("FREE"))}
                                        >
                                            {t("schedule.freeCount", { count: freeCount })}
                                        </Badge>
                                        <Badge
                                            variant="outline"
                                            className={cn("text-[10px]", scheduleSlotStatusClass("BOOKED"))}
                                        >
                                            {t("schedule.bookedCount", { count: bookedCount })}
                                        </Badge>
                                        <Button
                                            type="button"
                                            size="icon-sm"
                                            variant="outline"
                                            onClick={(e) => onDeleteSchedule(e, schedule)}
                                            aria-label={t("schedule.deleteSchedule")}
                                        >
                                            <FaTrash className="size-3.5" />
                                        </Button>
                                    </div>
                                </CardContent>
                            </Card>
                        )
                    })}
                </div>
            )}

            {canEdit && (
                <CreateScheduleModal
                    isOpen={isOpen}
                    setIsOpen={setIsOpen}
                />
            )}
        </div>
    )
}
