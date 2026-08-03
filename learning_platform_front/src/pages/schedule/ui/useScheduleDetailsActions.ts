import { useCallback } from "react"
import type { TFunction } from "i18next"
import type { NavigateFunction } from "react-router-dom"
import type { AppDispatch } from "@/app/providers/storeProvider"
import { getRouteSchedule } from "@/app/router/routePaths"
import {
    deleteLessonFromScheduleSlot,
    deleteSchedule,
    deleteScheduleSlot,
    type ScheduleData,
    type ScheduleSlotData,
} from "@/entities/schedule"
import { notificationActions } from "@/features/notifications"

type UseScheduleDetailsActionsParams = {
    dispatch: AppDispatch
    navigate: NavigateFunction
    t: TFunction
    schedule: ScheduleData | null
    setBindSlotId: (slotId: number | null) => void
    setIsBindOpen: (open: boolean) => void
}

export function useScheduleDetailsActions({
    dispatch,
    navigate,
    t,
    schedule,
    setBindSlotId,
    setIsBindOpen,
}: UseScheduleDetailsActionsParams) {
    const onDeleteSchedule = useCallback(async () => {
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
    }, [dispatch, navigate, schedule, t])

    const onDeleteSlot = useCallback(async (slot: ScheduleSlotData) => {
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
    }, [dispatch, t])

    const onUnbindLesson = useCallback(async (slot: ScheduleSlotData) => {
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
    }, [dispatch, t])

    const openBind = useCallback((slotId: number) => {
        setBindSlotId(slotId)
        setIsBindOpen(true)
    }, [setBindSlotId, setIsBindOpen])

    return {
        onDeleteSchedule,
        onDeleteSlot,
        onUnbindLesson,
        openBind,
    }
}
