import { Folder } from "lucide-react"
import { mimeToMediaKind } from "@/entities/material"
import { Checkbox } from "@/shared/ui/Checkbox"
import { cn } from "@/shared/lib/utils"
import { mediaIcons } from "../lib/mediaIcons"
import type { MaterialsViewProps } from "../model/types"
import { MaterialsItemMenu } from "./MaterialsItemMenu"

export function MaterialsGridView({
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
    const isMovingFolder = (id: number) =>
        movingItem?.kind === "folder" && movingItem.folder.id === id

    const isMovingFile = (id: number) =>
        movingItem?.kind === "material" && movingItem.file.id === id

    return (
        <div className="space-y-8">
            {folders.length > 0 && (
                <div className="grid auto-rows-fr gap-2 grid-cols-2 sm:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8">
                    {folders.map((folder) => {
                        const checked = selection.folders.has(folder.id)

                        return (
                            <div
                                key={folder.id}
                                className={cn(
                                    "group relative flex h-full flex-col items-center rounded-xl p-2 text-center transition-colors hover:bg-muted/50",
                                    isMovingFolder(folder.id) && "opacity-40",
                                    checked && "bg-muted/50 ring-1 ring-border"
                                )}
                            >
                                {selectionMode ? (
                                    <div className="absolute top-1 left-1 z-10">
                                        <Checkbox
                                            checked={checked}
                                            onCheckedChange={() => onToggleSelect({ kind: "folder", folder })}
                                        />
                                    </div>
                                ) : (
                                    <MaterialsItemMenu
                                        canEdit={canEdit}
                                        className="absolute top-1 right-1 z-10 opacity-0 transition-opacity group-hover:opacity-100 focus-within:opacity-100"
                                        onAction={(action) => onItemAction({ kind: "folder", folder }, action)}
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
                                    className="flex w-full cursor-pointer flex-col items-center gap-1.5"
                                >
                                    <div className="flex h-16 w-full items-center justify-center">
                                        <Folder
                                            className="size-14 shrink-0 fill-foreground/70 text-foreground/70 transition-transform group-hover:scale-[1.03]"
                                            strokeWidth={1}
                                        />
                                    </div>
                                    <div
                                        className="line-clamp-2 w-full overflow-hidden text-ellipsis text-sm font-medium break-words"
                                        title={folder.title}
                                    >
                                        {folder.title}
                                    </div>
                                </button>
                            </div>
                        )
                    })}
                </div>
            )}

            {files.length > 0 && (
                <div className="grid auto-rows-fr gap-2 grid-cols-2 sm:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8">
                    {files.map((file) => {
                        const kind = mimeToMediaKind(file.mimeType)
                        const Icon = mediaIcons[kind]
                        const checked = selection.materials.has(file.id)

                        return (
                            <div
                                key={file.id}
                                className={cn(
                                    "group relative flex h-full flex-col items-center rounded-xl p-2 text-center transition-colors hover:bg-muted/50",
                                    isMovingFile(file.id) && "opacity-40",
                                    checked && "bg-muted/50 ring-1 ring-border"
                                )}
                            >
                                {selectionMode ? (
                                    <div className="absolute top-1 left-1 z-10">
                                        <Checkbox
                                            checked={checked}
                                            onCheckedChange={() => onToggleSelect({ kind: "material", file })}
                                        />
                                    </div>
                                ) : (
                                    <MaterialsItemMenu
                                        canEdit={canEdit}
                                        isFile
                                        className="absolute top-1 right-1 z-10 opacity-0 transition-opacity group-hover:opacity-100 focus-within:opacity-100"
                                        onAction={(action) => onItemAction({ kind: "material", file }, action)}
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
                                    className="flex w-full cursor-pointer flex-col items-center gap-1.5"
                                >
                                    <div className="flex h-16 w-full items-center justify-center">
                                        <Icon
                                            className="size-12 shrink-0 text-foreground/75 transition-transform group-hover:scale-[1.03]"
                                            strokeWidth={1.25}
                                        />
                                    </div>
                                    <div
                                        className="line-clamp-2 w-full overflow-hidden text-ellipsis text-sm font-medium break-words"
                                        title={file.title}
                                    >
                                        {file.title}
                                    </div>
                                </button>
                            </div>
                        )
                    })}
                </div>
            )}
        </div>
    )
}
