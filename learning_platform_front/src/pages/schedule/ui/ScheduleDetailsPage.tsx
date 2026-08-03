import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { Link, useNavigate, useParams } from "react-router-dom"
import { format } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { FaArrowLeft, FaLink, FaPen, FaPlus, FaTrash, FaUnlink } from "react-icons/fa"
import { MonthlyCalendar, WeeklyCalendar } from "@/widgets/calendars"
import { AddScheduleSlotsModal } from "@/widgets/addScheduleSlotsModal"
import { BindLessonModal } from "@/widgets/bindLessonModal"
import { EditScheduleModal } from "@/widgets/editScheduleModal"
import { EditScheduleSlotModal } from "@/widgets/editScheduleSlotModal"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { cn } from "@/shared/lib/utils"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getRouteSchedule } from "@/app/router/routePaths"
import {
    deleteLessonFromScheduleSlot,
    deleteSchedule,
    deleteScheduleSlot,
    getScheduleById,
    getSchedules,
    getSchedulesIsLoading,
    mapSchedulesToCalendarEvents,
    scheduleSlotStatusClass,
    type ScheduleSlotData,
} from "@/entities/schedule"
import { useCanEdit } from "@/entities/user"
import { notificationActions } from "@/features/notifications"

type ScheduleTab = "week" | "month"

export default function ScheduleDetailsPage() {
    const { id } = useParams<{ id: string }>()
    const scheduleId = Number(id)
    const { t, i18n } = useTranslation()
    const navigate = useNavigate()
    const dispatch = useAppDispatch()
    const canEdit = useCanEdit()
    const schedules = useAppSelector(getSchedules)
    const isLoading = useAppSelector(getSchedulesIsLoading)

    const [tab, setTab] = useState<ScheduleTab>("month")
    const [isEditOpen, setIsEditOpen] = useState(false)
    const [isAddSlotsOpen, setIsAddSlotsOpen] = useState(false)
    const [isBindOpen, setIsBindOpen] = useState(false)
    const [bindSlotId, setBindSlotId] = useState<number | null>(null)
    const [editingSlot, setEditingSlot] = useState<ScheduleSlotData | null>(null)
    const [pickedDate, setPickedDate] = useState<{ scheduleId: number; date: Date } | null>(null)
    const dateLocale = i18n.language === "ru" ? ru : enUS

    const schedule = useMemo(
        () => schedules?.find((item) => item.id === scheduleId) ?? null,
        [schedules, scheduleId],
    )

    const freeSlots = useMemo(
        () => schedule?.slots.filter((slot) => slot.status === "FREE") ?? [],
        [schedule],
    )

    const scheduleStartDate = useMemo(
        () => (schedule ? new Date(schedule.startTime) : new Date()),
        [schedule?.startTime],
    )

    const selectedDate = pickedDate?.scheduleId === scheduleId
        ? pickedDate.date
        : scheduleStartDate

    const setSelectedDate = (date: Date) => {
        setPickedDate({ scheduleId, date })
    }

    useEffect(() => {
        if (!canEdit || !Number.isFinite(scheduleId) || scheduleId <= 0) return
        dispatch(getScheduleById(scheduleId))
    }, [canEdit, dispatch, scheduleId])

    const events = useMemo(
        () => mapSchedulesToCalendarEvents(schedule ? [schedule] : [], {
            freeSlot: t("schedule.freeSlot"),
            bookedLesson: (lessonId) => t("schedule.bookedLesson", { id: lessonId }),
        }),
        [schedule, t],
    )

    const tabs: { id: ScheduleTab; label: string }[] = [
        { id: "week", label: t("schedule.week") },
        { id: "month", label: t("schedule.month") },
    ]

    const selectedKey = format(selectedDate, "yyyy-MM-dd")
    const selectedEvents = events
        .filter((event) => event.date === selectedKey)
        .sort((a, b) => a.start - b.start)

    const onDeleteSchedule = async () => {
        if (!schedule) return
        const response = await dispatch(deleteSchedule(schedule.id))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteSuccess"),
                type: "success",
            }))
            navigate(getRouteSchedule())
        } else {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteError"),
                type: "error",
            }))
        }
    }

    const onDeleteSlot = async (slot: ScheduleSlotData) => {
        const response = await dispatch(deleteScheduleSlot(slot.id))

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteSlotSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("schedule.deleteSlotError"),
                type: "error",
            }))
        }
    }

    const onUnbindLesson = async (slot: ScheduleSlotData) => {
        const response = await dispatch(deleteLessonFromScheduleSlot(slot.id))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("schedule.unbindSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("schedule.unbindError"),
                type: "error",
            }))
        }
    }

    const openBind = (slotId: number) => {
        setBindSlotId(slotId)
        setIsBindOpen(true)
    }

    if (!canEdit) {
        return (
            <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-4">
                <Link to={getRouteSchedule()} className="inline-flex items-center gap-2 text-sm text-muted-foreground hover:text-foreground">
                    <FaArrowLeft className="size-3" />
                    {t("schedule.backToList")}
                </Link>
                <p className="text-sm text-muted-foreground">{t("schedule.studentHint")}</p>
            </div>
        )
    }

    if (!Number.isFinite(scheduleId) || scheduleId <= 0) {
        return (
            <div className="py-8 lg:py-10 px-6 lg:px-20">
                <p className="text-sm text-muted-foreground">{t("schedule.notFound")}</p>
            </div>
        )
    }

    return (
        <div className="flex flex-col py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
                <div className="space-y-2">
                    <Link
                        to={getRouteSchedule()}
                        className="inline-flex items-center gap-2 text-sm text-muted-foreground hover:text-foreground"
                    >
                        <FaArrowLeft className="size-3" />
                        {t("schedule.backToList")}
                    </Link>
                    <div className="space-y-1">
                        <Label className="text-xl lg:text-2xl">
                            {schedule?.title?.trim()
                                || t("schedule.scheduleItem", { id: scheduleId })}
                        </Label>
                        {schedule && (
                            <Label className="text-sm lg:text-base font-normal text-primary/50">
                                {format(new Date(schedule.startTime), "d MMM yyyy", { locale: dateLocale })}
                                {" – "}
                                {format(new Date(schedule.endTime), "d MMM yyyy", { locale: dateLocale })}
                            </Label>
                        )}
                    </div>
                </div>
                <div className="flex flex-wrap items-center gap-2 sm:justify-end">
                    <div className="flex rounded-lg border border-border p-0.5 bg-secondary/60">
                        {tabs.map((item) => (
                            <button
                                key={item.id}
                                type="button"
                                onClick={() => setTab(item.id)}
                                className={cn(
                                    "rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                                    item.id === tab
                                        ? "bg-primary text-primary-foreground shadow-sm"
                                        : "text-secondary-foreground hover:bg-muted"
                                )}
                            >
                                {item.label}
                            </button>
                        ))}
                    </div>
                    {schedule && (
                        <Button
                            type="button"
                            size="sm"
                            variant="outline"
                            onClick={() => setIsEditOpen(true)}
                        >
                            <FaPen className="size-3" />
                            {t("schedule.editPeriod")}
                        </Button>
                    )}
                    <Button
                        type="button"
                        size="sm"
                        variant="outline"
                        onClick={onDeleteSchedule}
                    >
                        <FaTrash className="size-3" />
                        {t("schedule.deleteSchedule")}
                    </Button>
                </div>
            </div>

            <div className="grid gap-4 lg:grid-cols-[minmax(0,1fr)_300px]">
                <Card className="overflow-hidden">
                    <CardContent className="p-3 lg:p-4">
                        {isLoading && !schedule ? (
                            <p className="py-10 text-center text-sm text-muted-foreground">
                                {t("common.loading")}
                            </p>
                        ) : !schedule ? (
                            <p className="py-10 text-center text-sm text-muted-foreground">
                                {t("schedule.notFound")}
                            </p>
                        ) : tab === "month" ? (
                            <MonthlyCalendar
                                events={events}
                                selectedDate={selectedDate}
                                onSelectDate={setSelectedDate}
                                periodStart={schedule.startTime}
                                periodEnd={schedule.endTime}
                            />
                        ) : (
                            <WeeklyCalendar
                                events={events}
                                selectedDate={selectedDate}
                                onSelectDate={setSelectedDate}
                                periodStart={schedule.startTime}
                                periodEnd={schedule.endTime}
                            />
                        )}
                    </CardContent>
                </Card>

                <div className="space-y-4">
                    <Card>
                        <CardHeader className="pb-2">
                            <div className="flex items-start justify-between gap-2">
                                <div className="min-w-0 space-y-1">
                                    <CardTitle className="text-sm">
                                        {t("schedule.dayDetails")}
                                    </CardTitle>
                                    <p className="text-xs text-muted-foreground capitalize">
                                        {format(selectedDate, "EEEE, d MMMM", { locale: dateLocale })}
                                    </p>
                                </div>
                                {schedule && (
                                    <Button
                                        type="button"
                                        size="xs"
                                        variant="outline"
                                        onClick={() => setIsAddSlotsOpen(true)}
                                    >
                                        <FaPlus className="size-3" />
                                        {t("schedule.addSlot")}
                                    </Button>
                                )}
                            </div>
                        </CardHeader>
                        <CardContent className="space-y-2">
                            {selectedEvents.length === 0 ? (
                                <p className="text-sm text-muted-foreground">
                                    {t("schedule.noEvents")}
                                </p>
                            ) : (
                                selectedEvents.map((event) => {
                                    const slot = schedule?.slots.find((item) => item.id === event.id)
                                    const isFree = slot?.status === "FREE"

                                    return (
                                        <div
                                            key={event.id}
                                            role={isFree ? "button" : undefined}
                                            tabIndex={isFree ? 0 : undefined}
                                            onClick={() => {
                                                if (slot && isFree) openBind(slot.id)
                                            }}
                                            onKeyDown={(e) => {
                                                if (!slot || !isFree) return
                                                if (e.key === "Enter" || e.key === " ") {
                                                    e.preventDefault()
                                                    openBind(slot.id)
                                                }
                                            }}
                                            className={cn(
                                                "rounded-lg border p-2.5 space-y-2 transition-colors",
                                                isFree && "border-dashed border-primary/40 bg-primary/5 cursor-pointer hover:bg-primary/10 hover:border-primary/60",
                                            )}
                                        >
                                            <div className="flex items-start justify-between gap-2">
                                                <div className="min-w-0 space-y-1">
                                                    <div className="text-sm font-medium">{event.title}</div>
                                                    <Badge
                                                        variant="outline"
                                                        className={cn(
                                                            "text-[10px]",
                                                            event.status && scheduleSlotStatusClass(event.status),
                                                        )}
                                                    >
                                                        {t(`scheduleSlotStatus.${event.status}`)}
                                                    </Badge>
                                                </div>
                                                {slot && (
                                                    <div
                                                        className="flex items-center gap-1 shrink-0"
                                                        onClick={(e) => e.stopPropagation()}
                                                    >
                                                        <Button
                                                            type="button"
                                                            size="icon-xs"
                                                            variant="outline"
                                                            onClick={() => setEditingSlot(slot)}
                                                            aria-label={t("schedule.editSlot")}
                                                        >
                                                            <FaPen className="size-3" />
                                                        </Button>
                                                        {isFree ? (
                                                            <Button
                                                                type="button"
                                                                size="icon-xs"
                                                                variant="outline"
                                                                onClick={() => onDeleteSlot(slot)}
                                                                aria-label={t("schedule.deleteSlot")}
                                                            >
                                                                <FaTrash className="size-3" />
                                                            </Button>
                                                        ) : (
                                                            <Button
                                                                type="button"
                                                                size="icon-xs"
                                                                variant="outline"
                                                                onClick={() => onUnbindLesson(slot)}
                                                                aria-label={t("schedule.unbindLesson")}
                                                            >
                                                                <FaUnlink className="size-3" />
                                                            </Button>
                                                        )}
                                                    </div>
                                                )}
                                            </div>
                                            <div className="flex items-end justify-between gap-2">
                                                <span className="text-xs text-muted-foreground">
                                                    {event.timeLabel}
                                                </span>
                                                {slot && isFree && (
                                                    <Button
                                                        type="button"
                                                        size="icon-xs"
                                                        variant="ghost"
                                                        className="-mb-1 -mr-1 text-primary hover:bg-primary/10 hover:text-primary"
                                                        onClick={(e) => {
                                                            e.stopPropagation()
                                                            openBind(slot.id)
                                                        }}
                                                        aria-label={t("schedule.bindLesson")}
                                                    >
                                                        <FaLink className="size-3" />
                                                    </Button>
                                                )}
                                            </div>
                                        </div>
                                    )
                                })
                            )}
                        </CardContent>
                    </Card>
                </div>
            </div>

            {schedule && (
                <>
                    <EditScheduleModal
                        isOpen={isEditOpen}
                        setIsOpen={setIsEditOpen}
                        schedule={schedule}
                    />
                    <AddScheduleSlotsModal
                        isOpen={isAddSlotsOpen}
                        setIsOpen={setIsAddSlotsOpen}
                        schedule={schedule}
                        defaultDate={selectedDate}
                    />
                    <EditScheduleSlotModal
                        isOpen={Boolean(editingSlot)}
                        setIsOpen={(open) => {
                            if (!open) setEditingSlot(null)
                        }}
                        slot={editingSlot}
                    />
                    <BindLessonModal
                        isOpen={isBindOpen}
                        setIsOpen={setIsBindOpen}
                        freeSlots={freeSlots}
                        preselectedSlotId={bindSlotId}
                    />
                </>
            )}
        </div>
    )
}
