import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { getUserData } from "@/entities/user"
import { useEffect } from "react"
import { Outlet } from "react-router-dom"

export function BoardLayout() {
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(getUserData())
    }, [dispatch])

    return (
        <div className="relative h-svh overflow-hidden bg-secondary">
            <Outlet />
        </div>
    )
}
