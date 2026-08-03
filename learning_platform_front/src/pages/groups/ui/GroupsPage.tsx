import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getAllGroups, loadGroupsByRole } from "@/entities/group"
import { getUserRole, useCanEdit } from "@/entities/user"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { CreateGroupModal } from "@/widgets/createGroupModal"
import { GroupsMenu } from "@/widgets/groupMenu"
import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { FaPlus } from "react-icons/fa"

export default function GroupsPage() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const groups = useAppSelector(getAllGroups)
    const role = useAppSelector(getUserRole)
    const isCanEdit = useCanEdit()

    const [isOpen, setIsOpen] = useState<boolean>(false)

    useEffect(() => {
        const action = loadGroupsByRole(role)
        if (action) {
            dispatch(action)
        }
    }, [dispatch, role])

    return (
        <div className="py-8 lg:py-10 px-6 lg:px-20 space-y-6">
            <div className="space-y-1">
                <div className="flex justify-between items-center">
                    <Label className="text-xl lg:text-2xl">
                        {t("groups.title")}
                    </Label>
                    {isCanEdit && (
                        <Button
                            size="sm"
                            onClick={() => setIsOpen(true)}
                            className="rounded-full md:h-8 md:gap-1.5 md:px-2.5 md:text-sm"
                        >
                            <FaPlus className="size-3 md:size-3.5" />
                            {t("groups.create")}
                        </Button>
                    )}
                </div>
                <Label className="text-sm lg:text-base font-normal text-primary/50">
                    {t("groups.subtitle")}
                </Label>
            </div>
            <GroupsMenu
                groups={groups}
            />
            <CreateGroupModal
                isOpen={isOpen}
                setIsOpen={setIsOpen}
            />
        </div>
    )
}
