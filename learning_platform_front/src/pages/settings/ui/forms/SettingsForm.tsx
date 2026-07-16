import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { updateUserSettings } from "@/entities/user/api/updateUserSettings";
import type { UserSettingsRequest } from "@/entities/user/types/types";
import { notificationActions } from "@/features/notifications";
import { enumToStringLanguage } from "@/features/registerForm/utils/utils";
import { UserLanguageEnum } from "@/shared/enums/user";
import { Field } from "@/shared/ui/Field";
import { Switch } from "@/shared/ui/Switch";
import { ConfirmButton } from "@/widgets/confirmButton";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { FaLanguage } from "react-icons/fa6";

type SettingsFormProps = {
    userIs2FaEnabled: boolean
    userIsNotificationEnabled: boolean
    userLanguage: UserLanguageEnum
}

export function SettingsForm({
    userIs2FaEnabled,
    userIsNotificationEnabled,
    userLanguage
}: SettingsFormProps) {
    const { i18n } = useTranslation();
    const [is2FaEnabled, setIs2FaEnabled] = useState<boolean>(userIs2FaEnabled)
    const [isNotificationEnabled, setIsNotificationEnabled] = useState<boolean>(userIsNotificationEnabled)
    const [language, setLanguage] = useState<UserLanguageEnum>(userLanguage)
    const dispatch = useAppDispatch()

    const onClickConfirm = async () => {
        const request: UserSettingsRequest = {
            is_notifications_enabled: isNotificationEnabled,
            is_2_fa_enabled: is2FaEnabled,
            language: language,
        }

        const response = await dispatch(updateUserSettings(request))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: 'Настройки обновлены',
                type: 'success',
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: 'Не удалось обновить насйтроки',
                type: 'error',
            }))
        }
    }

    return (
        <div className="flex flex-col border border-border bg-background p-5 lg:p-10 pb-5 rounded-lg space-y-5">
            <Field className="flex flex-row">
                <div className="flex items-center justify-between rounded-lg border p-4">
                    <div>
                        <h3 className="font-medium flex flex-row gap-2 items-center">
                            Язык<FaLanguage className="size-5 lg:size-8" />
                        </h3>
                    </div>
                    <select
                        className="border border-input rounded-lg p-2"
                        value={language}
                        onChange={(e) => {
                            i18n.changeLanguage((e.target.value as UserLanguageEnum).toLowerCase())
                            setLanguage(e.target.value as UserLanguageEnum)
                        }}
                    >
                        {Object.values(UserLanguageEnum).map((language) => (
                            <option key={language} value={language}>
                                {enumToStringLanguage(language)}
                            </option>
                        ))}
                    </select>
                </div>
            </Field>
            <Field className="flex flex-row">
                <div className="flex items-center justify-between rounded-lg border p-4">
                    <div>
                        <h3 className="font-medium">
                            Двухфакторная аутентификация
                        </h3>

                        <p className="text-sm text-muted-foreground">
                            Запрашивать дополнительный код при входе.
                        </p>
                    </div>
                    <Switch
                        checked={is2FaEnabled}
                        onCheckedChange={() => setIs2FaEnabled(!is2FaEnabled)}
                    />
                </div>
            </Field>
            <Field className="flex flex-row">
                <div className="flex items-center justify-between rounded-lg border p-4">
                    <div>
                        <h3 className="font-medium">
                            Уведомления
                        </h3>

                        <p className="text-sm text-muted-foreground">
                            Присылать уведомления о событиях
                        </p>
                    </div>
                    <Switch
                        checked={isNotificationEnabled}
                        onCheckedChange={() => setIsNotificationEnabled(!isNotificationEnabled)}
                    />
                </div>
            </Field>
            <ConfirmButton
                onClickConfirm={onClickConfirm}
                onClickCancel={() => null}
            />
        </div>
    )
}