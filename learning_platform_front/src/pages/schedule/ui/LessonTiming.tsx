import { useTranslation } from "react-i18next"
import type { LessonData } from "@/entities/lesson"
import { useCountdown } from "@/shared/lib/useCountdown"

type LessonTimingProps = {
    lesson: LessonData
    canEdit: boolean
}

export function LessonTiming({ lesson, canEdit }: LessonTimingProps) {
    const { t } = useTranslation()
    const { isPast, label } = useCountdown(
        lesson.status === "SCHEDULED" ? lesson.startTime : null,
    )

    if (lesson.status === "CANCELLED" || lesson.status === "COMPLETED") {
        return null
    }

    if (lesson.status === "IN_PROCESS") {
        return (
            <div className="text-xs text-amber-700/80 dark:text-amber-200/70">
                {t("lessons.inProcessHint")}
            </div>
        )
    }

    if (!isPast) {
        return (
            <div className="text-xs text-muted-foreground">
                {t("lessons.startsIn", { time: label })}
            </div>
        )
    }

    if (canEdit) {
        return (
            <div className="text-xs text-amber-700/80 dark:text-amber-200/70">
                {t("lessons.waitingStartTutor")}
            </div>
        )
    }

    return (
        <div className="text-xs text-amber-700/80 dark:text-amber-200/70">
            {t("lessons.waitingStartStudent")}
        </div>
    )
}
