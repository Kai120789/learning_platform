import { useTranslation } from "react-i18next"
import { ChevronDown } from "lucide-react"
import type { GroupData, GroupUser, ShortUserInfo } from "@/entities/group"
import { cn } from "@/shared/lib/utils"
import { SelectionMark } from "@/shared/ui/SelectionMark"
import { CandidateUserRow } from "@/widgets/addUsersToGroupModal"
import { toShortUser } from "../lib/userHelpers"

type LessonGroupsTabProps = {
    groups: GroupData[]
    selectedUserIDs: Set<number>
    expandedGroupIDs: number[]
    onToggleExpanded: (groupID: number) => void
    onToggleGroup: (groupID: number) => void
    onToggleUser: (user: GroupUser | ShortUserInfo) => void
}

export function LessonGroupsTab({
    groups,
    selectedUserIDs,
    expandedGroupIDs,
    onToggleExpanded,
    onToggleGroup,
    onToggleUser,
}: LessonGroupsTabProps) {
    const { t } = useTranslation()

    if (groups.length === 0) {
        return (
            <p className="py-6 text-center text-sm text-muted-foreground">
                {t("createLesson.noGroups")}
            </p>
        )
    }

    return (
        <div className="min-w-0 space-y-0.5">
            {groups.map((group) => {
                const members = group.users ?? []
                const selectedCount = members.filter((user) => selectedUserIDs.has(user.id)).length
                const allSelected = members.length > 0 && selectedCount === members.length
                const expanded = expandedGroupIDs.includes(group.id)

                return (
                    <div key={group.id} className="min-w-0">
                        <div className="flex min-w-0 items-center gap-2 overflow-hidden py-1.5">
                            <button
                                type="button"
                                className="flex min-w-0 flex-1 cursor-pointer items-center gap-2 overflow-hidden text-left text-sm"
                                onClick={() => onToggleExpanded(group.id)}
                            >
                                <ChevronDown
                                    className={cn(
                                        "size-3.5 shrink-0 text-muted-foreground transition-transform",
                                        !expanded && "-rotate-90",
                                    )}
                                />
                                <span className="truncate font-medium">{group.title}</span>
                                <span className="shrink-0 text-xs text-muted-foreground">
                                    {t("createLesson.groupMembersSelected", {
                                        selected: selectedCount,
                                        count: members.length,
                                    })}
                                </span>
                            </button>
                            <button
                                type="button"
                                className="shrink-0 cursor-pointer rounded-[4px] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50"
                                disabled={members.length === 0}
                                onClick={() => onToggleGroup(group.id)}
                                aria-label={t("createLesson.selectEntireGroup")}
                                aria-pressed={allSelected}
                            >
                                <SelectionMark checked={allSelected} />
                            </button>
                        </div>

                        {expanded && (
                            <div className="ml-7 min-w-0 overflow-x-hidden pb-1">
                                {members.length === 0 ? (
                                    <p className="py-2 text-xs text-muted-foreground">
                                        {t("createLesson.emptyGroup")}
                                    </p>
                                ) : (
                                    members.map((user) => (
                                        <CandidateUserRow
                                            key={user.id}
                                            user={toShortUser(user)}
                                            isSelected={selectedUserIDs.has(user.id)}
                                            onToggle={() => onToggleUser(user)}
                                        />
                                    ))
                                )}
                            </div>
                        )}
                    </div>
                )
            })}
        </div>
    )
}
