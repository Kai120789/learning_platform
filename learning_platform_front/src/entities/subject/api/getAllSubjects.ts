import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { SubjectData } from '../types/types';

export const getAllSubjects = createAsyncThunk<
    SubjectData[],
    void,
    { rejectValue: string }
>(
    'getAllSubjects',
    async (_, { rejectWithValue }) => {
        try {
            const response = await $api.get<SubjectData[]>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/subject`,
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