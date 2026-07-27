import { ChevronRight } from "lucide-react"
import { useTranslation } from "react-i18next"
import { cn } from "@/shared/lib/utils"
import type { MaterialFolderMock } from "@/shared/mocks"
import { getFolderLabel } from "../lib/folderHelpers"

type MaterialsBreadcrumbsProps = {
    breadcrumbs: MaterialFolderMock[]
    currentFolderId: string | null
    onNavigate: (folderId: string | null) => void
}

export function MaterialsBreadcrumbs({
    breadcrumbs,
    currentFolderId,
    onNavigate,
}: MaterialsBreadcrumbsProps) {
    const { t } = useTranslation()

    return (
        <div className="flex flex-wrap items-center gap-1 text-sm">
            <button
                type="button"
                onClick={() => onNavigate(null)}
                className={cn(
                    "rounded-md py-1 transition-colors hover:bg-muted",
                    currentFolderId == null && "font-medium"
                )}
            >
                {t("materials.root")}
            </button>
            {breadcrumbs.map((folder) => (
                <div key={folder.id} className="flex items-center gap-1">
                    <ChevronRight className="size-4 text-muted-foreground" />
                    <button
                        type="button"
                        onClick={() => onNavigate(folder.id)}
                        className={cn(
                            "rounded-md px-2 py-1 transition-colors hover:bg-muted",
                            currentFolderId === folder.id && "font-medium"
                        )}
                    >
                        {getFolderLabel(folder, t)}
                    </button>
                </div>
            ))}
        </div>
    )
}
