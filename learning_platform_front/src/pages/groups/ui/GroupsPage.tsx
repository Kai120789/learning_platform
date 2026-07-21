import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getGroupsByTutorId } from "@/entities/group/api/getGroupsByTutorId"
import { getAllGroups } from "@/entities/group/selectors/selectots"
import { Button } from "@/shared/ui/Button"
import { Label } from "@/shared/ui/Label"
import { CreateGroupModal } from "@/widgets/createGroupModal/ui/CreateGroupModal"
import { GroupsMenu } from "@/widgets/groupMenu"
import { useEffect, useState } from "react"
import { FaPlus } from "react-icons/fa";


export default function GroupsPage() {
    const dispatch = useAppDispatch()
    const groups = useAppSelector(getAllGroups)

    const [isOpen, setIsOpen] = useState<boolean>(false)

    const fetchUserGroups = async () => {
        await dispatch(getGroupsByTutorId())
    }

    useEffect(() => {
        fetchUserGroups()
    }, [])

    return (
        <div className="py-10 lg:py-15 px-10 lg:px-40 space-y-8">
            <div className="space-y-1">
                <div className="flex justify-between items-center">
                    <Label className="text-2xl lg:text-4xl">
                        Группы
                    </Label>
                    <Button onClick={() => setIsOpen(true)} size="lg" className="rounded-full">
                        <FaPlus className="size-3" />
                        Создать группу
                    </Button>
                </div>
                <Label className="text-md lg:text-xl font-normal text-primary/50">
                    Здесь отображаются ваши группы и их участники
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

