import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { Link, useNavigate, useParams } from "react-router-dom"
import { format } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { FaArrowLeft } from "react-icons/fa"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getRouteSchedule } from "@/app/router/routePaths"
import {
    getScheduleById,
    getSchedules,
    getSchedulesIsLoading,
    mapSchedulesToCalendarEvents,
    type ScheduleSlotData,
} from "@/entities/schedule"
import { useCanEdit } from "@/entities/user"
import { ScheduleDayDetails } from "./ScheduleDayDetails"
import { ScheduleDetailsCalendar } from "./ScheduleDetailsCalendar"
import { ScheduleDetailsHeader, type ScheduleTab } from "./ScheduleDetailsHeader"
import { ScheduleDetailsModals } from "./ScheduleDetailsModals"
import { useScheduleDetailsActions } from "./useScheduleDetailsActions"

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

    const selectedKey = format(selectedDate, "yyyy-MM-dd")
    const selectedEvents = events
        .filter((event) => event.date === selectedKey)
        .sort((a, b) => a.start - b.start)

    const {
        onDeleteSchedule,
        onDeleteSlot,
        onUnbindLesson,
        openBind,
    } = useScheduleDetailsActions({
        dispatch,
        navigate,
        t,
        schedule,
        setBindSlotId,
        setIsBindOpen,
    })

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
            <ScheduleDetailsHeader
                scheduleId={scheduleId}
                schedule={schedule}
                tab={tab}
                onTabChange={setTab}
                dateLocale={dateLocale}
                onEdit={() => setIsEditOpen(true)}
                onDeleteSchedule={onDeleteSchedule}
            />

            <div className="grid gap-4 lg:grid-cols-[minmax(0,1fr)_300px]">
                <ScheduleDetailsCalendar
                    isLoading={isLoading}
                    schedule={schedule}
                    tab={tab}
                    events={events}
                    selectedDate={selectedDate}
                    onSelectDate={setSelectedDate}
                />

                <ScheduleDayDetails
                    selectedDate={selectedDate}
                    dateLocale={dateLocale}
                    selectedEvents={selectedEvents}
                    schedule={schedule}
                    onAddSlots={() => setIsAddSlotsOpen(true)}
                    onEditSlot={setEditingSlot}
                    onDeleteSlot={onDeleteSlot}
                    onUnbindLesson={onUnbindLesson}
                    onBindSlot={openBind}
                />
            </div>

            {schedule && (
                <ScheduleDetailsModals
                    schedule={schedule}
                    isEditOpen={isEditOpen}
                    setIsEditOpen={setIsEditOpen}
                    isAddSlotsOpen={isAddSlotsOpen}
                    setIsAddSlotsOpen={setIsAddSlotsOpen}
                    selectedDate={selectedDate}
                    editingSlot={editingSlot}
                    setEditingSlot={setEditingSlot}
                    isBindOpen={isBindOpen}
                    setIsBindOpen={setIsBindOpen}
                    freeSlots={freeSlots}
                    bindSlotId={bindSlotId}
                />
            )}
        </div>
    )
}
