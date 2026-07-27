import { getRouteMain } from "@/app/router/routePaths";
import { getIsAuth } from "@/entities/user";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import { Outlet, useNavigate } from "react-router-dom";

export function AuthLayout() {
    const isAuth = useSelector(getIsAuth)
    const navigate = useNavigate()

    useEffect(() => {
        if (isAuth) {
            navigate(getRouteMain());
        }
    }, [isAuth, navigate]);

    return (
        <div className="min-h-[100vh] bg-secondary">
            <Outlet />
        </div>
    )
}
