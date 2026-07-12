export interface NotificationSchema {
    notifications: Notification[];
}

export interface Notification {
    id: number;
    message: string;
    type: "error" | "success" | "info" | "warning"
}