import { Folder } from "lucide-react"
import { useTranslation } from "react-i18next"
import { Badge } from "@/shared/ui/Badge"
import { Checkbox } from "@/shared/ui/Checkbox"
import { formatFileSize, mimeToMediaKind } from "@/entities/material"
import { cn } from "@/shared/lib/utils"
import { mediaIcons } from "../lib/mediaIcons"
import type { MaterialsViewProps } from "../model/types"
import { MaterialsItemMenu } from "./MaterialsItemMenu"

export function MaterialsListView({
    folders,
    files,
    canEdit,
    selectionMode,
    selection,
    movingItem,
    onOpenFolder,
    onOpenFile,
    onToggleSelect,
    onItemAction,
}: MaterialsViewProps) {
    const { t } = useTranslation()

    const gridClass = selectionMode
        ? "grid-cols-[auto_minmax(0,1fr)] sm:grid-cols-[auto_minmax(0,1fr)_120px] md:grid-cols-[auto_minmax(0,1fr)_120px_100px]"
        : "grid-cols-[minmax(0,1fr)_40px] sm:grid-cols-[minmax(0,1fr)_120px_40px] md:grid-cols-[minmax(0,1fr)_120px_100px_40px]"

    const isMovingFolder = (id: number) =>
        movingItem?.kind === "folder" && movingItem.folder.id === id

    const isMovingFile = (id: number) =>
        movingItem?.kind === "material" && movingItem.file.id === id

    return (
        <div className="overflow-hidden rounded-xl border bg-card">
            <div className={`grid ${gridClass} gap-3 border-b bg-muted/40 px-4 py-2 text-xs font-medium text-muted-foreground`}>
                {selectionMode && <span className="w-4" />}
                <span>{t("materials.name")}</span>
                <span className="hidden sm:block">{t("materials.type")}</span>
                <span className="hidden md:block">{t("materials.size")}</span>
                {!selectionMode && <span className="sr-only">{t("materials.menu.actions")}</span>}
            </div>

            <div className="divide-y">
                {folders.map((folder) => {
                    const checked = selection.folders.has(folder.id)

                    return (
                        <div
                            key={folder.id}
                            className={cn(
                                `grid ${gridClass} items-center gap-3 px-4 py-2.5 transition-colors hover:bg-muted/50`,
                                isMovingFolder(folder.id) && "opacity-40",
                                checked && "bg-muted/40"
                            )}
                        >
                            {selectionMode && (
                                <Checkbox
                                    checked={checked}
                                    onCheckedChange={() => onToggleSelect({ kind: "folder", folder })}
                                />
                            )}
                            <button
                                type="button"
                                onClick={() => {
                                    if (selectionMode) {
                                        onToggleSelect({ kind: "folder", folder })
                                        return
                                    }
                                    onOpenFolder(folder)
                                }}
                                className="flex min-w-0 cursor-pointer items-center gap-3 text-left"
                            >
                                <Folder
                                    className="size-5 shrink-0 fill-foreground/70 text-foreground/70"
                                    strokeWidth={1}
                                />
                                <span className="block min-w-0 truncate font-medium" title={folder.title}>
                                    {folder.title}
                                </span>
                            </button>
                            <span className="hidden items-center text-sm text-muted-foreground sm:flex">
                                {t("materials.folders")}
                            </span>
                            <span className="hidden items-center text-sm text-muted-foreground md:flex">
                                –
                            </span>
                            {!selectionMode && (
                                <MaterialsItemMenu
                                    canEdit={canEdit}
                                    onAction={(action) => onItemAction({ kind: "folder", folder }, action)}
                                />
                            )}
                        </div>
                    )
                })}

                {files.map((file) => {
                    const kind = mimeToMediaKind(file.mimeType)
                    const Icon = mediaIcons[kind]
                    const checked = selection.materials.has(file.id)

                    return (
                        <div
                            key={file.id}
                            className={cn(
                                `grid ${gridClass} items-center gap-3 px-4 py-2.5 transition-colors hover:bg-muted/50`,
                                isMovingFile(file.id) && "opacity-40",
                                checked && "bg-muted/40"
                            )}
                        >
                            {selectionMode && (
                                <Checkbox
                                    checked={checked}
                                    onCheckedChange={() => onToggleSelect({ kind: "material", file })}
                                />
                            )}
                            <button
                                type="button"
                                onClick={() => {
                                    if (selectionMode) {
                                        onToggleSelect({ kind: "material", file })
                                        return
                                    }
                                    onOpenFile(file)
                                }}
                                className="flex min-w-0 cursor-pointer items-center gap-3 text-left"
                            >
                                <Icon className="size-5 shrink-0 text-foreground/80" />
                                <span className="block min-w-0 truncate font-medium" title={file.title}>
                                    {file.title}
                                </span>
                            </button>
                            <span className="hidden items-center sm:flex">
                                <Badge variant="secondary">{t(`mediaType.${kind}`)}</Badge>
                            </span>
                            <span className="hidden items-center text-sm text-muted-foreground md:flex">
                                {formatFileSize(file.size)}
                            </span>
                            {!selectionMode && (
                                <MaterialsItemMenu
                                    canEdit={canEdit}
                                    isFile
                                    onAction={(action) => onItemAction({ kind: "material", file }, action)}
                                />
                            )}
                        </div>
                    )
                })}
            </div>
        </div>
    )
}
