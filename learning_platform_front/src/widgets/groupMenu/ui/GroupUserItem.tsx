import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { RiTelegramFill } from "react-icons/ri"
import { MdDelete } from "react-icons/md";
import type { GroupUser } from "@/entities/group";
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { removeUserFromGroup } from "@/entities/group";
import { notificationActions } from "@/features/notifications";
import { useTranslation } from "react-i18next";
import { useCanEdit } from "@/entities/user";

type GroupUserItemProps = {
    user: GroupUser
    groupID: number
}

export function GroupUserItem({
    user,
    groupID
}: GroupUserItemProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const isCanEdit = useCanEdit()

    const onClickRemoveUser = async () => {
        const response = await dispatch(removeUserFromGroup({ groupID: groupID, userID: user.id }))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("groups.removeUserSuccess"),
                type: "success"
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("groups.removeUserError"),
                type: "error"
            }))
        }
    }

    return (
        <div className="flex min-w-0 items-center justify-between gap-3 overflow-hidden py-3 transition-colors hover:bg-muted/40">
            <div className="flex min-w-0 flex-1 items-center gap-3 overflow-hidden">
                <Avatar className="shrink-0">
                    <AvatarFallback>
                        {user.name[0] + user.surname[0]}
                    </AvatarFallback>
                </Avatar>

                <div className="min-w-0 flex-1 overflow-hidden">
                    <p className="truncate text-base font-medium leading-tight">
                        {`${user.name} ${user.surname}`}
                    </p>

                    {user.tgUsername && (
                        <p className="mt-1 truncate text-sm text-muted-foreground">
                            {user.tgUsername}
                        </p>
                    )}
                </div>
            </div>

            <div className="flex shrink-0 flex-row items-center gap-1.5 pr-4">
                <RiTelegramFill className="cursor-pointer" size={24} />
                {isCanEdit && (
                    <MdDelete
                        onClick={onClickRemoveUser}
                        className="cursor-pointer text-destructive/70"
                        size={24}
                    />
                )}
            </div>
        </div>
    )
}
