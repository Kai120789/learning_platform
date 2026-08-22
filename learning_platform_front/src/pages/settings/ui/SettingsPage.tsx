import { cn } from "@/shared/lib/utils";
import { Label } from "@/shared/ui/Label";
import { useState } from "react";
import { getUserFullData, useCanEdit } from "@/entities/user";
import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks";
import { SectionTabs } from "../types/types";
import { ChangeEmaildForm } from "./forms/ChangeEmailForm";
import { ChangePasswordForm } from "./forms/ChangePasswordForm";
import { SettingsForm } from "./forms/SettingsForm";
import { UserDataForm } from "./forms/UserDataForm";
import { ChangeTgLink } from "./forms/ChangeTgLink";
import { useTranslation } from "react-i18next";
import { GraduationCap, Settings, Shield, User } from "lucide-react";
import type { LucideIcon } from "lucide-react";
import { TutorTeachingSettings } from "@/widgets/tutorTeachingSettings";

type SettingsSection = {
    id: SectionTabs
    titleKey: string
    icon: LucideIcon
}

const baseSections: SettingsSection[] = [
    { id: SectionTabs.PROFILE, titleKey: "settings.tabs.profile", icon: User },
    { id: SectionTabs.ACCOUNT, titleKey: "settings.tabs.account", icon: Shield },
    { id: SectionTabs.SETTINGS, titleKey: "settings.tabs.settings", icon: Settings },
]

const teachingSection: SettingsSection = {
    id: SectionTabs.TEACHING,
    titleKey: "settings.tabs.teaching",
    icon: GraduationCap,
}

export default function SettingsPage() {
    const { t } = useTranslation()
    const userData = useAppSelector(getUserFullData)
    const canEdit = useCanEdit()
    const sections = canEdit
        ? [
            baseSections[0],
            baseSections[1],
            teachingSection,
            baseSections[2],
        ]
        : baseSections

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
                    <div className="space-y-3">
                        <ChangeTgLink
                            userTgLink={userData?.userInfo.tgLink}
                        />
                        <ChangeEmaildForm
                            userEmail={userData?.user.email}
                        />
                        <ChangePasswordForm />
                    </div>
                )
            case SectionTabs.SETTINGS:
                return (
                    <SettingsForm
                        userIs2FaEnabled={userData.userSettings.is2FaEnabled}
                        userIsNotificationEnabled={userData.userSettings.isNotificationsEnabled}
                        userLanguage={userData.userSettings.language}
                    />
                )
            case SectionTabs.TEACHING:
                return <TutorTeachingSettings />
            default:
                return (
                    <UserDataForm
                        userData={userData}
                    />
                )
        }
    }

    return (
        <div className="flex flex-col py-8 lg:py-10 px-6 lg:px-20 space-y-5">
            <div className="space-y-1">
                <Label className="text-xl lg:text-2xl">
                    {t("settings.title")}
                </Label>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("settings.subtitle")}
                </Label>
            </div>
            <div className="flex flex-col lg:flex-row gap-4 lg:gap-8">
                <aside className="w-full lg:w-44 shrink-0">
                    <nav className="flex flex-row lg:flex-col gap-1 border border-border lg:border-none rounded-lg lg:sticky lg:top-20">
                        {sections.map((section) => {
                            const Icon = section.icon;

                            return (
                                <button
                                    key={section.id}
                                    type="button"
                                    onClick={() => setActive(section.id)}
                                    className={cn(
                                        "flex w-full items-center gap-2 rounded-md px-2.5 py-2 text-sm text-left transition-colors justify-center lg:justify-start",
                                        active === section.id
                                            ? "bg-primary text-primary-foreground"
                                            : "text-muted-foreground hover:bg-muted hover:text-foreground"
                                    )}
                                >
                                    <Icon className="size-[15px]" />
                                    <span className="hidden lg:flex">{t(section.titleKey)}</span>
                                </button>
                            );
                        })}
                    </nav>
                </aside>
                <div className="min-w-0 flex-1 space-y-3">
                    {renderFormByTab()}
                </div>
            </div>
        </div>
    )
}
