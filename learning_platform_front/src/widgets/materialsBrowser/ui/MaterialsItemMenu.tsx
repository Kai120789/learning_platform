import { Download, Info, MoreHorizontal, Pencil, Trash2, FolderInput, CheckSquare } from "lucide-react"
import { useTranslation } from "react-i18next"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/shared/ui/DropdownMenu"
import { cn } from "@/shared/lib/utils"

export type MaterialItemAction =
    | "info"
    | "download"
    | "rename"
    | "move"
    | "delete"
    | "select"

type MaterialsItemMenuProps = {
    canEdit: boolean
    isFile?: boolean
    onAction: (action: MaterialItemAction) => void
    className?: string
    triggerClassName?: string
}

export function MaterialsItemMenu({
    canEdit,
    isFile = false,
    onAction,
    className,
    triggerClassName,
}: MaterialsItemMenuProps) {
    const { t } = useTranslation()

    return (
        <div className={cn("shrink-0", className)} onClick={(e) => e.stopPropagation()}>
            <DropdownMenu modal={false}>
                <DropdownMenuTrigger
                    className={cn(
                        "flex size-7 cursor-pointer items-center justify-center rounded-md text-muted-foreground transition-colors hover:bg-muted hover:text-foreground",
                        triggerClassName
                    )}
                >
                    <MoreHorizontal className="size-4" />
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end" className="w-44 min-w-44 bg-background">
                    <DropdownMenuItem
                        className="cursor-pointer gap-2 text-sm"
                        onClick={() => onAction("info")}
                    >
                        <Info className="size-4" />
                        {t("materials.menu.info")}
                    </DropdownMenuItem>
                    {isFile && (
                        <DropdownMenuItem
                            className="cursor-pointer gap-2 text-sm"
                            onClick={() => onAction("download")}
                        >
                            <Download className="size-4" />
                            {t("materials.menu.download")}
                        </DropdownMenuItem>
                    )}
                    {canEdit && (
                        <>
                            <DropdownMenuItem
                                className="cursor-pointer gap-2 text-sm"
                                onClick={() => onAction("select")}
                            >
                                <CheckSquare className="size-4" />
                                {t("materials.menu.select")}
                            </DropdownMenuItem>
                            <DropdownMenuItem
                                className="cursor-pointer gap-2 text-sm"
                                onClick={() => onAction("rename")}
                            >
                                <Pencil className="size-4" />
                                {t("materials.menu.rename")}
                            </DropdownMenuItem>
                            <DropdownMenuItem
                                className="cursor-pointer gap-2 text-sm"
                                onClick={() => onAction("move")}
                            >
                                <FolderInput className="size-4" />
                                {t("materials.menu.move")}
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem
                                variant="destructive"
                                className="cursor-pointer gap-2 text-sm"
                                onClick={() => onAction("delete")}
                            >
                                <Trash2 className="size-4" />
                                {t("materials.menu.delete")}
                            </DropdownMenuItem>
                        </>
                    )}
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}
