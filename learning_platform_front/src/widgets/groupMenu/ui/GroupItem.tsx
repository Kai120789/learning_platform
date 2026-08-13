import { useState } from "react"
import type { GroupData } from "@/entities/group"

import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
} from "@/shared/ui/Card"

import { Badge } from "@/shared/ui/Badge"

import { Send, Users } from "lucide-react"
import { Tooltip, TooltipContent, TooltipTrigger } from "@/shared/ui/Tooltip"
import { GroupModal } from "@/widgets/groupModal"

type GroupItemProps = {
    group: GroupData
}

export function GroupItem({ group }: GroupItemProps) {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <div className="h-full min-h-0">
            <Card
                size="sm"
                onClick={() => setIsOpen((prev) => !prev)}
                className="h-full cursor-pointer transition-all hover:shadow-md"
            >
                <CardHeader className="flex h-full flex-col space-y-2">
                    <div className="flex items-start justify-between gap-2">
                        <CardTitle className="min-w-0">
                            <Tooltip>
                                <TooltipTrigger className="text-sm font-medium break-all line-clamp-1 text-left">
                                    {group.title}
                                </TooltipTrigger>
                                <TooltipContent className="text-xs">{group.title}</TooltipContent>
                            </Tooltip>
                        </CardTitle>

                        <Badge variant="secondary" className="shrink-0 text-[10px]">
                            <Users className="mr-1 size-3" />
                            {group.users?.length ?? 0}
                        </Badge>
                    </div>

                    <div className="flex min-h-5 flex-wrap gap-1.5">
                        <Badge variant="outline" className="bg-muted text-[10px]">
                            {group.subject.title}
                        </Badge>
                        <Badge variant="default" className="text-[10px]">
                            {group.subject.type}
                        </Badge>
                    </div>

                    <CardDescription className="min-h-8 flex-1 text-xs break-all line-clamp-2">
                        {group.description || "\u00A0"}
                    </CardDescription>

                    <div className="mt-auto min-h-5 flex items-center">
                        {group.tgGroupLink ? (
                            <a
                                href={group.tgGroupLink}
                                target="_blank"
                                rel="noreferrer"
                                className="flex w-fit items-center gap-1.5 text-xs text-primary hover:underline"
                                onClick={(e) => e.stopPropagation()}
                            >
                                <Send className="size-[18px]" />
                                Telegram
                            </a>
                        ) : (
                            <span className="text-xs text-transparent select-none">–</span>
                        )}
                    </div>
                </CardHeader>
            </Card>
            <GroupModal
                group={group}
                isOpen={isOpen}
                setIsOpen={setIsOpen}
            />
        </div>
    )
}
