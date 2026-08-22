import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { formatTutorSubject, type TutorOfferData } from "@/entities/tutor"
import type { SubjectData } from "@/entities/subject"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { NativeSelect } from "@/shared/ui/NativeSelect"
import { Separator } from "@/shared/ui/Separator"
import { Textarea } from "@/shared/ui/Textarea"

type TutorOfferDialogProps = {
    isOpen: boolean
    setIsOpen: (open: boolean) => void
    subjects: SubjectData[]
    offer?: TutorOfferData | null
    onSubmit: (values: {
        title: string
        description?: string | null
        subject_id: number
        price: number
        duration_minutes?: number | null
    }) => Promise<boolean>
}

export function TutorOfferDialog({
    isOpen,
    setIsOpen,
    subjects,
    offer,
    onSubmit,
}: TutorOfferDialogProps) {
    const { t } = useTranslation()
    const [title, setTitle] = useState(offer?.title ?? "")
    const [description, setDescription] = useState(offer?.description ?? "")
    const [subjectId, setSubjectId] = useState(offer?.subject.id ?? subjects[0]?.id ?? 0)
    const [price, setPrice] = useState(offer ? String(offer.price) : "")
    const [duration, setDuration] = useState(
        offer?.durationMinutes != null ? String(offer.durationMinutes) : "",
    )
    const [submitting, setSubmitting] = useState(false)

    useEffect(() => {
        if (!isOpen) return
        setTitle(offer?.title ?? "")
        setDescription(offer?.description ?? "")
        setSubjectId(offer?.subject.id ?? subjects[0]?.id ?? 0)
        setPrice(offer ? String(offer.price) : "")
        setDuration(offer?.durationMinutes != null ? String(offer.durationMinutes) : "")
    }, [isOpen, offer, subjects])

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        if (!subjectId) return
        setSubmitting(true)
        const ok = await onSubmit({
            title: title.trim(),
            description: description.trim() || null,
            subject_id: subjectId,
            price: Number(price),
            duration_minutes: duration ? Number(duration) : null,
        })
        setSubmitting(false)
        if (ok) {
            setIsOpen(false)
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent className="sm:max-w-md p-5">
                <DialogHeader>
                    <DialogTitle className="pr-10 text-left text-base">
                        {offer ? t("tutors.editOffer") : t("tutors.createOffer")}
                    </DialogTitle>
                </DialogHeader>
                <Separator className="my-1" />
                <form className="w-full" onSubmit={handleSubmit}>
                    <FieldGroup className="w-full gap-3">
                        <Field className="w-full">
                            <FieldLabel>{t("tutors.offerTitle")}</FieldLabel>
                            <Input
                                required
                                value={title}
                                onChange={(event) => setTitle(event.target.value)}
                            />
                        </Field>
                        <Field className="w-full">
                            <FieldLabel>{t("common.subject")}</FieldLabel>
                            <NativeSelect
                                required
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
                        <Field className="w-full">
                            <FieldLabel>{t("tutors.offerDescription")}</FieldLabel>
                            <Textarea
                                value={description}
                                onChange={(event) => setDescription(event.target.value)}
                            />
                        </Field>
                        <div className="grid grid-cols-2 gap-3">
                            <Field>
                                <FieldLabel>{t("tutors.offerPrice")}</FieldLabel>
                                <Input
                                    required
                                    type="number"
                                    min={0}
                                    value={price}
                                    onChange={(event) => setPrice(event.target.value)}
                                />
                            </Field>
                            <Field>
                                <FieldLabel>{t("tutors.offerDuration")}</FieldLabel>
                                <Input
                                    type="number"
                                    min={1}
                                    value={duration}
                                    onChange={(event) => setDuration(event.target.value)}
                                />
                            </Field>
                        </div>
                        <div className="flex justify-end">
                            <Button type="submit" size="sm" disabled={submitting || !subjectId}>
                                {offer ? t("common.save") : t("common.create")}
                            </Button>
                        </div>
                    </FieldGroup>
                </form>
            </DialogContent>
        </Dialog>
    )
}
