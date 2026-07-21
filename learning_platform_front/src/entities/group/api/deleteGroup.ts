import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';

export const deleteGroup = createAsyncThunk<
    void,
    number,
    { rejectValue: string }
>(
    'deleteGroup',
    async (groupID: number, { rejectWithValue }) => {
        try {
            const response = await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/group/${groupID}`,
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