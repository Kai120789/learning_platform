import { FiUser, FiSettings, FiLogOut, FiUsers } from "react-icons/fi";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/shared/ui/DropdownMenu"
import type { ReactElement } from "react"
import { useNavigate } from "react-router-dom";
import { getRouteGroups, getRouteProfile, getRouteSettings, getRouteWelcome } from "@/app/router/routePaths";
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { logout } from "../api/logout";
import { notificationActions } from "@/features/notifications";
import { useTranslation } from "react-i18next";

type DropdownMenuIconsProps = {
    trigger: ReactElement
}

export function DropdownMenuIcons({ trigger }: DropdownMenuIconsProps) {
    const { t } = useTranslation()
    const navigate = useNavigate()

    const dispatch = useAppDispatch()

    const onClickExit = async () => {
        const response = await dispatch(logout())

        if (response.meta.requestStatus == "fulfilled") {
            localStorage.removeItem("isAuth")
            navigate(getRouteWelcome())
        } else {
            dispatch(notificationActions.addNotification({
                message: 'Ошибка',
                type: 'error',
            }))
        }
    }

    return (
        <DropdownMenu>
            <DropdownMenuTrigger render={trigger} />
            <DropdownMenuContent className="bg-background min-w-50 py-2 px-3 space-y-1 font-medium">
                <DropdownMenuItem onClick={() => navigate(getRouteProfile())} className="text-md gap-2">
                    <FiUser className="size-5" />
                    {t("rightMenu.profile")}
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => navigate(getRouteSettings())} className="text-md gap-2">
                    <FiSettings className="size-5" />
                    {t("rightMenu.settings")}
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => navigate(getRouteGroups())} className="text-md gap-2">
                    <FiUsers className="size-5" />
                    {t("rightMenu.groups")}
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem onClick={onClickExit} variant="destructive" className="text-md gap-2">
                    <FiLogOut className="size-5" />
                    {t("rightMenu.logout")}
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}
