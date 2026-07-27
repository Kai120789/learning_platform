import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';

export const addUsersToGroup = createAsyncThunk<
    number[],
    {
        groupID: number,
        userIDs: number[]
    },
    { rejectValue: string }
>(
    'addUsersToGroup',
    async ({ groupID, userIDs }, { rejectWithValue }) => {
        try {
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
