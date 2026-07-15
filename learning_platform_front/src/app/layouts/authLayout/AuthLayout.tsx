import { Outlet } from "react-router-dom";

export function AuthLayout() {
    return (
        <div className="min-h-[100vh] bg-secondary">
            <Outlet />
        </div>
    )
}
