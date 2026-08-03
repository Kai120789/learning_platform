export type CreateLessonMediaItemRequest = {
    s3_link: string
    s3_preview: string
    type: "IMAGE" | "VIDEO"
}

export type CreateLessonRequest = {
    board_id?: number | null
    meet_link?: string | null
    start_time: string
    duration: number
    tutor_id: number
    media_items?: CreateLessonMediaItemRequest[]
    user_ids: number[]
}
