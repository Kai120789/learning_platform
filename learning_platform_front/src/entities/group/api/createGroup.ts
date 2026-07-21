import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { GroupResponse } from '../types/types';
import type { CreateGroupRequest } from '../types/createGroup';

export const createGroup = createAsyncThunk<
    GroupResponse,
    CreateGroupRequest,
    { rejectValue: string }
>(
    'createGroup',
    async (request: CreateGroupRequest, { rejectWithValue }) => {
        try {
            const response = await $api.post<GroupResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group`,
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