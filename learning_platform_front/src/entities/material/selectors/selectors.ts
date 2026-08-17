import type { StateSchema } from "@/app/providers/storeProvider"

export const getMaterialFolders = (state: StateSchema) => state.material.folders
export const getMaterialFiles = (state: StateSchema) => state.material.materials
export const getMaterialCurrentFolderId = (state: StateSchema) => state.material.currentFolderId
export const getMaterialIsLoading = (state: StateSchema) => state.material.isLoading
export const getMaterialError = (state: StateSchema) => state.material.error
