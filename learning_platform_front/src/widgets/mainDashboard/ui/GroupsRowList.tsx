import { useState } from "react"
import { useTranslation } from "react-i18next"
import { RiTelegramFill } from "react-icons/ri"
import { Users } from "lucide-react"
import { mockHomeGroups } from "@/shared/mocks"
import { SubjectTypeEnum } from "@/entities/subject"
import type { GroupData } from "@/entities/group"
import { Badge } from "@/shared/ui/Badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/shared/ui/Card"
import { GroupModal } from "@/widgets/groupModal"

export function GroupsRowList() {
    const { t } = useTranslation()
    const [selectedGroup, setSelectedGroup] = useState<GroupData | null>(null)
    const [isOpen, setIsOpen] = useState(false)

    const openGroup = (group: typeof mockHomeGroups[number]) => {
        setSelectedGroup({
            id: group.id,
            title: group.title,
            description: group.description,
            subject: {
                id: group.subject.id,
                code: group.subject.code,
                title: group.subject.title,
                type: group.subject.type as SubjectTypeEnum,
            },
            users: [...group.users],
            tutorId: group.tutorId,
            tgGroupLink: "tgGroupLink" in group ? group.tgGroupLink : undefined,
        })
        setIsOpen(true)
    }

    return (
        <>
            <Card className="h-full">
                <CardHeader>
                    <CardTitle>{t("main.groups")}</CardTitle>
                </CardHeader>
                <CardContent className="space-y-2">
                    {mockHomeGroups.map((group) => (
                        <button
                            key={group.id}
                            type="button"
                            onClick={() => openGroup(group)}
                            className="flex w-full items-center justify-between gap-3 rounded-xl border px-3 py-3 text-left transition-colors hover:bg-muted/50"
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
                                    {group.users.length}
                                </Badge>
                                {"tgGroupLink" in group && group.tgGroupLink && (
                                    <a
                                        href={group.tgGroupLink}
                                        target="_blank"
                                        rel="noreferrer"
                                        className="text-primary hover:opacity-80"
                                        onClick={(e) => e.stopPropagation()}
                                    >
                                        <RiTelegramFill size={22} />
                                    </a>
                                )}
                            </div>
                        </button>
                    ))}
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
