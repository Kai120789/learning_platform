import { useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { Search } from "lucide-react"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { addUsersToGroup } from "@/entities/group"
import type { GroupUser } from "@/entities/group"
import { notificationActions } from "@/features/notifications"
import { mockGroupCandidates } from "@/shared/mocks"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"
import { CandidateUserRow } from "./CandidateUserRow"

type AddUsersToGroupModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    groupID: number
    existingUsers?: GroupUser[]
}

export function AddUsersToGroupModal({
    isOpen,
    setIsOpen,
    groupID,
    existingUsers,
}: AddUsersToGroupModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()

    const [search, setSearch] = useState<string>("")
    const [selectedIDs, setSelectedIDs] = useState<number[]>([])

    const candidates = useMemo(() => {
        const existingIDs = new Set(existingUsers?.map((user) => user.id))
        const query = search.trim().toLowerCase()

        return mockGroupCandidates
            .filter((user) => !existingIDs.has(user.id))
            .filter((user) => {
                if (!query) return true
                const fullName = `${user.name} ${user.surname}`.toLowerCase()
                return (
                    fullName.includes(query) ||
                    user.tgUsername?.toLowerCase().includes(query)
                )
            })
    }, [existingUsers, search])

    const toggleUser = (userID: number) => {
        setSelectedIDs((prev) =>
            prev.includes(userID)
                ? prev.filter((id) => id !== userID)
                : [...prev, userID]
        )
    }

    const closeModal = (open: boolean) => {
        setIsOpen(open)
        if (!open) {
            setSearch("")
            setSelectedIDs([])
        }
    }

    const onClickAdd = async () => {
        const response = await dispatch(addUsersToGroup({
            groupID: groupID,
            userIDs: selectedIDs,
        }))
        if (response.meta.requestStatus == "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("groups.addUsersSuccess"),
                type: "success"
            }))
            closeModal(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("groups.addUsersError"),
                type: "error"
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={closeModal}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle className="text-xl text-left">
                        {t("groups.addUsersTitle")}
                    </DialogTitle>
                </DialogHeader>

                <div className="relative">
                    <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                    <Input
                        className="pl-9"
                        placeholder={t("groups.searchUsers")}
                        value={search}
                        onChange={(e) => setSearch(e.target.value)}
                    />
                </div>

                <div className="max-h-72 space-y-2 overflow-y-auto">
                    {candidates.length === 0
                        ? (
                            <div className="py-6 text-center text-muted-foreground">
                                {t("groups.noUsersFound")}
                            </div>
                        )
                        : candidates.map((user) => (
                            <CandidateUserRow
                                key={user.id}
                                user={user}
                                isSelected={selectedIDs.includes(user.id)}
                                onToggle={() => toggleUser(user.id)}
                            />
                        ))
                    }
                </div>

                <Separator className="my-1" />

                <div className="flex items-center justify-between">
                    <span className="text-sm text-muted-foreground">
                        {t("groups.selectedCount", { count: selectedIDs.length })}
                    </span>
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            onClick={() => closeModal(false)}
                        >
                            {t("common.cancel")}
                        </Button>
                        <Button
                            disabled={selectedIDs.length === 0}
                            onClick={onClickAdd}
                        >
                            {t("common.add")}
                        </Button>
                    </div>
                </div>
            </DialogContent>
        </Dialog>
    )
}
