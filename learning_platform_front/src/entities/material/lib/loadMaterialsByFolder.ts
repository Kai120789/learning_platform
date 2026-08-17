import { getMaterials } from "../api/getMaterials"

export function loadMaterialsByFolder(folderId: number | null) {
    return getMaterials(folderId ?? 0)
}
