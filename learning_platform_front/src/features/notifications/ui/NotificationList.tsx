import { useEffect } from 'react';
import { useSelector } from "react-redux";
import { toast } from 'react-toastify';
import { useAppDispatch } from '@/app/providers/storeProvider/hooks/hooks';
import { getNotifications } from '../selectors/notificationSelectors';
import { notificationActions } from '../slice/notificationSlice';

export const NotificationList = () => {
    const notifications = useSelector(getNotifications);
    const dispatch = useAppDispatch();

    useEffect(() => {
        notifications.forEach(notification => {
            const toastOptions = {
                toastId: notification.id,
                autoClose: 5000,
                onClose: () => dispatch(notificationActions.removeNotification(notification.id)),
            };

            const content = <span style={{ whiteSpace: 'pre-line' }}>{notification.message}</span>;

            switch (notification.type) {
                case 'success':
                    toast.success(content, toastOptions);
                    break;
                case 'error':
                    toast.error(content, toastOptions);
                    break;
                case 'warning':
                    toast.warn(content, toastOptions);
                    break;
                case 'info':
                    toast.info(content, toastOptions);
                    break;
                default:
                    toast(content, toastOptions);
            }
        });
    }, [notifications, dispatch]);

    return null;
};
