import type { MaterialData, MaterialFolderData, MaterialMediaKind } from "@/entities/material"
import type { MaterialItemAction } from "../ui/MaterialsItemMenu"

export type MaterialsViewMode = "list" | "grid"

export type MaterialBreadcrumb = {
    id: number
    title: string
}

export type MaterialsTarget =
    | { kind: "folder"; folder: MaterialFolderData }
    | { kind: "material"; file: MaterialData }

export type MaterialsMovingItem = MaterialsTarget

export type MaterialsSelection = {
    folders: Set<number>
    materials: Set<number>
}

export type MaterialsViewProps = {
    folders: MaterialFolderData[]
    files: MaterialData[]
    canEdit: boolean
    selectionMode: boolean
    selection: MaterialsSelection
    movingItem: MaterialsMovingItem | null
    onOpenFolder: (folder: MaterialFolderData) => void
    onOpenFile: (file: MaterialData) => void
    onToggleSelect: (target: MaterialsTarget) => void
    onItemAction: (target: MaterialsTarget, action: MaterialItemAction) => void
}

export type { MaterialMediaKind, MaterialItemAction }
