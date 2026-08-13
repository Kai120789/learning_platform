import { LearningLogo } from "@/shared/ui/LearningLogo"
import { Menu } from "lucide-react"
import { useTranslation } from "react-i18next"

type TopMenuProps = {
    onClick: () => void
}

export function TopMenu({ onClick }: TopMenuProps) {
    const { t } = useTranslation()

    return (
        <header className="sticky top-0 z-30 flex h-14 items-center gap-3 border-b border-border bg-background px-4 lg:hidden">
            <button
                type="button"
                onClick={onClick}
                className="rounded-md p-1 text-foreground hover:bg-muted"
                aria-label={t("sidebar.open")}
            >
                <Menu className="size-5" />
            </button>
            <LearningLogo />
        </header>
    )
}
