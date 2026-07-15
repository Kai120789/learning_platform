import { getRouteLogin, getRouteMain } from "@/app/router/routePaths"
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
import { Stepper, type Step } from "@/shared/ui/Stepper"
import { useState, type ReactNode } from "react"
import { CgProfile } from "react-icons/cg"
import { MdLockOutline } from "react-icons/md";

import { useNavigate } from "react-router-dom"
import { RegisterRoleEnum, type RegisterRequestDTO } from "../types/types"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { register } from "../api/register"
import { notificationActions } from "@/features/notifications"
import { PiStudent } from "react-icons/pi"
import { getUserData } from "@/entities/user"
import { UserGenderEnum, UserLanguageEnum } from "@/shared/enums/user"
import { FiUserCheck } from "react-icons/fi"
import { RoleAndLanguageStep } from "./steps/RoleAndLanguageStep"
import { UserDataStep } from "./steps/UserDataStep"
import { AuthDataStep } from "./steps/AuthDataStep"
import { ConfirmStep } from "./steps/ConfirmStep"

export function RegisterForm({
    className,
    ...props
}: React.ComponentProps<"div">) {
    const navigate = useNavigate()

    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [confirmPassword, setConfirmPassword] = useState<string>("");
    const [name, setName] = useState<string>("");
    const [surname, setSurname] = useState<string>("");
    const [patronymic, setPatronymic] = useState<string>("");
    const [role, setRole] = useState<RegisterRoleEnum>(RegisterRoleEnum.STUDENT);
    const [gender, setGender] = useState<UserGenderEnum>(UserGenderEnum.UNKNOWN);
    const [language, setLanguage] = useState<UserLanguageEnum>(UserLanguageEnum.RU);
    const [birthDate, setBirthDate] = useState<Date>()
    const [isChecked, setIsCheked] = useState<boolean>(false)

    const dispatch = useAppDispatch()

    const steps: Step[] = [
        { id: 1, icon: <PiStudent size={20} /> },
        { id: 2, icon: <CgProfile size={20} /> },
        { id: 3, icon: <MdLockOutline size={20} /> },
        { id: 4, icon: <FiUserCheck size={20} /> }
    ]

    const [currentStep, setCurrentStep] = useState<number>(1)

    const nextStep = () => {
        setCurrentStep(currentStep + 1)
    }

    const prevStep = () => {
        setCurrentStep(currentStep - 1)
    }

    const onSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault();

        const request: RegisterRequestDTO = {
            name: name,
            surname: surname,
            patronymic: patronymic,
            role: role,
            email: email,
            password: password,
            gender: gender,
            language: language,
            birth_date: birthDate || undefined,
        }
        const response = await dispatch(register(request))
        if (response.meta.requestStatus == "fulfilled") {
            localStorage.setItem("isAuth", "true")
            dispatch(notificationActions.addNotification({
                message: 'Успешная регистрация!',
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
                message: 'Не удалось зарегистрировать пользователя!',
                type: 'error',
            }))
        }
    }

    const checkConfirmPasswordCorrect = (): boolean => {
        if (confirmPassword !== password) {
            return false
        }
        return true
    }

    const checkNextStepDisabled = (): boolean => {
        switch (currentStep) {
            case 1:
                return (!role)
            case 2:
                return (!name || !surname)
            case 3:
                return (!email || !password || !checkConfirmPasswordCorrect())
            case 4:
                return (!isChecked)
            default:
                return true
        }
    }

    const renderFormByStep = (): ReactNode => {
        switch (currentStep) {
            case 1:
                return (
                    <RoleAndLanguageStep
                        language={language}
                        setLanguage={setLanguage}
                        role={role}
                        setRole={setRole}
                    />
                )
            case 2:
                return (
                    <UserDataStep
                        name={name}
                        setName={setName}
                        surname={surname}
                        setSurname={setSurname}
                        patronymic={patronymic}
                        setPatronymic={setPatronymic}
                        gender={gender}
                        setGender={setGender}
                        birthDate={birthDate}
                        setBirthDate={setBirthDate}
                    />
                )
            case 3:
                return (
                    <AuthDataStep
                        email={email}
                        setEmail={setEmail}
                        password={password}
                        setPassword={setPassword}
                        confirmPassword={confirmPassword}
                        setConfirmPassword={setConfirmPassword}
                    />
                )
            case 4:
                return (
                    <ConfirmStep
                        isCheked={isChecked}
                        setIsCheked={setIsCheked}
                        name={name}
                        surname={surname}
                        patronymic={patronymic}
                        gender={gender}
                        birthDate={birthDate}
                        email={email}
                        language={language}
                    />
                )
            default:
                return (
                    <>
                        <Field>
                            <FieldLabel htmlFor="email">Почта</FieldLabel>
                            <Input
                                id="email"
                                type="email"
                                placeholder="m@example.com"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                required
                            />
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="password">Пароль</FieldLabel>
                            <Input
                                id="password"
                                type="password"
                                required
                                value={password}
                                onChange={(e) => setPassword(e.target.value)} />
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="confirmPassword">
                                Повторите пароль
                            </FieldLabel>
                            <Input
                                id="confirmPassword"
                                type="password"
                                required
                                value={confirmPassword}
                                onChange={(e) => setConfirmPassword(e.target.value)}
                            />
                        </Field>
                    </>
                )
        }
    }

    return (
        <div className={cn("flex flex-col gap-6", className)} {...props}>
            <Card className="overflow-hidden p-0">
                <CardContent className="grid p-0 md:grid-cols-2 min-h-[60vh] items-center">
                    <div className="relative hidden bg-muted md:block h-full">
                        <img
                            src="/placeholder.svg"
                            alt="Image"
                            className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
                        />
                    </div>
                    <form className="p-6 md:p-8" onSubmit={onSubmit}>
                        <FieldGroup>
                            <div className="flex flex-col items-center gap-2 text-center">
                                <h1 className="text-2xl font-bold">Регистрация</h1>
                                <p className="text-sm text-balance text-muted-foreground">
                                    Чтобы продолжить необходимо создать аккаут
                                </p>
                            </div>
                            <Stepper
                                steps={steps}
                                currentStep={currentStep}
                                onStepClick={setCurrentStep}
                            />
                            {renderFormByStep()}
                            <Field className="flex flex-row">
                                <Field>
                                    <Button
                                        onClick={prevStep}
                                        className={"border-border"}
                                        variant="secondary"
                                        disabled={currentStep === 1}
                                    >
                                        Назад
                                    </Button>
                                </Field>
                                <Field>
                                    {currentStep === steps.length
                                        ? (
                                            <Field>
                                                <Button
                                                    type="submit"
                                                    disabled={checkNextStepDisabled()}
                                                >
                                                    Зарегистрироваться
                                                </Button>
                                            </Field>
                                        )
                                        : (
                                            <Button
                                                onClick={nextStep}
                                                disabled={checkNextStepDisabled()}
                                            >
                                                Далее
                                            </Button>
                                        )
                                    }
                                </Field>
                            </Field>
                            <FieldDescription className="text-center">
                                Уже есть аккаунт? <a className="cursor-pointer" onClick={() => navigate(getRouteLogin())}>Войти</a>
                            </FieldDescription>
                        </FieldGroup>
                    </form>
                </CardContent>
            </Card>
            <FieldDescription className="px-6 text-center">
                Нажимая продолжить вы соглашаетесь с <a href="#">Terms of Service</a>{" "}
                и <a href="#">Privacy Policy</a>.
            </FieldDescription>
        </div>
    )
}
