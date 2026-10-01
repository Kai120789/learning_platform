import { useAppDispatch, useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getRouteBoard } from "@/app/router/routePaths"
import { createBoard, deleteBoard, getBoards, getBoardsLoading, getTutorBoards, updateBoard, type Board } from "@/entities/board"
import { notificationActions } from "@/features/notifications"
import { Button } from "@/shared/ui/Button"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/shared/ui/Dialog"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/shared/ui/DropdownMenu"
import { Field, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { Label } from "@/shared/ui/Label"
import { Separator } from "@/shared/ui/Separator"
import { MoreHorizontal, Pen, Plus, Trash2 } from "lucide-react"
import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { useNavigate } from "react-router-dom"

type TitleDialog =
    | { mode: "create" }
    | { mode: "rename"; board: Board }
    | { mode: "delete"; board: Board }

export default function BoardsPage() {
    const { t } = useTranslation()
    const dispatch = useAppDispatch()
    const navigate = useNavigate()
    const boards = useAppSelector(getBoards)
    const isLoading = useAppSelector(getBoardsLoading)
    const [dialog, setDialog] = useState<TitleDialog | null>(null)
    const [title, setTitle] = useState("")
    const [isSubmitting, setIsSubmitting] = useState(false)

    useEffect(() => {
        dispatch(getTutorBoards())
    }, [dispatch])

    const openCreate = () => {
        setTitle("")
        setDialog({ mode: "create" })
    }

    const openRename = (board: Board) => {
        setTitle(board.title)
        setDialog({ mode: "rename", board })
    }

    const onSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        if (!dialog || dialog.mode === "delete") return
        const nextTitle = title.trim()
        if (!nextTitle) return

        setIsSubmitting(true)
        const response = dialog.mode === "create"
            ? await dispatch(createBoard({ title: nextTitle }))
            : await dispatch(updateBoard({ boardId: dialog.board.id, title: nextTitle }))
        setIsSubmitting(false)

        const ok = response.meta.requestStatus === "fulfilled"
        dispatch(notificationActions.addNotification({
            message: t(ok
                ? dialog.mode === "create" ? "boards.createSuccess" : "boards.updateSuccess"
                : dialog.mode === "create" ? "boards.createError" : "boards.updateError"),
            type: ok ? "success" : "error",
        }))
        if (ok) setDialog(null)
    }

    const onDelete = async () => {
        if (dialog?.mode !== "delete") return
        setIsSubmitting(true)
        const response = await dispatch(deleteBoard(dialog.board.id))
        setIsSubmitting(false)
        const ok = response.meta.requestStatus === "fulfilled"
        dispatch(notificationActions.addNotification({
            message: t(ok ? "boards.deleteSuccess" : "boards.deleteError"),
            type: ok ? "success" : "error",
        }))
        if (ok) setDialog(null)
    }

    const list = boards ?? []
    const showSkeleton = isLoading && boards === null

    return (
        <div className="space-y-6 px-6 py-8 lg:px-20 lg:py-10">
            <div className="space-y-1">
                <div className="flex items-center justify-between gap-3">
                    <div className="flex items-baseline gap-2">
                        <Label className="text-xl lg:text-2xl">
                            {t("boards.title")}
                        </Label>
                        {boards && (
                            <span className="text-sm text-muted-foreground">{boards.length}</span>
                        )}
                    </div>
                    <Button
                        size="sm"
                        onClick={openCreate}
                        className="rounded-full md:h-8 md:gap-1.5 md:px-2.5 md:text-sm"
                    >
                        <Plus className="size-3 md:size-3.5" />
                        {t("boards.create")}
                    </Button>
                </div>
                <Label className="text-sm font-normal text-primary/50 lg:text-base">
                    {t("boards.subtitle")}
                </Label>
            </div>

            {showSkeleton && (
                <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
                    {Array.from({ length: 4 }, (_, index) => (
                        <div key={index} className="overflow-hidden rounded-2xl bg-card ring-1 ring-foreground/10">
                            <div className="aspect-[16/10] animate-pulse bg-muted" />
                            <div className="h-12 animate-pulse bg-muted/60" />
                        </div>
                    ))}
                </div>
            )}

            {!showSkeleton && list.length === 0 && (
                <button
                    type="button"
                    onClick={openCreate}
                    className="flex w-full flex-col items-center justify-center gap-3 rounded-2xl border border-dashed border-foreground/15 bg-card px-6 py-16 text-center transition-colors hover:border-foreground/30 hover:bg-muted/40"
                >
                    <span className="flex size-12 items-center justify-center rounded-2xl bg-muted text-foreground">
                        <Plus className="size-5" />
                    </span>
                    <span className="text-base font-medium">{t("boards.empty")}</span>
                    <span className="max-w-sm text-sm text-muted-foreground">{t("boards.emptyText")}</span>
                </button>
            )}

            {list.length > 0 && (
                <div className="grid auto-rows-fr grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
                    {list.map((board) => (
                        <BoardCard
                            key={board.id}
                            board={board}
                            openLabel={t("boards.open")}
                            renameLabel={t("boards.rename")}
                            deleteLabel={t("common.delete")}
                            onOpen={() => navigate(getRouteBoard(board.id))}
                            onRename={() => openRename(board)}
                            onDelete={() => setDialog({ mode: "delete", board })}
                        />
                    ))}
                    <button
                        type="button"
                        onClick={openCreate}
                        className="flex h-full min-h-44 flex-col items-center justify-center gap-2 rounded-2xl border border-dashed border-foreground/15 bg-card text-sm text-muted-foreground transition-colors hover:border-[#0d99ff]/50 hover:bg-[#0d99ff]/5 hover:text-foreground"
                    >
                        <Plus className="size-5" />
                        {t("boards.create")}
                    </button>
                </div>
            )}

            <Dialog open={dialog !== null} onOpenChange={(open) => { if (!open) setDialog(null) }}>
                <DialogContent className="p-5 sm:max-w-md">
                    <DialogHeader>
                        <DialogTitle className="pr-10 text-left text-base">
                            {dialog?.mode === "rename"
                                ? t("boards.renameTitle")
                                : dialog?.mode === "delete"
                                    ? t("boards.deleteTitle")
                                    : t("boards.createTitle")}
                        </DialogTitle>
                    </DialogHeader>
                    <Separator className="my-1" />
                    {dialog?.mode === "delete" ? (
                        <div className="space-y-4">
                            <p className="text-sm text-muted-foreground">
                                {t("boards.deleteText", { title: dialog.board.title })}
                            </p>
                            <div className="flex justify-end gap-2">
                                <Button type="button" variant="outline" onClick={() => setDialog(null)}>
                                    {t("common.cancel")}
                                </Button>
                                <Button type="button" variant="destructive" disabled={isSubmitting} onClick={onDelete}>
                                    {isSubmitting ? t("common.loading") : t("common.delete")}
                                </Button>
                            </div>
                        </div>
                    ) : (
                        <form className="w-full space-y-4" onSubmit={onSubmit}>
                            <Field className="w-full">
                                <FieldLabel>{t("boards.name")}</FieldLabel>
                                <Input
                                    className="w-full"
                                    required
                                    autoFocus
                                    value={title}
                                    onChange={(event) => setTitle(event.target.value)}
                                />
                            </Field>
                            <Button type="submit" className="w-full" disabled={isSubmitting}>
                                {isSubmitting ? t("common.loading") : t("common.save")}
                            </Button>
                        </form>
                    )}
                </DialogContent>
            </Dialog>
        </div>
    )
}

function BoardCard({
    board,
    openLabel,
    renameLabel,
    deleteLabel,
    onOpen,
    onRename,
    onDelete,
}: {
    board: Board
    openLabel: string
    renameLabel: string
    deleteLabel: string
    onOpen: () => void
    onRename: () => void
    onDelete: () => void
}) {
    return (
        <article className="group flex h-full flex-col gap-2 rounded-2xl bg-card p-2 ring-1 ring-foreground/10 transition-shadow hover:shadow-md">
            <button
                type="button"
                onClick={onOpen}
                className="relative block aspect-[16/10] w-full overflow-hidden rounded-xl bg-secondary text-left ring-1 ring-foreground/10"
            >
                <span
                    className="absolute inset-0"
                    style={{
                        backgroundImage: "radial-gradient(circle, color-mix(in oklch, var(--foreground) 32%, transparent) 1.05px, transparent 1.15px)",
                        backgroundSize: "18px 18px",
                    }}
                />
                <BoardSketch id={board.id} />
                <span className="absolute right-3 bottom-3 rounded-full bg-background/90 px-2.5 py-1 text-xs font-medium text-foreground opacity-0 shadow-sm ring-1 ring-foreground/10 transition-opacity group-hover:opacity-100">
                    {openLabel}
                </span>
            </button>
            <div className="flex items-center gap-1 px-1.5 pb-1">
                <button
                    type="button"
                    onClick={onOpen}
                    className="min-w-0 flex-1 truncate text-left text-sm font-medium"
                >
                    {board.title}
                </button>
                <DropdownMenu modal={false}>
                    <DropdownMenuTrigger
                        aria-label={renameLabel}
                        className="flex size-7 items-center justify-center rounded-md text-muted-foreground hover:bg-muted hover:text-foreground"
                    >
                        <MoreHorizontal className="size-4" />
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end" className="w-44 min-w-44 bg-background">
                        <DropdownMenuItem className="cursor-pointer gap-2 text-sm" onClick={onRename}>
                            <Pen className="size-4" />
                            {renameLabel}
                        </DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem className="cursor-pointer gap-2 text-sm" variant="destructive" onClick={onDelete}>
                            <Trash2 className="size-4" />
                            {deleteLabel}
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        </article>
    )
}

function BoardSketch({ id }: { id: number }) {
    const variant = Math.abs(id) % 3
    return (
        <svg viewBox="0 0 320 200" className="absolute inset-0 h-full w-full text-foreground" aria-hidden>
            {variant === 0 && (
                <>
                    <rect x="28" y="36" width="118" height="78" rx="10" fill="#0d99ff" fillOpacity="0.12" stroke="#0d99ff" strokeWidth="2.5" />
                    <ellipse cx="214" cy="86" rx="42" ry="30" fill="none" stroke="#0d99ff" strokeWidth="2.5" />
                    <path d="M46 146 H250" stroke="currentColor" strokeOpacity="0.28" strokeWidth="2" strokeLinecap="round" />
                    <path d="M46 162 H168" stroke="currentColor" strokeOpacity="0.16" strokeWidth="2" strokeLinecap="round" />
                </>
            )}
            {variant === 1 && (
                <>
                    <path d="M40 150 L150 42" stroke="#0d99ff" strokeWidth="2.5" strokeLinecap="round" />
                    <path d="M150 42 L248 118" stroke="#0d99ff" strokeWidth="2.5" strokeLinecap="round" markerEnd={`url(#board-sketch-arrow-${id})`} />
                    <rect x="78" y="88" width="92" height="52" rx="8" fill="#0d99ff" fillOpacity="0.12" stroke="#0d99ff" strokeWidth="2.5" />
                    <path d="M40 170 H200" stroke="currentColor" strokeOpacity="0.2" strokeWidth="2" strokeLinecap="round" />
                </>
            )}
            {variant === 2 && (
                <>
                    <polygon points="70,34 132,138 18,138" fill="#0d99ff" fillOpacity="0.12" stroke="#0d99ff" strokeWidth="2.5" strokeLinejoin="round" />
                    <rect x="168" y="48" width="112" height="64" rx="10" fill="none" stroke="#0d99ff" strokeWidth="2.5" />
                    <path d="M168 140 H286" stroke="currentColor" strokeOpacity="0.28" strokeWidth="2" strokeLinecap="round" />
                    <path d="M168 156 H236" stroke="currentColor" strokeOpacity="0.16" strokeWidth="2" strokeLinecap="round" />
                </>
            )}
            <defs>
                <marker id={`board-sketch-arrow-${id}`} markerWidth="8" markerHeight="8" refX="7" refY="4" orient="auto">
                    <path d="M0,0 L8,4 L0,8 Z" fill="#0d99ff" />
                </marker>
            </defs>
        </svg>
    )
}
