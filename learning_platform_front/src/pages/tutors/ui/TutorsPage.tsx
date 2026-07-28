import { useTranslation } from "react-i18next"
import { Star } from "lucide-react"
import { mockTutors } from "@/shared/mocks"
import { Badge } from "@/shared/ui/Badge"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/shared/ui/Card"
import { Label } from "@/shared/ui/Label"
import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"

export default function TutorsPage() {
    const { t } = useTranslation()

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("tutors.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("tutors.subtitle")}
                </Label>
            </div>

            <div className="grid auto-rows-fr gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {mockTutors.map((tutor) => (
                    <Card key={tutor.id} size="sm" className="h-full">
                        <CardHeader className="space-y-2">
                            <div className="flex items-center gap-2.5">
                                <Avatar size="sm">
                                    <AvatarFallback>
                                        {tutor.name[0]}{tutor.surname[0]}
                                    </AvatarFallback>
                                </Avatar>
                                <div className="min-w-0">
                                    <CardTitle className="text-sm font-medium truncate">
                                        {tutor.name} {tutor.surname}
                                    </CardTitle>
                                    <CardDescription className="text-xs truncate min-h-4">
                                        {tutor.city || "\u00A0"}
                                    </CardDescription>
                                </div>
                            </div>
                            <CardDescription className="text-xs line-clamp-2 min-h-8">
                                {tutor.about || "\u00A0"}
                            </CardDescription>
                        </CardHeader>
                        <CardContent className="mt-auto space-y-2">
                            <div className="flex min-h-6 flex-wrap gap-1 content-start">
                                {tutor.subjects.map((subject) => (
                                    <Badge key={subject} variant="outline" className="text-[10px]">
                                        {subject}
                                    </Badge>
                                ))}
                            </div>
                            <div className="text-xs text-muted-foreground">
                                {t("tutors.experience", { years: tutor.experienceYears })}
                            </div>
                            <div className="flex items-center gap-1 text-xs">
                                <Star className="size-3.5 fill-current" />
                                {t("tutors.rating")}: {tutor.rating}
                            </div>
                        </CardContent>
                    </Card>
                ))}
            </div>
        </div>
    )
}
