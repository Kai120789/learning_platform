import { useEffect, useMemo, useState } from "react"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    getAllGroups,
    getGroupsByTutorId,
    type GroupUser,
    type ShortUserInfo,
} from "@/entities/group"
import { getUsersWithPagination } from "@/entities/user"
import type { GetUsersWithPaginationResponse } from "@/entities/user"
import { UserRoleEnum } from "@/shared/enums/user"
import { SEARCH_DEBOUNCE_MS, SEARCH_LIMIT, type StudentsTab } from "../lib/constants"
import { placeholderUser, toShortUser } from "../lib/userHelpers"

type UseLessonStudentsSelectionOptions = {
    isOpen: boolean
    seedUserIDs?: number[] | null
    seedKey?: number | null
}

export function useLessonStudentsSelection({
    isOpen,
    seedUserIDs = null,
    seedKey = null,
}: UseLessonStudentsSelectionOptions) {
    const dispatch = useAppDispatch()
    const groups = useAppSelector(getAllGroups) ?? []

    const [studentsTab, setStudentsTab] = useState<StudentsTab>("groups")
    const [selectedUsers, setSelectedUsers] = useState<ShortUserInfo[]>([])
    const [expandedGroupIDs, setExpandedGroupIDs] = useState<number[]>([])
    const [search, setSearch] = useState("")
    const [debouncedSearch, setDebouncedSearch] = useState("")
    const [searchUsers, setSearchUsers] = useState<ShortUserInfo[]>([])
    const [searchTotal, setSearchTotal] = useState(0)
    const [searchPage, setSearchPage] = useState(1)
    const [isSearchLoading, setIsSearchLoading] = useState(false)
    const [activeSeedKey, setActiveSeedKey] = useState<number | null>(null)

    useEffect(() => {
        if (!isOpen) return
        dispatch(getGroupsByTutorId())
    }, [dispatch, isOpen])

    useEffect(() => {
        if (!isOpen) return
        if (seedKey == null || seedUserIDs == null) return
        if (activeSeedKey === seedKey) return

        setActiveSeedKey(seedKey)
        setStudentsTab("groups")
        setExpandedGroupIDs([])
        setSearch("")
        setDebouncedSearch("")
        setSearchUsers([])
        setSearchTotal(0)
        setSearchPage(1)
        setSelectedUsers(seedUserIDs.map((id) => placeholderUser(id)))
    }, [activeSeedKey, isOpen, seedKey, seedUserIDs])

    useEffect(() => {
        if (!isOpen || groups.length === 0 || seedKey == null) return

        const usersFromGroups = new Map<number, ShortUserInfo>()
        groups.forEach((group) => {
            group.users?.forEach((user) => {
                usersFromGroups.set(user.id, toShortUser(user))
            })
        })

        setSelectedUsers((prev) => {
            let changed = false
            const next = prev.map((user) => {
                const fromGroup = usersFromGroups.get(user.id)
                if (!fromGroup) return user
                if (user.name === fromGroup.name && user.surname === fromGroup.surname) return user
                changed = true
                return fromGroup
            })
            return changed ? next : prev
        })
    }, [groups, isOpen, seedKey])

    useEffect(() => {
        if (!isOpen) return
        const timeoutId = window.setTimeout(() => {
            setDebouncedSearch(search.trim())
            setSearchPage(1)
        }, SEARCH_DEBOUNCE_MS)
        return () => window.clearTimeout(timeoutId)
    }, [search, isOpen])

    useEffect(() => {
        if (!isOpen) return

        let cancelled = false
        const fetchUsers = async () => {
            setIsSearchLoading(true)
            const response = await dispatch(getUsersWithPagination({
                search: debouncedSearch,
                page: searchPage,
                limit: SEARCH_LIMIT,
                role: UserRoleEnum.STUDENT,
            }))
            if (cancelled) return

            if (response.meta.requestStatus === "fulfilled") {
                const payload = response.payload as GetUsersWithPaginationResponse
                const nextUsers = payload.users ?? []
                setSearchTotal(payload.count ?? nextUsers.length)
                setSearchUsers((prev) => (
                    searchPage === 1 ? nextUsers : [...prev, ...nextUsers]
                ))

                if (seedKey != null && nextUsers.length > 0) {
                    setSelectedUsers((prev) => {
                        const map = new Map(prev.map((user) => [user.id, user]))
                        let changed = false
                        nextUsers.forEach((user) => {
                            const existing = map.get(user.id)
                            if (existing && existing.name.startsWith("#")) {
                                map.set(user.id, user)
                                changed = true
                            }
                        })
                        return changed ? Array.from(map.values()) : prev
                    })
                }
            } else {
                if (searchPage === 1) setSearchUsers([])
                setSearchTotal(0)
            }
            setIsSearchLoading(false)
        }

        void fetchUsers()
        return () => {
            cancelled = true
        }
    }, [debouncedSearch, dispatch, isOpen, searchPage, seedKey])

    const selectedUserIDs = useMemo(
        () => new Set(selectedUsers.map((user) => user.id)),
        [selectedUsers],
    )

    const searchCandidates = useMemo(() => (
        searchUsers.filter((user) => !selectedUserIDs.has(user.id))
    ), [searchUsers, selectedUserIDs])

    const canLoadMore = searchUsers.length < searchTotal

    const upsertUsers = (users: Array<GroupUser | ShortUserInfo>) => {
        setSelectedUsers((prev) => {
            const map = new Map(prev.map((user) => [user.id, user]))
            users.forEach((user) => map.set(user.id, toShortUser(user)))
            return Array.from(map.values())
        })
    }

    const removeUsers = (ids: number[]) => {
        const removeSet = new Set(ids)
        setSelectedUsers((prev) => prev.filter((user) => !removeSet.has(user.id)))
    }

    const toggleUser = (user: GroupUser | ShortUserInfo) => {
        if (selectedUserIDs.has(user.id)) {
            removeUsers([user.id])
            return
        }
        upsertUsers([user])
    }

    const toggleGroup = (groupID: number) => {
        const group = groups.find((item) => item.id === groupID)
        const members = group?.users ?? []
        if (members.length === 0) return

        const allSelected = members.every((user) => selectedUserIDs.has(user.id))
        if (allSelected) {
            removeUsers(members.map((user) => user.id))
        } else {
            upsertUsers(members)
            setExpandedGroupIDs((prev) => (
                prev.includes(groupID) ? prev : [...prev, groupID]
            ))
        }
    }

    const toggleExpanded = (groupID: number) => {
        setExpandedGroupIDs((prev) => (
            prev.includes(groupID)
                ? prev.filter((id) => id !== groupID)
                : [...prev, groupID]
        ))
    }

    const resetStudents = () => {
        setStudentsTab("groups")
        setSelectedUsers([])
        setExpandedGroupIDs([])
        setSearch("")
        setDebouncedSearch("")
        setSearchUsers([])
        setSearchTotal(0)
        setSearchPage(1)
        setActiveSeedKey(null)
    }

    return {
        groups,
        studentsTab,
        setStudentsTab,
        selectedUsers,
        selectedUserIDs,
        expandedGroupIDs,
        search,
        setSearch,
        searchCandidates,
        isSearchLoading,
        searchPage,
        canLoadMore,
        loadMore: () => setSearchPage((prev) => prev + 1),
        toggleUser,
        toggleGroup,
        toggleExpanded,
        removeUsers,
        resetStudents,
    }
}
