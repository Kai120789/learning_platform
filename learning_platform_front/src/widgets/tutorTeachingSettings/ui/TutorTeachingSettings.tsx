import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import { Clock, Pencil, Plus, Trash2 } from "lucide-react"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    addTutorOffer,
    deleteOneTutorOffer,
    deleteOneTutorStudent,
    formatPrice,
    formatTutorName,
    formatTutorSubject,
    getOneTutor,
    getTutorInitials,
    getTutorIsStudentsLoading,
    getTutorIsTeachingLoading,
    getTutorStudents,
    getTutorTeaching,
    updateTutorOffer,
    updateTutorSubjects,
    type TutorOfferData,
} from "@/entities/tutor"
import { getSubjects } from "@/entities/subject"
import { getUserFullData } from "@/entities/user"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Badge } from "@/shared/ui/Badge"
import { Avatar, AvatarFallback } from "@/shared/ui/Avatar"
import { TutorOfferDialog } from "./TutorOfferDialog"
import { TutorSubjectsDialog } from "./TutorSubjectsDialog"
import { TutorAddStudentsDialog } from "./TutorAddStudentsDialog"

export function TutorTeachingSettings() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const userId = useAppSelector(getUserFullData)?.user.userID
    const teaching = useAppSelector(getTutorTeaching)
    const isLoading = useAppSelector(getTutorIsTeachingLoading)
    const isStudentsLoading = useAppSelector(getTutorIsStudentsLoading)
    const allSubjects = useAppSelector(getSubjects) ?? []

    const [selectedSubjectIds, setSelectedSubjectIds] = useState<number[]>([])
    const [subjectsDialogOpen, setSubjectsDialogOpen] = useState(false)
    const [offerDialogOpen, setOfferDialogOpen] = useState(false)
    const [studentsDialogOpen, setStudentsDialogOpen] = useState(false)
    const [editingOffer, setEditingOffer] = useState<TutorOfferData | null>(null)

    useEffect(() => {
        if (userId) {
            dispatch(getOneTutor({ tutorId: userId, purpose: "teaching" }))
            dispatch(getTutorStudents({ page: 1, limit: 50 }))
        }
    }, [dispatch, userId])

    useEffect(() => {
        if (teaching) {
            setSelectedSubjectIds(teaching.subjectIds)
        }
    }, [teaching])

    const selectedSubjects = useMemo(
        () => allSubjects.filter((subject) => selectedSubjectIds.includes(subject.id)),
        [allSubjects, selectedSubjectIds],
    )

    const onSaveSubjects = async (subjectIds: number[]) => {
        const response = await dispatch(updateTutorSubjects(subjectIds))
        if (response.meta.requestStatus === "fulfilled") {
            setSelectedSubjectIds(subjectIds)
            dispatch(notificationActions.addNotification({
                message: t("tutors.subjectsSaveSuccess"),
                type: "success",
            }))
            return true
        }
        dispatch(notificationActions.addNotification({
            message: t("tutors.subjectsSaveError"),
            type: "error",
        }))
        return false
    }

    const onSubmitOffer = async (values: {
        title: string
        description?: string | null
        subject_id: number
        price: number
        duration_minutes?: number | null
    }) => {
        const response = editingOffer
            ? await dispatch(updateTutorOffer({ id: editingOffer.id, ...values }))
            : await dispatch(addTutorOffer(values))

        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: editingOffer ? t("tutors.offerUpdateSuccess") : t("tutors.offerCreateSuccess"),
                type: "success",
            }))
            return true
        }
        dispatch(notificationActions.addNotification({
            message: editingOffer ? t("tutors.offerUpdateError") : t("tutors.offerCreateError"),
            type: "error",
        }))
        return false
    }

    const onDeleteOffer = async (offerId: number) => {
        const response = await dispatch(deleteOneTutorOffer(offerId))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("tutors.offerDeleteSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("tutors.offerDeleteError"),
                type: "error",
            }))
        }
    }

    const onDeleteStudent = async (studentId: number) => {
        const response = await dispatch(deleteOneTutorStudent(studentId))
        if (response.meta.requestStatus === "fulfilled") {
            dispatch(notificationActions.addNotification({
                message: t("tutors.studentDeleteSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("tutors.studentDeleteError"),
                type: "error",
            }))
        }
    }

    return (
        <div className="space-y-5">
            <section className="overflow-hidden rounded-lg border border-border bg-background">
                <div className="flex items-start justify-between gap-3 px-4 py-3.5">
                    <div className="min-w-0 space-y-1">
                        <h3 className="text-sm font-medium">{t("tutors.mySubjects")}</h3>
                        <p className="text-xs text-muted-foreground">{t("tutors.mySubjectsHint")}</p>
                    </div>
                    <Button
                        type="button"
                        size="sm"
                        variant="outline"
                        onClick={() => setSubjectsDialogOpen(true)}
                    >
                        {t("tutors.editSubjects")}
                    </Button>
                </div>
                <div className="px-4 pb-4">
                    {selectedSubjects.length === 0 ? (
                        <p className="text-sm text-muted-foreground">{t("tutors.emptySubjects")}</p>
                    ) : (
                        <div className="flex flex-wrap gap-1.5">
                            {selectedSubjects.map((subject) => (
                                <Badge key={subject.id} variant="secondary" className="text-xs">
                                    {formatTutorSubject(subject)}
                                </Badge>
                            ))}
                        </div>
                    )}
                </div>
            </section>

            <section className="overflow-hidden rounded-lg border border-border bg-background">
                <div className="flex items-center justify-between gap-3 px-4 py-3.5">
                    <div className="space-y-1">
                        <h3 className="text-sm font-medium">{t("tutors.myOffers")}</h3>
                        <p className="text-xs text-muted-foreground">{t("tutors.myOffersHint")}</p>
                    </div>
                    <Button
                        type="button"
                        size="sm"
                        disabled={selectedSubjects.length === 0}
                        onClick={() => {
                            setEditingOffer(null)
                            setOfferDialogOpen(true)
                        }}
                    >
                        <Plus className="size-3.5" />
                        {t("tutors.addOffer")}
                    </Button>
                </div>
                {selectedSubjects.length === 0 ? (
                    <div className="px-4 pb-4 text-sm text-muted-foreground">
                        {t("tutors.offersNeedSubjects")}
                    </div>
                ) : (teaching?.offers.length ?? 0) === 0 ? (
                    <div className="px-4 pb-4 text-sm text-muted-foreground">
                        {t("tutors.emptyOffers")}
                    </div>
                ) : (
                    <div className="divide-y divide-border px-4">
                        {teaching?.offers.map((offer) => {
                            const subjectLabel = formatTutorSubject(offer.subject)
                            return (
                                <div
                                    key={offer.id}
                                    className="flex flex-col gap-2 py-3 sm:flex-row sm:items-center"
                                >
                                    <div className="min-w-0 flex-1">
                                        <div className="flex min-w-0 items-baseline gap-2">
                                            <span className="truncate text-sm font-medium">
                                                {offer.title}
                                            </span>
                                            {subjectLabel && (
                                                <span className="shrink-0 text-sm text-muted-foreground">
                                                    {subjectLabel}
                                                </span>
                                            )}
                                        </div>
                                        {offer.description && (
                                            <p className="mt-0.5 line-clamp-2 text-xs text-muted-foreground">
                                                {offer.description}
                                            </p>
                                        )}
                                    </div>
                                    <div className="flex shrink-0 items-center gap-3">
                                        <div className="min-w-20 sm:text-right">
                                            <div className="text-base font-semibold leading-tight">
                                                {t("tutors.price", { price: formatPrice(offer.price) })}
                                            </div>
                                            {offer.durationMinutes != null && (
                                                <div className="mt-0.5 inline-flex items-center gap-1 text-sm text-muted-foreground">
                                                    <Clock className="size-3.5" />
                                                    {t("common.minutes", { count: offer.durationMinutes })}
                                                </div>
                                            )}
                                        </div>
                                        <div className="flex gap-1">
                                            <Button
                                                type="button"
                                                size="icon-sm"
                                                variant="ghost"
                                                onClick={() => {
                                                    setEditingOffer(offer)
                                                    setOfferDialogOpen(true)
                                                }}
                                            >
                                                <Pencil className="size-3.5" />
                                            </Button>
                                            <Button
                                                type="button"
                                                size="icon-sm"
                                                variant="ghost"
                                                onClick={() => onDeleteOffer(offer.id)}
                                            >
                                                <Trash2 className="size-3.5" />
                                            </Button>
                                        </div>
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                )}
            </section>

            <section className="overflow-hidden rounded-lg border border-border bg-background">
                <div className="flex items-center justify-between gap-3 px-4 py-3.5">
                    <div className="space-y-1">
                        <h3 className="text-sm font-medium">{t("tutors.myStudents")}</h3>
                        <p className="text-xs text-muted-foreground">{t("tutors.myStudentsHint")}</p>
                    </div>
                    <Button
                        type="button"
                        size="sm"
                        onClick={() => setStudentsDialogOpen(true)}
                    >
                        <Plus className="size-3.5" />
                        {t("tutors.addStudent")}
                    </Button>
                </div>
                {isStudentsLoading && (teaching?.students.length ?? 0) === 0 ? (
                    <div className="px-4 pb-4 text-sm text-muted-foreground">
                        {t("common.loading")}
                    </div>
                ) : (teaching?.students.length ?? 0) === 0 ? (
                    <div className="px-4 pb-4 text-sm text-muted-foreground">
                        {t("tutors.emptyStudents")}
                    </div>
                ) : (
                    <div className="divide-y divide-border px-4 pb-2">
                        {teaching?.students.map(({ student, lastInteractedAt }) => (
                            <div
                                key={student.id}
                                className="flex items-center justify-between gap-3 py-3"
                            >
                                <div className="flex min-w-0 items-center gap-3">
                                    <Avatar className="size-8 shrink-0">
                                        <AvatarFallback>
                                            {getTutorInitials(student)}
                                        </AvatarFallback>
                                    </Avatar>
                                    <div className="min-w-0">
                                        <div className="truncate text-sm font-medium">
                                            {formatTutorName(student)}
                                        </div>
                                        <div className="truncate text-xs text-muted-foreground">
                                            {student.tgUsername
                                                ? `@${student.tgUsername}`
                                                : lastInteractedAt
                                                    ? t("tutors.lastInteracted", {
                                                        date: new Date(lastInteractedAt).toLocaleDateString(),
                                                    })
                                                    : t("tutors.noUsername")}
                                        </div>
                                    </div>
                                </div>
                                <Button
                                    type="button"
                                    size="icon-sm"
                                    variant="ghost"
                                    onClick={() => onDeleteStudent(student.id)}
                                >
                                    <Trash2 className="size-3.5" />
                                </Button>
                            </div>
                        ))}
                    </div>
                )}
            </section>

            <TutorSubjectsDialog
                isOpen={subjectsDialogOpen}
                setIsOpen={setSubjectsDialogOpen}
                subjects={allSubjects}
                selectedIds={selectedSubjectIds}
                isSaving={isLoading}
                onSave={onSaveSubjects}
            />
            <TutorOfferDialog
                isOpen={offerDialogOpen}
                setIsOpen={(open) => {
                    setOfferDialogOpen(open)
                    if (!open) setEditingOffer(null)
                }}
                subjects={selectedSubjects}
                offer={editingOffer}
                onSubmit={onSubmitOffer}
            />
            <TutorAddStudentsDialog
                isOpen={studentsDialogOpen}
                setIsOpen={setStudentsDialogOpen}
                existingIds={(teaching?.students ?? []).map((item) => item.student.id)}
            />
        </div>
    )
}
