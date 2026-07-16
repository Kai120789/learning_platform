import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';

export const updateUserAvatar = createAsyncThunk<
    undefined,
    string,
    { rejectValue: string }
>(
    'updateUserAvatar',
    async (avatar, { rejectWithValue }) => {
        try {
            const response = await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/user/avatar?avatar=${avatar}`,
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