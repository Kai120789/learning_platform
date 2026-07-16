import { createSlice } from "@reduxjs/toolkit";
import type { UserSchema } from "../types/types";
import { getUserData } from "@/entities/user";

const initialState: UserSchema = {
    data: null,
    isLoading: false,
    error: undefined
};

const userSlice = createSlice({
    name: 'notifications',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(getUserData.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getUserData.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getUserData.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ''
            state.data = {
                user: {
                    userID: action.payload.user_id,
                    email: action.payload.email,
                    role: action.payload.role,
                    status: action.payload.status,
                },
                userInfo: {
                    name: action.payload.user_info.name,
                    surname: action.payload.user_info.surname,
                    patronymic: action.payload.user_info.patronymic,
                    city: action.payload.user_info.city,
                    about: action.payload.user_info.about,
                    avatar: action.payload.user_info.avatar,
                    gender: action.payload.user_info.gender,
                    birthDate: action.payload.user_info.birth_date
                        ? new Date(action.payload.user_info.birth_date)
                        : undefined,
                },
                userSettings: {
                    is2FaEnabled: action.payload.user_settings.is_2_fa_enabled,
                    isNotificationsEnabled: action.payload.user_settings.is_notifications_enabled,
                    language: action.payload.user_settings.language,
                    theme: action.payload.user_settings.theme
                }
            }
        })
    }
});

export const { actions: userActions, reducer: userReducer } =
    userSlice;