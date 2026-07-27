import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { ShortUserInfo } from '../types/types';

export const addUsersToGroup = createAsyncThunk<
    number[],
    {
        groupID: number,
        users: ShortUserInfo[]
    },
    { rejectValue: string }
>(
    'addUsersToGroup',
    async ({ groupID, users }, { rejectWithValue }) => {
        try {
            const userIDs = users.map((user) => user.id)
            const response = await $api.post(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group/${groupID}/add-user`,
                userIDs,
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
