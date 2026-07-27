import { useSelector } from "react-redux"
import { getUserRole } from "@/entities/user"
import { UserRoleEnum } from "@/shared/enums/user"

export function useCanEdit() {
    const userRole = useSelector(getUserRole)
    return userRole === UserRoleEnum.ADMIN || userRole === UserRoleEnum.TUTOR
}
