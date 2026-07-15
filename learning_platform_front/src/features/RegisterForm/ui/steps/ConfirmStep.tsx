import type { UserGenderEnum, UserLanguageEnum } from "@/shared/enums/user"
import { Checkbox } from "@/shared/ui/Checkbox"
import {
    Field,
    FieldDescription,
    FieldLabel,
} from "@/shared/ui/Field"
import { format } from "date-fns"
import { ru } from "date-fns/locale"
import { enumToStringGender, enumToStringLanguage } from "../../utils/utils"

type ConfirmStepProps = {
    isCheked: boolean
    setIsCheked: (isCheked: boolean) => void

    name: string
    surname: string
    patronymic: string
    gender: UserGenderEnum
    birthDate: Date | undefined
    language: UserLanguageEnum
    email: string
}

export function ConfirmStep({
    isCheked,
    setIsCheked,
    name,
    surname,
    patronymic,
    gender,
    birthDate,
    language,
    email,
}: ConfirmStepProps) {
    return (
        <Field className="gap-6">
            <div className="rounded-xl border bg-card p-5">
                <h3 className="mb-4 text-lg font-semibold">
                    Проверьте введённые данные
                </h3>

                <div className="space-y-4">
                    <InfoRow label="Фамилия" value={surname} />
                    <InfoRow label="Имя" value={name} />
                    <InfoRow label="Отчество" value={patronymic} />
                    <InfoRow label="Дата рождения" value={birthDate
                        ? format(birthDate, "dd MMMM yyyy", { locale: ru })
                        : "Не выбрано"} />
                    <InfoRow label="Пол" value={enumToStringGender(gender)} />
                    <InfoRow label="Язык" value={enumToStringLanguage(language)} />
                    <InfoRow label="Почта" value={email} />
                </div>
            </div>

            <div className="flex items-start gap-3 rounded-lg border p-4">
                <Checkbox
                    id="confirm"
                    checked={isCheked}
                    onCheckedChange={(checked: boolean) =>
                        setIsCheked(checked)
                    }
                />

                <div className="space-y-1">
                    <FieldLabel
                        htmlFor="confirm"
                        className="cursor-pointer"
                    >
                        Подтверждаю корректность данных
                    </FieldLabel>

                    <FieldDescription>
                        Я подтверждаю, что вся указанная информация является
                        достоверной.
                    </FieldDescription>
                </div>
            </div>
        </Field>
    )
}

type InfoRowProps = {
    label: string
    value: string
}

function InfoRow({ label, value }: InfoRowProps) {
    return (
        <div className="flex items-center justify-between border-b pb-2 last:border-0 last:pb-0">
            <span className="text-sm text-muted-foreground">
                {label}
            </span>

            <span className="font-medium text-right">
                {value}
            </span>
        </div>
    )
}