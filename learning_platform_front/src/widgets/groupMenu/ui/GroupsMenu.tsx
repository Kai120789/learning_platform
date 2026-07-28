import type { GroupData } from "@/entities/group"
import { GroupItem } from "./GroupItem"

type GroupsMenuProps = {
    groups: GroupData[] | null
}

export function GroupsMenu({ groups }: GroupsMenuProps) {
    if (!groups?.length) return null

    return (
        <div className="grid auto-rows-fr gap-3 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5">
            {groups.map((group) => (
                <GroupItem key={group.id} group={group} />
            ))}
        </div>
    )
}
