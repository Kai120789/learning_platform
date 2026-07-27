import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { UserSchema } from "../types/types";
import { getUserData } from "@/entities/user";
import { login } from "@/features/loginForm";
import { register } from "@/features/registerForm/api/register";
import { logout } from "@/widgets/dropdownMenu/api/logout";

const initialState: UserSchema = {
    data: null,
    isAuth: false,
    isInitialized: false,
    isLoading: false,
    error: undefined
};

const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setIsAuth: (state, action: PayloadAction<boolean>) => {
            state.isAuth = action.payload
        }
    },
    extraReducers: (builder) => {
        builder.addCase(getUserData.pending, (state) => {
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(getUserData.rejected, (state, action) => {
            state.isLoading = false
            state.isInitialized = true
            state.isAuth = false
            state.error = action.payload as string
        })
        builder.addCase(getUserData.fulfilled, (state, action) => {
            state.isLoading = false
            state.isInitialized = true
            state.error = ''
            state.isAuth = true
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
                    tgLink: action.payload.user_info.tg_link,
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
        builder.addCase(login.pending, (state) => {
            state.isAuth = false
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(login.rejected, (state, action) => {
            state.isAuth = false
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(login.fulfilled, (state) => {
            state.isAuth = true
            state.isInitialized = true
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(register.pending, (state) => {
            state.isAuth = false
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(register.rejected, (state, action) => {
            state.isAuth = false
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(register.fulfilled, (state) => {
            state.isAuth = true
            state.isInitialized = true
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(logout.pending, (state) => {
            state.isAuth = false
            state.isLoading = true
            state.error = ''
        })
        builder.addCase(logout.rejected, (state, action) => {
            state.isAuth = false
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(logout.fulfilled, (state) => {
            state.isAuth = false
            state.isLoading = true
            state.error = ''
        })
    }
});

export const { actions: userActions, reducer: userReducer } =
    userSlice;