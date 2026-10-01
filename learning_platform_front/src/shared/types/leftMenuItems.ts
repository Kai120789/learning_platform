import { getRouteBoards, getRouteCourses, getRouteGroups, getRouteMain, getRouteMaterials, getRoutePractices, getRouteSchedule, getRouteTutors } from "@/app/router/routePaths"
import { useCanEdit } from "@/entities/user"
import { LeftMenuTabs } from "./leftMenuTabs"
import { BookOpen, Calendar, GraduationCap, Home, Library, PenLine, PenTool, ShoppingCart, UserRound, Users } from "lucide-react";
import type { LucideIcon } from "lucide-react";
import { useTranslation } from "react-i18next";

export interface LeftMenuItemTab {
    icon: LucideIcon
    path: string
    text: string
    field: LeftMenuTabs
    childrens?: LeftMenuItemTab[]
}

export function LeftMenuItemsType(): LeftMenuItemTab[] {
    const { t } = useTranslation()
    const canEdit = useCanEdit()

    const studyingChildren: LeftMenuItemTab[] = [
        {
            icon: Library,
            path: getRouteCourses(),
            text: t("tabs.courses"),
            field: LeftMenuTabs.COURSES
        },
        {
            icon: PenTool,
            path: getRoutePractices(),
            text: t("tabs.practices"),
            field: LeftMenuTabs.PRACTICES
        },
        {
            icon: BookOpen,
            path: getRouteMaterials(),
            text: t("tabs.materials"),
            field: LeftMenuTabs.MATERIALS
        },
    ]

    if (canEdit) {
        studyingChildren.push({
            icon: PenLine,
            path: getRouteBoards(),
            text: t("tabs.board"),
            field: LeftMenuTabs.BOARD
        })
    }

    return [
        {
            icon: Home,
            path: getRouteMain(),
            text: t("tabs.main"),
            field: LeftMenuTabs.MAIN
        },
        {
            icon: Calendar,
            path: getRouteSchedule(),
            text: t("tabs.schedule"),
            field: LeftMenuTabs.SCHEDULE
        },
        {
            icon: Users,
            path: getRouteGroups(),
            text: t("tabs.groups"),
            field: LeftMenuTabs.GROUPS
        },
        {
            icon: UserRound,
            path: getRouteTutors(),
            text: t("tabs.tutors"),
            field: LeftMenuTabs.TUTORS
        },
        {
            icon: ShoppingCart,
            path: "",
            text: t("tabs.services"),
            field: LeftMenuTabs.SERVICES,
            childrens: [
                {
                    icon: ShoppingCart,
                    path: "",
                    text: t("tabs.items"),
                    field: LeftMenuTabs.SERVICES_ITEMS
                },
                {
                    icon: ShoppingCart,
                    path: "",
                    text: t("tabs.orders"),
                    field: LeftMenuTabs.SERVICES_ORDERS
                },
            ]
        },
        {
            icon: GraduationCap,
            path: "",
            text: t("tabs.studying"),
            field: LeftMenuTabs.STUDYING,
            childrens: studyingChildren,
        },
    ]
}
