import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { useNavigate } from "react-router-dom"
import { Send, Users } from "lucide-react"
import { getRouteGroups } from "@/app/router/routePaths"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getAllGroups, loadGroupsByRole } from "@/entities/group"
import type { GroupData } from "@/entities/group"
import { getUserRole } from "@/entities/user"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { GroupModal } from "@/widgets/groupModal"

export function GroupsRowList() {
    const { t } = useTranslation()
    const navigate = useNavigate()
    const dispatch = useAppDispatch()
    const groups = useAppSelector(getAllGroups)
    const role = useAppSelector(getUserRole)

    const [selectedGroup, setSelectedGroup] = useState<GroupData | null>(null)
    const [isOpen, setIsOpen] = useState(false)

    useEffect(() => {
        const action = loadGroupsByRole(role)
        if (action) {
            dispatch(action)
        }
    }, [dispatch, role])

    const openGroup = (group: GroupData) => {
        setSelectedGroup(group)
        setIsOpen(true)
    }

    return (
        <>
            <Card className="h-full">
                <CardHeader>
                    <CardTitle
                        className="cursor-pointer transition-colors hover:text-primary"
                        onClick={() => navigate(getRouteGroups())}
                    >
                        {t("main.groups")}
                    </CardTitle>
                </CardHeader>
                <CardContent className="space-y-2">
                    {!groups?.length ? (
                        <div className="py-4 text-sm text-muted-foreground">
                            {t("common.empty")}
                        </div>
                    ) : (
                        groups.map((group) => (
                            <button
                                key={group.id}
                                type="button"
                                onClick={() => openGroup(group)}
                                className="flex w-full cursor-pointer items-center justify-between gap-3 rounded-xl border px-3 py-3 text-left transition-colors hover:bg-muted/30"
                            >
                                <div className="min-w-0 space-y-1">
                                    <div className="truncate font-medium">{group.title}</div>
                                    <div className="truncate text-sm text-muted-foreground">
                                        {group.subject.title} · {group.subject.type}
                                    </div>
                                </div>
                                <div className="flex shrink-0 items-center gap-3">
                                    <Badge variant="secondary">
                                        <Users className="mr-1 size-3.5" />
                                        {group.users?.length ?? 0}
                                    </Badge>
                                    {group.tgGroupLink && (
                                        <a
                                            href={group.tgGroupLink}
                                            target="_blank"
                                            rel="noreferrer"
                                            className="text-primary hover:opacity-80"
                                            onClick={(e) => e.stopPropagation()}
                                        >
                                            <Send className="size-[22px]" />
                                        </a>
                                    )}
                                </div>
                            </button>
                        ))
                    )}
                </CardContent>
            </Card>

            {selectedGroup && (
                <GroupModal
                    group={selectedGroup}
                    isOpen={isOpen}
                    setIsOpen={setIsOpen}
                />
            )}
        </>
    )
}
