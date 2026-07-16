import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { ConfirmButton } from "@/widgets/confirmButton"
import { useState } from "react"

type ChangeEmailFormProps = {
    userEmail: string
}

export function ChangeEmaildForm({
    userEmail
}: ChangeEmailFormProps) {
    const [email, setEmail] = useState<string>(userEmail);

    return (
        <div className="flex flex-col border border-border bg-background p-5 lg:p-10 pb-5 rounded-lg space-y-5">
            <div className="flex flex-row gap-4 justify-end">
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
            </div>
            <ConfirmButton
                onClickConfirm={() => null}
                onClickCancel={() => null}
            />
        </div>
    )
}