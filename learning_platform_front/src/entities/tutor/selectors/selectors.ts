import type { StateSchema } from "@/app/providers/storeProvider"

export const getTutorList = (state: StateSchema) => state.tutor.list
export const getTutorListCount = (state: StateSchema) => state.tutor.listCount
export const getCurrentTutor = (state: StateSchema) => state.tutor.current
export const getTutorReviewsList = (state: StateSchema) => state.tutor.reviews
export const getTutorReviewsCount = (state: StateSchema) => state.tutor.reviewsCount
export const getTutorTeaching = (state: StateSchema) => state.tutor.teaching
export const getTutorIsLoading = (state: StateSchema) => state.tutor.isLoading
export const getTutorIsReviewsLoading = (state: StateSchema) => state.tutor.isReviewsLoading
export const getTutorIsTeachingLoading = (state: StateSchema) => state.tutor.isTeachingLoading
export const getTutorIsStudentsLoading = (state: StateSchema) => state.tutor.isStudentsLoading
export const getTutorError = (state: StateSchema) => state.tutor.error
