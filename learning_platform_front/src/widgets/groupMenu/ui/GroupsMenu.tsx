import type { GroupData } from "@/entities/group"
import { GroupItem } from "./GroupItem"

type GroupsMenuProps = {
    groups: GroupData[] | null
}

export function GroupsMenu({ groups }: GroupsMenuProps) {
    if (!groups?.length) return null

    return (
        <div>
            <div className="grid gap-3 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4">
                {groups.map((group) => (
                    <GroupItem key={group.id} group={group} />
                ))}
            </div>
        </div>
    )
}