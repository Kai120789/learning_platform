import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { deleteGroup, updateGroup, getAllGroups } from "@/entities/group"
import type { GroupData } from "@/entities/group"
import { useCanEdit } from "@/entities/user"
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
    group: groupProp
}: GroupModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const groups = useAppSelector(getAllGroups)
    const group = groups?.find((item) => item.id === groupProp.id) ?? groupProp
    const [isEditMode, setIsEditMode] = useState<boolean>(false)
    const [isAddUsersOpen, setIsAddUsersOpen] = useState<boolean>(false)
    const [draftTitle, setDraftTitle] = useState<string>(group.title)
    const [draftDescription, setDraftDescription] = useState<string>(group.description)
    const isCanEdit = useCanEdit()

    const title = isEditMode ? draftTitle : group.title
    const description = isEditMode ? draftDescription : group.description

    const startEdit = () => {
        setDraftTitle(group.title)
        setDraftDescription(group.description)
        setIsEditMode(true)
    }

    const cancelEdit = () => {
        setIsEditMode(false)
    }

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
            title: draftTitle,
            description: draftDescription,
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
            <DialogContent className="sm:max-w-md p-5 overflow-x-hidden">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {isEditMode
                            ? (
                                <Input
                                    value={draftTitle}
                                    onChange={(e) => setDraftTitle(e.target.value)}
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
                        className="w-full break-words min-h-40 max-h-60"
                        value={description}
                        onChange={(e) => setDraftDescription(e.target.value)}
                    />
                </div>
                {!isEditMode && (
                    <>
                        <Separator className="my-1" />

                        <div className="flex flex-row items-center justify-between">
                            <div className="text-sm font-medium">
                                {t("groups.users")}
                            </div>
                            {isCanEdit && <Button size="xs" onClick={() => setIsAddUsersOpen(true)}>
                                {t("groups.addUser")}
                            </Button>}
                        </div>

                        <div className="max-h-52 overflow-x-hidden overflow-y-auto">
                            {group.users?.length
                                ? group.users.map((user) => (
                                    <GroupUserItem key={user.id} user={user} groupID={group.id} />
                                ))
                                : <div className="py-2 text-sm text-muted-foreground">{t("groups.emptyUsers")}</div>
                            }
                        </div>
                    </>
                )}
                {isCanEdit && <Separator className="my-1" />}

                {isCanEdit && (isEditMode
                    ? (
                        <div className="flex w-full justify-between gap-2">
                            <Button
                                variant="outline"
                                onClick={cancelEdit}
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
                                onClick={startEdit}
                            >
                                {t("common.edit")}
                            </Button>
                            <Button
                                variant="destructive"
                                className="border border-destructive/30"
                                onClick={onClickDelete}
                            >
                                {t("common.delete")}
                            </Button>
                        </div>
                    ))
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
