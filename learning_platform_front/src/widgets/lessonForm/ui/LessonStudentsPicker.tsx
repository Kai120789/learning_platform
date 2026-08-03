import { useTranslation } from "react-i18next"
import type { GroupData, GroupUser, ShortUserInfo } from "@/entities/group"
import { cn } from "@/shared/lib/utils"
import type { StudentsTab } from "../lib/constants"
import { LessonGroupsTab } from "./LessonGroupsTab"
import { LessonSearchTab } from "./LessonSearchTab"
import { SelectedStudentsChips } from "./SelectedStudentsChips"

type LessonStudentsPickerProps = {
    groups: GroupData[]
    studentsTab: StudentsTab
    onStudentsTabChange: (tab: StudentsTab) => void
    selectedUsers: ShortUserInfo[]
    selectedUserIDs: Set<number>
    expandedGroupIDs: number[]
    onToggleExpanded: (groupID: number) => void
    onToggleGroup: (groupID: number) => void
    onToggleUser: (user: GroupUser | ShortUserInfo) => void
    onRemoveUsers: (ids: number[]) => void
    search: string
    onSearchChange: (value: string) => void
    searchCandidates: ShortUserInfo[]
    isSearchLoading: boolean
    searchPage: number
    canLoadMore: boolean
    onLoadMore: () => void
}

export function LessonStudentsPicker({
    groups,
    studentsTab,
    onStudentsTabChange,
    selectedUsers,
    selectedUserIDs,
    expandedGroupIDs,
    onToggleExpanded,
    onToggleGroup,
    onToggleUser,
    onRemoveUsers,
    search,
    onSearchChange,
    searchCandidates,
    isSearchLoading,
    searchPage,
    canLoadMore,
    onLoadMore,
}: LessonStudentsPickerProps) {
    const { t } = useTranslation()

    return (
        <div className="space-y-3">
            <div
                className="flex gap-5 border-b border-border"
                role="tablist"
                aria-label={t("createLesson.studentsSource")}
            >
                {([
                    ["groups", "createLesson.groups"],
                    ["search", "createLesson.searchStudents"],
                ] as const).map(([tab, labelKey]) => {
                    const active = studentsTab === tab
                    return (
                        <button
                            key={tab}
                            type="button"
                            role="tab"
                            aria-selected={active}
                            className={cn(
                                "-mb-px cursor-pointer border-b-2 pb-2 text-sm transition-colors",
                                active
                                    ? "border-foreground font-medium text-foreground"
                                    : "border-transparent text-muted-foreground hover:text-foreground",
                            )}
                            onClick={() => onStudentsTabChange(tab)}
                        >
                            {t(labelKey)}
                        </button>
                    )
                })}
            </div>

            <div className="h-56 overflow-x-hidden overflow-y-auto">
                {studentsTab === "groups" ? (
                    <LessonGroupsTab
                        groups={groups}
                        selectedUserIDs={selectedUserIDs}
                        expandedGroupIDs={expandedGroupIDs}
                        onToggleExpanded={onToggleExpanded}
                        onToggleGroup={onToggleGroup}
                        onToggleUser={onToggleUser}
                    />
                ) : (
                    <LessonSearchTab
                        search={search}
                        onSearchChange={onSearchChange}
                        candidates={searchCandidates}
                        selectedUserIDs={selectedUserIDs}
                        isLoading={isSearchLoading}
                        searchPage={searchPage}
                        canLoadMore={canLoadMore}
                        onLoadMore={onLoadMore}
                        onToggleUser={onToggleUser}
                    />
                )}
            </div>

            <SelectedStudentsChips
                users={selectedUsers}
                onRemove={(userID) => onRemoveUsers([userID])}
            />
        </div>
    )
}
