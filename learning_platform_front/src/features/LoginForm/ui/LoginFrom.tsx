import { getRouteMain, getRouteRegister } from "@/app/router/routePaths"
import { cn } from "@/shared/lib/utils"
import { Button } from "@/shared/ui/Button"
import { Card, CardContent } from "@/shared/ui/Card"
import {
    Field,
    FieldDescription,
    FieldGroup,
    FieldLabel,
} from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { useNavigate } from "react-router-dom"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import type { LoginRequestDTO } from "../types/types"
import { login } from "../api/login"
import { notificationActions } from "@/features/notifications"
import { useState } from "react"
import { getUserData } from "@/entities/user"

export function LoginForm({
    className,
    ...props
}: React.ComponentProps<"div">) {
    const navigate = useNavigate()
    const dispatch = useAppDispatch()

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const onSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault();

        const request: LoginRequestDTO = {
            email: email,
            password: password,
        }
        const response = await dispatch(login(request))
        if (response.meta.requestStatus == "fulfilled") {
            localStorage.setItem("isAuth", "true")
            dispatch(notificationActions.addNotification({
                message: 'Успешный вход!',
                type: 'success',
            }))
            navigate(getRouteMain())

            const userRes = await dispatch(getUserData())
            if (userRes.meta.requestStatus != "fulfilled") {
                dispatch(notificationActions.addNotification({
                    message: 'Не удалось получить данные пользователя!',
                    type: 'error',
                }))
            }
        } else {
            dispatch(notificationActions.addNotification({
                message: 'Не удалось войти!',
                type: 'error',
            }))
        }
    }

    return (
        <div className={cn("flex flex-col gap-6", className)} {...props}>
            <Card className="overflow-hidden p-0">
                <CardContent className="grid p-0 md:grid-cols-2 min-h-[60vh] items-center">
                    <form className="p-6 md:p-8" onSubmit={onSubmit}>
                        <FieldGroup>
                            <div className="flex flex-col items-center gap-2 text-center">
                                <h1 className="text-2xl font-bold">Добро пожаловать</h1>
                                <p className="text-balance text-muted-foreground">
                                    Войдите в аккаунт, чтобы продолжить
                                </p>
                            </div>
                            <Field>
                                <FieldLabel htmlFor="email">Почта</FieldLabel>
                                <Input
                                    id="email"
                                    type="email"
                                    placeholder="m@example.com"
                                    required
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                />
                            </Field>
                            <Field>
                                <div className="flex items-center">
                                    <FieldLabel htmlFor="password">Пароль</FieldLabel>
                                    <a
                                        href="#"
                                        className="ml-auto text-sm underline-offset-2 hover:underline"
                                    >
                                        Забыли пароль?
                                    </a>
                                </div>
                                <Input
                                    id="password"
                                    type="password"
                                    required
                                    onChange={(e) => setPassword(e.target.value)}
                                />
                            </Field>
                            <Field>
                                <Button type="submit">Войти</Button>
                            </Field>
                            <FieldDescription className="text-center">
                                Нет аккаунта? <a className="cursor-pointer" onClick={() => navigate(getRouteRegister())}>Зарегистрироваться</a>
                            </FieldDescription>
                        </FieldGroup>
                    </form>
                    <div className="relative hidden bg-muted md:block h-full">
                        <img
                            src="/placeholder.svg"
                            alt="Image"
                            className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
                        />
                    </div>
                </CardContent>
            </Card>
            <FieldDescription className="px-6 text-center">
                Нажимая продолжить вы соглашаетесь с <a href="#">Terms of Service</a>{" "}
                и <a href="#">Privacy Policy</a>.
            </FieldDescription>
        </div>
    )
}
