import { getRouteCourses, getRouteMain, getRouteMaterials, getRoutePractices, getRouteSchedule, getRouteTutors } from "@/app/router/routePaths"
import { LeftMenuTabs } from "./leftMenuTabs"
import { AiOutlineHome, AiOutlineCalendar } from "react-icons/ai";
import { FiUsers, FiShoppingCart, FiBookOpen, FiPenTool } from "react-icons/fi";
import { PiStudent } from "react-icons/pi";
import type { IconType } from "react-icons/lib";
import { useTranslation } from "react-i18next";

export interface LeftMenuItemTab {
    icon: IconType
    path: string
    text: string
    field: LeftMenuTabs
    childrens?: LeftMenuItemTab[]
}

export const LeftMenuItemsType = (): LeftMenuItemTab[] => {
    const { t } = useTranslation()

    return [
        {
            icon: AiOutlineHome,
            path: getRouteMain(),
            text: t("tabs.main"),
            field: LeftMenuTabs.MAIN
        },
        {
            icon: AiOutlineCalendar,
            path: getRouteSchedule(),
            text: t("tabs.schedules"),
            field: LeftMenuTabs.SCHEDULE
        },
        {
            icon: FiUsers,
            path: getRouteTutors(),
            text: t("tabs.tutors"),
            field: LeftMenuTabs.TUTORS
        },
        {
            icon: FiShoppingCart,
            path: "",
            text: t("tabs.services"),
            field: LeftMenuTabs.SERVICES,
            childrens: [
                {
                    icon: FiShoppingCart,
                    path: "",
                    text: t("tabs.items"),
                    field: LeftMenuTabs.SERVICES_ITEMS
                },
                {
                    icon: FiShoppingCart,
                    path: "",
                    text: t("tabs.orders"),
                    field: LeftMenuTabs.SERVICES_ORDERS
                },
            ]
        },
        {
            icon: PiStudent,
            path: "",
            text: t("tabs.studying"),
            field: LeftMenuTabs.STUDYING,
            childrens: [
                {
                    icon: FiUsers,
                    path: getRouteCourses(),
                    text: t("tabs.courses"),
                    field: LeftMenuTabs.COURSES
                },
                {
                    icon: FiPenTool,
                    path: getRoutePractices(),
                    text: t("tabs.practices"),
                    field: LeftMenuTabs.PRACTICES
                },
                {
                    icon: FiBookOpen,
                    path: getRouteMaterials(),
                    text: t("tabs.materials"),
                    field: LeftMenuTabs.MATERIALS
                },
            ]
        },
    ]
}