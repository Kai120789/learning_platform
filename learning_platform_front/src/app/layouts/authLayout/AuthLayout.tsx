import { useState } from 'react'
import { LeftMenu } from '@/widgets/leftMenu'
import RightTopMenu from '@/features/rightTopMenu/ui/RightTopMenu'
import { Outlet } from 'react-router-dom'

const AuthLayout = () => {
    const [isOpen, setIsOpen] = useState<boolean>(false)

    const onClick = () => {
        setIsOpen(!isOpen)
    }

    return (
        <div>
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

export default AuthLayout
