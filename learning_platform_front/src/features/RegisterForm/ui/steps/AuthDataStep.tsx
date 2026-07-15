import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"

type AuthDataStepProps = {
    email: string
    setEmail: (email: string) => void
    password: string
    setPassword: (password: string) => void
    confirmPassword: string
    setConfirmPassword: (confirmPassword: string) => void
}

export function AuthDataStep({
    email,
    setEmail,
    password,
    setPassword,
    confirmPassword,
    setConfirmPassword,
}: AuthDataStepProps) {
    return (
        <>
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
        </>
    )
}