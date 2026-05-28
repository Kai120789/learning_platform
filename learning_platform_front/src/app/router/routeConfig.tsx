import React from 'react'
import { AppRoutes, getRouteLogin, getRouteRegistration, getRouteWelcome } from './routePaths'
import type { AppRoutesProps } from './AppRouter'
import { WelcomePage } from '@/pages/welcome'
import { LoginPage } from '@/pages/login'
import { RegistrationPage } from '@/pages/registration'


export const routeConfig: Record<AppRoutes, AppRoutesProps> = {
    [AppRoutes.WELCOME]: {
        path: getRouteWelcome(),
        element: (
            <WelcomePage />
        ),
        authOnly: false
    },
    [AppRoutes.LOGIN]: {
        path: getRouteLogin(),
        element: (
            <LoginPage />
        ),
        authOnly: false
    },
    [AppRoutes.REGISTRATION]: {
        path: getRouteRegistration(),
        element: (
            <RegistrationPage />
        ),
        authOnly: false
    }
}