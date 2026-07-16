import { UserGenderEnum } from "@/shared/enums/user"
import { cn } from "@/shared/lib/utils"
import { Button } from "@/shared/ui/Button"
import { Calendar } from "@/shared/ui/Calendar"
import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { PopoverTrigger, PopoverContent, Popover } from "@/shared/ui/Popover"
import { RadioGroup, RadioGroupItem } from "@/shared/ui/RadioGroup"
import { format } from "date-fns"
import { CalendarIcon } from "lucide-react"

import { ru } from "date-fns/locale"

type UserDataStepProps = {
    name: string
    setName: (name: string) => void
    surname: string
    setSurname: (surname: string) => void
    patronymic: string
    setPatronymic: (patronymic: string) => void
    gender: UserGenderEnum
    setGender: (gender: UserGenderEnum) => void
    birthDate: Date | undefined
    setBirthDate: (birthDate: Date | undefined) => void
}

export function UserDataStep({
    name,
    setName,
    surname,
    setSurname,
    patronymic,
    setPatronymic,
    gender,
    setGender,
    birthDate,
    setBirthDate,
}: UserDataStepProps) {
    return (
        <>
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
                <FieldLabel className="text-primary/60" htmlFor="patronymic">Отчество (не обязательно)</FieldLabel>
                <Input
                    id="patronymic"
                    type="patronymic"
                    value={patronymic}
                    onChange={(e) => setPatronymic(e.target.value)}
                />
            </Field>
            <Field>
                <FieldLabel>Дата рождения</FieldLabel>
                <Popover>
                    <PopoverTrigger asChild>
                        <Button
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
                            month={birthDate}
                            onMonthChange={setBirthDate}
                            selected={birthDate}
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
                    className="flex gap-6"
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
        </>
    )
}