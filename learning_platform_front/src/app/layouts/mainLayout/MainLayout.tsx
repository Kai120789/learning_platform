import { LeftMenu } from "@/widgets/leftMenu"
import { RightTopMenu } from "@/widgets/rightTopMenu"
import { useState } from "react"
import { Outlet } from "react-router-dom"

export function MainLayout() {
    const [isOpen, setIsOpen] = useState<boolean>(false)

    const onClick = () => {
        setIsOpen(!isOpen)
    }

    return (
        <div className="min-h-[100vh] bg-muted">
            <div className="flex flex-row items-start">
                <div className='flex flex-1 flex-col'>
                    <RightTopMenu onClick={onClick} />
                    <Outlet />
                </div>
            </div>
            <div
                onClick={onClick}
                className={`
                    fixed inset-0 z-40 bg-black/40
                    transition-opacity duration-300
                    ${isOpen ? "opacity-100" : "pointer-events-none opacity-0"}
                `}
            />

            <LeftMenu isOpen={isOpen} onClick={onClick} />
        </div>
    )
}
