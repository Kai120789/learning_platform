export enum AppRoutes {
    WELCOME = 'welcome',
    LOGIN = 'login',
    REGISTER = 'register',
    SCHEDULE = 'schedule',
    PROFILE = 'profile',
    COURSES = 'courses',
    MAIN = 'main',
    MATERIALS = 'materials',
    PRACTICES = 'practices',
    TUTORS = 'tutors',
    TUTOR = 'tutor',
    SETTINGS = 'settings',
    GROUPS = 'groups'
}

export const getRouteWelcome = () => '/welcome'
export const getRouteLogin = () => '/login'
export const getRouteRegister = () => '/register'
export const getRouteCourses = () => '/courses'
export const getRouteSchedule = () => '/schedule'
export const getRouteProfile = () => '/profile'
export const getRouteMain = () => '/'
export const getRouteMaterials = () => '/materials'
export const getRoutePractices = () => '/practices'
export const getRouteTutors = () => '/tutors'
export const getRouteTutor = (tutorId: number | string) => `/tutors/${tutorId}`
export const getRouteSettings = () => '/settings'
export const getRouteGroups = () => '/groups'
