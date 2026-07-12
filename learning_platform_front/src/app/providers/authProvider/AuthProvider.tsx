import { Navigate, useLocation } from 'react-router-dom';
import { type JSX } from 'react';
import { getRouteLogin } from '@/app/router/routePaths';


type AuthProviderProps = {
    children: JSX.Element;
};

export function AuthProvider({ children }: AuthProviderProps) {
    const isAuth = localStorage.getItem("isAuth")
    const location = useLocation()

    if (!isAuth || isAuth == "false") {
        return <Navigate to={getRouteLogin()} state={{ from: location }} replace />
    }

    return children
}
