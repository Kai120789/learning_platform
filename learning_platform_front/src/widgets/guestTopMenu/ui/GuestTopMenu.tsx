
import { getRouteLogin, getRouteRegister } from "@/app/router/routePaths";
import { Button } from "@/shared/ui/Button";
import { useNavigate } from "react-router-dom";
import { LearningLogo } from "@/shared/ui/LearningLogo";
import { useTranslation } from "react-i18next";

export function GuestTopMenu() {
    const { t } = useTranslation()
    const navigate = useNavigate()

    return (
        <div className="border-b-2 border-border bg-background">
            <div className="flex flex-row p-[20px] justify-between items-center">
                <LearningLogo className="text-base" />
                <div className="flex gap-4 lg:gap-8 items-center">
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            size="lg"
                            className="cursor-pointer"
                            onClick={() => navigate(getRouteRegister())}
                        >
                            {t("guest.register")}
                        </Button>
                        <Button
                            size="lg"
                            className="cursor-pointer"
                            onClick={() => navigate(getRouteLogin())}
                        >
                            {t("guest.login")}
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    )
}