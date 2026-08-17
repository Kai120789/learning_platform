import type { MaterialBreadcrumb, MaterialsViewMode } from "../model/types"

export const MATERIALS_VIEW_MODE_KEY = "materials-view-mode"
export const MATERIALS_BREADCRUMBS_KEY = "materials-breadcrumbs"

export function getMaterialFileUrl(mediaObjectId: string): string {
    const base = import.meta.env.VITE_S3_URL ?? ""
    if (!base) return mediaObjectId
    return `${base}${mediaObjectId}`
}

export function parseFolderIdParam(value: string | null): number | null {
    if (!value) return null
    const id = Number(value)
    if (!Number.isFinite(id) || id <= 0) return null
    return id
}

export function readStoredBreadcrumbs(folderId: number | null): MaterialBreadcrumb[] {
    if (folderId == null) return []
    try {
        const raw = window.sessionStorage.getItem(MATERIALS_BREADCRUMBS_KEY)
        if (!raw) return []
        const parsed = JSON.parse(raw) as unknown
        if (!Array.isArray(parsed)) return []
        const crumbs = parsed.filter((item): item is MaterialBreadcrumb => (
            typeof item === "object"
            && item != null
            && typeof (item as MaterialBreadcrumb).id === "number"
            && typeof (item as MaterialBreadcrumb).title === "string"
        ))
        const index = crumbs.findIndex((crumb) => crumb.id === folderId)
        if (index === -1) return []
        return crumbs.slice(0, index + 1)
    } catch {
        return []
    }
}

export function storeBreadcrumbs(breadcrumbs: MaterialBreadcrumb[]) {
    window.sessionStorage.setItem(MATERIALS_BREADCRUMBS_KEY, JSON.stringify(breadcrumbs))
}

export function readStoredViewMode(): MaterialsViewMode {
    const value = window.localStorage.getItem(MATERIALS_VIEW_MODE_KEY)
    return value === "grid" || value === "list" ? value : "list"
}

export function storeViewMode(viewMode: MaterialsViewMode) {
    window.localStorage.setItem(MATERIALS_VIEW_MODE_KEY, viewMode)
}
