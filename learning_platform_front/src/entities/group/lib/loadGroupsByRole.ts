import { UserRoleEnum } from "@/shared/enums/user"
import { getGroupsByTutorId } from "@/entities/group"
import { getGroupsByStudentId } from "@/entities/group"

export function loadGroupsByRole(role?: UserRoleEnum) {
    if (!role) return null

    if (role === UserRoleEnum.TUTOR || role === UserRoleEnum.ADMIN) {
        return getGroupsByTutorId()
    }

    return getGroupsByStudentId()
}
