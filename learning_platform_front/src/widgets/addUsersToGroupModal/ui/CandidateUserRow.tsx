import type { CandidateUserMock } from "@/shared/mocks"
import { cn } from "@/shared/lib/utils"
import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { Checkbox } from "@/shared/ui/Checkbox"

type CandidateUserRowProps = {
    user: CandidateUserMock
    isSelected: boolean
    onToggle: () => void
}

export function CandidateUserRow({
    user,
    isSelected,
    onToggle,
}: CandidateUserRowProps) {
    return (
        <div
            className={cn(
                "flex cursor-pointer items-center justify-between rounded-lg border p-3 transition-colors hover:bg-muted/50",
                isSelected && "border-primary bg-primary/5"
            )}
            onClick={onToggle}
        >
            <div className="flex items-center gap-3">
                <Avatar>
                    <AvatarFallback>
                        {user.name[0] + user.surname[0]}
                    </AvatarFallback>
                </Avatar>

                <div>
                    <p className="font-medium leading-none">
                        {`${user.name} ${user.surname}`}
                    </p>

                    <p className="text-sm text-muted-foreground">
                        {user.tgUsername}
                    </p>
                </div>
            </div>

            <Checkbox
                checked={isSelected}
                onCheckedChange={onToggle}
                onClick={(e) => e.stopPropagation()}
            />
        </div>
    )
}
