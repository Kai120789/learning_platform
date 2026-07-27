import { getRouteMain } from "@/app/router/routePaths";
import { getIsAuth } from "@/entities/user";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import { Outlet, useLocation, useNavigate } from "react-router-dom";

export function AuthLayout() {
    const isAuth = useSelector(getIsAuth)
    const navigate = useNavigate()
    const location = useLocation()

    useEffect(() => {
        if (isAuth) {
            const from = (location.state as { from?: { pathname?: string } } | null)?.from?.pathname
            navigate(from || getRouteMain(), { replace: true })
        }
    }, [isAuth, location.state, navigate]);

    return (
        <div className="min-h-[100vh] bg-secondary">
            <Outlet />
        </div>
    )
}
