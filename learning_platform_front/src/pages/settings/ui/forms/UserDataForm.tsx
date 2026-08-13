import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { updateUserInfo } from "@/entities/user"
import type { UserFullData, UserInfoRequest } from "@/entities/user"
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
import { CalendarIcon, Pencil, User } from "lucide-react"
import { useState } from "react"
import { useTranslation } from "react-i18next"

type UserDataFormProps = {
    userData: UserFullData
}

export function UserDataForm({
    userData
}: UserDataFormProps) {
    const { t } = useTranslation()
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
                message: t("settings.userUpdateSuccess"),
                type: 'success',
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("settings.userUpdateError"),
                type: 'error',
            }))
        }
    }

    return (
        <form className="border border-border bg-background p-4 lg:p-5 rounded-lg space-y-4">
            <div className="flex items-center gap-4 pb-4 border-b">
                <div className="group relative h-16 w-16 lg:h-20 lg:w-20 cursor-pointer overflow-hidden rounded-full border">
                    <div className="flex h-full w-full items-center justify-center bg-background">
                        <User className="size-8 lg:size-10" />
                    </div>
                    <div
                        className="
                            absolute inset-0 flex items-center justify-center rounded-full
                            bg-black/30 opacity-0 transition-opacity duration-200 group-hover:opacity-100
                        "
                    >
                        <Pencil className="size-[22px] text-white" />
                    </div>
                </div>
                <div className="space-y-0.5">
                    <h2 className="text-sm lg:text-base font-semibold">
                        {name} {surname}
                    </h2>
                    <p className="text-xs text-muted-foreground">
                        {userData.user.email}
                    </p>
                </div>
            </div>
            <div className="flex flex-col space-y-4">
                <div className="flex flex-col lg:flex-row gap-4">
                    <Field className="flex-1">
                        <FieldLabel htmlFor="surname">{t("settings.surname")}</FieldLabel>
                        <Input
                            id="surname"
                            type="surname"
                            required
                            value={surname}
                            onChange={(e) => setSurname(e.target.value)}
                        />
                    </Field>
                    <Field className="flex-1">
                        <FieldLabel htmlFor="name">{t("settings.name")}</FieldLabel>
                        <Input
                            id="name"
                            type="name"
                            required
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </Field>
                    <Field className="flex-1">
                        <FieldLabel className="text-primary/60" htmlFor="patronymic">{t("settings.patronymic")}</FieldLabel>
                        <Input
                            id="patronymic"
                            type="patronymic"
                            value={patronymic}
                            onChange={(e) => setPatronymic(e.target.value)}
                        />
                    </Field>
                </div>
                <Field>
                    <FieldLabel>{t("settings.birthDate")}</FieldLabel>
                    <Popover>
                        <PopoverTrigger asChild>
                            <Button
                                variant="outline"
                                className={cn(
                                    "justify-start text-left font-normal w-full",
                                    !birthDate && "text-muted-foreground"
                                )}
                            >
                                <CalendarIcon className="mr-2 h-4 w-4" />
                                {birthDate
                                    ? format(birthDate, "dd MMMM yyyy", { locale: ru })
                                    : t("settings.pickDate")}
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
                    <FieldLabel>{t("settings.gender")}</FieldLabel>
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
                        className="w-full break-words min-h-28 text-sm"
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