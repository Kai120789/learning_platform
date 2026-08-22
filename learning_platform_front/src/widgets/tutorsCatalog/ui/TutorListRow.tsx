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
            className="flex items-center gap-4 rounded-xl border border-border bg-card px-4 py-3.5 transition-colors hover:bg-muted/30"
        >
            <Avatar className="!size-14 shrink-0 text-sm">
                <AvatarFallback>
                    {getTutorInitials(tutor)}
                </AvatarFallback>
            </Avatar>
            <div className="min-w-0 flex-1 space-y-2">
                <div className="min-w-0">
                    <div className="truncate text-sm font-medium">
                        {formatTutorName(tutor)}
                    </div>
                    <div className="mt-0.5 truncate text-xs text-muted-foreground">
                        {tutor.tgUsername ? `@${tutor.tgUsername}` : t("tutors.noUsername")}
                    </div>
                </div>
                {tutor.subjects.length > 0 && (
                    <div className="flex flex-wrap gap-1">
                        {tutor.subjects.map((subject) => (
                            <Badge key={subject.id} variant="outline" className="text-[10px]">
                                {subject.title} · {subject.type}
                            </Badge>
                        ))}
                    </div>
                )}
            </div>
            <div className="hidden shrink-0 sm:block">
                <TutorRatingLabel
                    rating={tutor.rating}
                    reviewsCount={tutor.reviewsCount}
                    studentsCount={tutor.studentsCount}
                    compact
                />
            </div>
        </Link>
    )
}
