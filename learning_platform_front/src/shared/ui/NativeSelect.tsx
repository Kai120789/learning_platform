import * as React from "react"
import { ChevronDownIcon } from "lucide-react"

import { cn } from "@/shared/lib/utils"

function NativeSelect({
    className,
    containerClassName,
    children,
    ...props
}: React.ComponentProps<"select"> & {
    containerClassName?: string
}) {
    return (
        <div
            data-slot="native-select"
            className={cn("relative min-w-0 w-full", containerClassName)}
        >
            <select
                data-slot="native-select-control"
                className={cn(
                    "h-10 w-full min-w-0 appearance-none rounded-lg border border-input bg-transparent py-1 pl-2.5 pr-9 text-base transition-colors outline-none focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50 disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm dark:bg-input/30",
                    className
                )}
                {...props}
            >
                {children}
            </select>
            <ChevronDownIcon
                aria-hidden
                className="pointer-events-none absolute top-1/2 right-3 size-4 -translate-y-1/2 text-muted-foreground"
            />
        </div>
    )
}

export { NativeSelect }
