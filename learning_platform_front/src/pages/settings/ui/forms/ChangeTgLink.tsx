import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { ConfirmButton } from "@/widgets/confirmButton"
import { useState } from "react"
import { useTranslation } from "react-i18next"
import { RiTelegramFill } from "react-icons/ri"

// TODO: поменять link на username

type ChangeTgLinkProps = {
    userTgLink?: string
}

export function ChangeTgLink({
    userTgLink
}: ChangeTgLinkProps) {
    const { t } = useTranslation()
    const [tgLink, setTgLink] = useState<string>(userTgLink || "");

    return (
        <div className="flex flex-col border border-border bg-background p-4 lg:p-5 rounded-lg space-y-3">
            <Field>
                <FieldLabel htmlFor="tgLink" className="flex items-center gap-1.5">
                    <RiTelegramFill size={18} />
                    {t("settings.telegram")}
                </FieldLabel>
                <Input
                    id="tgLink"
                    type="tgLink"
                    placeholder="https://t.me/username"
                    value={tgLink}
                    onChange={(e) => setTgLink(e.target.value)}
                    required
                />
            </Field>
            <ConfirmButton
                onClickConfirm={() => null}
                onClickCancel={() => null}
            />
        </div>
    )
}
