import type { SubjectTypeEnum } from "@/entities/subject/types/types";

export interface GroupSchema {
    data: GroupData[] | null;
    isLoading: boolean;
    error?: string;
}

export type GroupData = {
    id: number
    title: string
    description: string
    subject: Subject
    users?: GroupUser[]
    tutorId: number
    tgGroupLink?: string
    tgChatId?: string
}


export type GroupResponse = {
    id: number
    title: string
    description: string
    subject: Subject
    users?: ShortUserInfo[]
    tutor_id: number
    tg_group_link?: string
    tg_chat_id?: string
}

export type Subject = {
    id: number
    code: string
    title: string
    type: SubjectTypeEnum
}

export type GroupUser = {
    id: number
    name: string
    surname: string
    patronymic?: string
    tgUsername?: string
}

export type ShortUserInfo = {
    id: number
    name: string
    surname: string
    patronymic?: string
    tg_username?: string
}