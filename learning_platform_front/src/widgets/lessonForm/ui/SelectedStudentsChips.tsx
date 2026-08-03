import { useTranslation } from "react-i18next"
import { X } from "lucide-react"
import type { ShortUserInfo } from "@/entities/group"
import { displayName } from "../lib/userHelpers"

type SelectedStudentsChipsProps = {
    users: ShortUserInfo[]
    onRemove: (userID: number) => void
}

export function SelectedStudentsChips({
    users,
    onRemove,
}: SelectedStudentsChipsProps) {
    const { t } = useTranslation()

    return (
        <div className="space-y-2">
            <p className="text-sm text-muted-foreground">
                {t("createLesson.selectedStudents", { count: users.length })}
            </p>
            {users.length === 0 ? (
                <p className="text-sm text-muted-foreground/80">
                    {t("createLesson.noSelectedStudents")}
                </p>
            ) : (
                <div className="flex flex-wrap gap-1.5">
                    {users.map((user) => (
                        <button
                            type="button"
                            key={user.id}
                            className="inline-flex cursor-pointer items-center gap-1 rounded-md bg-muted px-2 py-1 text-xs text-foreground transition-colors hover:bg-muted/80"
                            onClick={() => onRemove(user.id)}
                            title={t("createLesson.removeStudent")}
                        >
                            {displayName(user) || `#${user.id}`}
                            <X className="size-3 opacity-60" />
                        </button>
                    ))}
                </div>
            )}
        </div>
    )
}
