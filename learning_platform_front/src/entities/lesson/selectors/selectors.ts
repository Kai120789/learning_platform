import type { StateSchema } from "@/app/providers/storeProvider/config/StateSchema"

export const getLessons = (state: StateSchema) => state.lesson.data
export const getLessonsIsLoading = (state: StateSchema) => state.lesson.isLoading
export const getLessonsError = (state: StateSchema) => state.lesson.error
