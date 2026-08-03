import type { CreateLessonMediaItemRequest } from "./createLesson"

export type UpdateLessonRequest = {
    id: number
    board_id?: number | null
    meet_link?: string | null
    start_time: string
    duration: number
    media_items?: CreateLessonMediaItemRequest[]
    user_ids?: number[]
    deleted_user_ids?: number[]
    deleted_media_ids?: number[]
}
