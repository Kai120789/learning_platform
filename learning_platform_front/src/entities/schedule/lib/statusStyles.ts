import type { ScheduleSlotStatus } from "../types/types"

export function scheduleSlotStatusClass(status: ScheduleSlotStatus): string {
    switch (status) {
        case "FREE":
            return "border-current/50 bg-emerald-500/[0.05] text-emerald-700/65 dark:text-emerald-200/55"
        case "BOOKED":
            return "border-current/50 bg-sky-500/[0.05] text-sky-700/65 dark:text-sky-200/55"
        default:
            return ""
    }
}
