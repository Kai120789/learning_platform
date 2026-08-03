import type { ShortUserInfo } from "@/entities/group"
import { cn } from "@/shared/lib/utils"
import { SelectionMark } from "@/shared/ui/SelectionMark"

type CandidateUserRowProps = {
    user: ShortUserInfo
    isSelected: boolean
    onToggle: () => void
}

export function CandidateUserRow({
    user,
    isSelected,
    onToggle,
}: CandidateUserRowProps) {
    const initials = `${user.name?.[0] ?? ""}${user.surname?.[0] ?? ""}`.toUpperCase() || "?"

    return (
        <div
            role="button"
            tabIndex={0}
            className={cn(
                "flex min-w-0 cursor-pointer items-center justify-between gap-3 overflow-hidden py-2.5 transition-colors hover:bg-muted/40",
                isSelected && "bg-primary/5",
            )}
            onClick={onToggle}
            onKeyDown={(e) => {
                if (e.key === "Enter" || e.key === " ") {
                    e.preventDefault()
                    onToggle()
                }
            }}
        >
            <div className="flex min-w-0 flex-1 items-center gap-2.5 overflow-hidden">
                <span className="flex size-6 shrink-0 items-center justify-center rounded-full bg-muted text-[10px] font-medium text-muted-foreground">
                    {initials}
                </span>

                <div className="min-w-0 flex-1 overflow-hidden">
                    <p className="truncate text-sm font-medium leading-none">
                        {`${user.name} ${user.surname}`.trim()}
                    </p>

                    {user.tg_username && (
                        <p className="mt-1 truncate text-xs text-muted-foreground">
                            {user.tg_username}
                        </p>
                    )}
                </div>
            </div>

            <SelectionMark checked={isSelected} />
        </div>
    )
}
