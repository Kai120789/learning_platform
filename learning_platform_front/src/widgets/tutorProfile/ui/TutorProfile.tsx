import { useEffect } from "react"
import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { ArrowLeft, Clock, Star, Users } from "lucide-react"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getRouteTutors } from "@/app/router/routePaths"
import {
    formatPrice,
    formatRating,
    formatTutorName,
    formatTutorSubject,
    getCurrentTutor,
    getOneTutor,
    getTutorInitials,
    getTutorIsLoading,
    getTutorIsReviewsLoading,
    getTutorReviews,
    getTutorReviewsCount,
    getTutorReviewsList,
} from "@/entities/tutor"
import { getUserFullData } from "@/entities/user"
import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { TutorReviewCard } from "./TutorReviewCard"
import { TutorReviewForm } from "./TutorReviewForm"

const REVIEWS_LIMIT = 10

type TutorProfileProps = {
    tutorId: number
}

export function TutorProfile({ tutorId }: TutorProfileProps) {
    const { t, i18n } = useTranslation()
    const dispatch = useAppDispatch()
    const profile = useAppSelector(getCurrentTutor)
    const isLoading = useAppSelector(getTutorIsLoading)
    const reviews = useAppSelector(getTutorReviewsList)
    const reviewsCount = useAppSelector(getTutorReviewsCount)
    const isReviewsLoading = useAppSelector(getTutorIsReviewsLoading)
    const userData = useAppSelector(getUserFullData)
    const currentUserId = userData?.user.userID

    useEffect(() => {
        dispatch(getOneTutor({ tutorId }))
        dispatch(getTutorReviews({
            tutorId,
            page: 1,
            limit: REVIEWS_LIMIT,
        }))
    }, [dispatch, tutorId])

    const onLoadMoreReviews = () => {
        const nextPage = Math.floor(reviews.length / REVIEWS_LIMIT) + 1
        dispatch(getTutorReviews({
            tutorId,
            page: nextPage,
            limit: REVIEWS_LIMIT,
            append: true,
        }))
    }

    if (isLoading && !profile) {
        return (
            <div className="py-12 text-center text-sm text-muted-foreground">
                {t("common.loading")}
            </div>
        )
    }

    if (!profile) {
        return (
            <div className="space-y-4">
                <Link
                    to={getRouteTutors()}
                    className="inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground"
                >
                    <ArrowLeft className="size-4" />
                    {t("tutors.back")}
                </Link>
                <div className="rounded-2xl border border-dashed border-border px-4 py-12 text-center text-sm text-muted-foreground">
                    {t("tutors.notFound")}
                </div>
            </div>
        )
    }

    const { tutor, offers, myReview } = profile
    const isOwnProfile = currentUserId === tutor.id
    const about = isOwnProfile ? userData?.userInfo.about : undefined
    const otherReviews = reviews.filter((review) => review.id !== myReview?.id)

    return (
        <div className="space-y-5">
            <Link
                to={getRouteTutors()}
                className="inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground"
            >
                <ArrowLeft className="size-4" />
                {t("tutors.back")}
            </Link>

            <section className="rounded-2xl border border-border bg-card p-5 lg:p-6">
                <div className="flex flex-col gap-4 sm:flex-row sm:items-center">
                    <Avatar className="!size-24 shrink-0 text-xl">
                        <AvatarFallback className="bg-muted text-xl text-foreground">
                            {getTutorInitials(tutor)}
                        </AvatarFallback>
                    </Avatar>
                    <div className="min-w-0 flex-1 space-y-2">
                        <div>
                            <div className="text-xl font-medium leading-tight">
                                {formatTutorName(tutor, true)}
                            </div>
                            {tutor.tgUsername && (
                                <div className="mt-1 text-sm text-muted-foreground">
                                    @{tutor.tgUsername}
                                </div>
                            )}
                        </div>
                        {tutor.subjects.length > 0 && (
                            <div className="flex flex-wrap gap-1.5">
                                {tutor.subjects.map((subject) => (
                                    <Badge key={subject.id} variant="outline" className="text-xs">
                                        {formatTutorSubject(subject)}
                                    </Badge>
                                ))}
                            </div>
                        )}
                    </div>
                    <div className="flex shrink-0 gap-2 sm:w-32 sm:flex-col">
                        <div className="flex-1 rounded-xl bg-muted/70 px-3.5 py-2.5 text-sm">
                            <div className="flex items-center gap-1 font-medium">
                                <Star className={tutor.reviewsCount > 0 ? "size-3.5 fill-current" : "size-3.5"} />
                                {tutor.reviewsCount > 0 ? formatRating(tutor.rating) : "—"}
                            </div>
                            <div className="text-xs text-muted-foreground">
                                {tutor.reviewsCount > 0
                                    ? t("tutors.reviewsCount", { count: tutor.reviewsCount })
                                    : t("tutors.noReviewsYet")}
                            </div>
                        </div>
                        <div className="flex-1 rounded-xl bg-muted/70 px-3.5 py-2.5 text-sm">
                            <div className="flex items-center gap-1 font-medium">
                                <Users className="size-3.5" />
                                {tutor.studentsCount}
                            </div>
                            <div className="text-xs text-muted-foreground">
                                {t("common.students")}
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <div className="grid gap-5 lg:grid-cols-2">
                <section className="rounded-2xl border border-border bg-card p-5">
                    <div className="mb-3 text-sm font-medium">{t("tutors.about")}</div>
                    {about ? (
                        <p className="whitespace-pre-wrap text-sm leading-relaxed text-muted-foreground">
                            {about}
                        </p>
                    ) : (
                        <p className="text-sm text-muted-foreground">{t("tutors.emptyAbout")}</p>
                    )}
                </section>

                <section className="rounded-2xl border border-border bg-card p-5">
                    <div className="mb-3 text-sm font-medium">{t("tutors.offers")}</div>
                    {offers.length === 0 ? (
                        <p className="text-sm text-muted-foreground">{t("tutors.emptyOffers")}</p>
                    ) : (
                        <div className="divide-y divide-border">
                            {offers.map((offer) => {
                                const subjectLabel = formatTutorSubject(offer.subject)
                                return (
                                    <div
                                        key={offer.id}
                                        className="flex flex-col gap-2 py-3 first:pt-0 last:pb-0 sm:flex-row sm:items-center"
                                    >
                                        <div className="min-w-0 flex-1">
                                            <div className="flex min-w-0 items-baseline gap-2">
                                                <span className="truncate text-sm font-medium">
                                                    {offer.title}
                                                </span>
                                                {subjectLabel && (
                                                    <span className="shrink-0 text-sm text-muted-foreground">
                                                        {subjectLabel}
                                                    </span>
                                                )}
                                            </div>
                                            {offer.description && (
                                                <p className="mt-0.5 line-clamp-2 text-xs text-muted-foreground">
                                                    {offer.description}
                                                </p>
                                            )}
                                        </div>
                                        <div className="shrink-0 sm:text-right">
                                            <div className="text-base font-semibold leading-tight">
                                                {t("tutors.price", { price: formatPrice(offer.price) })}
                                            </div>
                                            {offer.durationMinutes != null && (
                                                <div className="mt-0.5 inline-flex items-center gap-1 text-sm text-muted-foreground">
                                                    <Clock className="size-3.5" />
                                                    {t("common.minutes", { count: offer.durationMinutes })}
                                                </div>
                                            )}
                                        </div>
                                    </div>
                                )
                            })}
                        </div>
                    )}
                </section>
            </div>

            <section className="rounded-2xl border border-border bg-card p-5">
                <div className="mb-3 text-sm font-medium">{t("tutors.reviews")}</div>
                <div className="space-y-2">
                    {!isOwnProfile && userData && (
                        <TutorReviewForm
                            tutorId={tutor.id}
                            subjects={tutor.subjects}
                            myReview={myReview}
                            author={{
                                name: userData.userInfo.name,
                                surname: userData.userInfo.surname,
                            }}
                        />
                    )}
                    {otherReviews.length === 0 && !isReviewsLoading && (isOwnProfile || !userData) ? (
                        <p className="text-sm text-muted-foreground">{t("tutors.emptyReviews")}</p>
                    ) : (
                        otherReviews.map((review) => (
                            <TutorReviewCard
                                key={review.id}
                                author={review.author}
                                subjectLabel={formatTutorSubject(review.subject)}
                                dateLabel={new Date(review.createdAt).toLocaleDateString(i18n.language)}
                                rating={review.rating}
                                text={review.text}
                            />
                        ))
                    )}
                </div>
                {reviews.length < reviewsCount && (
                    <div className="flex justify-center pt-3">
                        <Button
                            type="button"
                            variant="outline"
                            size="sm"
                            disabled={isReviewsLoading}
                            onClick={onLoadMoreReviews}
                        >
                            {t("tutors.loadMore")}
                        </Button>
                    </div>
                )}
            </section>
        </div>
    )
}
