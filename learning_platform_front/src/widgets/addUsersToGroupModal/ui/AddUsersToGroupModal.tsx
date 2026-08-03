import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { Search } from "lucide-react"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { addUsersToGroup } from "@/entities/group"
import type { GroupUser, ShortUserInfo } from "@/entities/group"
import { getUsersWithPagination } from "@/entities/user"
import type { GetUsersWithPaginationResponse } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { UserRoleEnum } from "@/shared/enums/user"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"
import { CandidateUserRow } from "./CandidateUserRow"

const SEARCH_LIMIT = 20
const SEARCH_DEBOUNCE_MS = 300

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
    const [debouncedSearch, setDebouncedSearch] = useState<string>("")
    const [users, setUsers] = useState<ShortUserInfo[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [error, setError] = useState<string | null>(null)
    const [selectedIDs, setSelectedIDs] = useState<number[]>([])

    useEffect(() => {
        if (!isOpen) return

        const timeoutId = window.setTimeout(() => {
            setDebouncedSearch(search.trim())
        }, SEARCH_DEBOUNCE_MS)

        return () => window.clearTimeout(timeoutId)
    }, [search, isOpen])

    useEffect(() => {
        if (!isOpen) return

        let cancelled = false

        const fetchUsers = async () => {
            setIsLoading(true)
            setError(null)

            const response = await dispatch(getUsersWithPagination({
                search: debouncedSearch,
                page: 1,
                limit: SEARCH_LIMIT,
                role: UserRoleEnum.STUDENT,
            }))

            if (cancelled) return

            if (response.meta.requestStatus === "fulfilled") {
                const payload = response.payload as GetUsersWithPaginationResponse
                setUsers(payload.users ?? [])
            } else {
                setUsers([])
                setError(t("groups.searchError"))
            }

            setIsLoading(false)
        }

        void fetchUsers()

        return () => {
            cancelled = true
        }
    }, [debouncedSearch, dispatch, isOpen, t])

    const candidates = useMemo(() => {
        const existingIDs = new Set(existingUsers?.map((user) => user.id))
        return users.filter((user) => !existingIDs.has(user.id))
    }, [existingUsers, users])

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
            setDebouncedSearch("")
            setUsers([])
            setSelectedIDs([])
            setError(null)
            setIsLoading(false)
        }
    }

    const onClickAdd = async () => {
        const selectedUsers = candidates.filter((user) => selectedIDs.includes(user.id))
        const response = await dispatch(addUsersToGroup({
            groupID: groupID,
            users: selectedUsers,
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

    const renderList = () => {
        if (isLoading) {
            return (
                <div className="py-6 text-center text-muted-foreground">
                    {t("groups.searching")}
                </div>
            )
        }

        if (error) {
            return (
                <div className="py-6 text-center text-destructive">
                    {error}
                </div>
            )
        }

        if (candidates.length === 0) {
            return (
                <div className="py-6 text-center text-muted-foreground">
                    {t("groups.noUsersFound")}
                </div>
            )
        }

        return candidates.map((user) => (
            <CandidateUserRow
                key={user.id}
                user={user}
                isSelected={selectedIDs.includes(user.id)}
                onToggle={() => toggleUser(user.id)}
            />
        ))
    }

    return (
        <Dialog open={isOpen} onOpenChange={closeModal}>
            <DialogContent className="sm:max-w-lg p-5 overflow-x-hidden">
                <DialogHeader>
                    <DialogTitle className="text-base text-left">
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

                <div className="max-h-64 overflow-x-hidden overflow-y-auto">
                    {renderList()}
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
