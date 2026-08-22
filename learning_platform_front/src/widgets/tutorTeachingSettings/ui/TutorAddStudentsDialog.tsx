import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { Search } from "lucide-react"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { addTutorStudents, type TutorUserData } from "@/entities/tutor"
import { getUsersWithPagination } from "@/entities/user"
import type { GetUsersWithPaginationResponse } from "@/entities/user"
import type { ShortUserInfo } from "@/entities/group"
import { notificationActions } from "@/features/notifications"
import { UserRoleEnum } from "@/shared/enums/user"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Input } from "@/shared/ui/Input"
import { Separator } from "@/shared/ui/Separator"
import { cn } from "@/shared/lib/utils"
import { SelectionMark } from "@/shared/ui/SelectionMark"

const SEARCH_LIMIT = 20
const SEARCH_DEBOUNCE_MS = 300

type TutorAddStudentsDialogProps = {
    isOpen: boolean
    setIsOpen: (open: boolean) => void
    existingIds: number[]
}

export function TutorAddStudentsDialog({
    isOpen,
    setIsOpen,
    existingIds,
}: TutorAddStudentsDialogProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const [search, setSearch] = useState("")
    const [debouncedSearch, setDebouncedSearch] = useState("")
    const [users, setUsers] = useState<ShortUserInfo[]>([])
    const [isLoading, setIsLoading] = useState(false)
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
                setError(t("tutors.studentsSearchError"))
            }
            setIsLoading(false)
        }

        void fetchUsers()
        return () => {
            cancelled = true
        }
    }, [debouncedSearch, dispatch, isOpen, t])

    const candidates = useMemo(() => {
        const existing = new Set(existingIds)
        return users.filter((user) => !existing.has(user.id))
    }, [existingIds, users])

    const closeModal = (open: boolean) => {
        setIsOpen(open)
        if (!open) {
            setSearch("")
            setDebouncedSearch("")
            setUsers([])
            setSelectedIDs([])
            setError(null)
        }
    }

    const onAdd = async () => {
        const selected = candidates
            .filter((user) => selectedIDs.includes(user.id))
            .map((user): TutorUserData => ({
                id: user.id,
                name: user.name,
                surname: user.surname,
                patronymic: user.patronymic,
                tgUsername: user.tg_username,
            }))
        const response = await dispatch(addTutorStudents(selected))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("tutors.studentsAddSuccess"),
                type: "success",
            }))
            closeModal(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("tutors.studentsAddError"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={closeModal}>
            <DialogContent className="sm:max-w-lg p-5">
                <DialogHeader>
                    <DialogTitle className="text-left text-base">
                        {t("tutors.addStudents")}
                    </DialogTitle>
                </DialogHeader>
                <div className="relative">
                    <Search className="absolute left-3 top-1/2 size-4 -translate-y-1/2 text-muted-foreground" />
                    <Input
                        className="pl-9"
                        placeholder={t("tutors.searchStudents")}
                        value={search}
                        onChange={(event) => setSearch(event.target.value)}
                    />
                </div>
                <div className="max-h-64 overflow-y-auto">
                    {isLoading ? (
                        <div className="py-6 text-center text-sm text-muted-foreground">
                            {t("common.loading")}
                        </div>
                    ) : error ? (
                        <div className="py-6 text-center text-sm text-destructive">{error}</div>
                    ) : candidates.length === 0 ? (
                        <div className="py-6 text-center text-sm text-muted-foreground">
                            {t("tutors.noStudentsFound")}
                        </div>
                    ) : (
                        candidates.map((user) => {
                            const selected = selectedIDs.includes(user.id)
                            return (
                                <button
                                    key={user.id}
                                    type="button"
                                    onClick={() => setSelectedIDs((current) => (
                                        current.includes(user.id)
                                            ? current.filter((id) => id !== user.id)
                                            : [...current, user.id]
                                    ))}
                                    className={cn(
                                        "flex w-full cursor-pointer items-center justify-between gap-3 py-2.5 text-left hover:bg-muted/40",
                                        selected && "bg-primary/5"
                                    )}
                                >
                                    <span className="min-w-0">
                                        <span className="block truncate text-sm font-medium">
                                            {user.name} {user.surname}
                                        </span>
                                        {user.tg_username && (
                                            <span className="mt-0.5 block truncate text-xs text-muted-foreground">
                                                @{user.tg_username}
                                            </span>
                                        )}
                                    </span>
                                    <SelectionMark checked={selected} />
                                </button>
                            )
                        })
                    )}
                </div>
                <Separator className="my-1" />
                <div className="flex items-center justify-between">
                    <span className="text-sm text-muted-foreground">
                        {t("tutors.selectedCount", { count: selectedIDs.length })}
                    </span>
                    <div className="flex gap-2">
                        <Button variant="outline" size="sm" onClick={() => closeModal(false)}>
                            {t("common.cancel")}
                        </Button>
                        <Button size="sm" disabled={selectedIDs.length === 0} onClick={onAdd}>
                            {t("common.add")}
                        </Button>
                    </div>
                </div>
            </DialogContent>
        </Dialog>
    )
}
