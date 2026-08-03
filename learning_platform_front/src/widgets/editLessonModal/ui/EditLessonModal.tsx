import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import { updateLesson, type LessonData } from "@/entities/lesson"
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

type EditLessonModalProps = {
    isOpen: boolean
    setIsOpen: (isOpen: boolean) => void
    lesson: LessonData | null
}

export function EditLessonModal({
    isOpen,
    setIsOpen,
    lesson,
}: EditLessonModalProps) {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()

    const [meetLink, setMeetLink] = useState("")
    const [startTime, setStartTime] = useState("")
    const [duration, setDuration] = useState(60)
    const [initialUserIDs, setInitialUserIDs] = useState<number[]>([])
    const [isSubmitting, setIsSubmitting] = useState(false)
    const [activeLessonId, setActiveLessonId] = useState<number | null>(null)

    const students = useLessonStudentsSelection({
        isOpen,
        seedKey: lesson?.id ?? null,
        seedUserIDs: lesson?.userIds ?? null,
    })

    useEffect(() => {
        if (!isOpen || !lesson) return
        if (activeLessonId === lesson.id) return

        setActiveLessonId(lesson.id)
        setMeetLink(lesson.meetLink ?? "")
        setStartTime(toDateTimeLocalValue(lesson.startTime))
        setDuration(lesson.duration)
        setInitialUserIDs(lesson.userIds)
        setIsSubmitting(false)
    }, [activeLessonId, isOpen, lesson])

    const onOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (!open) {
            setActiveLessonId(null)
            setInitialUserIDs([])
            students.resetStudents()
        }
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!lesson || !startTime || duration < 15) {
            dispatch(notificationActions.addNotification({
                message: t("editLesson.error"),
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

        const currentIDs = students.selectedUsers.map((user) => user.id)
        const initialSet = new Set(initialUserIDs)
        const currentSet = new Set(currentIDs)
        const user_ids = currentIDs.filter((id) => !initialSet.has(id))
        const deleted_user_ids = initialUserIDs.filter((id) => !currentSet.has(id))

        setIsSubmitting(true)
        const response = await dispatch(updateLesson({
            id: lesson.id,
            board_id: lesson.boardId ?? null,
            meet_link: meetLink.trim() || null,
            start_time: new Date(startTime).toISOString(),
            duration,
            user_ids,
            deleted_user_ids,
            media_items: [],
            deleted_media_ids: [],
        }))
        setIsSubmitting(false)

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("editLesson.success"),
                type: "success",
            }))
            onOpenChange(false)
        } else {
            dispatch(notificationActions.addNotification({
                message: t("editLesson.error"),
                type: "error",
            }))
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-xl p-5 max-h-[90vh] overflow-x-hidden overflow-y-auto">
                <DialogHeader>
                    <DialogTitle className="text-base text-left line-clamp-2 pr-10">
                        {t("editLesson.title")}
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

                    <Button type="submit" className="w-full" disabled={isSubmitting || !lesson}>
                        {isSubmitting ? t("common.loading") : t("common.save")}
                    </Button>
                </form>
            </DialogContent>
        </Dialog>
    )
}
