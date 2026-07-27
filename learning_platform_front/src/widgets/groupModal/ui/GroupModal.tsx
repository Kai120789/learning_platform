import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { deleteGroup, updateGroup } from "@/entities/group"
import type { GroupData } from "@/entities/group"
import { notificationActions } from "@/features/notifications"
import { Badge } from "@/shared/ui/Badge"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"
import { Textarea } from "@/shared/ui/Textarea"
import { AddUsersToGroupModal } from "@/widgets/addUsersToGroupModal"
import { GroupUserItem } from "@/widgets/groupMenu"
import { useState } from "react"
import { useTranslation } from "react-i18next"

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
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const [isEditMode, setIsEditMode] = useState<boolean>(false)
    const [isAddUsersOpen, setIsAddUsersOpen] = useState<boolean>(false)
    const [title, setTitle] = useState<string>(group.title)
    const [description, setDescription] = useState<string>(group.description)

    const onClickDelete = async () => {
        const response = await dispatch(deleteGroup(group.id))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("groups.deleteSuccess"),
                type: "success"
            }))
            setIsOpen(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("groups.deleteError"),
                type: "error"
            }))
        }
    }

    const onClickUpdate = async () => {
        const request = {
            title: title,
            description: description,
        }

        const response = await dispatch(updateGroup({
            groupID: group.id,
            request: request
        }))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("groups.updateSuccess"),
                type: "success"
            }))
            setIsEditMode(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("groups.updateError"),
                type: "error"
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle className="text-xl text-left line-clamp-2 pr-10">
                        {isEditMode
                            ? (
                                <Input
                                    value={title}
                                    onChange={(e) => setTitle(e.target.value)}
                                />
                            )
                            : (
                                <>
                                    {title}
                                </>
                            )
                        }
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
                        disabled={!isEditMode}
                        className="w-full break-words min-h-50"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                </div>
                {!isEditMode && (
                    <>
                        <Separator className="my-1" />

                        <div className="flex flex-row items-center justify-between">
                            <div className="font-medium">
                                {t("groups.users")}
                            </div>
                            <Button size="xs" onClick={() => setIsAddUsersOpen(true)}>
                                {t("groups.addUser")}
                            </Button>
                        </div>

                        <div className="space-y-2">
                            {group.users
                                ? group.users.map((user) => (
                                    <GroupUserItem key={user.id} user={user} groupID={group.id} />
                                ))
                                : <div className="text-muted-foreground">{t("groups.emptyUsers")}</div>
                            }
                        </div>
                    </>
                )}
                <Separator className="my-1" />

                {isEditMode
                    ? (
                        <div className="flex w-full justify-between gap-2">
                            <Button
                                variant="outline"
                                onClick={() => setIsEditMode(false)}
                            >
                                {t("common.cancel")}
                            </Button>
                            <Button
                                type="submit"
                                onClick={onClickUpdate}
                            >
                                {t("common.save")}
                            </Button>
                        </div>
                    )
                    : (
                        <div className="flex w-full justify-between gap-2">
                            <Button
                                variant="outline"
                                onClick={() => setIsEditMode(true)}
                            >
                                {t("common.edit")}
                            </Button>
                            <Button
                                variant="destructive"
                                className="border border-destructive"
                                onClick={onClickDelete}
                            >
                                {t("common.delete")}
                            </Button>
                        </div>
                    )
                }

                <AddUsersToGroupModal
                    isOpen={isAddUsersOpen}
                    setIsOpen={setIsAddUsersOpen}
                    groupID={group.id}
                    existingUsers={group.users}
                />
            </DialogContent>
        </Dialog>
    )
}
