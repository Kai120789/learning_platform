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
                <div className="flex flex-row sm:gap-2 border border-border sm:border-none rounded-xl">
                    <button
                        onClick={() => setViewMode("list")}
                        className={cn(
                            "flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm text-left transition-colors justify-center sm:justify-start",
                            viewMode === "list"
                                ? "bg-black text-white"
                                : "sm:bg-white sm:border sm:border-border hover:bg-muted"
                        )}
                    >
                        <List className="size-4" />
                        {t("materials.viewList")}
                    </button>
                    <button
                        onClick={() => setViewMode("grid")}
                        className={cn(
                            "flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm text-left transition-colors justify-center sm:justify-start",
                            viewMode === "grid"
                                ? "bg-black text-white"
                                : "sm:bg-white sm:border sm:border-border hover:bg-muted"
                        )}
                    >
                        <LayoutGrid className="size-4" />
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
