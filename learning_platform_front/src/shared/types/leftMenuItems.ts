import { getRouteCourses, getRouteMain, getRouteMaterials, getRoutePractices, getRouteSchedule, getRouteTutors } from "@/app/router/routePaths"
import { LeftMenuTabs } from "./leftMenuTabs"
import { AiOutlineHome, AiOutlineCalendar } from "react-icons/ai";
import { FiUsers, FiShoppingCart } from "react-icons/fi";
import { PiStudent } from "react-icons/pi";
import type { IconType } from "react-icons/lib";

export interface LeftMenuItemTab {
    icon: IconType
    path: string
    text: string
    field: LeftMenuTabs
    childrens?: LeftMenuItemTab[]
}

export const LeftMenuItemsType = (): LeftMenuItemTab[] => {
    return [
        {
            icon: AiOutlineHome,
            path: getRouteMain(),
            text: 'Главная',
            field: LeftMenuTabs.MAIN
        },
        {
            icon: AiOutlineCalendar,
            path: getRouteSchedule(),
            text: 'Расписания',
            field: LeftMenuTabs.SCHEDULE
        },
        {
            icon: FiUsers,
            path: getRouteTutors(),
            text: 'Репетиторы',
            field: LeftMenuTabs.TUTORS
        },
        {
            icon: FiShoppingCart,
            path: "",
            text: 'Услуги',
            field: LeftMenuTabs.SERVICES,
            childrens: [
                {
                    icon: FiShoppingCart,
                    path: "",
                    text: 'Предложения',
                    field: LeftMenuTabs.SERVICES_ITEMS
                },
                {
                    icon: FiShoppingCart,
                    path: "",
                    text: 'Заказы',
                    field: LeftMenuTabs.SERVICES_ORDERS
                },
            ]
        },
        {
            icon: PiStudent,
            path: "",
            text: 'Обучение',
            field: LeftMenuTabs.STUDY,
            childrens: [
                {
                    icon: FiUsers,
                    path: getRouteCourses(),
                    text: 'Курсы',
                    field: LeftMenuTabs.COURSES
                },
                {
                    icon: FiShoppingCart,
                    path: getRoutePractices(),
                    text: 'Задания',
                    field: LeftMenuTabs.PRACTICES
                },
                {
                    icon: FiShoppingCart,
                    path: getRouteMaterials(),
                    text: 'Материалы',
                    field: LeftMenuTabs.MATERIALS
                },
            ]
        },
    ]
}