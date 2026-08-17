import { useTranslation } from "react-i18next"
import { ChevronRight } from "lucide-react"
import type { MaterialBreadcrumb } from "../model/types"

type MaterialsBreadcrumbsProps = {
    breadcrumbs: MaterialBreadcrumb[]
    currentFolderId: number | null
    onNavigate: (folderId: number | null, index: number) => void
}

export function MaterialsBreadcrumbs({
    breadcrumbs,
    currentFolderId,
    onNavigate,
}: MaterialsBreadcrumbsProps) {
    const { t } = useTranslation()
    const isRoot = currentFolderId == null

    return (
        <nav className="flex min-w-0 flex-wrap items-center gap-x-1 text-sm text-muted-foreground">
            {isRoot ? (
                <span className="font-medium text-foreground">{t("materials.root")}</span>
            ) : (
                <button
                    type="button"
                    onClick={() => onNavigate(null, -1)}
                    className="cursor-pointer rounded px-1 py-0.5 hover:bg-muted hover:text-foreground"
                >
                    {t("materials.root")}
                </button>
            )}

            {breadcrumbs.map((folder, index) => {
                const isLast = index === breadcrumbs.length - 1

                return (
                    <div key={`${folder.id}-${index}`} className="flex min-w-0 items-center gap-x-1">
                        <ChevronRight className="size-3.5 shrink-0 opacity-50" />
                        {isLast ? (
                            <span className="truncate font-medium text-foreground" title={folder.title}>
                                {folder.title}
                            </span>
                        ) : (
                            <button
                                type="button"
                                onClick={() => onNavigate(folder.id, index)}
                                className="max-w-40 cursor-pointer truncate rounded px-1 py-0.5 hover:bg-muted hover:text-foreground"
                                title={folder.title}
                            >
                                {folder.title}
                            </button>
                        )}
                    </div>
                )
            })}
        </nav>
    )
}
