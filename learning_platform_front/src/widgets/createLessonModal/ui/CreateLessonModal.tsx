import { useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { createLesson } from "@/entities/lesson"
import { getUserFullData } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Separator } from "@/shared/ui/Separator"
import {
    LessonDetailsFields,
    LessonStudentsPicker,
    toDateTimeLocalValue,
    useLessonStudentsSelection,
} from "@/widgets/lessonForm"

const TUTOR_BOARDS_STUB: { id: number; title: string }[] = []

type CreateLessonModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
}

export function CreateLessonModal({
    isOpen,
    setIsOpen,
}: CreateLessonModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const userData = useAppSelector(getUserFullData)

    const [meetLink, setMeetLink] = useState("")
    const [startTime, setStartTime] = useState(toDateTimeLocalValue(new Date()))
    const [duration, setDuration] = useState(60)
    const [boardId, setBoardId] = useState<number | "">("")
    const [isSubmitting, setIsSubmitting] = useState(false)

    const students = useLessonStudentsSelection({ isOpen })

    const resetForm = () => {
        setMeetLink("")
        setStartTime(toDateTimeLocalValue(new Date()))
        setDuration(60)
        setBoardId("")
        setIsSubmitting(false)
        students.resetStudents()
    }

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (open) {
            resetForm()
        }
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (!userData?.user.userID || !startTime || duration < 15) {
            dispatch(notificationActions.addNotification({
                message: t("createLesson.error"),
                type: "error",
            }))
            return
        }

        if (students.selectedUsers.length === 0) {
            dispatch(notificationActions.addNotification({
                message: t("createLesson.studentsRequired"),
                type: "error",
            }))
            return
        }

        setIsSubmitting(true)
        const response = await dispatch(createLesson({
            tutor_id: userData.user.userID,
            board_id: boardId === "" ? null : boardId,
            meet_link: meetLink.trim() || null,
            start_time: new Date(startTime).toISOString(),
            duration,
            user_ids: students.selectedUsers.map((user) => user.id),
            media_items: [],
        }))
        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("createLesson.success"),
                type: "success",
            }))
            setIsOpen(false)
            resetForm()
        } else {
            dispatch(notificationActions.addNotification({
                message: t("createLesson.error"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-xl p-5 max-h-[90vh] overflow-x-hidden overflow-y-auto">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("createLesson.title")}
                    </DialogTitle>
                </DialogHeader>

                <Separator className="my-1" />

                <form className="w-full space-y-5" onSubmit={onSubmit}>
                    <LessonDetailsFields
                        meetLink={meetLink}
                        onMeetLinkChange={setMeetLink}
                        startTime={startTime}
                        onStartTimeChange={setStartTime}
                        duration={duration}
                        onDurationChange={setDuration}
                        boardId={boardId}
                        onBoardIdChange={setBoardId}
                        boards={TUTOR_BOARDS_STUB}
                        showBoard
                    />

                    <LessonStudentsPicker
                        groups={students.groups}
                        studentsTab={students.studentsTab}
                        onStudentsTabChange={students.setStudentsTab}
                        selectedUsers={students.selectedUsers}
                        selectedUserIDs={students.selectedUserIDs}
                        expandedGroupIDs={students.expandedGroupIDs}
                        onToggleExpanded={students.toggleExpanded}
                        onToggleGroup={students.toggleGroup}
                        onToggleUser={students.toggleUser}
                        onRemoveUsers={students.removeUsers}
                        search={students.search}
                        onSearchChange={students.setSearch}
                        searchCandidates={students.searchCandidates}
                        isSearchLoading={students.isSearchLoading}
                        searchPage={students.searchPage}
                        canLoadMore={students.canLoadMore}
                        onLoadMore={students.loadMore}
                    />

                    <Button type="submit" className="w-full" disabled={isSubmitting}>
                        {isSubmitting ? t("common.loading") : t("common.create")}
                    </Button>
                </form>
            </DialogContent>
        </Dialog>
    )
}
