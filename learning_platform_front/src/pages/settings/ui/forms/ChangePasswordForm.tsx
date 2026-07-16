import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { ConfirmButton } from "@/widgets/confirmButton";
import { useState } from "react"

export function ChangePasswordForm() {
    const [password, setPassword] = useState<string>("");
    const [confirmPassword, setConfirmPassword] = useState<string>("");

    return (
        <div className="flex flex-col border border-border bg-background p-5 lg:p-10 pb-5 rounded-lg space-y-5">
            <div className="flex flex-col lg:flex-row gap-4 justify-end">
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
            </div>
            <ConfirmButton
                onClickConfirm={() => null}
                onClickCancel={() => null}
            />
        </div>
    )
}