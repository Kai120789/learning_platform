export type LessonStatus = "SCHEDULED" | "IN_PROCESS" | "COMPLETED" | "CANCELLED"
export type LessonMediaType = "IMAGE" | "VIDEO"

export type LessonMediaItemData = {
    id: number
    lessonId: number
    s3Link: string
    s3Preview: string
    type: LessonMediaType
}

export type LessonData = {
    id: number
    boardId?: number | null
    meetLink?: string | null
    startTime: string
    duration: number
    tutorId: number
    status: LessonStatus
    userIds: number[]
    mediaItems: LessonMediaItemData[]
}

export type LessonMediaItemResponse = {
    id: number
    lesson_id: number
    s3_link: string
    s3_preview: string
    type: LessonMediaType
}

export type LessonResponse = {
    id: number
    board_id?: number | null
    meet_link?: string | null
    start_time: string
    duration: number
    tutor_id: number
    status: LessonStatus
    user_ids: number[]
    media_items: LessonMediaItemResponse[]
}

export type LessonSchema = {
    data: LessonData[] | null
    isLoading: boolean
    error?: string
}
