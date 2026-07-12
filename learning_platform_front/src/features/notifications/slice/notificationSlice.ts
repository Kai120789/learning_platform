import { createSlice, type PayloadAction } from '@reduxjs/toolkit';
import type { NotificationSchema, Notification } from '../types/notifications';


const initialState: NotificationSchema = {
    notifications: [],
};

const notificationSlice = createSlice({
    name: 'notifications',
    initialState,
    reducers: {
        addNotification: (state, action: PayloadAction<Omit<Notification, 'id'>>) => {
            state.notifications.push({
                id: Date.now(),
                ...action.payload,
            });
        },
        removeNotification: (state, action: PayloadAction<number>) => {
            state.notifications = state.notifications.filter((n) => n.id !== action.payload);
        },
    },
});

export const { actions: notificationActions, reducer: notificationReducer } =
    notificationSlice;
