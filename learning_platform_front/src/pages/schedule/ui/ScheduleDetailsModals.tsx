import { AddScheduleSlotsModal } from "@/widgets/addScheduleSlotsModal"
import { BindLessonModal } from "@/widgets/bindLessonModal"
import { EditScheduleModal } from "@/widgets/editScheduleModal"
import { EditScheduleSlotModal } from "@/widgets/editScheduleSlotModal"
import type { ScheduleData, ScheduleSlotData } from "@/entities/schedule"

type ScheduleDetailsModalsProps = {
    schedule: ScheduleData
    isEditOpen: boolean
    setIsEditOpen: (open: boolean) => void
    isAddSlotsOpen: boolean
    setIsAddSlotsOpen: (open: boolean) => void
    selectedDate: Date
    editingSlot: ScheduleSlotData | null
    setEditingSlot: (slot: ScheduleSlotData | null) => void
    isBindOpen: boolean
    setIsBindOpen: (open: boolean) => void
    freeSlots: ScheduleSlotData[]
    bindSlotId: number | null
}

export function ScheduleDetailsModals({
    schedule,
    isEditOpen,
    setIsEditOpen,
    isAddSlotsOpen,
    setIsAddSlotsOpen,
    selectedDate,
    editingSlot,
    setEditingSlot,
    isBindOpen,
    setIsBindOpen,
    freeSlots,
    bindSlotId,
}: ScheduleDetailsModalsProps) {
    return (
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
    )
}
