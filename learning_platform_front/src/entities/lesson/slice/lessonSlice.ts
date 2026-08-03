import { createSlice } from "@reduxjs/toolkit"
import type { LessonSchema } from "../types/types"
import { mapLessonResponse } from "../lib/mapLesson"
import { getLessonById } from "@/entities/lesson"
import { getLessonsByStudentId } from "@/entities/lesson"
import { getLessonsByTutorId } from "@/entities/lesson"
import { createLesson } from "@/entities/lesson"
import { updateLesson } from "@/entities/lesson"
import { updateLessonStatus } from "@/entities/lesson"

const initialState: LessonSchema = {
    data: null,
    isLoading: false,
    error: undefined,
}

function upsertLesson(
    state: LessonSchema,
    mapped: ReturnType<typeof mapLessonResponse>,
) {
    if (!state.data) {
        state.data = [mapped]
        return
    }
    const index = state.data.findIndex((item) => item.id === mapped.id)
    if (index >= 0) {
        state.data[index] = mapped
    } else {
        state.data.push(mapped)
    }
}

const lessonSlice = createSlice({
    name: "lesson",
    initialState,
    reducers: {
        clearLessons: (state) => {
            state.data = null
            state.error = undefined
            state.isLoading = false
        },
    },
    extraReducers: (builder) => {
        builder.addCase(getLessonsByStudentId.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(getLessonsByStudentId.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
            if (state.data === null) state.data = []
        })
        builder.addCase(getLessonsByStudentId.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.data = (action.payload ?? []).map(mapLessonResponse)
        })

        builder.addCase(getLessonsByTutorId.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(getLessonsByTutorId.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
            if (state.data === null) state.data = []
        })
        builder.addCase(getLessonsByTutorId.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.data = (action.payload ?? []).map(mapLessonResponse)
        })

        builder.addCase(getLessonById.fulfilled, (state, action) => {
            state.error = ""
            upsertLesson(state, mapLessonResponse(action.payload))
        })
        builder.addCase(getLessonById.rejected, (state, action) => {
            state.error = action.payload as string
        })

        builder.addCase(createLesson.fulfilled, (state, action) => {
            state.error = ""
            if (!state.data) state.data = []
            state.data.push(mapLessonResponse(action.payload))
        })
        builder.addCase(createLesson.rejected, (state, action) => {
            state.error = action.payload as string
        })

        builder.addCase(updateLesson.fulfilled, (state, action) => {
            state.error = ""
            upsertLesson(state, mapLessonResponse(action.payload))
        })
        builder.addCase(updateLesson.rejected, (state, action) => {
            state.error = action.payload as string
        })

        builder.addCase(updateLessonStatus.fulfilled, (state, action) => {
            state.error = ""
            const lesson = state.data?.find((item) => item.id === action.payload.lessonID)
            if (lesson) {
                lesson.status = action.payload.status
            }
        })
        builder.addCase(updateLessonStatus.rejected, (state, action) => {
            state.error = action.payload as string
        })
    },
})

export const { actions: lessonActions, reducer: lessonReducer } = lessonSlice
