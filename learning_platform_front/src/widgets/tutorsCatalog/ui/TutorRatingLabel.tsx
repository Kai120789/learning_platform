import { Star, Users } from "lucide-react"
import { formatRating } from "@/entities/tutor"
import { useTranslation } from "react-i18next"

type TutorRatingLabelProps = {
    rating: number
    reviewsCount: number
    studentsCount: number
    compact?: boolean
    inline?: boolean
}

export function TutorRatingLabel({
    rating,
    reviewsCount,
    studentsCount,
    compact = false,
    inline = false,
}: TutorRatingLabelProps) {
    const { t } = useTranslation()

    return (
        <div className={inline
            ? "flex flex-nowrap items-center gap-3 text-xs text-muted-foreground"
            : compact
                ? "flex flex-col items-start gap-1 text-xs text-muted-foreground sm:items-end"
                : "flex flex-wrap items-center gap-x-3 gap-y-1 text-xs text-muted-foreground"
        }>
            {reviewsCount > 0 ? (
                <span className="inline-flex shrink-0 items-center gap-1">
                    <Star className="size-3.5 fill-current text-foreground" />
                    {formatRating(rating)}
                    {!compact && !inline && (
                        <span>· {t("tutors.reviewsCount", { count: reviewsCount })}</span>
                    )}
                </span>
            ) : (
                <span className="shrink-0">{t("tutors.noReviewsYet")}</span>
            )}
            <span className="inline-flex shrink-0 items-center gap-1">
                <Users className="size-3.5" />
                {studentsCount}
            </span>
        </div>
    )
}
