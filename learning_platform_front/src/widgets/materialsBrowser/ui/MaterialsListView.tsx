import { Folder } from "lucide-react"
import { useTranslation } from "react-i18next"
import { Badge } from "@/shared/ui/Badge"
import { mediaIcons } from "../lib/mediaIcons"
import type { MaterialsViewProps } from "../model/types"

export function MaterialsListView({
    folders,
    files,
    onOpenFolder,
    onOpenFile,
    getFolderLabel,
}: MaterialsViewProps) {
    const { t } = useTranslation()

    const gridClass =
        "grid-cols-[minmax(0,1fr)] sm:grid-cols-[minmax(0,1fr)_120px] md:grid-cols-[minmax(0,1fr)_120px_100px] lg:grid-cols-[minmax(0,1fr)_120px_100px_110px]";

    return (
        <div className="overflow-hidden rounded-xl border bg-card">
            <div className={`grid ${gridClass} gap-3 border-b bg-muted/40 px-4 py-2 text-xs font-medium text-muted-foreground`}>
                <span>{t("materials.name")}</span>
                <span className="hidden sm:block">{t("materials.type")}</span>
                <span className="hidden md:block">{t("materials.size")}</span>
                <span className="hidden lg:block">{t("materials.modified")}</span>
            </div>

            <div className="divide-y">
                {folders.map((folder) => (
                    <button
                        key={folder.id}
                        type="button"
                        onClick={() => onOpenFolder(folder.id)}
                        className={`grid ${gridClass} 
                            w-full cursor-pointer gap-3 px-4 py-3 text-left transition-colors hover:bg-muted/50
                        `}
                    >
                        <div className="flex min-w-0 items-center gap-3">
                            <Folder
                                className="size-5 shrink-0 fill-foreground/70 text-foreground/70"
                                strokeWidth={1}
                            />
                            <div className="min-w-0">
                                <div className="truncate font-medium">{getFolderLabel(folder)}</div>
                                <div className="text-xs text-muted-foreground sm:hidden">
                                    {t("materials.items", { count: folder.itemCount })}
                                </div>
                            </div>
                        </div>
                        <span className="hidden items-center text-sm text-muted-foreground sm:flex">
                            {t("materials.folders")}
                        </span>
                        <span className="hidden items-center text-sm text-muted-foreground md:flex">
                            {t("materials.items", { count: folder.itemCount })}
                        </span>
                        <span className="hidden items-center text-sm text-muted-foreground lg:flex">
                            —
                        </span>
                    </button>
                ))}

                {files.map((file) => {
                    const Icon = mediaIcons[file.type]

                    return (
                        <button
                            key={file.id}
                            type="button"
                            onClick={() => onOpenFile(file)}
                            className={`grid ${gridClass} 
                                w-full cursor-pointer gap-3 px-4 py-3 text-left transition-colors hover:bg-muted/50
                            `}
                        >

                            <div className="flex min-w-0 items-center gap-3">
                                <Icon className="size-5 shrink-0 text-foreground/80" />
                                <div className="truncate font-medium">{file.title}</div>
                            </div>
                            <span className="hidden items-center sm:flex">
                                <Badge variant="secondary">{t(`mediaType.${file.type}`)}</Badge>
                            </span>
                            <span className="hidden items-center text-sm text-muted-foreground md:flex">
                                {file.sizeLabel}
                            </span>
                            <span className="hidden items-center text-sm text-muted-foreground lg:flex">
                                {new Date(file.updatedAt).toLocaleDateString()}
                            </span>
                        </button>
                    )
                })}
            </div>
        </div>
    )
}
