import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { GroupResponse } from '../types/types';

export const getGroupsByStudentId = createAsyncThunk<
    GroupResponse[],
    void,
    { rejectValue: string }
>(
    'getGroupsByStudentId',
    async (_, { rejectWithValue }) => {
        try {
            const response = await $api.get<GroupResponse[]>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group/student`,
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