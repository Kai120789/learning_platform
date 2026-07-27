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
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("tutors.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("tutors.subtitle")}
                </Label>
            </div>

            <div className="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
                {mockTutors.map((tutor) => (
                    <Card key={tutor.id}>
                        <CardHeader className="space-y-4">
                            <div className="flex items-center gap-3">
                                <Avatar size="lg">
                                    <AvatarFallback>
                                        {tutor.name[0]}{tutor.surname[0]}
                                    </AvatarFallback>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-lg">
                                        {tutor.name} {tutor.surname}
                                    </CardTitle>
                                    {tutor.city && (
                                        <CardDescription>{tutor.city}</CardDescription>
                                    )}
                                </div>
                            </div>
                            {tutor.about && (
                                <CardDescription>{tutor.about}</CardDescription>
                            )}
                        </CardHeader>
                        <CardContent className="space-y-3">
                            <div className="flex flex-wrap gap-2">
                                {tutor.subjects.map((subject) => (
                                    <Badge key={subject} variant="outline">
                                        {subject}
                                    </Badge>
                                ))}
                            </div>
                            <div className="text-sm text-muted-foreground">
                                {t("tutors.experience", { years: tutor.experienceYears })}
                            </div>
                            <div className="flex items-center gap-1 text-sm">
                                <Star className="size-4 fill-current" />
                                {t("tutors.rating")}: {tutor.rating}
                            </div>
                        </CardContent>
                    </Card>
                ))}
            </div>
        </div>
    )
}
