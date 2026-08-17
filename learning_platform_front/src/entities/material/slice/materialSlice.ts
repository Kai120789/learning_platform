import { createSlice } from "@reduxjs/toolkit"
import type { MaterialSchema } from "../types/types"
import { mapFolderResponse, mapMaterialResponse } from "../lib/mappers"
import { getMaterials } from "../api/getMaterials"
import { createFolder } from "../api/createFolder"
import { moveFolders } from "../api/moveFolders"
import { renameFolder } from "../api/renameFolder"
import { deleteOneFolder } from "@/entities/material"
import { deleteFolders } from "@/entities/material"
import { createMaterials } from "@/entities/material"
import { moveMaterials } from "@/entities/material"
import { renameMaterial } from "@/entities/material"
import { deleteOneMaterial } from "@/entities/material"
import { deleteMaterials } from "@/entities/material"
import { updateUsersMaterialsAccess } from "@/entities/material"
import { updateUsersFoldersAccess } from "@/entities/material"

const initialState: MaterialSchema = {
    folders: [],
    materials: [],
    currentFolderId: null,
    isLoading: false,
    error: undefined,
}

const materialSlice = createSlice({
    name: "material",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(getMaterials.pending, (state, action) => {
            state.isLoading = true
            state.error = ""
            state.currentFolderId = action.meta.arg === 0 ? null : action.meta.arg
        })
        builder.addCase(getMaterials.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getMaterials.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.folders = (action.payload.folders ?? []).map(mapFolderResponse)
            state.materials = (action.payload.materials ?? []).map(mapMaterialResponse)
        })

        builder.addCase(createFolder.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(createFolder.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(createFolder.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const folder = mapFolderResponse(action.payload)
            if (folder.parentFolderId === state.currentFolderId) {
                state.folders.push(folder)
            }
        })

        builder.addCase(moveFolders.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(moveFolders.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(moveFolders.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const ids = new Set(action.payload.folder_ids)
            state.folders = state.folders.filter((folder) => !ids.has(folder.id))
        })

        builder.addCase(renameFolder.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(renameFolder.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(renameFolder.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const folder = state.folders.find((item) => item.id === action.payload.folderID)
            if (folder) {
                folder.title = action.payload.title
            }
        })

        builder.addCase(deleteOneFolder.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(deleteOneFolder.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteOneFolder.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.folders = state.folders.filter((folder) => folder.id !== action.payload)
        })

        builder.addCase(deleteFolders.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(deleteFolders.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteFolders.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const ids = new Set(action.payload)
            state.folders = state.folders.filter((folder) => !ids.has(folder.id))
        })

        builder.addCase(createMaterials.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(createMaterials.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(createMaterials.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const targetFolderId = action.meta.arg.folderID === 0
                ? null
                : action.meta.arg.folderID
            if (targetFolderId === state.currentFolderId) {
                state.materials.push(...action.payload.map(mapMaterialResponse))
            }
        })

        builder.addCase(moveMaterials.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(moveMaterials.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(moveMaterials.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const ids = new Set(action.payload.material_ids)
            state.materials = state.materials.filter((material) => !ids.has(material.id))
        })

        builder.addCase(renameMaterial.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(renameMaterial.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(renameMaterial.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const material = state.materials.find((item) => item.id === action.payload.materialID)
            if (material) {
                material.title = action.payload.title
            }
        })

        builder.addCase(deleteOneMaterial.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(deleteOneMaterial.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteOneMaterial.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.materials = state.materials.filter((material) => material.id !== action.payload)
        })

        builder.addCase(deleteMaterials.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(deleteMaterials.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteMaterials.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const ids = new Set(action.payload)
            state.materials = state.materials.filter((material) => !ids.has(material.id))
        })

        builder.addCase(updateUsersMaterialsAccess.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(updateUsersMaterialsAccess.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateUsersMaterialsAccess.fulfilled, (state) => {
            state.isLoading = false
            state.error = ""
        })

        builder.addCase(updateUsersFoldersAccess.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(updateUsersFoldersAccess.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateUsersFoldersAccess.fulfilled, (state) => {
            state.isLoading = false
            state.error = ""
        })
    },
})

export const { actions: materialActions, reducer: materialReducer } = materialSlice
