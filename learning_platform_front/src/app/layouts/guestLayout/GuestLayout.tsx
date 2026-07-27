import { Outlet, useNavigate } from "react-router-dom";
import { GuestTopMenu } from "@/widgets/guestTopMenu"
import { getRouteMain } from "@/app/router/routePaths";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import { getIsAuth } from "@/entities/user";

export function GuestLayout() {
    const isAuth = useSelector(getIsAuth)
    const navigate = useNavigate()

    useEffect(() => {
        if (isAuth) {
            navigate(getRouteMain());
        }
    }, [isAuth, navigate]);

    return (
        <div className="min-h-[100vh] bg-secondary">
            <div className="flex flex-row items-start">
                <div className='flex flex-1 flex-col'>
                    <GuestTopMenu />
                    <Outlet />
                </div>
            </div>
        </div>
    )
}
