import { Folder } from "lucide-react"
import { useTranslation } from "react-i18next"
import { mediaIcons } from "../lib/mediaIcons"
import type { MaterialsViewProps } from "../model/types"

export function MaterialsGridView({
    folders,
    files,
    onOpenFolder,
    onOpenFile,
    getFolderLabel,
}: MaterialsViewProps) {
    const { t } = useTranslation()

    return (
        <div className="space-y-8">
            {folders.length > 0 && (
                <div className="space-y-3">
                    <div className="grid auto-rows-fr gap-2 grid-cols-2 sm:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8">
                        {folders.map((folder) => (
                            <button
                                key={folder.id}
                                type="button"
                                onClick={() => onOpenFolder(folder.id)}
                                className="group flex h-full min-h-36 flex-col items-center gap-2 rounded-xl p-2 text-center transition-colors hover:bg-muted/50"
                            >
                                <Folder
                                    className="size-20 shrink-0 fill-foreground/70 text-foreground/70 transition-transform group-hover:scale-[1.03]"
                                    strokeWidth={1}
                                />
                                <div className="mt-auto w-full space-y-1">
                                    <div className="line-clamp-2 min-h-10 text-sm font-medium">
                                        {getFolderLabel(folder)}
                                    </div>
                                    <div className="text-xs text-muted-foreground">
                                        {t("materials.items", { count: folder.itemCount })}
                                    </div>
                                </div>
                            </button>
                        ))}
                    </div>
                </div>
            )}

            {files.length > 0 && (
                <div className="space-y-3">
                    <div className="grid auto-rows-fr gap-2 grid-cols-2 sm:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8">
                        {files.map((file) => {
                            const Icon = mediaIcons[file.type]

                            return (
                                <button
                                    key={file.id}
                                    type="button"
                                    onClick={() => onOpenFile(file)}
                                    className="group flex h-full min-h-40 flex-col items-center gap-2 rounded-xl p-2 text-center transition-colors hover:bg-muted/50"
                                >
                                    <Icon
                                        className="size-16 shrink-0 text-foreground/75 transition-transform group-hover:scale-[1.03]"
                                        strokeWidth={1.25}
                                    />
                                    <div className="mt-auto w-full space-y-1">
                                        <div className="line-clamp-2 min-h-10 text-sm font-medium">
                                            {file.title}
                                        </div>
                                        <div className="text-xs text-muted-foreground">
                                            {file.sizeLabel}
                                        </div>
                                        <div className="text-xs text-muted-foreground">
                                            {new Date(file.updatedAt).toLocaleDateString()}
                                        </div>
                                    </div>
                                </button>
                            )
                        })}
                    </div>
                </div>
            )}
        </div>
    )
}
