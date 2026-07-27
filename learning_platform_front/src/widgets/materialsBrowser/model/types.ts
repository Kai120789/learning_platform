import type { MaterialFileMock, MaterialFolderMock, MediaType } from "@/shared/mocks"

export type MaterialsViewMode = "list" | "grid"

export type MaterialsViewProps = {
    folders: MaterialFolderMock[]
    files: MaterialFileMock[]
    onOpenFolder: (folderId: string) => void
    onOpenFile: (file: MaterialFileMock) => void
    getFolderLabel: (folder: MaterialFolderMock) => string
}

export type { MaterialFileMock, MaterialFolderMock, MediaType }
