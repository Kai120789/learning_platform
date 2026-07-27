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
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <div className="flex justify-between items-center">
                    <Label className="text-2xl lg:text-4xl">
                        {t("groups.title")}
                    </Label>
                    {isCanEdit && <Button onClick={() => setIsOpen(true)} size="lg" className="rounded-full">
                        <FaPlus className="size-3" />
                        {t("groups.create")}
                    </Button>}
                </div>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
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
