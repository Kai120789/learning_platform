import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { UserThemeEnum } from '@/shared/enums/user';

export const updateUserTheme = createAsyncThunk<
    undefined,
    UserThemeEnum,
    { rejectValue: string }
>(
    'updateUserTheme',
    async (theme, { rejectWithValue }) => {
        try {
            const response = await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/user/theme?theme=${theme}`,
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