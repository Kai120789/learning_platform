import type { GroupData } from "@/entities/group/types/types"
import { GroupItem } from "./GroupItem"

type GroupsMenuProps = {
    groups: GroupData[] | null
}

export function GroupsMenu({ groups }: GroupsMenuProps) {
    if (!groups?.length) return null

    return (
        <div>
            <div className="grid gap-5 md:grid-cols-2 xl:grid-cols-4">
                {groups.map((group) => (
                    <GroupItem key={group.id} group={group} />
                ))}
            </div>
        </div>
    )
}