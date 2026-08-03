import { useTranslation } from "react-i18next"
import { Search } from "lucide-react"
import type { GroupUser, ShortUserInfo } from "@/entities/group"
import { Button } from "@/shared/ui/Button"
import { Input } from "@/shared/ui/Input"
import { CandidateUserRow } from "@/widgets/addUsersToGroupModal"

type LessonSearchTabProps = {
    search: string
    onSearchChange: (value: string) => void
    candidates: ShortUserInfo[]
    selectedUserIDs: Set<number>
    isLoading: boolean
    searchPage: number
    canLoadMore: boolean
    onLoadMore: () => void
    onToggleUser: (user: GroupUser | ShortUserInfo) => void
}

export function LessonSearchTab({
    search,
    onSearchChange,
    candidates,
    selectedUserIDs,
    isLoading,
    searchPage,
    canLoadMore,
    onLoadMore,
    onToggleUser,
}: LessonSearchTabProps) {
    const { t } = useTranslation()

    return (
        <div className="flex h-full min-w-0 flex-col gap-2">
            <div className="relative shrink-0">
                <Search className="pointer-events-none absolute top-1/2 left-2.5 size-3.5 -translate-y-1/2 text-muted-foreground" />
                <Input
                    className="w-full pl-8"
                    value={search}
                    onChange={(e) => onSearchChange(e.target.value)}
                    placeholder={t("createLesson.searchPlaceholder")}
                />
            </div>
            <div className="min-h-0 min-w-0 flex-1 overflow-x-hidden overflow-y-auto">
                {isLoading && searchPage === 1 ? (
                    <p className="py-6 text-center text-sm text-muted-foreground">
                        {t("common.loading")}
                    </p>
                ) : candidates.length === 0 ? (
                    <p className="py-6 text-center text-sm text-muted-foreground">
                        {t("createLesson.noStudents")}
                    </p>
                ) : (
                    candidates.map((user) => (
                        <CandidateUserRow
                            key={user.id}
                            user={user}
                            isSelected={selectedUserIDs.has(user.id)}
                            onToggle={() => onToggleUser(user)}
                        />
                    ))
                )}
            </div>
            {canLoadMore && (
                <Button
                    type="button"
                    size="xs"
                    variant="ghost"
                    className="shrink-0 self-start"
                    disabled={isLoading}
                    onClick={onLoadMore}
                >
                    {isLoading ? t("common.loading") : t("createLesson.loadMore")}
                </Button>
            )}
        </div>
    )
}
