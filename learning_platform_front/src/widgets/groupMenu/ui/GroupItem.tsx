import { useState } from "react"
import type { GroupData } from "@/entities/group/types/types"

import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
} from "@/shared/ui/Card"

import { Badge } from "@/shared/ui/Badge"

import { RiTelegramFill } from "react-icons/ri"
import { Users } from "lucide-react"
import { Tooltip, TooltipContent, TooltipTrigger } from "@/shared/ui/Tooltip"
import { GroupModal } from "@/widgets/groupModal/ui/GroupModal"

type GroupItemProps = {
    group: GroupData
}

export function GroupItem({ group }: GroupItemProps) {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <div>
            <Card onClick={() => setIsOpen((prev) => !prev)} className="transition-all hover:shadow-lg cursor-pointer">
                <CardHeader className="space-y-3">
                    <div className="flex items-start justify-between">
                        <div className="flex flex-col space-y-2">
                            <CardTitle >
                                <Tooltip>
                                    <TooltipTrigger className="text-lg break-all line-clamp-1">{group.title}</TooltipTrigger>
                                    <TooltipContent className="text-xl">{group.title}</TooltipContent>
                                </Tooltip>
                            </CardTitle>
                        </div>

                        <Badge variant="secondary">
                            <Users className="mr-1 h-3.5 w-3.5" />
                            {group.users
                                ? group.users?.length
                                : 0
                            }
                        </Badge>
                    </div>

                    <div className="flex gap-2">
                        <Badge variant="outline" className="bg-muted">
                            {group.subject.title}
                        </Badge>
                        <Badge variant="default">
                            {group.subject.type}
                        </Badge>
                    </div>
                    <CardDescription className="mt-1 break-all line-clamp-1">
                        {group.description}
                    </CardDescription>
                    {group.tgGroupLink && (
                        <a
                            href={group.tgGroupLink}
                            target="_blank"
                            rel="noreferrer"
                            className="flex w-fit items-center gap-2 text-sm text-primary hover:underline"
                            onClick={(e) => e.stopPropagation()}
                        >
                            <RiTelegramFill size={25} />
                            Telegram
                        </a>
                    )}
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