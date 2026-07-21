import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { RiTelegramFill } from "react-icons/ri"
import { MdDelete } from "react-icons/md";
import type { GroupUser } from "@/entities/group/types/types";
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks";
import { removeUserFromGroup } from "@/entities/group/api/removeUserFromGroup";
import { notificationActions } from "@/features/notifications";


type GroupUserItemProps = {
    user: GroupUser
    groupID: number
}

export function GroupUserItem({
    user,
    groupID
}: GroupUserItemProps) {
    const dispatch = useAppDispatch()

    const onClickRemoveUser = async () => {
        const response = await dispatch(removeUserFromGroup({ groupID: groupID, userID: user.id }))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: "Пользователь успешно удален из группы",
                type: "success"
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: "Не удалось удалить пользователя из группы",
                type: "error"
            }))
        }
    }

    return (
        <div className="flex items-center justify-between rounded-lg border p-3 transition-colors hover:bg-muted/50">
            <div className="flex items-center gap-3">
                <Avatar>
                    <AvatarFallback>
                        {user.name[0] + user.surname[0]}
                    </AvatarFallback>
                </Avatar>

                <div>
                    <p className="font-medium leading-none">
                        {`${user.name} ${user.surname}`}
                    </p>

                    <p className="text-sm text-muted-foreground">
                        {user.tgUsername}
                    </p>
                </div>
            </div>

            <div className="flex flex-row gap-1">
                <RiTelegramFill className="cursor-pointer" size={25} />
                <MdDelete onClick={onClickRemoveUser} className="cursor-pointer text-destructive/70" size={25} />
            </div>
        </div>
    )
}
