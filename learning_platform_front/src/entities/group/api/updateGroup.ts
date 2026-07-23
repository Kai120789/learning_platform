import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { UpdateGroupRequest } from '../types/updateGroup';
import type { GroupResponse } from '../types/types';

export const updateGroup = createAsyncThunk<
    GroupResponse,
    { request: UpdateGroupRequest, groupID: number },
    { rejectValue: string }
>(
    'updateGroup',
    async ({ request, groupID }, { rejectWithValue }) => {
        try {
            const response = await $api.patch<GroupResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group/${groupID}`,
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