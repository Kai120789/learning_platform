import { Moon, Sun } from "lucide-react";
import { useTheme } from '@teispace/next-themes/client';


import { Switch } from "@/shared/ui/Switch";
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { updateUserTheme } from "@/entities/user";
import { UserThemeEnum } from "@/shared/enums/user";
import { notificationActions } from "@/features/notifications";

type ThemeSwitchProps = {
    compact?: boolean
}

export function ThemeSwitch({ compact = false }: ThemeSwitchProps) {
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
            {compact ? (
                <button
                    type="button"
                    onClick={toggleTheme}
                    className="flex size-8 items-center justify-center rounded-md text-muted-foreground hover:bg-muted hover:text-foreground"
                    aria-label="Theme"
                >
                    {theme === "dark" ? <Moon className="size-4" /> : <Sun className="size-4" />}
                </button>
            ) : (
                <>
                    <Sun className="size-3 shrink-0" />
                    <Switch
                        size="sm"
                        checked={theme == 'dark'}
                        onCheckedChange={toggleTheme}
                    />
                    <Moon className="size-3 shrink-0" />
                </>
            )}
        </div>
    )
}
