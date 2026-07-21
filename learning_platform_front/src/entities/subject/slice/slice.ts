import { createSlice } from "@reduxjs/toolkit";
import type { SubjectSchema } from "../types/types";
import { getAllSubjects } from "../api/getAllSubjects";

const initialState: SubjectSchema = {
    data: null,
    isLoading: false,
    error: undefined
};

const subjectSlice = createSlice({
    name: 'subject',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(getAllSubjects.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getAllSubjects.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getAllSubjects.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = action.payload
        })
    }
});

export const { actions: subjectActions, reducer: subjectReducer } =
    subjectSlice;