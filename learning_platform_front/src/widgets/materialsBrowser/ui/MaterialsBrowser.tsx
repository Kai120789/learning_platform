import { useEffect, useMemo, useRef, useState } from "react"
import { useTranslation } from "react-i18next"
import { useSearchParams } from "react-router-dom"
import { ClipboardPaste, FolderPlus, LayoutGrid, List, Trash2, Upload, X } from "lucide-react"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    createMaterials,
    deleteFolders,
    deleteMaterials,
    deleteOneFolder,
    deleteOneMaterial,
    getMaterialFiles,
    getMaterialFolders,
    getMaterialIsLoading,
    loadMaterialsByFolder,
    moveFolders,
    moveMaterials,
    renameFolder,
    renameMaterial,
    type MaterialData,
    type MaterialFolderData,
} from "@/entities/material"
import { useCanEdit } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { cn } from "@/shared/lib/utils"
import type {
    MaterialItemAction,
    MaterialBreadcrumb,
    MaterialsMovingItem,
    MaterialsSelection,
    MaterialsTarget,
    MaterialsViewMode,
} from "../model/types"
import {
    getMaterialFileUrl,
    parseFolderIdParam,
    readStoredBreadcrumbs,
    readStoredViewMode,
    storeBreadcrumbs,
    storeViewMode,
} from "../lib/materialsNavigation"
import { CreateFolderDialog } from "./CreateFolderDialog"
import { MaterialsBreadcrumbs } from "./MaterialsBreadcrumbs"
import { MaterialsDeleteDialog } from "./MaterialsDeleteDialog"
import { MaterialsDropOverlay } from "./MaterialsDropOverlay"
import { MaterialsEmptyDropzone } from "./MaterialsEmptyDropzone"
import { MaterialsGridView } from "./MaterialsGridView"
import { MaterialsInfoDialog } from "./MaterialsInfoDialog"
import { MaterialsListView } from "./MaterialsListView"
import { MaterialsRenameDialog } from "./MaterialsRenameDialog"

export function MaterialsBrowser() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const [searchParams, setSearchParams] = useSearchParams()
    const folders = useAppSelector(getMaterialFolders)
    const files = useAppSelector(getMaterialFiles)
    const isLoading = useAppSelector(getMaterialIsLoading)
    const canEdit = useCanEdit()

    const currentFolderId = useMemo(
        () => parseFolderIdParam(searchParams.get("folder")),
        [searchParams],
    )

    const [breadcrumbs, setBreadcrumbs] = useState<MaterialBreadcrumb[]>(() => (
        readStoredBreadcrumbs(parseFolderIdParam(new URLSearchParams(window.location.search).get("folder")))
    ))
    const [viewMode, setViewMode] = useState<MaterialsViewMode>(() => readStoredViewMode())
    const [createFolderOpen, setCreateFolderOpen] = useState(false)
    const [isDragging, setIsDragging] = useState(false)
    const [movingItem, setMovingItem] = useState<MaterialsMovingItem | null>(null)
    const [infoTarget, setInfoTarget] = useState<MaterialsTarget | null>(null)
    const [renameTarget, setRenameTarget] = useState<MaterialsTarget | null>(null)
    const [deleteNames, setDeleteNames] = useState<string[]>([])
    const [deleteTargets, setDeleteTargets] = useState<MaterialsTarget[]>([])
    const [selectionMode, setSelectionMode] = useState(false)
    const [selection, setSelection] = useState<MaterialsSelection>({
        folders: new Set(),
        materials: new Set(),
    })
    const fileInputRef = useRef<HTMLInputElement>(null)
    const dragDepth = useRef(0)

    const clearSelection = () => {
        setSelectionMode(false)
        setSelection({ folders: new Set(), materials: new Set() })
    }

    useEffect(() => {
        dispatch(loadMaterialsByFolder(currentFolderId))
    }, [dispatch, currentFolderId])

    useEffect(() => {
        storeViewMode(viewMode)
    }, [viewMode])

    useEffect(() => {
        const stored = readStoredBreadcrumbs(currentFolderId)
        setBreadcrumbs(stored)
        clearSelection()
    }, [currentFolderId])

    const navigateToFolder = (
        folderId: number | null,
        nextBreadcrumbs: MaterialBreadcrumb[],
        replace = false,
    ) => {
        setBreadcrumbs(nextBreadcrumbs)
        storeBreadcrumbs(nextBreadcrumbs)

        const params = new URLSearchParams()
        if (folderId != null) {
            params.set("folder", String(folderId))
        }
        setSearchParams(params, { replace })
    }

    const isEmpty = !isLoading && folders.length === 0 && files.length === 0

    const openFolder = (folder: MaterialFolderData) => {
        navigateToFolder(folder.id, [...breadcrumbs, { id: folder.id, title: folder.title }])
    }

    const goToBreadcrumb = (folderId: number | null, index: number) => {
        navigateToFolder(
            folderId,
            index < 0 ? [] : breadcrumbs.slice(0, index + 1),
        )
    }

    const openFile = (file: MaterialData) => {
        window.open(getMaterialFileUrl(file.mediaObjectId), "_blank", "noreferrer")
    }

    const uploadFiles = async (selected: File[]) => {
        if (selected.length === 0) return

        const response = await dispatch(createMaterials({
            folderID: currentFolderId ?? 0,
            files: selected,
        }))

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("materials.uploadSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("materials.uploadError"),
                type: "error",
            }))
        }
    }

    const onUploadClick = () => {
        fileInputRef.current?.click()
    }

    const onFilesSelected = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const selected = Array.from(e.target.files ?? [])
        e.target.value = ""
        await uploadFiles(selected)
    }

    const onDragEnter = (e: React.DragEvent) => {
        if (!canEdit) return
        e.preventDefault()
        e.stopPropagation()
        dragDepth.current += 1
        if (e.dataTransfer.types.includes("Files")) {
            setIsDragging(true)
        }
    }

    const onDragLeave = (e: React.DragEvent) => {
        if (!canEdit) return
        e.preventDefault()
        e.stopPropagation()
        dragDepth.current -= 1
        if (dragDepth.current <= 0) {
            dragDepth.current = 0
            setIsDragging(false)
        }
    }

    const onDragOver = (e: React.DragEvent) => {
        if (!canEdit) return
        e.preventDefault()
        e.stopPropagation()
        e.dataTransfer.dropEffect = "copy"
    }

    const onDrop = async (e: React.DragEvent) => {
        if (!canEdit) return
        e.preventDefault()
        e.stopPropagation()
        dragDepth.current = 0
        setIsDragging(false)
        await uploadFiles(Array.from(e.dataTransfer.files))
    }

    const onToggleSelect = (target: MaterialsTarget) => {
        setSelection((prev) => {
            const folders = new Set(prev.folders)
            const materials = new Set(prev.materials)
            if (target.kind === "folder") {
                if (folders.has(target.folder.id)) folders.delete(target.folder.id)
                else folders.add(target.folder.id)
            } else {
                if (materials.has(target.file.id)) materials.delete(target.file.id)
                else materials.add(target.file.id)
            }
            return { folders, materials }
        })
    }

    const selectedCount = selection.folders.size + selection.materials.size

    const onItemAction = async (target: MaterialsTarget, action: MaterialItemAction) => {
        if (action === "info") {
            setInfoTarget(target)
            return
        }
        if (action === "download") {
            if (target.kind !== "material") return
            console.log("скачивание...", target.file.title, target.file.mediaObjectId)
            return
        }
        if (action === "select") {
            setSelectionMode(true)
            onToggleSelect(target)
            return
        }
        if (action === "rename") {
            setRenameTarget(target)
            return
        }
        if (action === "move") {
            setMovingItem(target)
            return
        }
        if (action === "delete") {
            setDeleteTargets([target])
            setDeleteNames([
                target.kind === "folder" ? target.folder.title : target.file.title,
            ])
        }
    }

    const onDeleteSelectedClick = () => {
        if (selectedCount === 0) return

        const targets: MaterialsTarget[] = []
        const names: string[] = []

        folders.forEach((folder) => {
            if (selection.folders.has(folder.id)) {
                targets.push({ kind: "folder", folder })
                names.push(folder.title)
            }
        })
        files.forEach((file) => {
            if (selection.materials.has(file.id)) {
                targets.push({ kind: "material", file })
                names.push(file.title)
            }
        })

        setDeleteTargets(targets)
        setDeleteNames(names)
    }

    const onDeleteConfirm = async () => {
        if (deleteTargets.length === 0) return

        const folderIds = deleteTargets
            .filter((item) => item.kind === "folder")
            .map((item) => item.folder.id)
        const materialIds = deleteTargets
            .filter((item) => item.kind === "material")
            .map((item) => item.file.id)

        const results = await Promise.all([
            folderIds.length === 1
                ? dispatch(deleteOneFolder(folderIds[0]))
                : folderIds.length > 1
                    ? dispatch(deleteFolders(folderIds))
                    : null,
            materialIds.length === 1
                ? dispatch(deleteOneMaterial(materialIds[0]))
                : materialIds.length > 1
                    ? dispatch(deleteMaterials(materialIds))
                    : null,
        ])

        const failed = results.some(
            (result) => result != null && result.meta.requestStatus !== "fulfilled",
        )

        if (!failed) {
            if (
                movingItem
                && (
                    (movingItem.kind === "folder" && folderIds.includes(movingItem.folder.id))
                    || (movingItem.kind === "material" && materialIds.includes(movingItem.file.id))
                )
            ) {
                setMovingItem(null)
            }
            setDeleteTargets([])
            setDeleteNames([])
            clearSelection()
            dispatch(notificationActions.addNotification({
                message: t("materials.deleteSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("materials.deleteError"),
                type: "error",
            }))
        }
    }

    const onRenameSubmit = async (title: string) => {
        if (!renameTarget) return

        const response = renameTarget.kind === "folder"
            ? await dispatch(renameFolder({ folderID: renameTarget.folder.id, title }))
            : await dispatch(renameMaterial({ materialID: renameTarget.file.id, title }))

        if (response.meta.requestStatus === "fulfilled") {
            if (movingItem?.kind === "folder" && renameTarget.kind === "folder" && movingItem.folder.id === renameTarget.folder.id) {
                setMovingItem({ kind: "folder", folder: { ...movingItem.folder, title } })
            }
            if (movingItem?.kind === "material" && renameTarget.kind === "material" && movingItem.file.id === renameTarget.file.id) {
                setMovingItem({ kind: "material", file: { ...movingItem.file, title } })
            }
            if (renameTarget.kind === "folder") {
                const nextBreadcrumbs = breadcrumbs.map((crumb) => (
                    crumb.id === renameTarget.folder.id ? { ...crumb, title } : crumb
                ))
                navigateToFolder(currentFolderId, nextBreadcrumbs, true)
            }
            setRenameTarget(null)
            dispatch(notificationActions.addNotification({
                message: t("materials.renameSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("materials.renameError"),
                type: "error",
            }))
        }
    }

    const onPaste = async () => {
        if (!movingItem) return

        const response = movingItem.kind === "folder"
            ? await dispatch(moveFolders({
                folder_ids: [movingItem.folder.id],
                parent_folder_id: currentFolderId,
            }))
            : await dispatch(moveMaterials({
                material_ids: [movingItem.file.id],
                folder_id: currentFolderId,
            }))

        if (response.meta.requestStatus === "fulfilled") {
            setMovingItem(null)
            await dispatch(loadMaterialsByFolder(currentFolderId))
            dispatch(notificationActions.addNotification({
                message: t("materials.moveSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("materials.moveError"),
                type: "error",
            }))
        }
    }

    const renameInitialTitle = renameTarget?.kind === "folder"
        ? renameTarget.folder.title
        : renameTarget?.file.title ?? ""

    return (
        <div className="space-y-6">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-center justify-between">
                <MaterialsBreadcrumbs
                    breadcrumbs={breadcrumbs}
                    currentFolderId={currentFolderId}
                    onNavigate={goToBreadcrumb}
                />
                <div className="flex flex-wrap items-center gap-2">
                    {canEdit && selectionMode && (
                        <>
                            <Button
                                type="button"
                                size="sm"
                                variant="destructive"
                                disabled={selectedCount === 0}
                                onClick={onDeleteSelectedClick}
                            >
                                <Trash2 className="size-3.5" />
                                {t("materials.deleteSelected", { count: selectedCount })}
                            </Button>
                            <Button
                                type="button"
                                size="sm"
                                variant="outline"
                                onClick={clearSelection}
                            >
                                <X className="size-3.5" />
                                {t("materials.cancelMove")}
                            </Button>
                        </>
                    )}
                    {canEdit && movingItem && !selectionMode && (
                        <>
                            <Button type="button" size="sm" onClick={onPaste}>
                                <ClipboardPaste className="size-3.5" />
                                {t("materials.paste")}
                            </Button>
                            <Button
                                type="button"
                                size="sm"
                                variant="outline"
                                onClick={() => setMovingItem(null)}
                            >
                                <X className="size-3.5" />
                                {t("materials.cancelMove")}
                            </Button>
                        </>
                    )}
                    {canEdit && !selectionMode && (
                        <>
                            <Button
                                type="button"
                                size="sm"
                                variant="outline"
                                onClick={() => setCreateFolderOpen(true)}
                            >
                                <FolderPlus className="size-3.5" />
                                {t("materials.createFolder")}
                            </Button>
                            <Button type="button" size="sm" onClick={onUploadClick}>
                                <Upload className="size-3.5" />
                                {t("materials.upload")}
                            </Button>
                            <input
                                ref={fileInputRef}
                                type="file"
                                multiple
                                className="hidden"
                                onChange={onFilesSelected}
                            />
                        </>
                    )}
                    <div className="flex rounded-lg border border-border p-0.5 bg-secondary/60">
                        <button
                            type="button"
                            onClick={() => setViewMode("list")}
                            className={cn(
                                "flex cursor-pointer items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
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
                                "flex cursor-pointer items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
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
            </div>

            <div
                className="relative min-h-56"
                onDragEnter={onDragEnter}
                onDragLeave={onDragLeave}
                onDragOver={onDragOver}
                onDrop={onDrop}
            >
                <div
                    className={cn(
                        "min-h-56 transition-opacity duration-200",
                        isDragging ? "opacity-0" : "opacity-100"
                    )}
                >
                    {isLoading ? (
                        <div className="flex min-h-56 items-center justify-center rounded-xl border border-dashed text-muted-foreground">
                            {t("materials.loading")}
                        </div>
                    ) : isEmpty ? (
                        <MaterialsEmptyDropzone canUpload={canEdit} />
                    ) : (
                        <div className="min-h-56">
                            {viewMode === "list" ? (
                                <MaterialsListView
                                    folders={folders}
                                    files={files}
                                    canEdit={canEdit}
                                    selectionMode={selectionMode}
                                    selection={selection}
                                    movingItem={movingItem}
                                    onOpenFolder={openFolder}
                                    onOpenFile={openFile}
                                    onToggleSelect={onToggleSelect}
                                    onItemAction={onItemAction}
                                />
                            ) : (
                                <MaterialsGridView
                                    folders={folders}
                                    files={files}
                                    canEdit={canEdit}
                                    selectionMode={selectionMode}
                                    selection={selection}
                                    movingItem={movingItem}
                                    onOpenFolder={openFolder}
                                    onOpenFile={openFile}
                                    onToggleSelect={onToggleSelect}
                                    onItemAction={onItemAction}
                                />
                            )}
                        </div>
                    )}
                </div>

                {canEdit && <MaterialsDropOverlay visible={isDragging} />}
            </div>

            {canEdit && (
                <CreateFolderDialog
                    open={createFolderOpen}
                    onOpenChange={setCreateFolderOpen}
                    parentFolderId={currentFolderId}
                />
            )}

            <MaterialsInfoDialog
                open={infoTarget != null}
                onOpenChange={(open) => {
                    if (!open) setInfoTarget(null)
                }}
                target={infoTarget}
            />

            <MaterialsRenameDialog
                open={renameTarget != null}
                onOpenChange={(open) => {
                    if (!open) setRenameTarget(null)
                }}
                initialTitle={renameInitialTitle}
                onSubmit={onRenameSubmit}
            />

            <MaterialsDeleteDialog
                open={deleteNames.length > 0}
                onOpenChange={(open) => {
                    if (!open) {
                        setDeleteNames([])
                        setDeleteTargets([])
                    }
                }}
                names={deleteNames}
                onConfirm={onDeleteConfirm}
            />
        </div>
    )
}
