import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { GroupResponse } from '../types/types';

export const getUserGroups = createAsyncThunk<
    GroupResponse[],
    void,
    { rejectValue: string }
>(
    'getUserGroups',
    async (_, { rejectWithValue }) => {
        try {
            const response = await $api.get<GroupResponse[]>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group/user`,
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