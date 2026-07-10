import { FiUser, FiSettings, FiLogOut } from "react-icons/fi";

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/shared/ui/dropdown-menu"
import type { ReactElement } from "react"

type DropdownMenuIconsProps = {
    trigger: ReactElement
}

export function DropdownMenuIcons({ trigger }: DropdownMenuIconsProps) {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger render={trigger} />
            <DropdownMenuContent className="min-w-50 p-3 space-y-2">
                <DropdownMenuItem className="text-lg gap-2">
                    <FiUser className="size-6" />
                    Профиль
                </DropdownMenuItem>
                <DropdownMenuItem className="text-lg gap-2">
                    <FiSettings className="size-6" />
                    Настройки
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem variant="destructive" className="text-lg gap-2">
                    <FiLogOut className="size-6" />
                    Выйти
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}
