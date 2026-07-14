import { Outlet } from "react-router-dom";
import { GuestTopMenu } from "@/widgets/guestTopMenu"

export function GuestLayout() {
    return (
        <div className="min-h-[100vh] bg-muted">
            <div className="flex flex-row items-start">
                <div className='flex flex-1 flex-col'>
                    <GuestTopMenu />
                    <Outlet />
                </div>
            </div>
        </div>
    )
}
