import { Star, Users } from "lucide-react"
import { formatRating } from "@/entities/tutor"
import { useTranslation } from "react-i18next"

type TutorRatingLabelProps = {
    rating: number
    reviewsCount: number
    studentsCount: number
    compact?: boolean
}

export function TutorRatingLabel({
    rating,
    reviewsCount,
    studentsCount,
    compact = false,
}: TutorRatingLabelProps) {
    const { t } = useTranslation()

    return (
        <div className={compact
            ? "flex flex-col items-end gap-1 text-xs text-muted-foreground"
            : "flex flex-wrap items-center gap-x-3 gap-y-1 text-xs text-muted-foreground"
        }>
            {reviewsCount > 0 ? (
                <span className="inline-flex items-center gap-1">
                    <Star className="size-3.5 fill-current text-foreground" />
                    {formatRating(rating)}
                    {!compact && (
                        <span>· {t("tutors.reviewsCount", { count: reviewsCount })}</span>
                    )}
                </span>
            ) : (
                <span>{t("tutors.noReviewsYet")}</span>
            )}
            <span className="inline-flex items-center gap-1">
                <Users className="size-3.5" />
                {studentsCount}
            </span>
        </div>
    )
}
