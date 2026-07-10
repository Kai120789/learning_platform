import {
    AppRoutes,
    getRouteCourses,
    getRouteLessons,
    getRouteLogin, getRouteMain, getRouteMaterials, getRoutePractices, getRouteProfile,
    getRouteRegistration, getRouteSchedule, getRouteTutors,
    getRouteWelcome
} from './routePaths'
import type { AppRoutesProps } from './AppRouter'
import { WelcomePage } from '@/pages/welcome'
import { LoginPage } from '@/pages/login'
import { RegistrationPage } from '@/pages/registration'
import { LessonsPage } from "@/pages/lessons";
import { MainPage } from "@/pages/main";
import { MaterialsPage } from "@/pages/materials";
import { PracticesPage } from "@/pages/practices";
import { ProfilePage } from "@/pages/profile";
import { SchedulePage } from "@/pages/schedule";
import { CoursesPage } from "@/pages/courses";
import { TutorsPage } from "@/pages/tutors";

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
    },
    [AppRoutes.LESSONS]: {
        path: getRouteLessons(),
        element: (
            <LessonsPage />
        ),
        authOnly: true
    },
    [AppRoutes.COURSES]: {
        path: getRouteCourses(),
        element: (
            <CoursesPage />
        ),
        authOnly: true
    },
    [AppRoutes.MAIN]: {
        path: getRouteMain(),
        element: (
            <MainPage />
        ),
        authOnly: true
    },
    [AppRoutes.MATERIALS]: {
        path: getRouteMaterials(),
        element: (
            <MaterialsPage />
        ),
        authOnly: true
    },
    [AppRoutes.PRACTICES]: {
        path: getRoutePractices(),
        element: (
            <PracticesPage />
        ),
        authOnly: true
    },
    [AppRoutes.PROFILE]: {
        path: getRouteProfile(),
        element: (
            <ProfilePage />
        ),
        authOnly: true
    },
    [AppRoutes.SCHEDULE]: {
        path: getRouteSchedule(),
        element: (
            <SchedulePage />
        ),
        authOnly: true
    },
    [AppRoutes.TUTORS]: {
        path: getRouteTutors(),
        element: (
            <TutorsPage />
        ),
        authOnly: true
    }
}
