import { useTranslation } from "react-i18next"
import { useNavigate } from "react-router-dom"
import { CalendarDays, Mail, MapPin, Pencil } from "lucide-react"
import { getRouteSettings } from "@/app/router/routePaths"
import { Avatar, AvatarFallback, AvatarImage } from "@/shared/ui/Avatar"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Card } from "@/shared/ui/Card"
import { Separator } from "@/shared/ui/Separator"

type ProfileHeaderCardProps = {
    name: string
    surname: string
    email: string
    role: string
    city: string
    about: string
    avatar?: string
    memberSince: Date
    subjects: string[]
}

export function ProfileHeaderCard({
    name,
    surname,
    email,
    role,
    city,
    about,
    avatar,
    memberSince,
    subjects,
}: ProfileHeaderCardProps) {
    const { t, i18n } = useTranslation()
    const navigate = useNavigate()

    const memberSinceLabel = memberSince.toLocaleDateString(i18n.language, {
        month: "long",
        year: "numeric",
    })

    return (
        <Card className="p-0 gap-5">
            <div className="h-28 bg-primary/25 lg:h-36" />

            <div className="px-6 pb-6">
                <div className="flex flex-wrap items-end justify-between gap-4">
                    <div className="flex items-end gap-4">
                        <Avatar className="-mt-12 size-24 ring-4 ring-background lg:-mt-14 lg:size-28">
                            {avatar && <AvatarImage src={avatar} alt={name} />}
                            <AvatarFallback className="text-2xl">
                                {name[0]}{surname[0]}
                            </AvatarFallback>
                        </Avatar>
                        <div className="space-y-1 pb-1">
                            <div className="flex flex-wrap items-center gap-2">
                                <span className="text-xl font-semibold lg:text-2xl">
                                    {name} {surname}
                                </span>
                                <Badge>{t(`roles.${role}`)}</Badge>
                            </div>
                            <div className="flex flex-wrap items-center gap-x-4 gap-y-1 text-sm text-muted-foreground">
                                <span className="flex items-center gap-1">
                                    <MapPin className="size-3.5" />
                                    {city}
                                </span>
                                <span className="flex items-center gap-1">
                                    <Mail className="size-3.5" />
                                    {email}
                                </span>
                                <span className="flex items-center gap-1">
                                    <CalendarDays className="size-3.5" />
                                    {t("profile.memberSince", { date: memberSinceLabel })}
                                </span>
                            </div>
                        </div>
                    </div>

                    <Button
                        variant="outline"
                        onClick={() => navigate(getRouteSettings())}
                    >
                        <Pencil className="size-4" />
                        {t("profile.edit")}
                    </Button>
                </div>

                <Separator className="my-4" />

                <div className="space-y-3">
                    <p className="max-w-3xl text-sm text-muted-foreground">
                        {about}
                    </p>
                    <div className="flex flex-wrap items-center gap-2">
                        <span className="text-sm font-medium">
                            {t("profile.subjects")}:
                        </span>
                        {subjects.map((subject) => (
                            <Badge key={subject} variant="outline" className="bg-muted">
                                {subject}
                            </Badge>
                        ))}
                    </div>
                </div>
            </div>
        </Card>
    )
}
