import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';

export const removeUserFromGroup = createAsyncThunk<
    void,
    {
        groupID: number,
        userID: number
    },
    { rejectValue: string }
>(
    'removeUserFromGroup',
    async ({ groupID, userID }, { rejectWithValue }) => {
        try {
            const response = await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group/${groupID}/remove-user/${userID}`,
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