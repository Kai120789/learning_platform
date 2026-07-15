import { UserLanguageEnum } from "@/shared/enums/user";
import { cn } from "@/shared/lib/utils";
import { Field, FieldLabel } from "@/shared/ui/Field";
import { IoCheckmarkCircle } from "react-icons/io5";
import { PiChalkboardTeacher, PiStudent } from "react-icons/pi";
import { RegisterRoleEnum } from "../../types/types";
import { enumToStringLanguage } from "../../utils/utils";

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
    return (
        <>
            <Field>
                <FieldLabel>Кто вы?</FieldLabel>

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
                            <IoCheckmarkCircle className="absolute right-3 top-3 h-5 w-5 text-primary" />
                        )}

                        <PiStudent className="mb-4 h-10 w-10 text-primary" />

                        <span className="font-semibold">
                            Ученик
                        </span>

                        <span className="mt-2 text-sm text-muted-foreground">
                            Проходить курсы, выполнять задания и отслеживать прогресс.
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
                            <IoCheckmarkCircle className="absolute right-3 top-3 h-5 w-5 text-primary" />
                        )}

                        <PiChalkboardTeacher className="mb-4 h-10 w-10 text-primary" />

                        <span className="font-semibold">
                            Преподаватель
                        </span>

                        <span className="mt-2 text-sm text-muted-foreground">
                            Создавать курсы, публиковать материалы и проверять задания.
                        </span>
                    </button>
                </div>
            </Field>
            <Field>
                <FieldLabel htmlFor="language">Язык</FieldLabel>
                <select
                    className="border border-input rounded-lg p-2"
                    value={language}
                    onChange={(e) => setLanguage(e.target.value as UserLanguageEnum)}
                >
                    {Object.values(UserLanguageEnum).map((language) => (
                        <option key={language} value={language}>
                            {enumToStringLanguage(language)}
                        </option>
                    ))}
                </select>
            </Field>
        </>
    )
}