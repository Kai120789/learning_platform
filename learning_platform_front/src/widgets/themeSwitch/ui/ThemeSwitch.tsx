import { IoMoonOutline } from "react-icons/io5";
import { FiSun } from "react-icons/fi";
import { useTheme } from '@teispace/next-themes/client';


import { Switch } from "@/shared/ui/Switch";
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { updateUserTheme } from "@/entities/user";
import { UserThemeEnum } from "@/shared/enums/user";
import { notificationActions } from "@/features/notifications";

export function ThemeSwitch() {
    const { theme, setTheme } = useTheme();
    const dispatch = useAppDispatch()

    const toggleTheme = async () => {
        setTheme(theme === 'dark' ? 'light' : 'dark')
        const response = await dispatch(updateUserTheme(theme === 'dark' ? UserThemeEnum.LIGHT : UserThemeEnum.DARK))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: 'Тема обновлена',
                type: 'success',
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: 'Не удалось обновить тему',
                type: 'error',
            }))
        }
    };

    return (
        <div className="flex items-center gap-1.5">
            <FiSun className="size-3 shrink-0" />
            <Switch
                size="sm"
                checked={theme == 'dark'}
                onCheckedChange={toggleTheme}
            />
            <IoMoonOutline className="size-3 shrink-0" />
        </div>
    )
}
