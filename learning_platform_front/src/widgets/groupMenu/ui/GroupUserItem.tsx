import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { Button, buttonVariants } from "@/shared/ui/Button"
import { Send, Trash2 } from "lucide-react"
import type { GroupUser } from "@/entities/group"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { removeUserFromGroup } from "@/entities/group"
import { notificationActions } from "@/features/notifications"
import { useTranslation } from "react-i18next"
import { useCanEdit } from "@/entities/user"
import { cn } from "@/shared/lib/utils"

type GroupUserItemProps = {
    user: GroupUser
    groupID: number
}

function telegramHandle(username?: string) {
    if (!username) return ""
    return username.replace(/^@/, "")
}

export function GroupUserItem({
    user,
    groupID,
}: GroupUserItemProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const isCanEdit = useCanEdit()
    const handle = telegramHandle(user.tgUsername)

    const onClickRemoveUser = async () => {
        const response = await dispatch(removeUserFromGroup({ groupID: groupID, userID: user.id }))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("groups.removeUserSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("groups.removeUserError"),
                type: "error",
            }))
        }
    }

    return (
        <div className="flex items-center justify-between gap-3 py-2.5">
            <div className="flex min-w-0 items-center gap-3">
                <Avatar className="size-8 shrink-0">
                    <AvatarFallback>
                        {`${user.name[0] ?? ""}${user.surname[0] ?? ""}`.toUpperCase()}
                    </AvatarFallback>
                </Avatar>
                <div className="min-w-0">
                    <div className="truncate text-sm font-medium">
                        {`${user.name} ${user.surname}`}
                    </div>
                    <div className="truncate text-xs text-muted-foreground">
                        {handle ? `@${handle}` : t("groups.noUsername")}
                    </div>
                </div>
            </div>
            <div className="flex shrink-0 items-center">
                {handle && (
                    <a
                        href={`https://t.me/${handle}`}
                        target="_blank"
                        rel="noreferrer"
                        className={cn(buttonVariants({ variant: "ghost", size: "icon-sm" }))}
                    >
                        <Send className="size-3.5" />
                    </a>
                )}
                {isCanEdit && (
                    <Button
                        type="button"
                        size="icon"
                        variant="ghost"
                        className="text-destructive hover:bg-destructive/10 hover:text-destructive"
                        onClick={onClickRemoveUser}
                    >
                        <Trash2 className="size-4" />
                    </Button>
                )}
            </div>
        </div>
    )
}
