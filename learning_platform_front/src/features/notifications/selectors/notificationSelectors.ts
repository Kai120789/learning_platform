import type { StateSchema } from "@/app/providers/storeProvider/config/StateSchema";

export const getNotifications = (state: StateSchema) => state.notifications.notifications
