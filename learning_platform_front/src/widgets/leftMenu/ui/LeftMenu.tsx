import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getUserFullData } from "@/entities/user"
import { LeftMenuItemsType } from "@/shared/types/leftMenuItems"
import { cn } from "@/shared/lib/utils"
import { Avatar, AvatarFallback, AvatarImage } from "@/shared/ui/Avatar"
import { LearningLogo } from "@/shared/ui/LearningLogo"
import { DropdownMenuIcons } from "@/widgets/dropdownMenu"
import { ThemeSwitch } from "@/widgets/themeSwitch"
import { LeftMenuItem } from "./LeftMenuItem"
import { ChevronsLeft, ChevronsRight, X } from "lucide-react"
import { useTranslation } from "react-i18next"

type LeftMenuProps = {
    className?: string
    collapsed?: boolean
    onCollapsedChange?: (collapsed: boolean) => void
    onClose?: () => void
    showClose?: boolean
}

export function LeftMenu({
    className,
    collapsed = false,
    onCollapsedChange,
    onClose,
    showClose = false,
}: LeftMenuProps) {
    const { t } = useTranslation()
    const userData = useAppSelector(getUserFullData)
    const items = LeftMenuItemsType()

    const name = userData?.userInfo.name
    const surname = userData?.userInfo.surname
    const displayName = [name, surname].filter(Boolean).join(" ")
    const initials = `${name?.[0] ?? ""}${surname?.[0] ?? ""}` || "U"
    const role = userData?.user.role

    const userButton = (
        <button
            type="button"
            className={cn(
                "flex w-full flex-col items-center rounded-lg text-center outline-none transition-colors hover:bg-muted focus-visible:bg-muted",
                collapsed ? "py-2" : "gap-2 px-2 py-3"
            )}
        >
            <Avatar className={collapsed ? "size-9" : "size-14"}>
                {userData?.userInfo.avatar && (
                    <AvatarImage src={userData.userInfo.avatar} alt={displayName} />
                )}
                <AvatarFallback className={collapsed ? "text-xs" : "text-base"}>
                    {initials}
                </AvatarFallback>
            </Avatar>
            {!collapsed && (
                <span className="min-w-0 w-full">
                    <span className="block truncate text-sm font-medium leading-tight">
                        {displayName || userData?.user.email}
                    </span>
                    {role && (
                        <span className="mt-0.5 block truncate text-xs text-muted-foreground">
                            {t(`roles.${role}`)}
                        </span>
                    )}
                </span>
            )}
        </button>
    )

    return (
        <aside
            className={cn(
                "h-full shrink-0 flex-col overflow-hidden border-r border-border bg-background",
                collapsed ? "w-16" : "w-52",
                className
            )}
        >
            <div
                className={cn(
                    "flex h-14 shrink-0 items-center",
                    collapsed ? "justify-center px-2" : "justify-between px-4"
                )}
            >
                <LearningLogo compact={collapsed} />
                {showClose && (
                    <button
                        type="button"
                        onClick={onClose}
                        className="rounded-md p-1 text-muted-foreground hover:bg-muted hover:text-foreground"
                        aria-label={t("sidebar.close")}
                    >
                        <X className="size-4" />
                    </button>
                )}
                {onCollapsedChange && !showClose && !collapsed && (
                    <button
                        type="button"
                        onClick={() => onCollapsedChange(true)}
                        className="rounded-md p-1 text-muted-foreground hover:bg-muted hover:text-foreground"
                        aria-label={t("sidebar.collapse")}
                    >
                        <ChevronsLeft className="size-4" />
                    </button>
                )}
            </div>

            <div className={cn("shrink-0 px-2 pb-1", collapsed && "px-1.5")}>
                <DropdownMenuIcons
                    side="right"
                    align="start"
                    trigger={userButton}
                />
            </div>

            <nav className={cn("flex-1 space-y-0.5 overflow-y-auto px-2 pb-2", collapsed && "px-1.5")}>
                {items.map((item) => (
                    <LeftMenuItem
                        key={item.field}
                        item={item}
                        collapsed={collapsed}
                        onNavigate={onClose}
                        onExpandSidebar={() => onCollapsedChange?.(false)}
                    />
                ))}
            </nav>

            <div className={cn("mt-auto shrink-0 px-2 pb-5 pt-2", collapsed && "px-1.5")}>
                <div className="mx-2 h-px bg-border" />
                <div className={cn("mt-5 flex items-center", collapsed ? "justify-center" : "justify-between px-2 text-sm text-muted-foreground")}>
                    {!collapsed && t("theme")}
                    <ThemeSwitch compact={collapsed} />
                </div>

                {onCollapsedChange && collapsed && (
                    <button
                        type="button"
                        onClick={() => onCollapsedChange(false)}
                        className="flex h-10 w-full items-center justify-center rounded-md text-muted-foreground hover:bg-muted/60 hover:text-foreground"
                        aria-label={t("sidebar.expand")}
                    >
                        <ChevronsRight className="size-4" />
                    </button>
                )}
            </div>
        </aside>
    )
}
