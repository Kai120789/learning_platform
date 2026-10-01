import type { BoardItem, BoardPoint } from "@/entities/board"
import { useEffect, useRef, useState, type PointerEvent as ReactPointerEvent, type ReactNode } from "react"
import {
    appendInk,
    axisDelta,
    cloneItem,
    constrainPoint,
    finishInk,
    idsInMarquee,
    itemFromClick,
    itemFromDrag,
    itemFromInk,
    moveEndpointRotated,
    moveItem,
    paintOrder,
    PEN_STROKE_WIDTH,
    resizeRotated,
    rotateItem,
    type BoardTool,
    type ResizeHandle,
} from "../lib/items"

const MIN_ZOOM = 0.1
const MAX_ZOOM = 4
const HANDLE = 7
const ACCENT = "#0d99ff"

type Camera = {
    zoom: number
    x: number
    y: number
}

type Point = { x: number; y: number }
type Mods = { shift: boolean; alt: boolean }
type DrawTool = Exclude<BoardTool, "select" | "text" | "image" | "pen">

type MoveDrag = {
    mode: "move"
    originals: BoardItem[]
    origin: Point
    point: Point
    clientX: number
    clientY: number
    changed: boolean
    copyIds: string[]
    collapseTo: string | null
}

type ResizeDrag = {
    mode: "resize"
    item: BoardItem
    handle: ResizeHandle
    point: Point
    clientX: number
    clientY: number
    changed: boolean
}

type EndpointDrag = {
    mode: "endpoint"
    item: BoardItem
    index: number
    point: Point
    clientX: number
    clientY: number
    changed: boolean
}

type DrawDrag = {
    mode: "draw"
    tool: DrawTool
    origin: Point
    point: Point
    shift: boolean
    alt: boolean
}

type MarqueeDrag = {
    mode: "marquee"
    origin: Point
    point: Point
    clientX: number
    clientY: number
    additive: boolean
    base: string[]
    changed: boolean
    emitted: string
}

type PenDrag = {
    mode: "pen"
    points: BoardPoint[]
    point: Point
}

type RotateDrag = {
    mode: "rotate"
    item: BoardItem
    offset: number
    point: Point
    clientX: number
    clientY: number
    changed: boolean
}

type Drag = MoveDrag | ResizeDrag | EndpointDrag | DrawDrag | MarqueeDrag | PenDrag | RotateDrag

type BoardCanvasProps = {
    items: BoardItem[]
    tool: BoardTool
    authorId: number
    color: string
    fill: string
    text: string
    imageUrl: string
    selectedIds: string[]
    onSelect: (ids: string[]) => void
    onAdd: (item: BoardItem) => void
    onAddItems: (items: BoardItem[]) => void
    onUpdate: (item: BoardItem) => void
    onDeselect: () => void
    background?: "dots" | "grid" | "none"
}

const handles: { id: ResizeHandle; cursor: string; px: number; py: number }[] = [
    { id: "nw", cursor: "nwse-resize", px: 0, py: 0 },
    { id: "n", cursor: "ns-resize", px: 0.5, py: 0 },
    { id: "ne", cursor: "nesw-resize", px: 1, py: 0 },
    { id: "e", cursor: "ew-resize", px: 1, py: 0.5 },
    { id: "se", cursor: "nwse-resize", px: 1, py: 1 },
    { id: "s", cursor: "ns-resize", px: 0.5, py: 1 },
    { id: "sw", cursor: "nesw-resize", px: 0, py: 1 },
    { id: "w", cursor: "ew-resize", px: 0, py: 0.5 },
]

function isAdditive(event: { shiftKey: boolean; metaKey: boolean; ctrlKey: boolean }) {
    return event.shiftKey || event.metaKey || event.ctrlKey
}

function screenMoved(clientX: number, clientY: number, event: { clientX: number; clientY: number }) {
    return Math.hypot(event.clientX - clientX, event.clientY - clientY) > 3
}

export function BoardCanvas({
    items,
    tool,
    authorId,
    color,
    fill,
    text,
    imageUrl,
    selectedIds,
    onSelect,
    onAdd,
    onAddItems,
    onUpdate,
    onDeselect,
    background = "dots",
}: BoardCanvasProps) {
    const svgRef = useRef<SVGSVGElement>(null)
    const dragRef = useRef<Drag | null>(null)
    const ghostsRef = useRef<BoardItem[] | null>(null)
    const ghostExtraRef = useRef(false)
    const cameraRef = useRef<Camera>({ zoom: 1, x: 80, y: 72 })
    const syncRef = useRef<(point: Point, mods: Mods) => void>(() => {})
    const onDeselectRef = useRef(onDeselect)
    const onSelectRef = useRef(onSelect)
    onDeselectRef.current = onDeselect
    onSelectRef.current = onSelect
    const [camera, setCamera] = useState<Camera>(cameraRef.current)
    const [draft, setDraft] = useState<DrawDrag | null>(null)
    const [ink, setInk] = useState<BoardPoint[] | null>(null)
    const [marquee, setMarquee] = useState<MarqueeDrag | null>(null)
    const [ghosts, setGhosts] = useState<BoardItem[] | null>(null)
    const [ghostExtra, setGhostExtra] = useState(false)

    const commitCamera = (next: Camera) => {
        cameraRef.current = next
        setCamera(next)
    }

    const zoomAt = (factor: number, sx: number, sy: number) => {
        const current = cameraRef.current
        const zoom = Math.min(MAX_ZOOM, Math.max(MIN_ZOOM, current.zoom * factor))
        const worldX = (sx - current.x) / current.zoom
        const worldY = (sy - current.y) / current.zoom
        commitCamera({
            zoom,
            x: sx - worldX * zoom,
            y: sy - worldY * zoom,
        })
    }

    useEffect(() => {
        const node = svgRef.current
        if (!node) return

        const onWheel = (event: WheelEvent) => {
            event.preventDefault()
            const rect = node.getBoundingClientRect()
            const sx = event.clientX - rect.left
            const sy = event.clientY - rect.top
            if (event.ctrlKey || event.metaKey) {
                zoomAt(Math.exp(-event.deltaY * 0.002), sx, sy)
                return
            }
            const current = cameraRef.current
            commitCamera({
                ...current,
                x: current.x - event.deltaX,
                y: current.y - event.deltaY,
            })
        }

        node.addEventListener("wheel", onWheel, { passive: false })
        return () => node.removeEventListener("wheel", onWheel)
    }, [])

    const rememberGhosts = (next: BoardItem[] | null, extra: boolean) => {
        ghostsRef.current = next
        ghostExtraRef.current = extra
        setGhosts(next)
        setGhostExtra(extra)
    }

    const clearGesture = () => {
        dragRef.current = null
        setDraft(null)
        setInk(null)
        setMarquee(null)
        rememberGhosts(null, false)
    }

    const worldFromClient = (clientX: number, clientY: number) => {
        const rect = svgRef.current?.getBoundingClientRect()
        const current = cameraRef.current
        if (!rect) return { left: 0, top: 0 }
        return {
            left: (clientX - rect.left - current.x) / current.zoom,
            top: (clientY - rect.top - current.y) / current.zoom,
        }
    }

    const samplesOf = (event: ReactPointerEvent) => {
        const native = event.nativeEvent
        const list = native instanceof PointerEvent && native.getCoalescedEvents().length
            ? native.getCoalescedEvents()
            : [event]
        return list.map((sample) => worldFromClient(sample.clientX, sample.clientY))
    }

    const inkSpacing = () => Math.max(0.6, 1.4 / cameraRef.current.zoom)

    const worldPoint = (event: ReactPointerEvent) => {
        const rect = svgRef.current?.getBoundingClientRect()
        const current = cameraRef.current
        if (!rect) return { x: 0, y: 0 }
        return {
            x: (event.clientX - rect.left - current.x) / current.zoom,
            y: (event.clientY - rect.top - current.y) / current.zoom,
        }
    }

    const capture = (event: ReactPointerEvent) => {
        svgRef.current?.setPointerCapture(event.pointerId)
    }

    syncRef.current = (point, mods) => {
        const drag = dragRef.current
        if (!drag) return
        if (drag.mode === "pen") return
        if (drag.mode === "draw") {
            const next = { ...drag, point, shift: mods.shift, alt: mods.alt }
            dragRef.current = next
            setDraft(next)
            return
        }
        if (drag.mode === "marquee") {
            const next = { ...drag, point }
            if (next.changed) {
                const hit = idsInMarquee(items, next.origin.x, next.origin.y, next.point.x, next.point.y)
                const ids = next.additive ? [...new Set([...next.base, ...hit])] : hit
                const emitted = ids.join("\0")
                if (emitted !== next.emitted) {
                    next.emitted = emitted
                    onSelect(ids)
                }
            }
            dragRef.current = next
            setMarquee(next)
            return
        }
        drag.point = point
        if (!drag.changed) {
            rememberGhosts(null, false)
            return
        }
        if (drag.mode === "move") {
            const delta = axisDelta(point.x - drag.origin.x, point.y - drag.origin.y, mods.shift)
            const copies = mods.alt
            const moved = drag.originals.map((item, index) => {
                const next = moveItem(item, delta.dx, delta.dy)
                if (!copies) return next
                return { ...next, object_id: drag.copyIds[index], author_id: authorId }
            })
            rememberGhosts(moved, copies)
            return
        }
        if (drag.mode === "resize") {
            rememberGhosts([resizeRotated(drag.item, drag.handle, point.x, point.y, mods)], false)
            return
        }
        if (drag.mode === "rotate") {
            rememberGhosts([rotateItem(drag.item, point.x, point.y, drag.offset, mods.shift)], false)
            return
        }
        rememberGhosts([moveEndpointRotated(drag.item, drag.index, point.x, point.y, mods.shift)], false)
    }

    useEffect(() => {
        const onKey = (event: KeyboardEvent) => {
            const target = event.target
            if (target instanceof HTMLInputElement || target instanceof HTMLTextAreaElement) return
            if (event.key === "Escape") {
                const drag = dragRef.current
                if (drag) {
                    event.preventDefault()
                    if (drag.mode === "marquee") onSelectRef.current(drag.base)
                    clearGesture()
                    return
                }
                onDeselectRef.current()
                return
            }
            if (event.key !== "Shift" && event.key !== "Alt") return
            const drag = dragRef.current
            if (!drag || drag.mode === "marquee") return
            syncRef.current(drag.point, { shift: event.shiftKey, alt: event.altKey })
        }
        window.addEventListener("keydown", onKey)
        window.addEventListener("keyup", onKey)
        return () => {
            window.removeEventListener("keydown", onKey)
            window.removeEventListener("keyup", onKey)
        }
    }, [])

    const onBackgroundPointerDown = (event: ReactPointerEvent<SVGRectElement>) => {
        if (event.button !== 0) return
        const point = worldPoint(event)
        if (tool === "select") {
            const next: MarqueeDrag = {
                mode: "marquee",
                origin: point,
                point,
                clientX: event.clientX,
                clientY: event.clientY,
                additive: isAdditive(event),
                base: selectedIds,
                changed: false,
                emitted: "*",
            }
            dragRef.current = next
            setMarquee(next)
            capture(event)
            return
        }
        if (tool === "text" || tool === "image") {
            onAdd(itemFromClick({
                tool,
                authorId,
                x: point.x,
                y: point.y,
                color,
                fill,
                text,
                imageUrl,
            }))
            return
        }

        if (tool === "pen") {
            const next: PenDrag = {
                mode: "pen",
                points: [{ left: point.x, top: point.y }],
                point,
            }
            dragRef.current = next
            setInk(next.points)
            capture(event)
            return
        }

        const next: DrawDrag = {
            mode: "draw",
            tool,
            origin: point,
            point,
            shift: event.shiftKey,
            alt: event.altKey,
        }
        dragRef.current = next
        setDraft(next)
        capture(event)
    }

    const onItemPointerDown = (event: ReactPointerEvent, item: BoardItem) => {
        event.stopPropagation()
        if (event.button !== 0 || tool !== "select") return
        const point = worldPoint(event)
        const additive = isAdditive(event)
        let nextIds = selectedIds
        let collapseTo: string | null = null
        if (additive) {
            nextIds = selectedIds.includes(item.object_id)
                ? selectedIds.filter((id) => id !== item.object_id)
                : [...selectedIds, item.object_id]
            onSelect(nextIds)
        } else if (!selectedIds.includes(item.object_id)) {
            nextIds = [item.object_id]
            onSelect(nextIds)
        } else if (selectedIds.length > 1) {
            collapseTo = item.object_id
        }

        const originals = items.filter((one) => nextIds.includes(one.object_id)).map(cloneItem)
        if (!originals.length) return
        dragRef.current = {
            mode: "move",
            originals,
            origin: point,
            point,
            clientX: event.clientX,
            clientY: event.clientY,
            changed: false,
            copyIds: originals.map(() => crypto.randomUUID()),
            collapseTo,
        }
        capture(event)
    }

    const onResizePointerDown = (event: ReactPointerEvent, item: BoardItem, handle: ResizeHandle) => {
        event.stopPropagation()
        if (event.button !== 0) return
        const point = worldPoint(event)
        dragRef.current = {
            mode: "resize",
            item: cloneItem(item),
            handle,
            point,
            clientX: event.clientX,
            clientY: event.clientY,
            changed: false,
        }
        capture(event)
    }

    const onEndpointPointerDown = (event: ReactPointerEvent, item: BoardItem, index: number) => {
        event.stopPropagation()
        if (event.button !== 0) return
        const point = worldPoint(event)
        dragRef.current = {
            mode: "endpoint",
            item: cloneItem(item),
            index,
            point,
            clientX: event.clientX,
            clientY: event.clientY,
            changed: false,
        }
        capture(event)
    }

    const onRotatePointerDown = (event: ReactPointerEvent, item: BoardItem) => {
        event.stopPropagation()
        if (event.button !== 0) return
        const point = worldPoint(event)
        const cx = item.left + item.width / 2
        const cy = item.top + item.height / 2
        const pointer = Math.atan2(point.y - cy, point.x - cx) * 180 / Math.PI
        dragRef.current = {
            mode: "rotate",
            item: cloneItem(item),
            offset: item.angle - pointer,
            point,
            clientX: event.clientX,
            clientY: event.clientY,
            changed: false,
        }
        capture(event)
    }

    const onPointerMove = (event: ReactPointerEvent<SVGSVGElement>) => {
        const drag = dragRef.current
        if (!drag) return
        if (drag.mode === "pen") {
            const points = appendInk(drag.points, samplesOf(event), inkSpacing())
            if (points === drag.points) return
            dragRef.current = { ...drag, points, point: { x: points[points.length - 1].left, y: points[points.length - 1].top } }
            setInk(points)
            return
        }
        if (drag.mode === "move" || drag.mode === "resize" || drag.mode === "endpoint" || drag.mode === "marquee" || drag.mode === "rotate") {
            drag.changed = drag.changed || screenMoved(drag.clientX, drag.clientY, event)
        }
        syncRef.current(worldPoint(event), { shift: event.shiftKey, alt: event.altKey })
    }

    const onPointerUp = (event: ReactPointerEvent<SVGSVGElement>) => {
        const drag = dragRef.current
        if (!drag) return
        dragRef.current = null

        if (drag.mode === "marquee") {
            const moved = drag.changed || screenMoved(drag.clientX, drag.clientY, event)
            if (!moved) {
                if (!drag.additive && !isAdditive(event)) onSelect([])
            } else {
                const hit = idsInMarquee(items, drag.origin.x, drag.origin.y, drag.point.x, drag.point.y)
                const additive = drag.additive || isAdditive(event)
                onSelect(additive ? [...new Set([...drag.base, ...hit])] : hit)
            }
            setMarquee(null)
            return
        }

        if (drag.mode === "pen") {
            const item = itemFromInk({
                authorId,
                points: finishInk(drag.points, samplesOf(event), inkSpacing()),
                color,
            })
            if (item) onAdd(item)
            setInk(null)
            return
        }

        if (drag.mode === "draw") {
            const adjusted = constrainPoint(drag.origin.x, drag.origin.y, drag.point.x, drag.point.y, {
                shift: drag.shift,
                alt: drag.alt,
                kind: drag.tool === "line" || drag.tool === "arrow" ? "line" : "box",
            })
            const item = itemFromDrag({
                tool: drag.tool,
                authorId,
                startX: adjusted.startX,
                startY: adjusted.startY,
                x: adjusted.x,
                y: adjusted.y,
                color,
                fill,
            })
            if (item) onAdd(item)
            setDraft(null)
            return
        }

        if (drag.mode === "move") {
            if (!drag.changed) {
                if (drag.collapseTo) onSelect([drag.collapseTo])
            } else if (ghostExtraRef.current && ghostsRef.current) {
                onAddItems(ghostsRef.current)
            } else if (ghostsRef.current) {
                ghostsRef.current.forEach((item) => onUpdate(item))
            }
            rememberGhosts(null, false)
            return
        }

        if (drag.changed && ghostsRef.current?.[0]) onUpdate(ghostsRef.current[0])
        rememberGhosts(null, false)
    }

    const preview = draft
        ? itemFromDrag({
            tool: draft.tool,
            authorId,
            ...constrainPoint(draft.origin.x, draft.origin.y, draft.point.x, draft.point.y, {
                shift: draft.shift,
                alt: draft.alt,
                kind: draft.tool === "line" || draft.tool === "arrow" ? "line" : "box",
            }),
            color,
            fill,
        })
        : null

    const ghostById = !ghostExtra && ghosts ? new Map(ghosts.map((item) => [item.object_id, item])) : null
    const visible = paintOrder(items.map((item) => ghostById?.get(item.object_id) ?? item))
    const painted = ghostExtra && ghosts ? [...visible, ...ghosts] : visible
    const chromeItems = ghostExtra && ghosts?.length
        ? ghosts
        : selectedIds.flatMap((id) => {
            const item = visible.find((one) => one.object_id === id)
            return item ? [item] : []
        })
    const marqueeBox = marquee && {
        x: Math.min(marquee.origin.x, marquee.point.x),
        y: Math.min(marquee.origin.y, marquee.point.y),
        w: Math.abs(marquee.point.x - marquee.origin.x),
        h: Math.abs(marquee.point.y - marquee.origin.y),
    }

    const inkItem: BoardItem | null = ink?.length
        ? {
            object_id: "ink",
            author_id: authorId,
            type: "LINE",
            left: 0,
            top: 0,
            width: 0,
            height: 0,
            angle: 0,
            points: ink,
            color,
            stroke_width: PEN_STROKE_WIDTH,
        }
        : null

    const zoomBy = (factor: number) => {
        const rect = svgRef.current?.getBoundingClientRect()
        zoomAt(factor, (rect?.width ?? 0) / 2, (rect?.height ?? 0) / 2)
    }

    return (
        <>
            <svg
                ref={svgRef}
                className="absolute inset-0 h-full w-full touch-none select-none"
                onPointerMove={onPointerMove}
                onPointerUp={onPointerUp}
                onPointerCancel={onPointerUp}
                onContextMenu={(event) => event.preventDefault()}
            >
                <defs>
                    <pattern id="board-dots" width="24" height="24" patternUnits="userSpaceOnUse">
                        <circle cx="1" cy="1" r="1.15" className="fill-foreground/45" />
                    </pattern>
                    <pattern id="board-grid" width="24" height="24" patternUnits="userSpaceOnUse">
                        <path d="M 24 0 L 0 0 0 24" fill="none" className="stroke-foreground/25" strokeWidth="1" />
                    </pattern>
                    {painted.map((item) => item.figure_type === "ARROW" ? (
                        <marker
                            key={item.object_id}
                            id={markerId(item.object_id)}
                            markerWidth="8"
                            markerHeight="8"
                            refX="7"
                            refY="4"
                            orient="auto"
                        >
                            <path d="M0,0 L8,4 L0,8 Z" fill={item.color || "#111827"} />
                        </marker>
                    ) : null)}
                    {preview?.figure_type === "ARROW" && (
                        <marker id={markerId("preview")} markerWidth="8" markerHeight="8" refX="7" refY="4" orient="auto">
                            <path d="M0,0 L8,4 L0,8 Z" fill={preview.color || "#111827"} />
                        </marker>
                    )}
                </defs>
                <g transform={`translate(${camera.x} ${camera.y}) scale(${camera.zoom})`}>
                    <rect
                        x={-8000}
                        y={-8000}
                        width={16000}
                        height={16000}
                        fill={background === "grid" ? "url(#board-grid)" : background === "dots" ? "url(#board-dots)" : "transparent"}
                        className={tool === "select" ? "cursor-default" : "cursor-crosshair"}
                        onPointerDown={onBackgroundPointerDown}
                    />
                    <g style={{ pointerEvents: tool === "select" ? "auto" : "none" }}>
                        {visible.map((item) => (
                            <BoardShape
                                key={item.object_id}
                                item={item}
                                onPointerDown={(event) => onItemPointerDown(event, item)}
                                onDoubleClick={() => {
                                    if (item.type !== "TEXT") return
                                    const next = window.prompt("Текст", item.text || "")
                                    if (next == null) return
                                    onUpdate({ ...item, text: next })
                                }}
                            />
                        ))}
                    </g>
                    {ghostExtra && ghosts?.map((item) => (
                        <g key={item.object_id} pointerEvents="none">
                            <BoardShape item={item} />
                        </g>
                    ))}
                    {preview && <BoardShape item={{ ...preview, object_id: "preview" }} />}
                    {inkItem && (
                        <g pointerEvents="none">
                            <BoardShape item={inkItem} />
                        </g>
                    )}
                    {marqueeBox && (
                        <rect
                            x={marqueeBox.x}
                            y={marqueeBox.y}
                            width={marqueeBox.w}
                            height={marqueeBox.h}
                            fill="rgba(13,153,255,0.12)"
                            stroke={ACCENT}
                            strokeWidth={1}
                            vectorEffect="non-scaling-stroke"
                            pointerEvents="none"
                        />
                    )}
                </g>
                {tool === "select" && chromeItems.map((item) => (
                    <SelectionChrome
                        key={item.object_id}
                        item={item}
                        camera={camera}
                        interactive={chromeItems.length === 1}
                        onResize={(event, handle) => onResizePointerDown(event, item, handle)}
                        onEndpoint={(event, index) => onEndpointPointerDown(event, item, index)}
                        onRotate={(event) => onRotatePointerDown(event, item)}
                    />
                ))}
            </svg>
            <div className="absolute right-4 bottom-4 z-20 flex items-center gap-0.5 rounded-xl border border-border bg-background p-1 shadow-lg">
                <ZoomButton label="−" onClick={() => zoomBy(1 / 1.15)} />
                <button
                    type="button"
                    className="h-8 min-w-14 rounded-lg px-2 text-xs font-medium text-foreground hover:bg-accent"
                    onClick={() => commitCamera({ zoom: 1, x: 80, y: 72 })}
                >
                    {Math.round(camera.zoom * 100)}%
                </button>
                <ZoomButton label="+" onClick={() => zoomBy(1.15)} />
            </div>
        </>
    )
}

function ZoomButton({ label, onClick }: { label: string; onClick: () => void }) {
    return (
        <button
            type="button"
            onClick={onClick}
            className="flex size-8 items-center justify-center rounded-lg text-base text-foreground hover:bg-accent"
        >
            {label}
        </button>
    )
}

function SelectionChrome({
    item,
    camera,
    interactive,
    onResize,
    onEndpoint,
    onRotate,
}: {
    item: BoardItem
    camera: Camera
    interactive: boolean
    onResize: (event: ReactPointerEvent, handle: ResizeHandle) => void
    onEndpoint: (event: ReactPointerEvent, index: number) => void
    onRotate: (event: ReactPointerEvent) => void
}) {
    const straight = item.figure_type === "ARROW" || (item.type === "LINE" && (item.points?.length ?? 2) === 2)
    const cx = (item.left + item.width / 2) * camera.zoom + camera.x
    const cy = (item.top + item.height / 2) * camera.zoom + camera.y
    const top = item.top * camera.zoom + camera.y
    const handleY = Math.min(top, cy) - 28

    let body: ReactNode
    if (straight) {
        const points = item.points?.length
            ? item.points
            : [
                { left: item.left, top: item.top },
                { left: item.left + item.width, top: item.top + item.height },
            ]
        const screen = points.map((point) => ({
            x: point.left * camera.zoom + camera.x,
            y: point.top * camera.zoom + camera.y,
        }))
        body = (
            <>
                {screen.length >= 2 && (
                    <line
                        x1={screen[0].x}
                        y1={screen[0].y}
                        x2={screen[1].x}
                        y2={screen[1].y}
                        stroke={ACCENT}
                        strokeWidth="1.5"
                        pointerEvents="none"
                    />
                )}
                {interactive && screen.map((point, index) => (
                    <circle
                        key={index}
                        cx={point.x}
                        cy={point.y}
                        r="5.5"
                        fill="#fff"
                        stroke={ACCENT}
                        strokeWidth="1.5"
                        className="cursor-crosshair"
                        onPointerDown={(event) => onEndpoint(event, index)}
                    />
                ))}
            </>
        )
    } else {
        const box = {
            x: item.left * camera.zoom + camera.x,
            y: item.top * camera.zoom + camera.y,
            w: Math.max(item.width, 1) * camera.zoom,
            h: Math.max(item.height, 1) * camera.zoom,
        }
        body = (
            <>
                <rect
                    x={box.x}
                    y={box.y}
                    width={box.w}
                    height={box.h}
                    fill="none"
                    stroke={ACCENT}
                    strokeWidth="1.5"
                    pointerEvents="none"
                />
                {interactive && handles.map((handle) => {
                    const hx = box.x + box.w * handle.px
                    const hy = box.y + box.h * handle.py
                    return (
                        <rect
                            key={handle.id}
                            x={hx - HANDLE / 2}
                            y={hy - HANDLE / 2}
                            width={HANDLE}
                            height={HANDLE}
                            rx="1.5"
                            fill="#fff"
                            stroke={ACCENT}
                            strokeWidth="1.5"
                            className={cursorClass(handle.cursor)}
                            onPointerDown={(event) => onResize(event, handle.id)}
                        />
                    )
                })}
            </>
        )
    }

    return (
        <g transform={item.angle ? `rotate(${item.angle} ${cx} ${cy})` : undefined}>
            {body}
            {interactive && (
                <>
                    <line
                        x1={cx}
                        y1={Math.min(top, cy)}
                        x2={cx}
                        y2={handleY}
                        stroke={ACCENT}
                        strokeWidth="1.5"
                        pointerEvents="none"
                    />
                    <circle
                        cx={cx}
                        cy={handleY}
                        r="6"
                        fill="#fff"
                        stroke={ACCENT}
                        strokeWidth="1.5"
                        className="cursor-grab"
                        onPointerDown={onRotate}
                    />
                </>
            )}
        </g>
    )
}

function cursorClass(cursor: string) {
    if (cursor === "nwse-resize") return "cursor-nwse-resize"
    if (cursor === "nesw-resize") return "cursor-nesw-resize"
    if (cursor === "ns-resize") return "cursor-ns-resize"
    return "cursor-ew-resize"
}

function markerId(objectId: string) {
    return `arrow-${objectId}`
}

function inkPath(points: BoardPoint[]) {
    const first = points[0]
    let d = `M ${first.left} ${first.top}`
    if (points.length === 2) return `${d} L ${points[1].left} ${points[1].top}`
    for (let i = 1; i < points.length - 1; i++) {
        const curr = points[i]
        const next = points[i + 1]
        d += ` Q ${curr.left} ${curr.top} ${(curr.left + next.left) / 2} ${(curr.top + next.top) / 2}`
    }
    const last = points[points.length - 1]
    return `${d} L ${last.left} ${last.top}`
}

function BoardShape({
    item,
    onPointerDown,
    onDoubleClick,
}: {
    item: BoardItem
    onPointerDown?: (event: ReactPointerEvent) => void
    onDoubleClick?: () => void
}) {
    const shape = (
        <BoardShapeBody
            item={item}
            onPointerDown={onPointerDown}
            onDoubleClick={onDoubleClick}
        />
    )
    if (!item.angle) return shape
    const cx = item.left + item.width / 2
    const cy = item.top + item.height / 2
    return <g transform={`rotate(${item.angle} ${cx} ${cy})`}>{shape}</g>
}

function BoardShapeBody({
    item,
    onPointerDown,
    onDoubleClick,
}: {
    item: BoardItem
    onPointerDown?: (event: ReactPointerEvent) => void
    onDoubleClick?: () => void
}) {
    const stroke = item.color || "#111827"
    const shapeFill = item.fill || "transparent"
    const strokeWidth = item.stroke_width || 2
    const common = {
        onPointerDown,
        onDoubleClick,
        className: onPointerDown ? "cursor-move" : undefined,
    }

    if (item.type === "LINE" && item.points && item.points.length > 2) {
        const d = inkPath(item.points)
        return (
            <g {...common}>
                <path d={d} fill="none" stroke="transparent" strokeWidth={Math.max(16, strokeWidth)} strokeLinecap="round" strokeLinejoin="round" />
                <path d={d} fill="none" stroke={stroke} strokeWidth={strokeWidth} strokeLinecap="round" strokeLinejoin="round" />
            </g>
        )
    }

    if (item.type === "LINE" && item.points?.length === 1) {
        const point = item.points[0]
        const radius = Math.max(strokeWidth / 2, 1.5)
        return (
            <g {...common}>
                <circle cx={point.left} cy={point.top} r={Math.max(radius, 8)} fill="transparent" />
                <circle cx={point.left} cy={point.top} r={radius} fill={stroke} />
            </g>
        )
    }

    if (item.type === "LINE" || item.figure_type === "ARROW") {
        const [start, end] = item.points ?? [
            { left: item.left, top: item.top },
            { left: item.left + item.width, top: item.top + item.height },
        ]
        return (
            <g {...common}>
                <line
                    x1={start?.left}
                    y1={start?.top}
                    x2={end?.left}
                    y2={end?.top}
                    stroke="transparent"
                    strokeWidth="16"
                />
                <line
                    x1={start?.left}
                    y1={start?.top}
                    x2={end?.left}
                    y2={end?.top}
                    stroke={stroke}
                    strokeWidth={strokeWidth}
                    markerEnd={item.figure_type === "ARROW" ? `url(#${markerId(item.object_id)})` : undefined}
                />
            </g>
        )
    }

    if (item.type === "TEXT") {
        const fontSize = item.font_size || 20
        return (
            <g {...common}>
                <rect x={item.left} y={item.top} width={item.width} height={item.height} fill="transparent" />
                <text x={item.left} y={item.top + fontSize} fontSize={fontSize} fill={stroke}>
                    {item.text}
                </text>
            </g>
        )
    }

    if (item.type === "IMAGE") {
        return (
            <g {...common}>
                {item.image_url ? (
                    <image
                        href={item.image_url}
                        x={item.left}
                        y={item.top}
                        width={item.width}
                        height={item.height}
                        preserveAspectRatio="xMidYMid slice"
                    />
                ) : (
                    <rect
                        x={item.left}
                        y={item.top}
                        width={item.width}
                        height={item.height}
                        fill="#f3f4f6"
                        stroke={stroke}
                        strokeDasharray="6 4"
                    />
                )}
            </g>
        )
    }

    if (item.figure_type === "ELLIPSE") {
        return (
            <ellipse
                {...common}
                cx={item.left + item.width / 2}
                cy={item.top + item.height / 2}
                rx={Math.max(item.width / 2, 1)}
                ry={Math.max(item.height / 2, 1)}
                fill={shapeFill}
                stroke={stroke}
                strokeWidth={strokeWidth}
            />
        )
    }

    if (item.figure_type === "TRIANGLE") {
        const points = [
            `${item.left + item.width / 2},${item.top}`,
            `${item.left + item.width},${item.top + item.height}`,
            `${item.left},${item.top + item.height}`,
        ].join(" ")
        return (
            <polygon
                {...common}
                points={points}
                fill={shapeFill}
                stroke={stroke}
                strokeWidth={strokeWidth}
            />
        )
    }

    return (
        <rect
            {...common}
            x={item.left}
            y={item.top}
            width={item.width}
            height={item.height}
            fill={shapeFill}
            stroke={stroke}
            strokeWidth={strokeWidth}
        />
    )
}
