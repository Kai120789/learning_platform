import type { LessonStatus, PracticeStatus } from "@/shared/mocks"

export function lessonStatusClass(status: LessonStatus): string {
    switch (status) {
        case "SCHEDULED":
            return "border-current/50 bg-sky-500/[0.05] text-sky-700/65 dark:text-sky-200/55"
        case "IN_PROCESS":
            return "border-current/50 bg-amber-500/[0.05] text-amber-700/65 dark:text-amber-200/55"
        case "COMPLETED":
            return "border-current/50 bg-emerald-500/[0.05] text-emerald-700/65 dark:text-emerald-200/55"
        case "CANCELLED":
            return "border-current/50 bg-rose-500/[0.05] text-rose-700/65 dark:text-rose-200/55"
        default:
            return ""
    }
}

export function practiceStatusClass(status: PracticeStatus): string {
    switch (status) {
        case "opened":
            return "border-current/50 bg-sky-500/[0.05] text-sky-700/65 dark:text-sky-200/55"
        case "pending":
            return "border-current/50 bg-amber-500/[0.05] text-amber-700/65 dark:text-amber-200/55"
        case "completed":
            return "border-current/50 bg-emerald-500/[0.05] text-emerald-700/65 dark:text-emerald-200/55"
        case "closed":
            return "border-current/50 bg-stone-500/[0.05] text-stone-600/65 dark:text-stone-200/55"
        default:
            return ""
    }
}
