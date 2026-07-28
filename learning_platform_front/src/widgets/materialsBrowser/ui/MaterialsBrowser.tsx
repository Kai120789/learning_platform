import { useState } from "react"
import { useTranslation } from "react-i18next"
import { LayoutGrid, List } from "lucide-react"
import type { MaterialFileMock } from "@/shared/mocks"
import {
    getBreadcrumbs,
    getCurrentFolders,
    getFilesInFolder,
    getFolderLabel,
} from "../lib/folderHelpers"
import type { MaterialsViewMode } from "../model/types"
import { MaterialsBreadcrumbs } from "./MaterialsBreadcrumbs"
import { MaterialsGridView } from "./MaterialsGridView"
import { MaterialsListView } from "./MaterialsListView"
import { cn } from "@/shared/lib/utils"

export function MaterialsBrowser() {
    const { t } = useTranslation()
    const [viewMode, setViewMode] = useState<MaterialsViewMode>("list")
    const [currentFolderId, setCurrentFolderId] = useState<string | null>(null)

    const currentFolders = getCurrentFolders(currentFolderId)
    const currentFiles = getFilesInFolder(currentFolderId)
    const breadcrumbs = getBreadcrumbs(currentFolderId)
    const isEmpty = currentFolders.length === 0 && currentFiles.length === 0

    const openFolder = (folderId: string) => {
        setCurrentFolderId(folderId)
    }

    const goToBreadcrumb = (folderId: string | null) => {
        setCurrentFolderId(folderId)
    }

    const openFile = (file: MaterialFileMock) => {
        window.open(file.s3Link, "_blank", "noreferrer")
    }

    return (
        <div className="space-y-6">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-center justify-between">
                <MaterialsBreadcrumbs
                    breadcrumbs={breadcrumbs}
                    currentFolderId={currentFolderId}
                    onNavigate={goToBreadcrumb}
                />
                <div className="flex rounded-lg border border-border p-0.5 bg-secondary/60">
                    <button
                        type="button"
                        onClick={() => setViewMode("list")}
                        className={cn(
                            "flex items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                            viewMode === "list"
                                ? "bg-primary text-primary-foreground shadow-sm"
                                : "text-secondary-foreground hover:bg-muted"
                        )}
                    >
                        <List className="size-3.5" />
                        {t("materials.viewList")}
                    </button>
                    <button
                        type="button"
                        onClick={() => setViewMode("grid")}
                        className={cn(
                            "flex items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                            viewMode === "grid"
                                ? "bg-primary text-primary-foreground shadow-sm"
                                : "text-secondary-foreground hover:bg-muted"
                        )}
                    >
                        <LayoutGrid className="size-3.5" />
                        {t("materials.viewGrid")}
                    </button>
                </div>
            </div>

            {isEmpty ? (
                <div className="rounded-xl border border-dashed py-16 text-center text-muted-foreground">
                    {t("materials.emptyFolder")}
                </div>
            ) : viewMode === "list" ? (
                <MaterialsListView
                    folders={currentFolders}
                    files={currentFiles}
                    onOpenFolder={openFolder}
                    onOpenFile={openFile}
                    getFolderLabel={(folder) => getFolderLabel(folder, t)}
                />
            ) : (
                <MaterialsGridView
                    folders={currentFolders}
                    files={currentFiles}
                    onOpenFolder={openFolder}
                    onOpenFile={openFile}
                    getFolderLabel={(folder) => getFolderLabel(folder, t)}
                />
            )}
        </div>
    )
}
