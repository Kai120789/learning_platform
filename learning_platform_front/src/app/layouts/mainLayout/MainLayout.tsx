import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { getAllSubjects } from "@/entities/subject"
import { getUserData } from "@/entities/user"
import { cn } from "@/shared/lib/utils"
import { LeftMenu } from "@/widgets/leftMenu"
import { TopMenu } from "@/widgets/topMenu"
import { useEffect, useState } from "react"
import { Outlet } from "react-router-dom"

const COLLAPSED_KEY = "sidebar-collapsed"

export function MainLayout() {
    const [mobileOpen, setMobileOpen] = useState(false)
    const [collapsed, setCollapsed] = useState(() => {
        if (typeof window === "undefined") return false
        return window.localStorage.getItem(COLLAPSED_KEY) === "true"
    })

    const dispatch = useAppDispatch()

    const onCollapsedChange = (value: boolean) => {
        setCollapsed(value)
        window.localStorage.setItem(COLLAPSED_KEY, String(value))
    }

    useEffect(() => {
        dispatch(getAllSubjects())
        dispatch(getUserData())
    }, [dispatch])

    useEffect(() => {
        if (!mobileOpen) return

        const onKeyDown = (event: KeyboardEvent) => {
            if (event.key === "Escape") setMobileOpen(false)
        }

        window.addEventListener("keydown", onKeyDown)
        return () => window.removeEventListener("keydown", onKeyDown)
    }, [mobileOpen])

    return (
        <div className="flex min-h-svh bg-secondary">
            <LeftMenu
                className="sticky top-0 hidden h-svh lg:flex"
                collapsed={collapsed}
                onCollapsedChange={onCollapsedChange}
            />

            <div className="flex min-w-0 flex-1 flex-col">
                <TopMenu onClick={() => setMobileOpen(true)} />
                <Outlet />
            </div>

            <div
                onClick={() => setMobileOpen(false)}
                className={cn(
                    "fixed inset-0 z-40 bg-black/40 transition-opacity duration-300 lg:hidden",
                    mobileOpen ? "opacity-100" : "pointer-events-none opacity-0"
                )}
            />

            <LeftMenu
                className={cn(
                    "fixed inset-y-0 left-0 z-50 flex lg:hidden transition-transform duration-300 ease-in-out",
                    mobileOpen ? "translate-x-0" : "-translate-x-full"
                )}
                showClose
                onClose={() => setMobileOpen(false)}
            />
        </div>
    )
}
