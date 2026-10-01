import { useBoardSocket, type BoardItem } from "@/entities/board"
import { getRouteBoards, getRouteSchedule } from "@/app/router/routePaths"
import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getUserFullData, useCanEdit } from "@/entities/user"
import { cn } from "@/shared/lib/utils"
import { Input } from "@/shared/ui/Input"
import { Popover, PopoverContent, PopoverTrigger } from "@/shared/ui/Popover"
import { Tooltip, TooltipContent, TooltipTrigger } from "@/shared/ui/Tooltip"
import {
    Ban,
    ArrowLeft,
    ArrowUpRight,
    BringToFront,
    ChevronsDown,
    ChevronsUp,
    Circle,
    Grid3x3,
    Grip,
    Image as ImageIcon,
    MousePointer2,
    PanelBottom,
    PanelLeft,
    PanelRight,
    Pencil,
    SendToBack,
    Slash,
    Square,
    Trash2,
    Triangle,
    Type,
    type LucideIcon,
} from "lucide-react"
import { useCallback, useEffect, useRef, useState } from "react"
import { useTranslation } from "react-i18next"
import { Link, useNavigate, useParams } from "react-router-dom"
import { nextZIndex, restack, roundItem, type BoardTool, type StackMove } from "../lib/items"
import { BoardCanvas } from "./BoardCanvas"

const tools: { id: BoardTool; icon: LucideIcon; key?: string }[] = [
    { id: "select", icon: MousePointer2, key: "V" },
    { id: "line", icon: Slash, key: "L" },
    { id: "pen", icon: Pencil, key: "P" },
    { id: "rect", icon: Square, key: "R" },
    { id: "ellipse", icon: Circle, key: "O" },
    { id: "triangle", icon: Triangle },
    { id: "arrow", icon: ArrowUpRight },
    { id: "text", icon: Type, key: "T" },
    { id: "image", icon: ImageIcon },
]

type ToolDock = "left" | "right" | "bottom"

const DOCK_KEY = "board-tool-dock"

const strokePalette = [
    "#111827", "#6b7280", "#ffffff",
    "#ef4444", "#f97316", "#eab308",
    "#22c55e", "#14b8a6", "#0d99ff",
    "#6366f1", "#a855f7", "#ec4899",
]

const fillPalette = ["transparent", ...strokePalette]

const dockOrder: ToolDock[] = ["left", "right", "bottom"]

const dockIcon: Record<ToolDock, LucideIcon> = {
    left: PanelLeft,
    right: PanelRight,
    bottom: PanelBottom,
}

type BoardBackground = "dots" | "grid" | "none"

const BACKGROUND_KEY = "board-background"
const backgroundOrder: BoardBackground[] = ["dots", "grid", "none"]

const backgroundIcon: Record<BoardBackground, LucideIcon> = {
    dots: Grip,
    grid: Grid3x3,
    none: Ban,
}

const stackActions: { id: StackMove; icon: LucideIcon; shortcut: string }[] = [
    { id: "back", icon: SendToBack, shortcut: "⌘[" },
    { id: "backward", icon: ChevronsDown, shortcut: "[" },
    { id: "forward", icon: ChevronsUp, shortcut: "]" },
    { id: "front", icon: BringToFront, shortcut: "⌘]" },
]

const toolByKey: Record<string, BoardTool> = {
    v: "select",
    l: "line",
    p: "pen",
    r: "rect",
    o: "ellipse",
    t: "text",
}

export default function BoardPage() {
    const { t } = useTranslation()
    const { boardId: rawBoardId } = useParams()
    const boardId = Number(rawBoardId)
    const boardReady = Number.isInteger(boardId) && boardId > 0
    const navigate = useNavigate()
    const canEdit = useCanEdit()
    const userId = useAppSelector(getUserFullData)?.user.userID ?? 0
    const { items, send } = useBoardSocket(boardReady ? boardId : 0)
    const [tool, setTool] = useState<BoardTool>("select")
    const [color, setColor] = useState("#111827")
    const [fill, setFill] = useState("#dbeafe")
    const [text, setText] = useState("Текст")
    const [imageUrl, setImageUrl] = useState("")
    const [selectedIds, setSelectedIds] = useState<string[]>([])
    const [dock, setDock] = useState<ToolDock>(() => readDock())
    const [background, setBackground] = useState<BoardBackground>(() => readBackground())
    const selectedIdsRef = useRef(selectedIds)
    selectedIdsRef.current = selectedIds
    const itemsRef = useRef(items)
    itemsRef.current = items
    const selectedItems = items.filter((item) => selectedIds.includes(item.object_id))
    const selected = selectedItems.length === 1 ? selectedItems[0] : undefined

    useEffect(() => {
        setSelectedIds((prev) => {
            const next = prev.filter((id) => items.some((item) => item.object_id === id))
            return next.length === prev.length ? prev : next
        })
    }, [items])

    useEffect(() => {
        if (boardReady) return
        navigate(canEdit ? getRouteBoards() : getRouteSchedule(), { replace: true })
    }, [boardReady, canEdit, navigate])

    const sendItem = (type: "add" | "update", item: BoardItem) => {
        const stamped = type === "add" ? { ...item, z_index: nextZIndex(itemsRef.current) } : item
        const ok = send({
            envelope: { type, board_id: boardId, version: 0 },
            item: roundItem(stamped),
        })
        if (ok && type === "add") setSelectedIds([item.object_id])
        return ok
    }

    const addItems = (next: BoardItem[]) => {
        const ids: string[] = []
        const known = itemsRef.current.slice()
        for (const item of next) {
            const stamped = { ...item, z_index: nextZIndex(known) }
            const ok = send({
                envelope: { type: "add", board_id: boardId, version: 0 },
                item: roundItem(stamped),
            })
            if (!ok) continue
            ids.push(item.object_id)
            known.push(stamped)
        }
        if (ids.length) setSelectedIds(ids)
    }

    const remove = useCallback((objectIds: string[]) => {
        if (!objectIds.length) return
        const ok = send({
            envelope: { type: "delete", board_id: boardId, version: 0 },
            object_ids: objectIds,
        })
        if (ok) setSelectedIds([])
    }, [send, boardId])

    const paintSelected = (patch: Partial<BoardItem>) => {
        selectedItems.forEach((item) => sendItem("update", { ...item, ...patch }))
    }

    const arrange = useCallback((move: StackMove) => {
        const ids = selectedIdsRef.current
        if (!ids.length) return
        restack(itemsRef.current, ids, move).forEach((item) => {
            send({
                envelope: { type: "update", board_id: boardId, version: 0 },
                item: roundItem(item),
            })
        })
    }, [send, boardId])

    useEffect(() => {
        const onKeyDown = (event: KeyboardEvent) => {
            const target = event.target
            if (target instanceof HTMLInputElement || target instanceof HTMLTextAreaElement) return
            const nextTool = toolByKey[event.key.toLowerCase()]
            if (nextTool && !event.metaKey && !event.ctrlKey && !event.altKey && !event.shiftKey) {
                setTool(nextTool)
                return
            }
            if ((event.metaKey || event.ctrlKey) && event.key.toLowerCase() === "a") {
                event.preventDefault()
                setTool("select")
                setSelectedIds(itemsRef.current.map((item) => item.object_id))
                return
            }
            if (event.code === "BracketRight" || event.code === "BracketLeft") {
                if (!selectedIdsRef.current.length) return
                event.preventDefault()
                const forward = event.code === "BracketRight"
                const far = event.metaKey || event.ctrlKey
                arrange(far ? (forward ? "front" : "back") : (forward ? "forward" : "backward"))
                return
            }
            if (event.key !== "Delete" && event.key !== "Backspace") return
            const ids = selectedIdsRef.current
            if (!ids.length) return
            event.preventDefault()
            remove(ids)
        }

        window.addEventListener("keydown", onKeyDown)
        return () => window.removeEventListener("keydown", onKeyDown)
    }, [remove, arrange])

    const placeDock = (next: ToolDock) => {
        setDock(next)
        window.localStorage.setItem(DOCK_KEY, next)
    }

    const tooltipSide = dock === "bottom" ? "top" : dock === "right" ? "left" : "right"
    const DockIcon = dockIcon[dock]
    const BackgroundIcon = backgroundIcon[background]

    const cycleBackground = () => {
        const next = backgroundOrder[(backgroundOrder.indexOf(background) + 1) % backgroundOrder.length]
        setBackground(next)
        window.localStorage.setItem(BACKGROUND_KEY, next)
    }

    const showText = tool === "text" || selected?.type === "TEXT"
    const showImage = tool === "image" || selected?.type === "IMAGE"
    const showFill = selectedItems.length
        ? selectedItems.every((item) => item.type === "FIGURE" && item.figure_type !== "ARROW")
        : tool !== "line" && tool !== "pen" && tool !== "arrow" && tool !== "text"

    if (!boardReady) return null

    return (
            <div className="absolute inset-0">
                <BoardCanvas
                    items={items}
                    tool={tool}
                    authorId={userId}
                    color={selected?.color || color}
                    fill={selected?.fill || fill}
                    text={text}
                    imageUrl={imageUrl}
                    selectedIds={selectedIds}
                    onSelect={setSelectedIds}
                    onAdd={(item) => sendItem("add", item)}
                    onAddItems={addItems}
                    onUpdate={(item) => sendItem("update", item)}
                    onDeselect={() => {
                        setSelectedIds([])
                        setTool("select")
                    }}
                    background={background}
                />

                <div className="pointer-events-none absolute inset-x-0 top-0 z-20 flex items-start justify-between gap-3 p-3">
                    <Link
                        to={canEdit ? getRouteBoards() : getRouteSchedule()}
                        className="pointer-events-auto flex h-10 items-center gap-2 rounded-xl border border-border bg-background px-3 text-sm shadow-lg hover:bg-accent"
                    >
                        <ArrowLeft className="size-4" />
                        {t("board.back")}
                    </Link>

                    {selectedIds.length > 0 && (
                        <div className="pointer-events-auto absolute top-3 left-1/2 flex h-10 -translate-x-1/2 items-center gap-0.5 rounded-xl border border-border bg-background px-1 shadow-lg">
                            {stackActions.map((action, index) => {
                                const Icon = action.icon
                                return (
                                    <span key={action.id} className="flex items-center">
                                        {index === 2 && <span className="mx-0.5 h-5 w-px bg-border" />}
                                        <Tooltip>
                                            <TooltipTrigger
                                                type="button"
                                                aria-label={t(`board.stack.${action.id}`)}
                                                onClick={() => arrange(action.id)}
                                                className="flex size-8 items-center justify-center rounded-lg text-foreground hover:bg-accent"
                                            >
                                                <Icon className="size-4" />
                                            </TooltipTrigger>
                                            <TooltipContent side="bottom" sideOffset={8}>
                                                {t(`board.stack.${action.id}`)}
                                                {` · ${action.shortcut}`}
                                            </TooltipContent>
                                        </Tooltip>
                                    </span>
                                )
                            })}
                        </div>
                    )}

                    <div className="pointer-events-auto flex h-10 max-w-[70%] items-center gap-2 rounded-xl border border-border bg-background px-2 shadow-lg">
                        <div className="flex items-center gap-1">
                        <ColorPicker
                            label={t("board.stroke")}
                            value={selected?.color || color}
                            colors={strokePalette}
                            onChange={(next) => {
                                setColor(next)
                                paintSelected({ color: next })
                            }}
                        />
                        {showFill && (
                            <ColorPicker
                                label={t("board.fill")}
                                value={selected?.fill || fill}
                                colors={fillPalette}
                                onChange={(next) => {
                                    setFill(next)
                                    paintSelected({ fill: next })
                                }}
                            />
                        )}
                        </div>
                        {showText && (
                            <Input
                                value={selected?.type === "TEXT" ? selected.text || "" : text}
                                onChange={(event) => {
                                    const next = event.target.value
                                    if (selected?.type === "TEXT") paintSelected({ text: next })
                                    else setText(next)
                                }}
                                placeholder={t("board.text")}
                                className="h-8 w-36"
                            />
                        )}
                        {showImage && (
                            <Input
                                value={selected?.type === "IMAGE" ? selected.image_url || "" : imageUrl}
                                onChange={(event) => {
                                    const next = event.target.value
                                    if (selected?.type === "IMAGE") paintSelected({ image_url: next })
                                    else setImageUrl(next)
                                }}
                                placeholder={t("board.imageUrl")}
                                className="h-8 w-52"
                            />
                        )}
                    </div>
                </div>

                <aside className={cn(
                    "absolute z-20 flex gap-1 rounded-2xl border border-border bg-background p-1.5 shadow-lg",
                    dock === "left" && "top-1/2 left-3 -translate-y-1/2 flex-col",
                    dock === "right" && "top-1/2 right-3 -translate-y-1/2 flex-col",
                    dock === "bottom" && "bottom-4 left-1/2 -translate-x-1/2 flex-row",
                )}>
                    {tools.map((oneTool) => {
                        const Icon = oneTool.icon
                        const active = tool === oneTool.id
                        return (
                            <Tooltip key={oneTool.id}>
                                <TooltipTrigger
                                    type="button"
                                    aria-label={t(`board.tools.${oneTool.id}`)}
                                    aria-pressed={active}
                                    onClick={() => setTool(oneTool.id)}
                                    className={cn(
                                        "flex size-9 items-center justify-center rounded-xl transition-colors",
                                        active
                                            ? "bg-[#0d99ff] text-white"
                                            : "text-foreground hover:bg-accent"
                                    )}
                                >
                                    <Icon className="size-4" />
                                </TooltipTrigger>
                                <TooltipContent side={tooltipSide} sideOffset={8}>
                                    {t(`board.tools.${oneTool.id}`)}
                                    {oneTool.key ? ` · ${oneTool.key}` : ""}
                                </TooltipContent>
                            </Tooltip>
                        )
                    })}
                    <span className={cn("bg-border", dock === "bottom" ? "mx-0.5 my-1.5 w-px self-stretch" : "mx-1.5 my-0.5 h-px")} />
                    <Tooltip>
                        <TooltipTrigger
                            type="button"
                            aria-label={t("board.delete")}
                            disabled={!selectedIds.length}
                            onClick={() => remove(selectedIds)}
                            className="flex size-9 items-center justify-center rounded-xl text-foreground hover:bg-destructive/10 hover:text-destructive disabled:pointer-events-none disabled:opacity-30"
                        >
                            <Trash2 className="size-4" />
                        </TooltipTrigger>
                        <TooltipContent side={tooltipSide} sideOffset={8}>
                            {t("board.delete")}
                        </TooltipContent>
                    </Tooltip>
                    <span className={cn("bg-border", dock === "bottom" ? "mx-0.5 my-1.5 w-px self-stretch" : "mx-1.5 my-0.5 h-px")} />
                    <Tooltip>
                        <TooltipTrigger
                            type="button"
                            aria-label={t(`board.dock.${dock}`)}
                            onClick={() => placeDock(dockOrder[(dockOrder.indexOf(dock) + 1) % dockOrder.length])}
                            className="flex size-9 items-center justify-center rounded-xl text-foreground hover:bg-accent"
                        >
                            <DockIcon className="size-4" />
                        </TooltipTrigger>
                        <TooltipContent side={tooltipSide} sideOffset={8}>
                            {t(`board.dock.${dock}`)}
                        </TooltipContent>
                    </Tooltip>
                    <Tooltip>
                        <TooltipTrigger
                            type="button"
                            aria-label={t(`board.background.${background}`)}
                            onClick={cycleBackground}
                            className="flex size-9 items-center justify-center rounded-xl text-foreground hover:bg-accent"
                        >
                            <BackgroundIcon className="size-4" />
                        </TooltipTrigger>
                        <TooltipContent side={tooltipSide} sideOffset={8}>
                            {t(`board.background.${background}`)}
                        </TooltipContent>
                    </Tooltip>
                </aside>
            </div>
    )
}

function ColorPicker({
    label,
    value,
    colors,
    onChange,
}: {
    label: string
    value: string
    colors: string[]
    onChange: (value: string) => void
}) {
    const { t } = useTranslation()
    const [hex, setHex] = useState(value === "transparent" ? "" : value)

    const applyHex = () => {
        const next = parseHex(hex)
        if (!next) {
            setHex(value)
            return
        }
        setHex(next)
        onChange(next)
    }

    return (
        <Popover onOpenChange={(open) => { if (open) setHex(value === "transparent" ? "" : value) }}>
            <PopoverTrigger
                type="button"
                aria-label={label}
                className="flex size-8 items-center justify-center rounded-full hover:bg-accent"
            >
                <ColorSwatch color={value} />
            </PopoverTrigger>
            <PopoverContent align="end" className="w-auto gap-2 p-2.5">
                <div className="text-xs font-medium">{label}</div>
                <div className="grid grid-cols-6 gap-1.5">
                    {colors.map((color) => {
                        const active = color.toLowerCase() === value.toLowerCase()
                        return (
                            <button
                                key={color}
                                type="button"
                                aria-label={color}
                                onClick={() => {
                                    setHex(color === "transparent" ? "" : color)
                                    onChange(color)
                                }}
                                className="flex size-6 items-center justify-center"
                            >
                                <ColorSwatch color={color} active={active} />
                            </button>
                        )
                    })}
                </div>
                <div className="flex items-center gap-1.5">
                    <label className="relative size-7 shrink-0 cursor-pointer overflow-hidden rounded-full border border-black/10">
                        <span
                            className="absolute inset-0"
                            style={{ background: "conic-gradient(red, yellow, lime, cyan, blue, magenta, red)" }}
                        />
                        <input
                            type="color"
                            aria-label={t("board.customColor")}
                            value={normalizeColor(value)}
                            onChange={(event) => {
                                setHex(event.target.value)
                                onChange(event.target.value)
                            }}
                            className="absolute inset-0 cursor-pointer opacity-0"
                        />
                    </label>
                    <Input
                        value={hex}
                        onChange={(event) => setHex(event.target.value)}
                        onBlur={applyHex}
                        onKeyDown={(event) => {
                            if (event.key === "Enter") applyHex()
                        }}
                        className="h-8 w-28 font-mono text-xs uppercase"
                    />
                </div>
            </PopoverContent>
        </Popover>
    )
}

function ColorSwatch({ color, active = false }: { color: string; active?: boolean }) {
    const empty = !color || color === "transparent"
    return (
        <span
            className="block size-6 shrink-0 rounded-full border border-black/15 bg-white"
            style={{
                backgroundColor: empty ? undefined : color,
                backgroundImage: empty
                    ? "linear-gradient(135deg, transparent 46%, #bdbdbd 46% 54%, transparent 54%)"
                    : undefined,
                boxShadow: active ? "inset 0 0 0 1.5px #fff, 0 0 0 1.5px #0d99ff" : undefined,
            }}
        />
    )
}

function readBackground(): BoardBackground {
    const saved = window.localStorage.getItem(BACKGROUND_KEY)
    if (saved === "dots" || saved === "grid" || saved === "none") return saved
    return "dots"
}

function readDock(): ToolDock {
    const saved = window.localStorage.getItem(DOCK_KEY)
    if (saved === "left" || saved === "right" || saved === "bottom") return saved
    return "left"
}

function parseHex(value: string) {
    const raw = value.trim()
    if (/^#[0-9a-fA-F]{6}$/.test(raw)) return raw.toLowerCase()
    if (/^#[0-9a-fA-F]{3}$/.test(raw)) {
        const [hash, r, g, b] = raw
        return `${hash}${r}${r}${g}${g}${b}${b}`.toLowerCase()
    }
    return null
}

function normalizeColor(value: string) {
    return parseHex(value) ?? "#111827"
}
