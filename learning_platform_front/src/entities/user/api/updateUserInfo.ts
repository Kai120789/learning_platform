import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import type { UserInfoRequest, UserInfoResponse } from '../types/types';
import axios from 'axios';

export const updateUserInfo = createAsyncThunk<
    UserInfoResponse,
    UserInfoRequest,
    { rejectValue: string }
>(
    'updateUserInfo',
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.put<UserInfoResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/user/info`,
                request
            )

            return response.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(
                    error.response?.data ?? error.message
                );
            }

            return rejectWithValue("Неизвестная ошибка");
        }
    }
)