import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { updateUserSettings } from "@/entities/user";
import type { UserSettingsRequest } from "@/entities/user";
import { notificationActions } from "@/features/notifications";
import { languageToTranslationKey } from "@/features/registerForm/utils/utils";
import { UserLanguageEnum } from "@/shared/enums/user";
import { Switch } from "@/shared/ui/Switch";
import { NativeSelect } from "@/shared/ui/NativeSelect";
import { ConfirmButton } from "@/widgets/confirmButton";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { Languages } from "lucide-react";

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
    const { t, i18n } = useTranslation();
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
                message: t("settings.updateSuccess"),
                type: 'success',
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("settings.updateError"),
                type: 'error',
            }))
        }
    }

    return (
        <div className="flex flex-col border border-border bg-background rounded-lg overflow-hidden">
            <div className="divide-y divide-border">
                <div className="flex items-center justify-between gap-4 px-4 py-3.5">
                    <div className="min-w-0 space-y-0.5">
                        <h3 className="text-sm font-medium flex items-center gap-1.5">
                            {t("settings.language")}
                            <Languages className="size-4 text-muted-foreground" />
                        </h3>
                    </div>
                    <NativeSelect
                        containerClassName="w-auto shrink-0"
                        className="h-auto w-auto min-w-28 py-1.5"
                        value={language}
                        onChange={(e) => {
                            i18n.changeLanguage((e.target.value as UserLanguageEnum).toLowerCase())
                            setLanguage(e.target.value as UserLanguageEnum)
                        }}
                    >
                        {Object.values(UserLanguageEnum).map((item) => (
                            <option key={item} value={item}>
                                {t(languageToTranslationKey(item))}
                            </option>
                        ))}
                    </NativeSelect>
                </div>

                <div className="flex items-center justify-between gap-4 px-4 py-3.5">
                    <div className="min-w-0 space-y-0.5">
                        <h3 className="text-sm font-medium">
                            {t("settings.twoFa")}
                        </h3>
                        <p className="text-xs text-muted-foreground">
                            {t("settings.twoFaHint")}
                        </p>
                    </div>
                    <Switch
                        checked={is2FaEnabled}
                        onCheckedChange={() => setIs2FaEnabled(!is2FaEnabled)}
                    />
                </div>

                <div className="flex items-center justify-between gap-4 px-4 py-3.5">
                    <div className="min-w-0 space-y-0.5">
                        <h3 className="text-sm font-medium">
                            {t("settings.notifications")}
                        </h3>
                        <p className="text-xs text-muted-foreground">
                            {t("settings.notificationsHint")}
                        </p>
                    </div>
                    <Switch
                        checked={isNotificationEnabled}
                        onCheckedChange={() => setIsNotificationEnabled(!isNotificationEnabled)}
                    />
                </div>
            </div>

            <div className="border-t border-border px-4 py-3">
                <ConfirmButton
                    onClickConfirm={onClickConfirm}
                    onClickCancel={() => null}
                />
            </div>
        </div>
    )
}
