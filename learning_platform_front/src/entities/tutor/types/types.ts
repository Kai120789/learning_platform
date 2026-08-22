import type { SubjectTypeEnum } from "@/entities/subject"

export interface TutorSchema {
    list: TutorCardData[]
    listCount: number
    current: TutorProfileData | null
    reviews: TutorReviewData[]
    reviewsCount: number
    teaching: TutorTeachingData | null
    isLoading: boolean
    isReviewsLoading: boolean
    isTeachingLoading: boolean
    isStudentsLoading: boolean
    error?: string
}

export type TutorTeachingData = {
    subjectIds: number[]
    offers: TutorOfferData[]
    students: TutorStudentData[]
    studentsCount: number
}

export type TutorUserData = {
    id: number
    name: string
    surname: string
    patronymic?: string
    tgUsername?: string
}

export type TutorStudentData = {
    student: TutorUserData
    lastInteractedAt?: string
}

export type TutorSubjectData = {
    id: number
    code: string
    title: string
    type: SubjectTypeEnum
}

export type TutorCardData = {
    id: number
    name: string
    surname: string
    patronymic?: string
    tgUsername?: string
    rating: number
    reviewsCount: number
    studentsCount: number
    subjects: TutorSubjectData[]
}

export type TutorOfferData = {
    id: number
    tutorId: number
    subject: TutorSubjectData
    title: string
    description?: string
    price: number
    durationMinutes?: number
}

export type TutorReviewData = {
    id: number
    tutor: TutorUserData
    author: TutorUserData
    subject: TutorSubjectData
    text: string
    rating: number
    createdAt: string
    updatedAt: string
}

export type TutorMyReviewData = {
    id: number
    tutorId: number
    authorId: number
    subjectId: number
    text: string
    rating: number
    createdAt: string
    updatedAt: string
}

export type TutorProfileData = {
    tutor: TutorCardData
    offers: TutorOfferData[]
    myReview: TutorMyReviewData | null
}

export type TutorUserResponse = {
    id: number
    name: string
    surname: string
    patronymic?: string
    tg_username?: string
}

export type TutorSubjectResponse = {
    id: number
    code: string
    title: string
    type: SubjectTypeEnum
}

export type TutorShortInfoResponse = {
    tutor: TutorUserResponse
    rating: number
    reviews_count: number
    students_count: number
    subjects: TutorSubjectResponse[] | null
}

export type TutorOfferResponse = {
    id: number
    tutor_id: number
    subject: TutorSubjectResponse
    title: string
    description?: string | null
    price: number
    duration_minutes?: number | null
}

export type TutorMyReviewResponse = {
    id: number
    tutor_id: number
    author_id: number
    subject_id: number
    text: string
    rating: number
    created_at: string
    updated_at: string
}

export type TutorReviewResponse = {
    id: number
    tutor: TutorUserResponse
    author: TutorUserResponse
    subject: TutorSubjectResponse
    text: string
    rating: number
    created_at: string
    updated_at: string
}

export type OneTutorResponse = {
    tutor_info: TutorShortInfoResponse
    offers: TutorOfferResponse[] | null
    my_review: TutorMyReviewResponse | null
}

export type GetTutorsResponse = {
    tutors: TutorShortInfoResponse[] | null
    count: number
}

export type GetTutorReviewsResponse = {
    reviews: TutorReviewResponse[] | null
    count: number
}

export type GetTutorsRequest = {
    page: number
    limit: number
    subjectId?: number
    append?: boolean
}

export type GetTutorReviewsRequest = {
    tutorId: number
    page: number
    limit: number
    append?: boolean
}

export type GetTutorStudentsRequest = {
    page: number
    limit: number
    interactedWithinDays?: number
    append?: boolean
}

export type GetTutorStudentsResponse = {
    students: TutorStudentResponse[] | null
    count: number
}

export type TutorStudentResponse = {
    student: TutorUserResponse
    last_interacted_at?: string | null
}

export type NewTutorOfferRequest = {
    subject_id: number
    title: string
    description?: string | null
    price: number
    duration_minutes?: number | null
}

export type UpdateTutorOfferRequest = {
    id: number
    subject_id: number
    title: string
    description?: string | null
    price: number
    duration_minutes?: number | null
}

export type NewTutorReviewRequest = {
    text: string
    tutor_id: number
    subject_id: number
    rating: number
}

export type UpdateTutorReviewRequest = {
    id: number
    text: string
    subject_id: number
    rating: number
}
