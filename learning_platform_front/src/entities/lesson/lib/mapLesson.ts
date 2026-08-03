import type {
    LessonData,
    LessonMediaItemData,
    LessonMediaItemResponse,
    LessonResponse,
} from "../types/types"

export function mapLessonMediaItemResponse(
    item: LessonMediaItemResponse,
): LessonMediaItemData {
    return {
        id: item.id,
        lessonId: item.lesson_id,
        s3Link: item.s3_link,
        s3Preview: item.s3_preview,
        type: item.type,
    }
}

export function mapLessonResponse(lesson: LessonResponse): LessonData {
    return {
        id: lesson.id,
        boardId: lesson.board_id ?? null,
        meetLink: lesson.meet_link ?? null,
        startTime: lesson.start_time,
        duration: lesson.duration,
        tutorId: lesson.tutor_id,
        status: lesson.status,
        userIds: lesson.user_ids ?? [],
        mediaItems: (lesson.media_items ?? []).map(mapLessonMediaItemResponse),
    }
}
