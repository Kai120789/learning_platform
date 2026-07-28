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
import { useTranslation } from "react-i18next";
import { userActions } from "@/entities/user";

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
            dispatch(userActions.setIsAuth(false))
            navigate(getRouteWelcome())
        }
    }

    return (
        <DropdownMenu>
            <DropdownMenuTrigger render={trigger} />
            <DropdownMenuContent className="bg-background min-w-44 py-1.5 px-2 space-y-0.5 text-sm font-medium">
                <DropdownMenuItem onClick={() => navigate(getRouteProfile())} className="text-sm gap-2 cursor-pointer">
                    <FiUser className="size-4" />
                    {t("rightMenu.profile")}
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => navigate(getRouteSettings())} className="text-sm gap-2 cursor-pointer">
                    <FiSettings className="size-4" />
                    {t("rightMenu.settings")}
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => navigate(getRouteGroups())} className="text-sm gap-2 cursor-pointer">
                    <FiUsers className="size-4" />
                    {t("rightMenu.groups")}
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem onClick={onClickExit} variant="destructive" className="text-sm gap-2 cursor-pointer">
                    <FiLogOut className="size-4" />
                    {t("rightMenu.logout")}
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}
