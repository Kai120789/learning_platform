import type { LessonStatus } from "@/entities/lesson"
import type { PracticeStatus } from "@/shared/mocks"

/** Soft status chips — muted chroma in dark (GitHub / VS Code label style). */
export function lessonStatusClass(status: LessonStatus): string {
    switch (status) {
        case "SCHEDULED":
            return "border-sky-600/25 bg-sky-500/[0.07] text-sky-800/70 dark:border-sky-400/20 dark:bg-sky-400/[0.08] dark:text-sky-300/70"
        case "IN_PROCESS":
            return "border-amber-600/25 bg-amber-500/[0.07] text-amber-800/70 dark:border-amber-400/20 dark:bg-amber-400/[0.08] dark:text-amber-200/65"
        case "COMPLETED":
            return "border-emerald-600/25 bg-emerald-500/[0.07] text-emerald-800/70 dark:border-emerald-400/20 dark:bg-emerald-400/[0.08] dark:text-emerald-300/70"
        case "CANCELLED":
            return "border-muted-foreground/25 bg-muted/40 text-muted-foreground"
        default:
            return ""
    }
}

export function practiceStatusClass(status: PracticeStatus): string {
    switch (status) {
        case "opened":
            return "border-sky-600/25 bg-sky-500/[0.07] text-sky-800/70 dark:border-sky-400/20 dark:bg-sky-400/[0.08] dark:text-sky-300/70"
        case "pending":
            return "border-amber-600/25 bg-amber-500/[0.07] text-amber-800/70 dark:border-amber-400/20 dark:bg-amber-400/[0.08] dark:text-amber-200/65"
        case "completed":
            return "border-emerald-600/25 bg-emerald-500/[0.07] text-emerald-800/70 dark:border-emerald-400/20 dark:bg-emerald-400/[0.08] dark:text-emerald-300/70"
        case "closed":
            return "border-border bg-muted/50 text-muted-foreground"
        default:
            return ""
    }
}
