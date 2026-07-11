import { Outlet } from "react-router-dom";

export function GuestLayout() {
    return (
        <div className="min-h-[100vh]">
            <Outlet />
        </div>
    )
}
