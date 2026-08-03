import type { AxiosInstance } from 'axios';
import { getRouteLogin } from '@/app/router/routePaths';
import { notificationActions } from "@/features/notifications";
import type { AppStore } from '@/app/providers/storeProvider';
import { resetStore } from '@/shared/lib/resetStore';
import { t } from 'i18next';

export function interceptor(api: AxiosInstance, store: AppStore) {
    api.interceptors.response.use(
        (response) => response,
        (error) => {
            if (window.location.pathname !== getRouteLogin()) {
                if (error.response) {
                    const status = error.response.status;

                    if (status === 401 && store.getState().user.isAuth) {
                        console.error(`Auth error (${status}):`, error.response.data);
                        store.dispatch(resetStore());
                        store.dispatch(notificationActions.addNotification({
                            type: "error",
                            message: t("unauthorized"),
                        }))
                    }
                } else {
                    console.error('Network or CORS error:', error);
                }
            }

            return Promise.reject(error);
        }
    );
}
