import { Link } from "react-router-dom"
import { getRouteTutor } from "@/app/router/routePaths"
import {
    formatTutorName,
    getTutorInitials,
    type TutorCardData,
} from "@/entities/tutor"
import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { Badge } from "@/shared/ui/Badge"
import { useTranslation } from "react-i18next"
import { TutorRatingLabel } from "./TutorRatingLabel"

type TutorListRowProps = {
    tutor: TutorCardData
}

export function TutorListRow({ tutor }: TutorListRowProps) {
    const { t } = useTranslation()

    return (
        <Link
            to={getRouteTutor(tutor.id)}
            className="flex flex-col gap-3 rounded-xl border border-border bg-card px-4 py-3.5 transition-colors hover:bg-muted/30 lg:flex-row lg:items-center lg:gap-4"
        >
            <div className="flex min-w-0 items-center gap-3 lg:w-56 lg:shrink-0">
                <Avatar className="!size-14 shrink-0 text-sm">
                    <AvatarFallback>
                        {getTutorInitials(tutor)}
                    </AvatarFallback>
                </Avatar>
                <div className="min-w-0">
                    <div className="truncate text-sm font-medium">
                        {formatTutorName(tutor)}
                    </div>
                    <div className="mt-0.5 truncate text-xs text-muted-foreground">
                        {tutor.tgUsername ? `@${tutor.tgUsername}` : t("tutors.noUsername")}
                    </div>
                </div>
            </div>

            <div className="flex min-h-5 min-w-0 flex-1 flex-wrap content-start gap-1">
                {tutor.subjects.map((subject) => (
                    <Badge
                        key={subject.id}
                        variant="outline"
                        className="h-auto max-w-full min-w-0 whitespace-normal text-left text-[10px] leading-tight"
                    >
                        {subject.title} · {subject.type}
                    </Badge>
                ))}
            </div>

            <div className="shrink-0 border-t border-border pt-2 lg:border-t-0 lg:pt-0">
                <TutorRatingLabel
                    rating={tutor.rating}
                    reviewsCount={tutor.reviewsCount}
                    studentsCount={tutor.studentsCount}
                    inline
                />
            </div>
        </Link>
    )
}
