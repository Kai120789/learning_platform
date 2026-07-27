import type { UserGenderEnum, UserLanguageEnum } from "@/shared/enums/user"
import { Checkbox } from "@/shared/ui/Checkbox"
import {
    Field,
    FieldDescription,
    FieldLabel,
} from "@/shared/ui/Field"
import { format } from "date-fns"
import { enUS, ru } from "date-fns/locale"
import { useTranslation } from "react-i18next"
import { genderToTranslationKey, languageToTranslationKey } from "../../utils/utils"
import { InfoRow } from "@/shared/ui/InfoRow"

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
    const { t, i18n } = useTranslation()
    const dateLocale = i18n.language === "ru" ? ru : enUS

    return (
        <Field className="gap-6">
            <div className="rounded-xl border bg-card p-5">
                <h3 className="mb-4 text-lg font-semibold">
                    {t("auth.register.steps.checkData")}
                </h3>

                <div className="space-y-4">
                    <InfoRow label={t("auth.register.steps.surname")} value={surname} />
                    <InfoRow label={t("auth.register.steps.name")} value={name} />
                    <InfoRow label={t("auth.register.steps.patronymic")} value={patronymic} />
                    <InfoRow label={t("auth.register.steps.birthDate")} value={birthDate
                        ? format(birthDate, "dd MMMM yyyy", { locale: dateLocale })
                        : t("auth.register.steps.notChosen")} />
                    <InfoRow label={t("auth.register.steps.gender")} value={t(genderToTranslationKey(gender))} />
                    <InfoRow label={t("auth.register.steps.language")} value={t(languageToTranslationKey(language))} />
                    <InfoRow label={t("auth.register.steps.email")} value={email} />
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
                        {t("auth.register.steps.confirmCorrect")}
                    </FieldLabel>

                    <FieldDescription>
                        {t("auth.register.steps.confirmHint")}
                    </FieldDescription>
                </div>
            </div>
        </Field>
    )
}
