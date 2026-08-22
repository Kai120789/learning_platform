import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-i18next"
import type { SubjectData } from "@/entities/subject"
import { Button } from "@/shared/ui/Button"
import { Checkbox } from "@/shared/ui/Checkbox"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import { Separator } from "@/shared/ui/Separator"
import { cn } from "@/shared/lib/utils"

type TutorSubjectsDialogProps = {
    isOpen: boolean
    setIsOpen: (open: boolean) => void
    subjects: SubjectData[]
    selectedIds: number[]
    isSaving: boolean
    onSave: (subjectIds: number[]) => Promise<boolean>
}

export function TutorSubjectsDialog({
    isOpen,
    setIsOpen,
    subjects,
    selectedIds,
    isSaving,
    onSave,
}: TutorSubjectsDialogProps) {
    const { t } = useTranslation()
    const [draftIds, setDraftIds] = useState<number[]>(selectedIds)

    useEffect(() => {
        if (isOpen) {
            setDraftIds(selectedIds)
        }
    }, [isOpen, selectedIds])

    const grouped = useMemo(() => {
        const groups = new Map<string, SubjectData[]>()
        for (const subject of subjects) {
            const list = groups.get(subject.type) ?? []
            list.push(subject)
            groups.set(subject.type, list)
        }
        return [...groups.entries()]
    }, [subjects])

    const toggleSubject = (subjectId: number) => {
        setDraftIds((current) => (
            current.includes(subjectId)
                ? current.filter((id) => id !== subjectId)
                : [...current, subjectId]
        ))
    }

    const onSubmit = async () => {
        const ok = await onSave(draftIds)
        if (ok) {
            setIsOpen(false)
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogContent className="sm:max-w-lg p-5">
                <DialogHeader>
                    <DialogTitle className="pr-10 text-left text-base">
                        {t("tutors.selectSubjects")}
                    </DialogTitle>
                </DialogHeader>
                <Separator className="my-1" />
                <div className="max-h-80 space-y-3 overflow-y-auto pr-1">
                    {grouped.map(([type, items]) => (
                        <div key={type} className="space-y-1.5">
                            <div className="text-[11px] font-medium uppercase tracking-wide text-muted-foreground">
                                {type}
                            </div>
                            <div className="grid grid-cols-2 gap-1.5">
                                {items.map((subject) => {
                                    const checked = draftIds.includes(subject.id)
                                    return (
                                        <button
                                            key={subject.id}
                                            type="button"
                                            onClick={() => toggleSubject(subject.id)}
                                            className={cn(
                                                "flex min-w-0 cursor-pointer items-center gap-1.5 rounded-lg border px-2 py-1.5 text-left transition-colors",
                                                checked
                                                    ? "border-foreground/20 bg-muted/70"
                                                    : "border-border hover:bg-muted/40"
                                            )}
                                        >
                                            <Checkbox
                                                checked={checked}
                                                tabIndex={-1}
                                                className="pointer-events-none size-3.5"
                                            />
                                            <span className="truncate text-xs font-medium">
                                                {subject.title}
                                            </span>
                                        </button>
                                    )
                                })}
                            </div>
                        </div>
                    ))}
                </div>
                <div className="flex items-center justify-between gap-3 pt-1">
                    <span className="text-xs text-muted-foreground">
                        {t("tutors.selectedCount", { count: draftIds.length })}
                    </span>
                    <div className="flex gap-2">
                        <Button
                            type="button"
                            variant="outline"
                            size="sm"
                            onClick={() => setIsOpen(false)}
                        >
                            {t("common.cancel")}
                        </Button>
                        <Button
                            type="button"
                            size="sm"
                            disabled={isSaving}
                            onClick={onSubmit}
                        >
                            {t("common.save")}
                        </Button>
                    </div>
                </div>
            </DialogContent>
        </Dialog>
    )
}
