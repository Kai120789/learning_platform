import { cn } from "@/shared/lib/utils";
import { Label } from "@/shared/ui/Label";
import { useState } from "react";
import { getUserFullData } from "@/entities/user";
import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks";
import { SectionTabs } from "../types/types";
import { ChangeEmaildForm } from "./forms/ChangeEmailForm";
import { ChangePasswordForm } from "./forms/ChangePasswordForm";
import { SettingsForm } from "./forms/SettingsForm";
import { UserDataForm } from "./forms/UserDataForm";
import { ChangeTgLink } from "./forms/ChangeTgLink";
import { useTranslation } from "react-i18next";
import { FiSettings, FiShield, FiUser } from "react-icons/fi";
import type { IconType } from "react-icons/lib";

type SettingsSection = {
    id: SectionTabs
    titleKey: string
    icon: IconType
}

const sections: SettingsSection[] = [
    { id: SectionTabs.PROFILE, titleKey: "settings.tabs.profile", icon: FiUser },
    { id: SectionTabs.ACCOUNT, titleKey: "settings.tabs.account", icon: FiShield },
    { id: SectionTabs.SETTINGS, titleKey: "settings.tabs.settings", icon: FiSettings },
]

export default function SettingsPage() {
    const { t } = useTranslation()
    const userData = useAppSelector(getUserFullData)

    const [active, setActive] = useState<SectionTabs>(SectionTabs.PROFILE);

    const renderFormByTab = () => {
        if (!userData) {
            return <></>
        }

        switch (active) {
            case SectionTabs.PROFILE:
                return (
                    <UserDataForm
                        userData={userData}
                    />
                )
            case SectionTabs.ACCOUNT:
                return (
                    <>
                        <ChangeTgLink
                            userTgLink={userData?.userInfo.tgLink}
                        />
                        <ChangeEmaildForm
                            userEmail={userData?.user.email}
                        />
                        <ChangePasswordForm />
                    </>
                )
            case SectionTabs.SETTINGS:
                return (
                    <SettingsForm
                        userIs2FaEnabled={userData.userSettings.is2FaEnabled}
                        userIsNotificationEnabled={userData.userSettings.isNotificationsEnabled}
                        userLanguage={userData.userSettings.language}
                    />
                )
            default:
                return (
                    <UserDataForm
                        userData={userData}
                    />
                )
        }
    }

    return (
        <div className="flex flex-col py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <Label className="text-2xl lg:text-4xl">
                    {t("settings.title")}
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    {t("settings.subtitle")}
                </Label>
            </div>
            <div className="flex flex-col lg:flex-row gap-5 lg:gap-10">
                <div>
                    <div className="w-full lg:w-50 shrink-0">
                        <div className="flex flex-row lg:flex-col border border-border lg:border-none rounded-xl">
                            {sections.map((section) => {
                                const Icon = section.icon;

                                return (
                                    <button
                                        key={section.id}
                                        onClick={() => setActive(section.id)}
                                        className={cn(
                                            "flex w-full items-center gap-3 rounded-lg px-4 py-3 text-left transition-colors justify-center lg:justify-start hover:border hover:border-border",
                                            active === section.id
                                                ? "bg-black text-white"
                                                : "hover:bg-muted"
                                        )}
                                    >
                                        <Icon size={18} />
                                        <span className="hidden lg:flex">{t(section.titleKey)}</span>
                                    </button>
                                );
                            })}
                        </div>
                    </div>
                </div>
                <div className="flex flex-col space-y-5 w-full">
                    {renderFormByTab()}
                </div>
            </div>
        </div>
    )
}
