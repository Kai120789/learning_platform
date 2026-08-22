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

type TutorGridCardProps = {
    tutor: TutorCardData
}

export function TutorGridCard({ tutor }: TutorGridCardProps) {
    const { t } = useTranslation()

    return (
        <Link
            to={getRouteTutor(tutor.id)}
            className="flex h-full min-h-44 flex-col rounded-xl border border-border bg-card p-4 transition-colors hover:bg-muted/30"
        >
            <div className="flex items-center gap-3">
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

            <div className="mt-3 flex min-h-14 flex-1 flex-wrap content-start gap-1">
                {tutor.subjects.map((subject) => (
                    <Badge key={subject.id} variant="outline" className="text-[10px]">
                        {subject.title} · {subject.type}
                    </Badge>
                ))}
            </div>

            <div className="mt-3 border-t border-border pt-3">
                <TutorRatingLabel
                    rating={tutor.rating}
                    reviewsCount={tutor.reviewsCount}
                    studentsCount={tutor.studentsCount}
                />
            </div>
        </Link>
    )
}
