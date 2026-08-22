import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { Pencil, Star, Trash2 } from "lucide-react"
import { useAppDispatch } from "@/app/providers/storeProvider/hooks/hooks"
import {
    addTutorReview,
    deleteTutorReview,
    formatTutorSubject,
    updateTutorReview,
    type TutorMyReviewData,
    type TutorSubjectData,
} from "@/entities/tutor"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { NativeSelect } from "@/shared/ui/NativeSelect"
import { Textarea } from "@/shared/ui/Textarea"
import { cn } from "@/shared/lib/utils"
import { TutorReviewCard } from "./TutorReviewCard"

type ReviewAuthor = {
    name: string
    surname: string
}

type TutorReviewFormProps = {
    tutorId: number
    subjects: TutorSubjectData[]
    myReview: TutorMyReviewData | null
    author: ReviewAuthor
}

export function TutorReviewForm({
    tutorId,
    subjects,
    myReview,
    author,
}: TutorReviewFormProps) {
    const { t, i18n } = useTranslation()
    const dispatch = useAppDispatch()
    const [editing, setEditing] = useState(false)
    const [text, setText] = useState(myReview?.text ?? "")
    const [rating, setRating] = useState(myReview?.rating ?? 5)
    const [subjectId, setSubjectId] = useState(
        myReview?.subjectId ?? subjects[0]?.id ?? 0,
    )
    const [submitting, setSubmitting] = useState(false)

    useEffect(() => {
        setText(myReview?.text ?? "")
        setRating(myReview?.rating ?? 5)
        setSubjectId(myReview?.subjectId ?? subjects[0]?.id ?? 0)
        if (!myReview) {
            setEditing(false)
        }
    }, [myReview, subjects])

    const onSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        if (!subjectId) return
        setSubmitting(true)

        const response = myReview
            ? await dispatch(updateTutorReview({
                id: myReview.id,
                text,
                subject_id: subjectId,
                rating,
            }))
            : await dispatch(addTutorReview({
                text,
                tutor_id: tutorId,
                subject_id: subjectId,
                rating,
            }))

        setSubmitting(false)
        if (response.meta.requestStatus === "fulfilled") {
            setEditing(false)
            dispatch(notificationActions.addNotification({
                message: myReview ? t("tutors.reviewUpdateSuccess") : t("tutors.reviewCreateSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: myReview ? t("tutors.reviewUpdateError") : t("tutors.reviewCreateError"),
                type: "error",
            }))
        }
    }

    const onDelete = async () => {
        if (!myReview) return
        setSubmitting(true)
        const response = await dispatch(deleteTutorReview(myReview.id))
        setSubmitting(false)
        if (response.meta.requestStatus === "fulfilled") {
            setEditing(false)
            setText("")
            setRating(5)
            dispatch(notificationActions.addNotification({
                message: t("tutors.reviewDeleteSuccess"),
                type: "success",
            }))
        } else {
            dispatch(notificationActions.addNotification({
                message: t("tutors.reviewDeleteError"),
                type: "error",
            }))
        }
    }

    const onCancel = () => {
        setEditing(false)
        setText(myReview?.text ?? "")
        setRating(myReview?.rating ?? 5)
        setSubjectId(myReview?.subjectId ?? subjects[0]?.id ?? 0)
    }

    if (subjects.length === 0) {
        return null
    }

    if (myReview && !editing) {
        const subject = subjects.find((item) => item.id === myReview.subjectId)
        return (
            <TutorReviewCard
                author={author}
                subjectLabel={subject ? formatTutorSubject(subject) : ""}
                dateLabel={new Date(myReview.createdAt).toLocaleDateString(i18n.language)}
                rating={myReview.rating}
                text={myReview.text}
                muted
                badge={t("tutors.yourReview")}
                actions={(
                    <div className="-mr-1 flex">
                        <Button
                            type="button"
                            size="icon-sm"
                            variant="ghost"
                            disabled={submitting}
                            onClick={() => setEditing(true)}
                        >
                            <Pencil className="size-3.5" />
                        </Button>
                        <Button
                            type="button"
                            size="icon-sm"
                            variant="ghost"
                            disabled={submitting}
                            onClick={onDelete}
                        >
                            <Trash2 className="size-3.5" />
                        </Button>
                    </div>
                )}
            />
        )
    }

    return (
        <form
            onSubmit={onSubmit}
            className="rounded-xl bg-muted p-3.5 ring-1 ring-foreground/12"
        >
            <div className="mb-3 text-sm font-medium">
                {myReview ? t("tutors.editReview") : t("tutors.writeReview")}
            </div>
            <FieldGroup className="gap-3">
                <Field>
                    <FieldLabel>{t("common.subject")}</FieldLabel>
                    <NativeSelect
                        value={subjectId}
                        onChange={(event) => setSubjectId(Number(event.target.value))}
                    >
                        {subjects.map((subject) => (
                            <option key={subject.id} value={subject.id}>
                                {formatTutorSubject(subject)}
                            </option>
                        ))}
                    </NativeSelect>
                </Field>
                <Field>
                    <FieldLabel>{t("tutors.rating")}</FieldLabel>
                    <div className="flex gap-1">
                        {[1, 2, 3, 4, 5].map((value) => (
                            <button
                                key={value}
                                type="button"
                                onClick={() => setRating(value)}
                                className="cursor-pointer rounded-md p-0.5 text-foreground"
                            >
                                <Star
                                    className={cn(
                                        "size-5",
                                        value <= rating && "fill-current",
                                    )}
                                />
                            </button>
                        ))}
                    </div>
                </Field>
                <Field>
                    <FieldLabel>{t("tutors.reviewText")}</FieldLabel>
                    <Textarea
                        required
                        value={text}
                        onChange={(event) => setText(event.target.value)}
                    />
                </Field>
                <div className="flex justify-end gap-2">
                    {myReview && (
                        <Button
                            type="button"
                            variant="outline"
                            size="sm"
                            disabled={submitting}
                            onClick={onCancel}
                        >
                            {t("common.cancel")}
                        </Button>
                    )}
                    <Button type="submit" size="sm" disabled={submitting || !subjectId}>
                        {myReview ? t("common.save") : t("tutors.sendReview")}
                    </Button>
                </div>
            </FieldGroup>
        </form>
    )
}
