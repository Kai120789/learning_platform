import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { ConfirmButton } from "@/widgets/confirmButton"
import { useState } from "react"
import { RiTelegramFill } from "react-icons/ri"

// TODO: поменять link на username

type ChangeTgLinkProps = {
    userTgLink?: string
}

export function ChangeTgLink({
    userTgLink
}: ChangeTgLinkProps) {
    const [tgLink, setTgLink] = useState<string>(userTgLink || "");

    return (
        <div className="flex flex-col border border-border bg-background p-5 lg:p-10 pb-5 rounded-lg space-y-5">
            <div className="flex flex-row gap-4 justify-end">
                <Field>
                    <FieldLabel htmlFor="tgLink"><RiTelegramFill size={25} />Телеграм</FieldLabel>
                    <Input
                        id="tgLink"
                        type="tgLink"
                        placeholder="https://t.me/username"
                        value={tgLink}
                        onChange={(e) => setTgLink(e.target.value)}
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