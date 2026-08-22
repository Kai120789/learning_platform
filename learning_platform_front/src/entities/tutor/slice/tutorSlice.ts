import { createSlice } from "@reduxjs/toolkit"
import type { TutorSchema, TutorTeachingData } from "../types/types"
import { mapTutorCard, mapTutorOffer, mapTutorProfile, mapTutorReview, mapTutorStudent, sortTutorOffers } from "../lib/mappers"
import { getTutors } from "../api/getTutors"
import { getOneTutor } from "@/entities/tutor"
import { getTutorReviews } from "@/entities/tutor"
import { getTutorOffers } from "@/entities/tutor"
import { addTutorOffer } from "@/entities/tutor"
import { updateTutorOffer } from "@/entities/tutor"
import { deleteOneTutorOffer } from "@/entities/tutor"
import { updateTutorSubjects } from "@/entities/tutor"
import { addTutorReview } from "@/entities/tutor"
import { updateTutorReview } from "@/entities/tutor"
import { deleteTutorReview } from "@/entities/tutor"
import { getTutorStudents } from "@/entities/tutor"
import { addTutorStudents } from "@/entities/tutor"
import { deleteOneTutorStudent } from "@/entities/tutor"

function emptyTeaching(): TutorTeachingData {
    return {
        subjectIds: [],
        offers: [],
        students: [],
        studentsCount: 0,
    }
}

function withTeaching(
    teaching: TutorTeachingData | null,
    patch: Partial<TutorTeachingData>,
): TutorTeachingData {
    return {
        ...(teaching ?? emptyTeaching()),
        ...patch,
    }
}

const initialState: TutorSchema = {
    list: [],
    listCount: 0,
    current: null,
    reviews: [],
    reviewsCount: 0,
    teaching: null,
    isLoading: false,
    isReviewsLoading: false,
    isTeachingLoading: false,
    isStudentsLoading: false,
    error: undefined,
}

const tutorSlice = createSlice({
    name: "tutor",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(getTutors.pending, (state, action) => {
            state.isLoading = true
            state.error = ""
            if (!action.meta.arg.append) {
                state.list = []
            }
        })
        builder.addCase(getTutors.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getTutors.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const tutors = (action.payload.tutors ?? []).map(mapTutorCard)
            state.list = action.meta.arg.append ? [...state.list, ...tutors] : tutors
            state.listCount = action.payload.count
        })

        builder.addCase(getOneTutor.pending, (state, action) => {
            if (action.meta.arg.purpose === "teaching") {
                state.isTeachingLoading = true
            } else {
                state.isLoading = true
                state.current = null
            }
            state.error = ""
        })
        builder.addCase(getOneTutor.rejected, (state, action) => {
            if (action.meta.arg.purpose === "teaching") {
                state.isTeachingLoading = false
                state.teaching = withTeaching(state.teaching, {
                    subjectIds: [],
                    offers: [],
                })
            } else {
                state.isLoading = false
            }
            state.error = action.payload as string
        })
        builder.addCase(getOneTutor.fulfilled, (state, action) => {
            const profile = mapTutorProfile(action.payload)
            if (action.meta.arg.purpose === "teaching") {
                state.isTeachingLoading = false
                state.teaching = withTeaching(state.teaching, {
                    subjectIds: profile.tutor.subjects.map((subject) => subject.id),
                    offers: profile.offers,
                })
            } else {
                state.isLoading = false
                state.current = profile
            }
            state.error = ""
        })

        builder.addCase(getTutorReviews.pending, (state) => {
            state.isReviewsLoading = true
            state.error = ""
        })
        builder.addCase(getTutorReviews.rejected, (state, action) => {
            state.isReviewsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getTutorReviews.fulfilled, (state, action) => {
            state.isReviewsLoading = false
            state.error = ""
            const reviews = (action.payload.reviews ?? []).map(mapTutorReview)
            state.reviews = action.meta.arg.append ? [...state.reviews, ...reviews] : reviews
            state.reviewsCount = action.payload.count
        })

        builder.addCase(getTutorOffers.pending, (state) => {
            state.isTeachingLoading = true
            state.error = ""
        })
        builder.addCase(getTutorOffers.rejected, (state, action) => {
            state.isTeachingLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getTutorOffers.fulfilled, (state, action) => {
            state.isTeachingLoading = false
            state.error = ""
            const offers = action.payload.map(mapTutorOffer)
            if (!state.teaching) {
                state.teaching = withTeaching(null, { offers })
            } else {
                state.teaching.offers = offers
            }
        })

        builder.addCase(addTutorOffer.pending, (state) => {
            state.isTeachingLoading = true
            state.error = ""
        })
        builder.addCase(addTutorOffer.rejected, (state, action) => {
            state.isTeachingLoading = false
            state.error = action.payload as string
        })
        builder.addCase(addTutorOffer.fulfilled, (state, action) => {
            state.isTeachingLoading = false
            state.error = ""
            const offer = mapTutorOffer(action.payload)
            if (!state.teaching) {
                state.teaching = withTeaching(null, { offers: [offer] })
            } else {
                state.teaching.offers.push(offer)
                sortTutorOffers(state.teaching.offers)
            }
            if (state.current && state.current.tutor.id === offer.tutorId) {
                state.current.offers.push(offer)
                sortTutorOffers(state.current.offers)
            }
        })

        builder.addCase(updateTutorOffer.pending, (state) => {
            state.isTeachingLoading = true
            state.error = ""
        })
        builder.addCase(updateTutorOffer.rejected, (state, action) => {
            state.isTeachingLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateTutorOffer.fulfilled, (state, action) => {
            state.isTeachingLoading = false
            state.error = ""
            const offer = mapTutorOffer(action.payload)
            if (state.teaching) {
                state.teaching.offers = sortTutorOffers(
                    state.teaching.offers.map((item) => (
                        item.id === offer.id ? offer : item
                    )),
                )
            }
            if (state.current) {
                state.current.offers = sortTutorOffers(
                    state.current.offers.map((item) => (
                        item.id === offer.id ? offer : item
                    )),
                )
            }
        })

        builder.addCase(deleteOneTutorOffer.pending, (state) => {
            state.isTeachingLoading = true
            state.error = ""
        })
        builder.addCase(deleteOneTutorOffer.rejected, (state, action) => {
            state.isTeachingLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteOneTutorOffer.fulfilled, (state, action) => {
            state.isTeachingLoading = false
            state.error = ""
            if (state.teaching) {
                state.teaching.offers = state.teaching.offers.filter((item) => item.id !== action.payload)
            }
            if (state.current) {
                state.current.offers = state.current.offers.filter((item) => item.id !== action.payload)
            }
        })

        builder.addCase(updateTutorSubjects.pending, (state) => {
            state.isTeachingLoading = true
            state.error = ""
        })
        builder.addCase(updateTutorSubjects.rejected, (state, action) => {
            state.isTeachingLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateTutorSubjects.fulfilled, (state, action) => {
            state.isTeachingLoading = false
            state.error = ""
            if (!state.teaching) {
                state.teaching = withTeaching(null, { subjectIds: action.meta.arg })
            } else {
                state.teaching.subjectIds = action.meta.arg
            }
        })

        builder.addCase(addTutorReview.pending, (state) => {
            state.isReviewsLoading = true
            state.error = ""
        })
        builder.addCase(addTutorReview.rejected, (state, action) => {
            state.isReviewsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(addTutorReview.fulfilled, (state, action) => {
            state.isReviewsLoading = false
            state.error = ""
            const review = mapTutorReview(action.payload)
            state.reviews = [review, ...state.reviews]
            state.reviewsCount += 1
            if (state.current) {
                state.current.myReview = {
                    id: review.id,
                    tutorId: review.tutor.id,
                    authorId: review.author.id,
                    subjectId: review.subject.id,
                    text: review.text,
                    rating: review.rating,
                    createdAt: review.createdAt,
                    updatedAt: review.updatedAt,
                }
                state.current.tutor.reviewsCount += 1
            }
        })

        builder.addCase(updateTutorReview.pending, (state) => {
            state.isReviewsLoading = true
            state.error = ""
        })
        builder.addCase(updateTutorReview.rejected, (state, action) => {
            state.isReviewsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(updateTutorReview.fulfilled, (state, action) => {
            state.isReviewsLoading = false
            state.error = ""
            const review = mapTutorReview(action.payload)
            state.reviews = state.reviews.map((item) => (item.id === review.id ? review : item))
            if (state.current?.myReview?.id === review.id) {
                state.current.myReview = {
                    ...state.current.myReview,
                    subjectId: review.subject.id,
                    text: review.text,
                    rating: review.rating,
                    updatedAt: review.updatedAt,
                }
            }
        })

        builder.addCase(deleteTutorReview.pending, (state) => {
            state.isReviewsLoading = true
            state.error = ""
        })
        builder.addCase(deleteTutorReview.rejected, (state, action) => {
            state.isReviewsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteTutorReview.fulfilled, (state, action) => {
            state.isReviewsLoading = false
            state.error = ""
            state.reviews = state.reviews.filter((item) => item.id !== action.payload)
            state.reviewsCount = Math.max(0, state.reviewsCount - 1)
            if (state.current) {
                state.current.myReview = null
                state.current.tutor.reviewsCount = Math.max(0, state.current.tutor.reviewsCount - 1)
            }
        })
        builder.addCase(getTutorStudents.pending, (state, action) => {
            state.isStudentsLoading = true
            state.error = ""
            if (!action.meta.arg.append && state.teaching) {
                state.teaching.students = []
            }
        })
        builder.addCase(getTutorStudents.rejected, (state, action) => {
            state.isStudentsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getTutorStudents.fulfilled, (state, action) => {
            state.isStudentsLoading = false
            state.error = ""
            const students = (action.payload.students ?? []).map(mapTutorStudent)
            state.teaching = withTeaching(state.teaching, {
                students: action.meta.arg.append
                    ? [...(state.teaching?.students ?? []), ...students]
                    : students,
                studentsCount: action.payload.count,
            })
        })

        builder.addCase(addTutorStudents.pending, (state) => {
            state.isStudentsLoading = true
            state.error = ""
        })
        builder.addCase(addTutorStudents.rejected, (state, action) => {
            state.isStudentsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(addTutorStudents.fulfilled, (state, action) => {
            state.isStudentsLoading = false
            state.error = ""
            const existing = new Set((state.teaching?.students ?? []).map((item) => item.student.id))
            const added = action.payload
                .filter((student) => !existing.has(student.id))
                .map((student) => ({ student }))
            state.teaching = withTeaching(state.teaching, {
                students: [...added, ...(state.teaching?.students ?? [])],
                studentsCount: (state.teaching?.studentsCount ?? 0) + added.length,
            })
        })

        builder.addCase(deleteOneTutorStudent.pending, (state) => {
            state.isStudentsLoading = true
            state.error = ""
        })
        builder.addCase(deleteOneTutorStudent.rejected, (state, action) => {
            state.isStudentsLoading = false
            state.error = action.payload as string
        })
        builder.addCase(deleteOneTutorStudent.fulfilled, (state, action) => {
            state.isStudentsLoading = false
            state.error = ""
            if (!state.teaching) return
            state.teaching.students = state.teaching.students.filter(
                (item) => item.student.id !== action.payload,
            )
            state.teaching.studentsCount = Math.max(0, state.teaching.studentsCount - 1)
        })
    },
})

export const { actions: tutorActions, reducer: tutorReducer } = tutorSlice
