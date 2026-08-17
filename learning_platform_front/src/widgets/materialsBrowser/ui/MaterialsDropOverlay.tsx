import { FileText, Image, Video } from "lucide-react"
import { useTranslation } from "react-i18next"
import { cn } from "@/shared/lib/utils"

type MaterialsDropOverlayProps = {
    visible: boolean
}

export function MaterialsDropOverlay({ visible }: MaterialsDropOverlayProps) {
    const { t } = useTranslation()

    return (
        <div
            className={cn(
                "pointer-events-none absolute inset-0 z-20 flex items-center justify-center rounded-xl transition-all duration-200",
                visible
                    ? "opacity-100 backdrop-blur-[2px]"
                    : "opacity-0"
            )}
        >
            <div
                className={cn(
                    "absolute inset-0 rounded-xl border-2 border-dashed transition-colors duration-200",
                    visible ? "border-primary/60 bg-background/70" : "border-transparent bg-transparent"
                )}
            />

            <div className="relative flex h-40 w-full max-w-sm items-center justify-center">
                <Image
                    className={cn(
                        "absolute size-14 text-foreground/25 transition-all duration-300 ease-out",
                        visible
                            ? "-translate-x-20 -translate-y-4 -rotate-[18deg] scale-100 opacity-100"
                            : "translate-x-0 translate-y-3 scale-75 opacity-0"
                    )}
                    strokeWidth={1.2}
                />
                <FileText
                    className={cn(
                        "absolute size-16 text-foreground/30 transition-all duration-300 ease-out delay-75",
                        visible
                            ? "-translate-y-1 rotate-6 scale-100 opacity-100"
                            : "translate-y-4 scale-75 opacity-0"
                    )}
                    strokeWidth={1.2}
                />
                <Video
                    className={cn(
                        "absolute size-14 text-foreground/25 transition-all duration-300 ease-out delay-100",
                        visible
                            ? "translate-x-20 translate-y-3 rotate-[14deg] scale-100 opacity-100"
                            : "translate-x-0 translate-y-3 scale-75 opacity-0"
                    )}
                    strokeWidth={1.2}
                />
            </div>

            <p
                className={cn(
                    "absolute bottom-10 text-sm font-medium text-foreground/80 transition-all duration-200",
                    visible ? "translate-y-0 opacity-100" : "translate-y-2 opacity-0"
                )}
            >
                {t("materials.dropHere")}
            </p>
        </div>
    )
}
