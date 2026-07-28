import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { ConfirmButton } from "@/widgets/confirmButton";
import { useState } from "react"
import { useTranslation } from "react-i18next"

export function ChangePasswordForm() {
    const { t } = useTranslation()
    const [password, setPassword] = useState<string>("");
    const [confirmPassword, setConfirmPassword] = useState<string>("");

    return (
        <div className="flex flex-col border border-border bg-background p-4 lg:p-5 rounded-lg space-y-3">
            <div className="grid gap-4 sm:grid-cols-2">
                <Field>
                    <FieldLabel htmlFor="password">{t("settings.password")}</FieldLabel>
                    <Input
                        id="password"
                        type="password"
                        required
                        value={password}
                        onChange={(e) => setPassword(e.target.value)} />
                </Field>
                <Field>
                    <FieldLabel htmlFor="confirmPassword">
                        {t("settings.repeatPassword")}
                    </FieldLabel>
                    <Input
                        id="confirmPassword"
                        type="password"
                        required
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                    />
                </Field>
            </div>
            <ConfirmButton
                onClickConfirm={() => null}
                onClickCancel={() => null}
            />
        </div>
    )
}
