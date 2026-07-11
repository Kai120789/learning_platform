export enum AppRoutes {
    WELCOME = 'welcome',
    LOGIN = 'login',
    REGISTER = 'register',
    LESSONS = 'lessons',
    SCHEDULE = 'schedule',
    PROFILE = 'profile',
    COURSES = 'courses',
    MAIN = 'main',
    MATERIALS = 'materials',
    PRACTICES = 'practices',
    TUTORS = 'tutors',
    SETTINGS = 'settings'
}

export const getRouteWelcome = () => '/welcome'
export const getRouteLogin = () => '/login'
export const getRouteRegister = () => '/register'
export const getRouteCourses = () => '/courses'
export const getRouteLessons = () => '/lessons'
export const getRouteSchedule = () => '/schedule'
export const getRouteProfile = () => '/profile'
export const getRouteMain = () => '/'
export const getRouteMaterials = () => '/materials'
export const getRoutePractices = () => '/practices'
export const getRouteTutors = () => '/tutors'
export const getRouteSettings = () => '/settings'