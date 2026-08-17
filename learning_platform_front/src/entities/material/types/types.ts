export interface MaterialSchema {
    folders: MaterialFolderData[]
    materials: MaterialData[]
    currentFolderId: number | null
    isLoading: boolean
    error?: string
}

export type MaterialFolderData = {
    id: number
    title: string
    parentFolderId: number | null
    tutorId: number
}

export type MaterialData = {
    id: number
    title: string
    size: number
    folderId: number | null
    tutorId: number
    mimeType: string
    mediaObjectId: string
}

export type MaterialFolderResponse = {
    id: number
    title: string
    parent_folder_id: number | null
    tutor_id: number
}

export type MaterialResponse = {
    id: number
    title: string
    size: number
    folder_id: number | null
    tutor_id: number
    mime_type: string
    media_object_id: string
}

export type FoldersAndMaterialsResponse = {
    folders: MaterialFolderResponse[] | null
    materials: MaterialResponse[] | null
}

export type MaterialMediaKind = "VIDEO" | "IMAGE" | "DOCUMENT" | "LINK"
