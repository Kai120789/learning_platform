import { cn } from "@/shared/lib/utils";
import { Label } from "@/shared/ui/Label";
import { useState } from "react";
import { getUserFullData } from "@/entities/user/selectors/userSelectors";
import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks";
import { Sections, SectionTabs } from "../types/types";
import { ChangeEmaildForm } from "./forms/ChangeEmailForm";
import { ChangePasswordForm } from "./forms/ChangePasswordForm";
import { SettingsForm } from "./forms/SettingsForm";
import { UserDataForm } from "./forms/UserDataForm";

export default function SettingsPage() {
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
            case SectionTabs.SECURITY:
                return (
                    <>
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
                    Редактирование профиля
                </Label>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    Управляйте своими личными данными и настройками
                </Label>
            </div>
            <div className="flex flex-col lg:flex-row gap-5 lg:gap-10">
                <div>
                    <div className="w-full lg:w-50 shrink-0">
                        <div className="flex flex-row lg:flex-col border border-border lg:border-none rounded-xl">
                            {Sections.map((section) => {
                                const Icon = section.Icon;

                                return (
                                    <button
                                        key={section.ID}
                                        onClick={() => setActive(section.ID)}
                                        className={cn(
                                            "flex w-full items-center gap-3 rounded-lg px-4 py-3 text-left transition-colors justify-center lg:justify-start hover:border hover:border-border",
                                            active === section.ID
                                                ? "bg-black text-white"
                                                : "hover:bg-muted"
                                        )}
                                    >
                                        <Icon size={18} />
                                        <span className="hidden lg:flex">{section.Title}</span>
                                    </button>
                                );
                            })}
                        </div>
                    </div>
                </div>
                <div className="flex flex-col space-y-5 w-full">
                    {renderFormByTab()}
                </div>
            </div >
        </div >
    )
}
