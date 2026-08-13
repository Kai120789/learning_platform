import { cn } from "@/shared/lib/utils"

type LearningLogoProps = {
    className?: string
    compact?: boolean
}

export function LearningLogo({ className, compact }: LearningLogoProps) {
    if (compact) {
        return (
            <span
                className={cn(
                    "flex size-8 select-none items-center justify-center rounded-md border border-border text-sm font-semibold italic text-foreground",
                    className
                )}
            >
                L
            </span>
        )
    }

    return (
        <span
            className={cn(
                "select-none text-sm font-semibold italic tracking-[0.22em] text-foreground",
                className
            )}
        >
            LEARNING
        </span>
    )
}
