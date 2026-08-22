import type { GroupData } from "@/entities/group"
import { GroupItem } from "./GroupItem"

type GroupsMenuProps = {
    groups: GroupData[] | null
}

export function GroupsMenu({ groups }: GroupsMenuProps) {
    if (!groups?.length) return null

    return (
        <div className="grid auto-rows-fr grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
            {groups.map((group) => (
                <GroupItem key={group.id} group={group} />
            ))}
        </div>
    )
}
