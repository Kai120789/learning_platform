import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { deleteGroup } from "@/entities/group/api/deleteGroup"
import type { GroupData } from "@/entities/group/types/types"
import { notificationActions } from "@/features/notifications"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Separator } from "@/shared/ui/Separator"
import { Textarea } from "@/shared/ui/Textarea"
import { GroupUserItem } from "@/widgets/groupMenu"

type GroupModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    group: GroupData
}

export function GroupModal({
    isOpen,
    setIsOpen,
    group
}: GroupModalProps) {
    const dispatch = useAppDispatch()

    const onClickDelete = async () => {
        const response = await dispatch(deleteGroup(group.id))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: "Группа успешно удалена",
                type: "success"
            }))
            setIsOpen(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: "Не удалось удалить группу",
                type: "error"
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent>

                <DialogHeader>
                    <DialogTitle className="text-xl text-left line-clamp-2 pr-10">
                        {group.title}
                    </DialogTitle>
                </DialogHeader>

                <div className="flex gap-2">
                    <Badge variant="outline" className="bg-muted">
                        {group.subject.title}
                    </Badge>
                    <Badge variant="default">
                        {group.subject.type}
                    </Badge>
                </div>

                <div className="mt-1">
                    <Textarea
                        disabled
                        className="w-full break-words min-h-50"
                    >
                        {group.description}
                    </Textarea>

                </div>
                <Separator className="my-1" />

                <div className="flex flex-row items-center justify-between">
                    <div className="font-medium">
                        Пользователи
                    </div>
                    <Button>
                        Добавить
                    </Button>
                </div>

                <div className="space-y-2">
                    {group.users
                        ? group.users.map((user) => (
                            <GroupUserItem key={user.id} user={user} groupID={group.id} />
                        ))
                        : <div className="text-muted-foreground">Добавьте первых учеников в группу</div>
                    }
                </div>
                <Separator className="my-1" />

                <div className="flex w-full justify-between gap-2">
                    <Button variant="outline">Редактировать</Button>
                    <Button
                        variant="destructive"
                        className="border border-destructive"
                        onClick={onClickDelete}
                    >
                        Удалить
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    )
}