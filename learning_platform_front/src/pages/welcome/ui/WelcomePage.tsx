import { useTranslation } from "react-i18next"
import { useNavigate } from "react-router-dom"
import { BookOpen, Calendar, Users } from "lucide-react"
import { getRouteLogin, getRouteRegister } from "@/app/router/routePaths"
import { Button } from "@/shared/ui/Button"
import {
    Card,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/shared/ui/Card"
import { Label } from "@/shared/ui/Label"

export default function WelcomePage() {
    const { t } = useTranslation()
    const navigate = useNavigate()

    const features = [
        {
            icon: Calendar,
            title: t("welcome.features.schedule.title"),
            text: t("welcome.features.schedule.text"),
        },
        {
            icon: Users,
            title: t("welcome.features.groups.title"),
            text: t("welcome.features.groups.text"),
        },
        {
            icon: BookOpen,
            title: t("welcome.features.practice.title"),
            text: t("welcome.features.practice.text"),
        },
    ]

    return (
        <div className="flex flex-col py-8 lg:py-10 px-6 lg:px-20 space-y-8">
            <div className="space-y-4 max-w-3xl">
                <Label className="text-2xl lg:text-3xl">
                    {t("welcome.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("welcome.subtitle")}
                </Label>
                <div className="flex flex-col sm:flex-row gap-3 pt-2">
                    <Button size="lg" onClick={() => navigate(getRouteRegister())}>
                        {t("welcome.ctaRegister")}
                    </Button>
                    <Button size="lg" variant="outline" onClick={() => navigate(getRouteLogin())}>
                        {t("welcome.ctaLogin")}
                    </Button>
                </div>
            </div>

            <div className="grid auto-rows-fr gap-5 md:grid-cols-3">
                {features.map((feature) => {
                    const Icon = feature.icon

                    return (
                        <Card key={feature.title} className="h-full">
                            <CardHeader className="space-y-3">
                                <Icon className="size-6" />
                                <CardTitle>{feature.title}</CardTitle>
                                <CardDescription className="line-clamp-3">
                                    {feature.text}
                                </CardDescription>
                            </CardHeader>
                        </Card>
                    )
                })}
            </div>
        </div>
    )
}
