import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';
import type { GetUsersWithPaginationRequest, GetUsersWithPaginationResponse } from '../types/usersWithPagination';

export const getUsersWithPagination = createAsyncThunk<
    GetUsersWithPaginationResponse,
    GetUsersWithPaginationRequest,
    { rejectValue: string }
>(
    'getUsersWithPagination',
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.get<GetUsersWithPaginationResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/user/`,
                {
                    params: {
                        search: request.search,
                        page: request.page,
                        limit: request.limit,
                        role: request.role,
                    },
                },
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
