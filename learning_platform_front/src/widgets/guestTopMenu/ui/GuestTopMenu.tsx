
import { getRouteLogin, getRouteRegister } from "@/app/router/routePaths";
import { Button } from "@/shared/ui/Button";
import { useNavigate } from "react-router-dom";
import { FaRegFontAwesomeLogoFull } from "react-icons/fa";
import { ThemeSwitch } from "@/widgets/themeSwitch";
import { useTranslation } from "react-i18next";

export function GuestTopMenu() {
    const { t } = useTranslation()
    const navigate = useNavigate()

    return (
        <div className="border-b-2 border-border">
            <div className="flex flex-row p-[20px] justify-between">
                <FaRegFontAwesomeLogoFull className="h-[40px] w-[150px]" />
                <div className="flex gap-10 items-center">
                    <ThemeSwitch />
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