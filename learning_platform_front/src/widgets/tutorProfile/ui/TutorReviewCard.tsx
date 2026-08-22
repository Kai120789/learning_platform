import type { ReactNode } from "react"
import { Star } from "lucide-react"
import {
    formatTutorName,
    getTutorInitials,
} from "@/entities/tutor"
import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { cn } from "@/shared/lib/utils"

type ReviewAuthor = {
    name: string
    surname: string
}

type TutorReviewCardProps = {
    author: ReviewAuthor
    subjectLabel: string
    dateLabel: string
    rating: number
    text: string
    muted?: boolean
    badge?: string
    actions?: ReactNode
}

export function TutorReviewCard({
    author,
    subjectLabel,
    dateLabel,
    rating,
    text,
    muted,
    badge,
    actions,
}: TutorReviewCardProps) {
    return (
        <article
            className={cn(
                "rounded-xl p-3.5",
                muted
                    ? "bg-muted ring-1 ring-foreground/12"
                    : "border border-border bg-card",
            )}
        >
            <div className="flex items-start gap-3">
                <Avatar className="size-9 shrink-0">
                    <AvatarFallback className="text-xs">
                        {getTutorInitials(author)}
                    </AvatarFallback>
                </Avatar>
                <div className="min-w-0 flex-1">
                    <div className="flex items-start justify-between gap-3">
                        <div className="min-w-0">
                            <div className="flex min-w-0 flex-wrap items-center gap-x-2 gap-y-0.5">
                                <span className="truncate text-sm font-medium">
                                    {formatTutorName(author)}
                                </span>
                                {badge && (
                                    <span className="rounded-md bg-foreground/8 px-1.5 py-px text-[11px] font-medium text-foreground/80">
                                        {badge}
                                    </span>
                                )}
                            </div>
                            <div className="mt-0.5 truncate text-xs text-muted-foreground">
                                {subjectLabel}
                                {subjectLabel && dateLabel ? " · " : ""}
                                {dateLabel}
                            </div>
                        </div>
                        <div className="flex shrink-0 items-center gap-1">
                            <div className="inline-flex items-center gap-1 text-sm">
                                <Star className="size-3.5 fill-current" />
                                {rating}
                            </div>
                            {actions}
                        </div>
                    </div>
                    <p className="mt-2 text-sm whitespace-pre-wrap text-muted-foreground">
                        {text}
                    </p>
                </div>
            </div>
        </article>
    )
}
