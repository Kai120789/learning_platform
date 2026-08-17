import type { LeftMenuItemTab } from "@/shared/types/leftMenuItems"
import { cn } from "@/shared/lib/utils"
import { Tooltip, TooltipContent, TooltipTrigger } from "@/shared/ui/Tooltip"
import { ChevronDown } from "lucide-react"
import { useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"

type LeftMenuItemProps = {
    item: LeftMenuItemTab
    collapsed?: boolean
    nested?: boolean
    onNavigate?: () => void
    onExpandSidebar?: () => void
}

function isPathActive(pathname: string, path: string) {
    if (!path) return false
    if (path === "/") return pathname === "/"
    return pathname === path || pathname.startsWith(`${path}/`)
}

function isItemActive(item: LeftMenuItemTab, pathname: string): boolean {
    if (isPathActive(pathname, item.path)) return true
    return item.childrens?.some((child) => isItemActive(child, pathname)) ?? false
}

export function LeftMenuItem({
    item,
    collapsed = false,
    nested = false,
    onNavigate,
    onExpandSidebar,
}: LeftMenuItemProps) {
    const navigate = useNavigate()
    const { pathname } = useLocation()
    const hasChildren = Boolean(item.childrens?.length)
    const active = isItemActive(item, pathname)
    const childActive = item.childrens?.some((child) => isItemActive(child, pathname)) ?? false
    const [open, setOpen] = useState(childActive)

    useEffect(() => {
        if (childActive) setOpen(true)
    }, [childActive])

    const onClickItem = () => {
        if (hasChildren) {
            if (collapsed) {
                onExpandSidebar?.()
                setOpen(true)
                return
            }
            setOpen((prev) => !prev)
            return
        }

        if (!item.path) return
        navigate(item.path)
        onNavigate?.()
    }

    const rowClassName = cn(
        "relative flex w-full items-center rounded-md text-sm transition-colors",
        collapsed ? "h-10 justify-center px-0" : "h-9 gap-2.5 px-2.5",
        nested && !collapsed && "h-8 text-[13px]",
        active
            ? "bg-sidebar-accent font-medium text-sidebar-accent-foreground"
            : "text-sidebar-foreground/65 hover:bg-sidebar-accent/70 hover:text-sidebar-accent-foreground"
    )

    const content = (
        <>
            {active && !nested && (
                <span className="absolute left-0.5 h-4 w-0.5 rounded-full bg-sidebar-primary" />
            )}
            <item.icon className="size-4 shrink-0" />
            {!collapsed && (
                <>
                    <span className="min-w-0 flex-1 truncate text-left">{item.text}</span>
                    {hasChildren && (
                        <ChevronDown
                            className={cn(
                                "size-3.5 shrink-0 text-sidebar-foreground/50 transition-transform",
                                open && "rotate-180"
                            )}
                        />
                    )}
                </>
            )}
        </>
    )

    return (
        <div className="w-full">
            {collapsed ? (
                <Tooltip>
                    <TooltipTrigger
                        type="button"
                        onClick={onClickItem}
                        className={rowClassName}
                    >
                        {content}
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={8}>
                        {item.text}
                    </TooltipContent>
                </Tooltip>
            ) : (
                <button type="button" onClick={onClickItem} className={rowClassName}>
                    {content}
                </button>
            )}

            {hasChildren && open && !collapsed && (
                <div className="mt-0.5 ml-4 space-y-0.5 border-l border-sidebar-border pl-2">
                    {item.childrens?.map((child) => (
                        <LeftMenuItem
                            key={child.field}
                            item={child}
                            nested
                            onNavigate={onNavigate}
                        />
                    ))}
                </div>
            )}
        </div>
    )
}
