
import { getRouteLogin, getRouteRegister } from "@/app/router/routePaths";
import { Button } from "@/shared/ui/Button";
import { useNavigate } from "react-router-dom";
import Logo from '@/shared/assets/Logo-Test.png'

type GuestTopMenuProps = {}

export function GuestTopMenu(props: GuestTopMenuProps) {
    const navigate = useNavigate()

    return (
        <div className="border-b-2 border-[#D9D9D9]">
            <div className="flex flex-row p-[20px] justify-between">
                <img className="h-[40px] w-[100px]" src={Logo} />
                <div className="flex gap-2">
                    <Button
                        variant="outline"
                        size="lg"
                        className="cursor-pointer"
                        onClick={() => navigate(getRouteRegister())}
                    >
                        Зарегистрироваться
                    </Button>
                    <Button
                        size="lg"
                        className="cursor-pointer"
                        onClick={() => navigate(getRouteLogin())}
                    >
                        Войти
                    </Button>
                </div>
            </div>
        </div>
    )
}