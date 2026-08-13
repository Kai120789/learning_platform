import { LogOut, Settings, User } from "lucide-react";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/shared/ui/DropdownMenu"
import type { ReactElement } from "react"
import { useNavigate } from "react-router-dom";
import { getRouteProfile, getRouteSettings, getRouteWelcome } from "@/app/router/routePaths";
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { logout } from "../api/logout";
import { useTranslation } from "react-i18next";
import { resetStore } from "@/shared/lib/resetStore";

type DropdownMenuIconsProps = {
    trigger: ReactElement
    side?: "top" | "bottom" | "left" | "right"
    align?: "start" | "center" | "end"
}

export function DropdownMenuIcons({
    trigger,
    side = "bottom",
    align = "start",
}: DropdownMenuIconsProps) {
    const { t } = useTranslation()
    const navigate = useNavigate()

    const dispatch = useAppDispatch()

    const onClickExit = async () => {
        const response = await dispatch(logout())

        if (response.meta.requestStatus == "fulfilled") {
            dispatch(resetStore())
            navigate(getRouteWelcome())
        }
    }

    return (
        <DropdownMenu modal={false}>
            <DropdownMenuTrigger render={trigger} />
            <DropdownMenuContent
                side={side}
                align={align}
                className="w-44 min-w-44 bg-background py-1.5 px-1 space-y-0.5 text-sm font-medium"
            >
                <DropdownMenuItem onClick={() => navigate(getRouteProfile())} className="text-sm gap-2 cursor-pointer">
                    <User className="size-4" />
                    {t("rightMenu.profile")}
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => navigate(getRouteSettings())} className="text-sm gap-2 cursor-pointer">
                    <Settings className="size-4" />
                    {t("rightMenu.settings")}
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem onClick={onClickExit} variant="destructive" className="text-sm gap-2 cursor-pointer">
                    <LogOut className="size-4" />
                    {t("rightMenu.logout")}
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}
