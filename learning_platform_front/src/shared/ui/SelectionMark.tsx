import { Check } from "lucide-react"
import { cn } from "@/shared/lib/utils"

type SelectionMarkProps = {
    checked: boolean
    className?: string
}

export function SelectionMark({ checked, className }: SelectionMarkProps) {
    return (
        <span className={cn("shrink-0 pr-4", className)}>
            <span
                aria-hidden
                className={cn(
                    "flex size-4 items-center justify-center rounded-[4px] border border-input",
                    checked && "border-primary bg-primary text-primary-foreground",
                )}
            >
                {checked && <Check className="size-3" strokeWidth={3} />}
            </span>
        </span>
    )
}
