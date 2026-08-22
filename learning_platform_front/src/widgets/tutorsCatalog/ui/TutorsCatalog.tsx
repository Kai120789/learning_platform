import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { LayoutGrid, List } from "lucide-react"
import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import {
    getTutorIsLoading,
    getTutorList,
    getTutorListCount,
    getTutors,
} from "@/entities/tutor"
import { getSubjects } from "@/entities/subject"
import { Button } from "@/shared/ui/Button"
import { NativeSelect } from "@/shared/ui/NativeSelect"
import { cn } from "@/shared/lib/utils"
import { TutorGridCard } from "./TutorGridCard"
import { TutorListRow } from "./TutorListRow"

const PAGE_LIMIT = 12
const VIEW_MODE_KEY = "tutors-view-mode"

type TutorsViewMode = "grid" | "list"

function readStoredViewMode(): TutorsViewMode {
    if (typeof window === "undefined") return "grid"
    return window.localStorage.getItem(VIEW_MODE_KEY) === "list" ? "list" : "grid"
}

export function TutorsCatalog() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const tutors = useAppSelector(getTutorList)
    const tutorsCount = useAppSelector(getTutorListCount)
    const isLoading = useAppSelector(getTutorIsLoading)
    const subjects = useAppSelector(getSubjects)

    const [subjectId, setSubjectId] = useState<number | undefined>()
    const [page, setPage] = useState(1)
    const [viewMode, setViewMode] = useState<TutorsViewMode>(() => readStoredViewMode())

    useEffect(() => {
        dispatch(getTutors({
            page: 1,
            limit: PAGE_LIMIT,
            subjectId,
        }))
        setPage(1)
    }, [dispatch, subjectId])

    const onViewModeChange = (mode: TutorsViewMode) => {
        setViewMode(mode)
        window.localStorage.setItem(VIEW_MODE_KEY, mode)
    }

    const onLoadMore = () => {
        const nextPage = page + 1
        dispatch(getTutors({
            page: nextPage,
            limit: PAGE_LIMIT,
            subjectId,
            append: true,
        }))
        setPage(nextPage)
    }

    return (
        <div className="space-y-4">
            <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                <NativeSelect
                    containerClassName="sm:max-w-64"
                    value={subjectId ?? ""}
                    onChange={(event) => {
                        const value = event.target.value
                        setSubjectId(value ? Number(value) : undefined)
                    }}
                >
                    <option value="">{t("tutors.allSubjects")}</option>
                    {(subjects ?? []).map((subject) => (
                        <option key={subject.id} value={subject.id}>
                            {subject.title} · {subject.type}
                        </option>
                    ))}
                </NativeSelect>
                <div className="flex rounded-lg border border-border bg-secondary/60 p-0.5">
                    <button
                        type="button"
                        onClick={() => onViewModeChange("list")}
                        className={cn(
                            "flex cursor-pointer items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                            viewMode === "list"
                                ? "bg-primary text-primary-foreground shadow-sm"
                                : "text-secondary-foreground hover:bg-muted"
                        )}
                    >
                        <List className="size-3.5" />
                        {t("tutors.viewList")}
                    </button>
                    <button
                        type="button"
                        onClick={() => onViewModeChange("grid")}
                        className={cn(
                            "flex cursor-pointer items-center gap-1.5 rounded-md px-2.5 py-1.5 text-xs font-medium transition-colors",
                            viewMode === "grid"
                                ? "bg-primary text-primary-foreground shadow-sm"
                                : "text-secondary-foreground hover:bg-muted"
                        )}
                    >
                        <LayoutGrid className="size-3.5" />
                        {t("tutors.viewGrid")}
                    </button>
                </div>
            </div>

            {isLoading && tutors.length === 0 ? (
                <div className="py-12 text-center text-sm text-muted-foreground">
                    {t("common.loading")}
                </div>
            ) : tutors.length === 0 ? (
                <div className="rounded-2xl border border-dashed border-border px-4 py-12 text-center text-sm text-muted-foreground">
                    {t("tutors.empty")}
                </div>
            ) : viewMode === "grid" ? (
                <div className="grid auto-rows-fr gap-3 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3">
                    {tutors.map((tutor) => (
                        <TutorGridCard key={tutor.id} tutor={tutor} />
                    ))}
                </div>
            ) : (
                <div className="space-y-2">
                    {tutors.map((tutor) => (
                        <TutorListRow key={tutor.id} tutor={tutor} />
                    ))}
                </div>
            )}

            {tutors.length < tutorsCount && (
                <div className="flex justify-center">
                    <Button
                        type="button"
                        variant="outline"
                        size="sm"
                        disabled={isLoading}
                        onClick={onLoadMore}
                    >
                        {t("tutors.loadMore")}
                    </Button>
                </div>
            )}
        </div>
    )
}
