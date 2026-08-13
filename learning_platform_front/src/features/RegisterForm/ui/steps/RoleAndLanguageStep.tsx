import { UserLanguageEnum } from "@/shared/enums/user";
import { cn } from "@/shared/lib/utils";
import { Field, FieldLabel } from "@/shared/ui/Field";
import { NativeSelect } from "@/shared/ui/NativeSelect";
import { useTranslation } from "react-i18next";
import { CheckCircle2, GraduationCap, Presentation } from "lucide-react";
import { RegisterRoleEnum } from "../../types/types";
import { languageToTranslationKey } from "../../utils/utils";

type RoleAndLanguageStepProps = {
    language: UserLanguageEnum
    setLanguage: (language: UserLanguageEnum) => void
    role: RegisterRoleEnum
    setRole: (role: RegisterRoleEnum) => void
}

export function RoleAndLanguageStep({
    language,
    setLanguage,
    role,
    setRole
}: RoleAndLanguageStepProps) {
    const { t } = useTranslation()

    return (
        <>
            <Field>
                <FieldLabel>{t("auth.register.steps.whoAreYou")}</FieldLabel>

                <div className="grid grid-cols-2 gap-4">
                    <button
                        type="button"
                        onClick={() => setRole(RegisterRoleEnum.STUDENT)}
                        className={cn(
                            "relative flex flex-col rounded-xl border p-5 text-left transition-all duration-200",
                            "hover:border-primary hover:bg-accent",
                            role === RegisterRoleEnum.STUDENT &&
                            "border-primary bg-primary/5 ring-2 ring-primary/20"
                        )}
                    >
                        {role === RegisterRoleEnum.STUDENT && (
                            <CheckCircle2 className="absolute right-3 top-3 h-5 w-5 text-primary" />
                        )}

                        <GraduationCap className="mb-4 h-10 w-10 text-primary" />

                        <span className="font-semibold">
                            {t("auth.register.steps.student")}
                        </span>

                        <span className="mt-2 text-sm text-muted-foreground">
                            {t("auth.register.steps.studentHint")}
                        </span>
                    </button>

                    <button
                        type="button"
                        onClick={() => setRole(RegisterRoleEnum.TUTOR)}
                        className={cn(
                            "relative flex flex-col rounded-xl border p-5 text-left transition-all duration-200",
                            "hover:border-primary hover:bg-accent",
                            role === RegisterRoleEnum.TUTOR &&
                            "border-primary bg-primary/5 ring-2 ring-primary/20"
                        )}
                    >
                        {role === RegisterRoleEnum.TUTOR && (
                            <CheckCircle2 className="absolute right-3 top-3 h-5 w-5 text-primary" />
                        )}

                        <Presentation className="mb-4 h-10 w-10 text-primary" />

                        <span className="font-semibold">
                            {t("auth.register.steps.tutor")}
                        </span>

                        <span className="mt-2 text-sm text-muted-foreground">
                            {t("auth.register.steps.tutorHint")}
                        </span>
                    </button>
                </div>
            </Field>
            <Field>
                <FieldLabel htmlFor="language">{t("auth.register.steps.language")}</FieldLabel>
                <NativeSelect
                    id="language"
                    value={language}
                    onChange={(e) => setLanguage(e.target.value as UserLanguageEnum)}
                >
                    {Object.values(UserLanguageEnum).map((language) => (
                        <option key={language} value={language}>
                            {t(languageToTranslationKey(language))}
                        </option>
                    ))}
                </NativeSelect>
            </Field>
        </>
    )
}
