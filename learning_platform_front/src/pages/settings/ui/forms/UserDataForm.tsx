import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { updateUserInfo } from "@/entities/user/api/updateUserInfo"
import type { UserFullData, UserInfoRequest } from "@/entities/user/types/types"
import { notificationActions } from "@/features/notifications"
import { UserGenderEnum } from "@/shared/enums/user"
import { cn } from "@/shared/lib/utils"
import { Button } from "@/shared/ui/Button"
import { Calendar } from "@/shared/ui/Calendar"
import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Popover, PopoverContent, PopoverTrigger } from "@/shared/ui/Popover"
import { RadioGroup, RadioGroupItem } from "@/shared/ui/RadioGroup"
import { Textarea } from "@/shared/ui/Textarea"
import { ConfirmButton } from "@/widgets/confirmButton"
import { format } from "date-fns"
import { ru } from "date-fns/locale"
import { CalendarIcon } from "lucide-react"
import { useState } from "react"
import { CiEdit, CiUser } from "react-icons/ci"

type UserDataFormProps = {
    userData: UserFullData
}

export function UserDataForm({
    userData
}: UserDataFormProps) {
    const dispatch = useAppDispatch()
    const [about, setAbout] = useState<string>(userData?.userInfo.about || "");
    const [name, setName] = useState<string>(userData?.userInfo.name || "");
    const [surname, setSurname] = useState<string>(userData?.userInfo.surname || "");
    const [patronymic, setPatronymic] = useState<string>(userData?.userInfo.patronymic || "");
    const [gender, setGender] = useState<UserGenderEnum>(userData?.userInfo.gender || UserGenderEnum.UNKNOWN);
    const [birthDate, setBirthDate] = useState<Date | undefined>(userData?.userInfo.birthDate)

    const onClickConfirm = async () => {
        const request: UserInfoRequest = {
            name: name,
            surname: surname,
            patronymic: patronymic,
            gender: gender,
            birth_date: birthDate,
            about: about
        }

        const response = await dispatch(updateUserInfo(request))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: 'Данные пользователя обновлены',
                type: 'success',
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: 'Не удалось обновить данные пользователя',
                type: 'error',
            }))
        }
    }

    return (
        <form className="border border-border bg-background p-5 lg:p-10 pb-5 rounded-lg space-y-5">
            <div className="flex items-center gap-6 pb-8 border-b">
                <div className="group relative h-20 w-20 lg:h-36 lg:w-36 cursor-pointer overflow-hidden rounded-full border-2">
                    <div className="flex h-full w-full items-center justify-center bg-background">
                        <CiUser className="size-15 lg:size-25" />
                    </div>
                    <div
                        className="
                            absolute inset-0 flex items-center justify-center rounded-full
                            bg-black/30 opacity-0 transition-opacity duration-200 group-hover:opacity-100
                        "
                    >
                        <CiEdit className="text-white" size={42} />
                    </div>
                </div>
                <div className="space-y-1">
                    <h2 className="text-lg lg:text-2xl font-semibold">
                        {name} {surname}
                    </h2>
                    <p className="text-xs lg:text-md text-muted-foreground">
                        {userData.user.email}
                    </p>
                </div>
            </div>
            <div className="flex flex-col space-y-6">
                <div className="flex flex-col lg:flex-row gap-4">
                    <Field>
                        <FieldLabel htmlFor="surname">Фамилия</FieldLabel>
                        <Input
                            id="surname"
                            type="surname"
                            required
                            value={surname}
                            onChange={(e) => setSurname(e.target.value)}
                        />
                    </Field>
                    <Field>
                        <FieldLabel htmlFor="name">Имя</FieldLabel>
                        <Input
                            id="name"
                            type="name"
                            required
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </Field>
                    <Field>
                        <FieldLabel className="text-primary/60" htmlFor="patronymic">Отчество</FieldLabel>
                        <Input
                            id="patronymic"
                            type="patronymic"
                            value={patronymic}
                            onChange={(e) => setPatronymic(e.target.value)}
                        />
                    </Field>
                </div>
                <Field>
                    <FieldLabel>Дата рождения</FieldLabel>
                    <Popover>
                        <PopoverTrigger asChild>
                            <Button
                                size="lg"
                                variant="outline"
                                className={cn(
                                    "justify-start text-left font-normal",
                                    !birthDate && "text-muted-foreground"
                                )}
                            >
                                <CalendarIcon className="mr-2 h-4 w-4" />
                                {birthDate
                                    ? format(birthDate, "dd MMMM yyyy", { locale: ru })
                                    : "Выберите дату"}
                            </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-auto p-0">
                            <Calendar
                                mode="single"
                                selected={birthDate}
                                month={birthDate}
                                onMonthChange={setBirthDate}
                                onSelect={setBirthDate}
                                captionLayout="dropdown"
                                disabled={(date) => date > new Date()}
                            />
                        </PopoverContent>
                    </Popover>
                </Field>
                <Field>
                    <FieldLabel>Пол</FieldLabel>
                    <RadioGroup
                        value={gender}
                        onValueChange={(value) =>
                            setGender(value as UserGenderEnum)
                        }
                        className="flex flex-col lg:flex-row gap-2 lg:gap-6"
                    >
                        <div className="flex items-center space-x-2">
                            <RadioGroupItem
                                value={UserGenderEnum.MALE}
                                id="male"
                            />
                            <label htmlFor="male">Мужской</label>
                        </div>
                        <div className="flex items-center space-x-2">
                            <RadioGroupItem
                                value={UserGenderEnum.FEMALE}
                                id="female"
                            />
                            <label htmlFor="female">Женский</label>
                        </div>
                        <div className="flex items-center space-x-2">
                            <RadioGroupItem
                                value={UserGenderEnum.UNKNOWN}
                                id="unknown"
                            />
                            <label htmlFor="unknown">
                                Не указывать
                            </label>
                        </div>
                    </RadioGroup>
                </Field>
                <Field>
                    <FieldLabel>О себе</FieldLabel>
                    <Textarea
                        placeholder="Расскажите о себе..."
                        value={about}
                        onChange={(e) => setAbout(e.target.value)}
                        className="w-full break-words min-h-50"
                    />
                </Field>
            </div>
            <ConfirmButton
                onClickConfirm={onClickConfirm}
                onClickCancel={() => null}
            />
        </form>
    )
}