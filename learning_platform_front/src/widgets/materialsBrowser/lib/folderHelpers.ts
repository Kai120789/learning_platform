import {
    buildMaterialFolders,
    mockMaterials,
    type MaterialFileMock,
    type MaterialFolderMock,
} from "@/shared/mocks"

export const materialFolders = buildMaterialFolders(mockMaterials)

export function getFolderLabel(
    folder: MaterialFolderMock,
    t: (key: string, options?: Record<string, unknown>) => string
) {
    if (folder.kind === "lesson" && folder.lessonId != null) {
        return t("materials.lesson", { id: folder.lessonId })
    }
    return folder.name
}

export function getFilesInFolder(folderId: string | null): MaterialFileMock[] {
    if (!folderId) return []

    const folder = materialFolders.find((item) => item.id === folderId)
    if (!folder || folder.kind !== "lesson" || folder.lessonId == null) {
        return []
    }

    const parts = folder.id.split(":")
    const subjectTitle = parts[1]
    const groupTitle = parts[2]

    return mockMaterials.filter(
        (file) =>
            file.lessonId === folder.lessonId
            && file.subjectTitle === subjectTitle
            && file.groupTitle === groupTitle
    )
}

export function getBreadcrumbs(currentFolderId: string | null): MaterialFolderMock[] {
    const trail: MaterialFolderMock[] = []
    let cursor = currentFolderId

    while (cursor) {
        const folder = materialFolders.find((item) => item.id === cursor)
        if (!folder) break
        trail.unshift(folder)
        cursor = folder.parentId
    }

    return trail
}

export function getCurrentFolders(currentFolderId: string | null) {
    return materialFolders.filter((folder) => folder.parentId === currentFolderId)
}
