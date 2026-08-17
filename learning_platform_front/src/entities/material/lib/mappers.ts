import type {
    MaterialData,
    MaterialFolderData,
    MaterialFolderResponse,
    MaterialResponse,
} from "../types/types"

export function mapFolderResponse(folder: MaterialFolderResponse): MaterialFolderData {
    return {
        id: folder.id,
        title: folder.title,
        parentFolderId: folder.parent_folder_id,
        tutorId: folder.tutor_id,
    }
}

export function mapMaterialResponse(material: MaterialResponse): MaterialData {
    return {
        id: material.id,
        title: material.title,
        size: material.size,
        folderId: material.folder_id,
        tutorId: material.tutor_id,
        mimeType: material.mime_type,
        mediaObjectId: material.media_object_id,
    }
}
