import { Navigate, useLocation } from 'react-router-dom';
import { type JSX } from 'react';
import { getRouteLogin } from '@/app/router/routePaths';
import { useSelector } from 'react-redux';
import { getIsAuth, getIsInitialized } from '@/entities/user';


type AuthProviderProps = {
    children: JSX.Element;
};

export function AuthProvider({ children }: AuthProviderProps) {
    const isAuth = useSelector(getIsAuth)
    const isInitialized = useSelector(getIsInitialized)
    const location = useLocation()

    if (!isInitialized) {
        return null
    }

    if (!isAuth) {
        return <Navigate to={getRouteLogin()} state={{ from: location }} replace />
    }

    return children
}
