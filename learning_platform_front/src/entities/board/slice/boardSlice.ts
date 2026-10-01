import { createSlice } from "@reduxjs/toolkit"
import { createBoard, deleteBoard, getTutorBoards, updateBoard } from "../api/board"
import type { Board, BoardResponse, BoardSchema } from "../model/types"

const initialState: BoardSchema = {
    data: null,
    isLoading: false,
    error: undefined,
}

function mapBoard(board: BoardResponse): Board {
    return {
        id: board.id,
        title: board.title,
        tutorId: board.tutor_id,
    }
}

const boardSlice = createSlice({
    name: "board",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(getTutorBoards.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(getTutorBoards.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload
            state.data = state.data ?? []
        })
        builder.addCase(getTutorBoards.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.data = action.payload.map(mapBoard)
        })
        builder.addCase(createBoard.fulfilled, (state, action) => {
            state.error = ""
            const next = mapBoard(action.payload)
            state.data = state.data ? [...state.data, next] : [next]
        })
        builder.addCase(updateBoard.fulfilled, (state, action) => {
            state.error = ""
            const next = mapBoard(action.payload)
            state.data = state.data?.map((board) => board.id === next.id ? next : board) ?? null
        })
        builder.addCase(deleteBoard.fulfilled, (state, action) => {
            state.error = ""
            state.data = state.data?.filter((board) => board.id !== action.payload) ?? null
        })
    },
})

export const { actions: boardActions, reducer: boardReducer } = boardSlice
