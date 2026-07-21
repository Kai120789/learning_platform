import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { getAllSubjects } from "@/entities/subject/api/getAllSubjects"
import { getUserData } from "@/entities/user"
import { cn } from "@/shared/lib/utils"
import { LeftMenu } from "@/widgets/leftMenu"
import { TopMenu } from "@/widgets/topMenu"
import { useTheme } from "@teispace/next-themes/client"
import { useEffect, useState } from "react"
import { Outlet } from "react-router-dom"

export function MainLayout() {
    const { theme } = useTheme()
    const [isOpen, setIsOpen] = useState<boolean>(false)

    const onClick = () => {
        setIsOpen(!isOpen)
    }

    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(getAllSubjects())
        dispatch(getUserData())
    }, [])

    return (
        <div className="min-h-[100vh] bg-secondary">
            <div className="flex flex-row items-start">
                <div className='flex flex-1 flex-col'>
                    <TopMenu onClick={onClick} />
                    <Outlet />
                </div>
            </div>
            <div
                onClick={onClick}
                className={cn(
                    "fixed inset-0 z-40 transition-opacity duration-300",
                    theme === "dark"
                        ? "bg-white/10 backdrop-brightness-50"
                        : "bg-black/40",
                    isOpen ? "opacity-100" : "pointer-events-none opacity-0"
                )}
            />

            <LeftMenu isOpen={isOpen} onClick={onClick} />
        </div>
    )
}
