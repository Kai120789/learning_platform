import {
    AppRoutes,
    getRouteCourses,
    getRouteLessons,
    getRouteLogin, getRouteMain, getRouteMaterials, getRoutePractices, getRouteProfile,
    getRouteRegister, getRouteSchedule, getRouteSettings, getRouteTutors,
    getRouteWelcome
} from './routePaths'
import type { AppRoutesProps } from './AppRouter'
import { WelcomePage } from '@/pages/welcome'
import { LoginPage } from '@/pages/login'
import { LessonsPage } from "@/pages/lessons";
import { MainPage } from "@/pages/main";
import { MaterialsPage } from "@/pages/materials";
import { PracticesPage } from "@/pages/practices";
import { ProfilePage } from "@/pages/profile";
import { SchedulePage } from "@/pages/schedule";
import { CoursesPage } from "@/pages/courses";
import { TutorsPage } from "@/pages/tutors";
import { SettingsPage } from '@/pages/settings'
import { AuthLayout, GuestLayout, MainLayout } from '../layouts'
import RegisterPage from '@/pages/registration/ui/RegisterPage'

export const routeConfig: Record<AppRoutes, AppRoutesProps> = {
    [AppRoutes.WELCOME]: {
        path: getRouteWelcome(),
        element: (
            <WelcomePage />
        ),
        authOnly: false,
        layout: <GuestLayout />
    },
    [AppRoutes.LOGIN]: {
        path: getRouteLogin(),
        element: (
            <LoginPage />
        ),
        authOnly: false,
        layout: <AuthLayout />
    },
    [AppRoutes.REGISTER]: {
        path: getRouteRegister(),
        element: (
            <RegisterPage />
        ),
        authOnly: false,
        layout: <AuthLayout />
    },
    [AppRoutes.LESSONS]: {
        path: getRouteLessons(),
        element: (
            <LessonsPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.COURSES]: {
        path: getRouteCourses(),
        element: (
            <CoursesPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.MAIN]: {
        path: getRouteMain(),
        element: (
            <MainPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.MATERIALS]: {
        path: getRouteMaterials(),
        element: (
            <MaterialsPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.PRACTICES]: {
        path: getRoutePractices(),
        element: (
            <PracticesPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.PROFILE]: {
        path: getRouteProfile(),
        element: (
            <ProfilePage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.SCHEDULE]: {
        path: getRouteSchedule(),
        element: (
            <SchedulePage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.TUTORS]: {
        path: getRouteTutors(),
        element: (
            <TutorsPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    },
    [AppRoutes.SETTINGS]: {
        path: getRouteSettings(),
        element: (
            <SettingsPage />
        ),
        authOnly: true,
        layout: <MainLayout />
    }
}
