import type {
    OneTutorResponse,
    TutorCardData,
    TutorMyReviewData,
    TutorMyReviewResponse,
    TutorOfferData,
    TutorOfferResponse,
    TutorProfileData,
    TutorReviewData,
    TutorReviewResponse,
    TutorShortInfoResponse,
    TutorSubjectData,
    TutorSubjectResponse,
    TutorUserData,
    TutorUserResponse,
    TutorStudentData,
    TutorStudentResponse,
} from "../types/types"

export function mapTutorUser(user: TutorUserResponse): TutorUserData {
    return {
        id: user.id,
        name: user.name,
        surname: user.surname,
        patronymic: user.patronymic ?? undefined,
        tgUsername: user.tg_username ?? undefined,
    }
}

export function mapTutorSubject(subject: TutorSubjectResponse): TutorSubjectData {
    return {
        id: subject.id,
        code: subject.code,
        title: subject.title,
        type: subject.type,
    }
}

export function mapTutorCard(tutor: TutorShortInfoResponse): TutorCardData {
    return {
        id: tutor.tutor.id,
        name: tutor.tutor.name,
        surname: tutor.tutor.surname,
        patronymic: tutor.tutor.patronymic ?? undefined,
        tgUsername: tutor.tutor.tg_username ?? undefined,
        rating: tutor.rating,
        reviewsCount: tutor.reviews_count,
        studentsCount: tutor.students_count,
        subjects: (tutor.subjects ?? []).map(mapTutorSubject),
    }
}

export function mapTutorOffer(offer: TutorOfferResponse): TutorOfferData {
    return {
        id: offer.id,
        tutorId: offer.tutor_id,
        subject: mapTutorSubject(offer.subject),
        title: offer.title,
        description: offer.description ?? undefined,
        price: offer.price,
        durationMinutes: offer.duration_minutes ?? undefined,
    }
}

export function sortTutorOffers(offers: TutorOfferData[]) {
    return offers.sort((a, b) => {
        if (a.subject.id !== b.subject.id) {
            return a.subject.id - b.subject.id
        }
        return a.price - b.price
    })
}

export function mapTutorMyReview(review: TutorMyReviewResponse): TutorMyReviewData {
    return {
        id: review.id,
        tutorId: review.tutor_id,
        authorId: review.author_id,
        subjectId: review.subject_id,
        text: review.text,
        rating: review.rating,
        createdAt: review.created_at,
        updatedAt: review.updated_at,
    }
}

export function mapTutorReview(review: TutorReviewResponse): TutorReviewData {
    return {
        id: review.id,
        tutor: mapTutorUser(review.tutor),
        author: mapTutorUser(review.author),
        subject: mapTutorSubject(review.subject),
        text: review.text,
        rating: review.rating,
        createdAt: review.created_at,
        updatedAt: review.updated_at,
    }
}

export function mapTutorProfile(payload: OneTutorResponse): TutorProfileData {
    return {
        tutor: mapTutorCard(payload.tutor_info),
        offers: (payload.offers ?? []).map(mapTutorOffer),
        myReview: payload.my_review ? mapTutorMyReview(payload.my_review) : null,
    }
}

export function mapTutorStudent(student: TutorStudentResponse): TutorStudentData {
    return {
        student: mapTutorUser(student.student),
        lastInteractedAt: student.last_interacted_at ?? undefined,
    }
}
