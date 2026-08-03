import { useTranslation } from "react-i18next"
import { FaTrash } from "react-icons/fa"
import { Button } from "@/shared/ui/Button"
import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"

export type SlotDraft = {
    key: string
    time: string
    duration: number
}

type ScheduleSlotDraftRowProps = {
    slot: SlotDraft
    index: number
    canRemove: boolean
    onUpdate: (patch: Partial<SlotDraft>) => void
    onRemove: () => void
}

export function ScheduleSlotDraftRow({
    slot,
    index,
    canRemove,
    onUpdate,
    onRemove,
}: ScheduleSlotDraftRowProps) {
    const { t } = useTranslation()

    return (
        <div className="grid grid-cols-[1fr_88px_auto] gap-2 items-end rounded-lg border p-2.5">
            <Field className="w-full">
                <FieldLabel className="text-xs text-muted-foreground">
                    {t("createSchedule.slotStart", { index: index + 1 })}
                </FieldLabel>
                <Input
                    className="w-full"
                    required
                    type="time"
                    value={slot.time}
                    onChange={(e) => onUpdate({ time: e.target.value })}
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
                    onChange={(e) => onUpdate({ duration: Number(e.target.value) })}
                />
            </Field>
            <Button
                type="button"
                size="icon-sm"
                variant="outline"
                disabled={!canRemove}
                onClick={onRemove}
                aria-label={t("createSchedule.removeSlot")}
            >
                <FaTrash className="size-3" />
            </Button>
        </div>
    )
}
