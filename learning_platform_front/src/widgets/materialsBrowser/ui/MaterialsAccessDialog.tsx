import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { Search } from "lucide-react"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import type { ShortUserInfo } from "@/entities/group"
import {
    updateUsersFoldersAccess,
    updateUsersMaterialsAccess,
} from "@/entities/material"
import { getUsersWithPagination } from "@/entities/user"
import type { GetUsersWithPaginationResponse } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { UserRoleEnum } from "@/shared/enums/user"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"
import { CandidateUserRow } from "@/widgets/addUsersToGroupModal"
import type { MaterialsTarget } from "../model/types"

const SEARCH_LIMIT = 20
const SEARCH_DEBOUNCE_MS = 300

type MaterialsAccessDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    target: MaterialsTarget | null
}

export function MaterialsAccessDialog({
    open,
    onOpenChange,
    target,
}: MaterialsAccessDialogProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()

    const [search, setSearch] = useState("")
    const [debouncedSearch, setDebouncedSearch] = useState("")
    const [users, setUsers] = useState<ShortUserInfo[]>([])
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)
    const [selectedIDs, setSelectedIDs] = useState<number[]>([])
    const [isSubmitting, setIsSubmitting] = useState(false)

    useEffect(() => {
        if (!open) return

        const timeoutId = window.setTimeout(() => {
            setDebouncedSearch(search.trim())
        }, SEARCH_DEBOUNCE_MS)

        return () => window.clearTimeout(timeoutId)
    }, [search, open])

    useEffect(() => {
        if (!open) return

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
                setError(t("materials.accessSearchError"))
            }

            setIsLoading(false)
        }

        void fetchUsers()

        return () => {
            cancelled = true
        }
    }, [debouncedSearch, dispatch, open, t])

    const toggleUser = (userID: number) => {
        setSelectedIDs((prev) => (
            prev.includes(userID)
                ? prev.filter((id) => id !== userID)
                : [...prev, userID]
        ))
    }

    const closeModal = (nextOpen: boolean) => {
        onOpenChange(nextOpen)
        if (!nextOpen) {
            setSearch("")
            setDebouncedSearch("")
            setUsers([])
            setSelectedIDs([])
            setError(null)
            setIsLoading(false)
            setIsSubmitting(false)
        }
    }

    const onSubmit = async () => {
        if (!target || selectedIDs.length === 0) return

        setIsSubmitting(true)

        const response = target.kind === "folder"
            ? await dispatch(updateUsersFoldersAccess({
                folder_ids: [target.folder.id],
                user_ids: selectedIDs,
            }))
            : await dispatch(updateUsersMaterialsAccess({
                material_ids: [target.file.id],
                user_ids: selectedIDs,
            }))

        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("materials.accessSuccess"),
                type: "success",
            }))
            closeModal(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("materials.accessError"),
                type: "error",
            }))
        }
    }

    const renderList = () => {
        if (isLoading) {
            return (
                <div className="py-6 text-center text-muted-foreground">
                    {t("materials.accessSearching")}
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

        if (users.length === 0) {
            return (
                <div className="py-6 text-center text-muted-foreground">
                    {t("materials.accessNoUsers")}
                </div>
            )
        }

        return users.map((user) => (
            <CandidateUserRow
                key={user.id}
                user={user}
                isSelected={selectedIDs.includes(user.id)}
                onToggle={() => toggleUser(user.id)}
            />
        ))
    }

    return (
        <Dialog open={open} onOpenChange={closeModal}>
            <DialogContent className="sm:max-w-lg overflow-x-hidden p-5">
                <DialogHeader>
                    <DialogTitle className="text-left text-base">
                        {t("materials.accessTitle")}
                    </DialogTitle>
                </DialogHeader>

                <div className="relative">
                    <Search className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                    <Input
                        className="pl-9"
                        placeholder={t("materials.accessSearch")}
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
                        {t("materials.accessSelected", { count: selectedIDs.length })}
                    </span>
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            onClick={() => closeModal(false)}
                        >
                            {t("common.cancel")}
                        </Button>
                        <Button
                            disabled={selectedIDs.length === 0 || isSubmitting}
                            onClick={onSubmit}
                        >
                            {t("materials.accessSubmit")}
                        </Button>
                    </div>
                </div>
            </DialogContent>
        </Dialog>
    )
}
