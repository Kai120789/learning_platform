import type { BoardItem, BoardPoint, FigureType } from "@/entities/board"

export type BoardTool = "select" | "line" | "pen" | "rect" | "ellipse" | "triangle" | "arrow" | "text" | "image"

export const PEN_STROKE_WIDTH = 3

export function roundItem(item: BoardItem): BoardItem {
    return {
        ...item,
        left: Math.round(item.left),
        top: Math.round(item.top),
        width: Math.round(Math.max(item.width, 0)),
        height: Math.round(Math.max(item.height, 0)),
        angle: Math.round(item.angle),
        z_index: item.z_index == null ? undefined : Math.round(item.z_index),
        stroke_width: item.stroke_width == null ? undefined : Math.round(item.stroke_width),
        font_size: item.font_size == null ? undefined : Math.round(item.font_size),
        points: item.points?.map((point) => ({
            left: Math.round(point.left),
            top: Math.round(point.top),
        })),
    }
}

export type ResizeHandle = "n" | "ne" | "e" | "se" | "s" | "sw" | "w" | "nw"

type ResizeMods = {
    shift?: boolean
    alt?: boolean
}

export function cloneItem(item: BoardItem): BoardItem {
    return {
        ...item,
        points: item.points?.map((point) => ({ ...point })),
    }
}

export function snapAngle(x0: number, y0: number, x: number, y: number) {
    const dx = x - x0
    const dy = y - y0
    const length = Math.hypot(dx, dy)
    if (length === 0) return { x, y }
    const step = Math.PI / 4
    const angle = Math.round(Math.atan2(dy, dx) / step) * step
    return {
        x: x0 + Math.cos(angle) * length,
        y: y0 + Math.sin(angle) * length,
    }
}

export function axisDelta(dx: number, dy: number, shift: boolean) {
    if (!shift) return { dx, dy }
    if (Math.abs(dx) >= Math.abs(dy)) return { dx, dy: 0 }
    return { dx: 0, dy }
}

export function constrainPoint(
    startX: number,
    startY: number,
    x: number,
    y: number,
    options: { shift?: boolean; alt?: boolean; kind: "line" | "box" },
) {
    let nx = x
    let ny = y
    if (options.shift) {
        if (options.kind === "line") {
            const snapped = snapAngle(startX, startY, x, y)
            nx = snapped.x
            ny = snapped.y
        } else {
            const dx = x - startX
            const dy = y - startY
            const size = Math.max(Math.abs(dx), Math.abs(dy))
            nx = startX + Math.sign(dx || 1) * size
            ny = startY + Math.sign(dy || 1) * size
        }
    }
    if (options.alt) {
        return {
            startX: startX - (nx - startX),
            startY: startY - (ny - startY),
            x: nx,
            y: ny,
        }
    }
    return { startX, startY, x: nx, y: ny }
}

export type WorldBox = { left: number; top: number; right: number; bottom: number }

export function itemBox(item: BoardItem): WorldBox {
    let left = item.left
    let top = item.top
    let right = item.left + item.width
    let bottom = item.top + item.height
    if ((item.type === "LINE" || item.figure_type === "ARROW") && item.points?.length) {
        const xs = item.points.map((point) => point.left)
        const ys = item.points.map((point) => point.top)
        left = Math.min(...xs)
        right = Math.max(...xs)
        top = Math.min(...ys)
        bottom = Math.max(...ys)
    }
    if (item.angle) {
        const turned = rotatedBounds(left, top, right, bottom, item.angle)
        left = turned.left
        top = turned.top
        right = turned.right
        bottom = turned.bottom
    }
    const pad = 6
    if (right - left < pad) {
        const mid = (left + right) / 2
        left = mid - pad / 2
        right = mid + pad / 2
    }
    if (bottom - top < pad) {
        const mid = (top + bottom) / 2
        top = mid - pad / 2
        bottom = mid + pad / 2
    }
    return { left, top, right, bottom }
}

export function idsInMarquee(items: BoardItem[], x0: number, y0: number, x: number, y: number) {
    const box: WorldBox = {
        left: Math.min(x0, x),
        top: Math.min(y0, y),
        right: Math.max(x0, x),
        bottom: Math.max(y0, y),
    }
    return items
        .filter((item) => {
            const next = itemBox(item)
            return next.left <= box.right && next.right >= box.left && next.top <= box.bottom && next.bottom >= box.top
        })
        .map((item) => item.object_id)
}

export function resizeItem(item: BoardItem, handle: ResizeHandle, x: number, y: number, options: ResizeMods = {}): BoardItem {
    const min = 8
    const left0 = item.left
    const top0 = item.top
    const right0 = left0 + item.width
    const bottom0 = top0 + item.height
    const ratio = item.width > 0 && item.height > 0 ? item.width / item.height : 1
    const west = handle === "nw" || handle === "w" || handle === "sw"
    const east = handle === "ne" || handle === "e" || handle === "se"
    const north = handle === "nw" || handle === "n" || handle === "ne"
    const south = handle === "sw" || handle === "s" || handle === "se"
    const corner = (west || east) && (north || south)

    if (options.alt) {
        const cx = (left0 + right0) / 2
        const cy = (top0 + bottom0) / 2
        let halfW = west || east ? Math.abs(x - cx) : item.width / 2
        let halfH = north || south ? Math.abs(y - cy) : item.height / 2
        if (options.shift) {
            if (corner) {
                if (halfW >= halfH * ratio) halfH = halfW / ratio
                else halfW = halfH * ratio
            } else if (west || east) halfH = halfW / ratio
            else halfW = halfH * ratio
        }
        halfW = Math.max(halfW, min / 2)
        halfH = Math.max(halfH, min / 2)
        return fitStroke(item, roundItem({
            ...item,
            left: cx - halfW,
            top: cy - halfH,
            width: halfW * 2,
            height: halfH * 2,
        }))
    }

    const anchorX = west ? right0 : east ? left0 : (left0 + right0) / 2
    const anchorY = north ? bottom0 : south ? top0 : (top0 + bottom0) / 2
    let width = west || east ? (west ? anchorX - x : x - anchorX) : item.width
    let height = north || south ? (north ? anchorY - y : y - anchorY) : item.height

    if (options.shift) {
        if (corner) {
            if (Math.abs(width) >= Math.abs(height) * ratio) height = Math.sign(height || 1) * Math.abs(width) / ratio
            else width = Math.sign(width || 1) * Math.abs(height) * ratio
        } else if (west || east) height = Math.sign(height || 1) * Math.abs(width) / ratio
        else width = Math.sign(width || 1) * Math.abs(height) * ratio
    }

    width = Math.max(Math.abs(width), min) * Math.sign(width || 1)
    height = Math.max(Math.abs(height), min) * Math.sign(height || 1)

    let left = west ? anchorX - width : east ? anchorX : anchorX - width / 2
    let top = north ? anchorY - height : south ? anchorY : anchorY - height / 2
    if (width < 0) {
        left += width
        width = -width
    }
    if (height < 0) {
        top += height
        height = -height
    }

    return fitStroke(item, roundItem({ ...item, left, top, width, height }))
}

function fitStroke(source: BoardItem, next: BoardItem): BoardItem {
    if (!source.points || source.points.length <= 2) return next
    const points = source.points.map((point) => ({
        left: source.width === 0
            ? next.left + next.width / 2
            : next.left + ((point.left - source.left) / source.width) * next.width,
        top: source.height === 0
            ? next.top + next.height / 2
            : next.top + ((point.top - source.top) / source.height) * next.height,
    }))
    return roundItem({ ...next, points })
}

export function moveEndpoint(item: BoardItem, index: number, x: number, y: number, shift = false): BoardItem {
    const base = item.points?.length
        ? item.points.map((point) => ({ ...point }))
        : [
            { left: item.left, top: item.top },
            { left: item.left + item.width, top: item.top + item.height },
        ]
    let px = x
    let py = y
    if (shift && base.length > 1) {
        const other = base[index === 0 ? 1 : 0]
        const snapped = snapAngle(other.left, other.top, x, y)
        px = snapped.x
        py = snapped.y
    }
    const points = base.map((point, pointIndex) => pointIndex === index ? { left: px, top: py } : point)
    const xs = points.map((point) => point.left)
    const ys = points.map((point) => point.top)
    const left = Math.min(...xs)
    const top = Math.min(...ys)

    return roundItem({
        ...item,
        points,
        left,
        top,
        width: Math.max(...xs) - left,
        height: Math.max(...ys) - top,
    })
}

export function moveItem(item: BoardItem, dx: number, dy: number): BoardItem {
    return roundItem({
        ...item,
        left: item.left + dx,
        top: item.top + dy,
        points: item.points?.map((point) => ({
            left: point.left + dx,
            top: point.top + dy,
        })),
    })
}

export function paintOrder(items: BoardItem[]) {
    return items
        .map((item, index) => ({ item, index }))
        .sort((a, b) => (a.item.z_index ?? 0) - (b.item.z_index ?? 0) || a.index - b.index)
        .map((entry) => entry.item)
}

export function nextZIndex(items: BoardItem[]) {
    return items.reduce((max, item) => Math.max(max, item.z_index ?? 0), 0) + 1
}

export type StackMove = "front" | "forward" | "backward" | "back"

export function restack(items: BoardItem[], ids: string[], move: StackMove) {
    const selected = new Set(ids)
    if (!selected.size) return []
    let order = paintOrder(items)
    if (move === "front" || move === "back") {
        const picked = order.filter((item) => selected.has(item.object_id))
        const rest = order.filter((item) => !selected.has(item.object_id))
        order = move === "front" ? [...rest, ...picked] : [...picked, ...rest]
    } else if (move === "forward") {
        order = order.slice()
        for (let i = order.length - 2; i >= 0; i--) {
            if (selected.has(order[i].object_id) && !selected.has(order[i + 1].object_id)) {
                const above = order[i + 1]
                order[i + 1] = order[i]
                order[i] = above
            }
        }
    } else {
        order = order.slice()
        for (let i = 1; i < order.length; i++) {
            if (selected.has(order[i].object_id) && !selected.has(order[i - 1].object_id)) {
                const below = order[i - 1]
                order[i - 1] = order[i]
                order[i] = below
            }
        }
    }

    const changed: BoardItem[] = []
    let prevZ = Number.NEGATIVE_INFINITY
    for (const item of order) {
        const z = item.z_index ?? 0
        if (z > prevZ) {
            prevZ = z
            continue
        }
        const zIndex = prevZ + 1
        changed.push({ ...item, z_index: zIndex })
        prevZ = zIndex
    }
    return changed
}

export function toLocal(item: BoardItem, x: number, y: number) {
    if (!item.angle) return { x, y }
    const cx = item.left + item.width / 2
    const cy = item.top + item.height / 2
    const rad = (-item.angle * Math.PI) / 180
    const cos = Math.cos(rad)
    const sin = Math.sin(rad)
    const dx = x - cx
    const dy = y - cy
    return {
        x: cx + dx * cos - dy * sin,
        y: cy + dx * sin + dy * cos,
    }
}

export function rotateItem(item: BoardItem, x: number, y: number, offset: number, shift = false) {
    const cx = item.left + item.width / 2
    const cy = item.top + item.height / 2
    let angle = Math.atan2(y - cy, x - cx) * 180 / Math.PI + offset
    if (shift) angle = Math.round(angle / 15) * 15
    angle = Math.round(angle)
    angle = ((angle + 180) % 360 + 360) % 360 - 180
    return { ...item, angle }
}

export function resizeRotated(item: BoardItem, handle: ResizeHandle, x: number, y: number, options: ResizeMods = {}) {
    const local = toLocal(item, x, y)
    if (!item.angle || options.alt) return resizeItem(item, handle, local.x, local.y, options)
    const anchor = anchorPoint(item, handle)
    const fixed = toWorld(item, anchor.x, anchor.y)
    const resized = resizeItem(item, handle, local.x, local.y, options)
    const nextAnchor = anchorPoint(resized, handle)
    const moved = toWorld(resized, nextAnchor.x, nextAnchor.y)
    return moveItem(resized, fixed.x - moved.x, fixed.y - moved.y)
}

export function moveEndpointRotated(item: BoardItem, index: number, x: number, y: number, shift = false) {
    const local = toLocal(item, x, y)
    if (!item.angle) return moveEndpoint(item, index, local.x, local.y, shift)
    const base = item.points ?? []
    const otherIndex = index === 0 ? base.length - 1 : 0
    const other = base[otherIndex]
    if (!other) return moveEndpoint(item, index, local.x, local.y, shift)
    const fixed = toWorld(item, other.left, other.top)
    const next = moveEndpoint(item, index, local.x, local.y, shift)
    const nextOther = next.points?.[otherIndex]
    if (!nextOther) return next
    const moved = toWorld(next, nextOther.left, nextOther.top)
    return moveItem(next, fixed.x - moved.x, fixed.y - moved.y)
}

function toWorld(item: BoardItem, x: number, y: number) {
    if (!item.angle) return { x, y }
    const cx = item.left + item.width / 2
    const cy = item.top + item.height / 2
    const rad = (item.angle * Math.PI) / 180
    const cos = Math.cos(rad)
    const sin = Math.sin(rad)
    const dx = x - cx
    const dy = y - cy
    return {
        x: cx + dx * cos - dy * sin,
        y: cy + dx * sin + dy * cos,
    }
}

function anchorPoint(item: BoardItem, handle: ResizeHandle) {
    const left = item.left
    const top = item.top
    const right = left + item.width
    const bottom = top + item.height
    const mx = (left + right) / 2
    const my = (top + bottom) / 2
    if (handle === "e") return { x: left, y: my }
    if (handle === "w") return { x: right, y: my }
    if (handle === "n") return { x: mx, y: bottom }
    if (handle === "s") return { x: mx, y: top }
    if (handle === "nw") return { x: right, y: bottom }
    if (handle === "ne") return { x: left, y: bottom }
    if (handle === "se") return { x: left, y: top }
    return { x: right, y: top }
}

function rotatedBounds(left: number, top: number, right: number, bottom: number, angle: number) {
    const cx = (left + right) / 2
    const cy = (top + bottom) / 2
    const rad = (angle * Math.PI) / 180
    const cos = Math.cos(rad)
    const sin = Math.sin(rad)
    const corners = [
        [left, top],
        [right, top],
        [right, bottom],
        [left, bottom],
    ]
    const xs = corners.map(([x, y]) => cx + (x - cx) * cos - (y - cy) * sin)
    const ys = corners.map(([x, y]) => cy + (x - cx) * sin + (y - cy) * cos)
    return {
        left: Math.min(...xs),
        top: Math.min(...ys),
        right: Math.max(...xs),
        bottom: Math.max(...ys),
    }
}

export function bounds(startX: number, startY: number, x: number, y: number) {
    return {
        left: Math.min(startX, x),
        top: Math.min(startY, y),
        width: Math.abs(x - startX),
        height: Math.abs(y - startY),
    }
}

function pointsOf(startX: number, startY: number, x: number, y: number): BoardPoint[] {
    return [
        { left: startX, top: startY },
        { left: x, top: y },
    ]
}

export function appendInk(points: BoardPoint[], samples: BoardPoint[], minDistance: number) {
    let next = points
    for (const sample of samples) {
        const last = next[next.length - 1]
        if (!last || Math.hypot(sample.left - last.left, sample.top - last.top) >= minDistance) {
            if (next === points) next = points.slice()
            next.push(sample)
        }
    }
    return next
}

export function finishInk(points: BoardPoint[], samples: BoardPoint[], minDistance: number) {
    let next = appendInk(points, samples, minDistance)
    const tip = samples[samples.length - 1]
    const last = next[next.length - 1]
    if (tip && last && Math.hypot(tip.left - last.left, tip.top - last.top) >= minDistance * 0.35) {
        if (next === points) next = points.slice()
        next.push(tip)
    }
    return simplifyInk(next, Math.min(0.8, minDistance * 0.55))
}

function simplifyInk(points: BoardPoint[], epsilon: number) {
    if (points.length <= 2) return points
    const keep = new Array<boolean>(points.length).fill(false)
    keep[0] = true
    keep[points.length - 1] = true
    const stack: [number, number][] = [[0, points.length - 1]]
    while (stack.length) {
        const [start, end] = stack.pop()!
        let maxDist = 0
        let index = -1
        for (let i = start + 1; i < end; i++) {
            const dist = segmentDistance(points[i], points[start], points[end])
            if (dist > maxDist) {
                maxDist = dist
                index = i
            }
        }
        if (index !== -1 && maxDist > epsilon) {
            keep[index] = true
            stack.push([start, index], [index, end])
        }
    }
    return points.filter((_, index) => keep[index])
}

function segmentDistance(point: BoardPoint, start: BoardPoint, end: BoardPoint) {
    const dx = end.left - start.left
    const dy = end.top - start.top
    const length = dx * dx + dy * dy
    if (length === 0) return Math.hypot(point.left - start.left, point.top - start.top)
    const t = Math.max(0, Math.min(1, ((point.left - start.left) * dx + (point.top - start.top) * dy) / length))
    return Math.hypot(point.left - (start.left + t * dx), point.top - (start.top + t * dy))
}

export function itemFromInk(options: {
    authorId: number
    points: BoardPoint[]
    color: string
}): BoardItem | null {
    if (!options.points.length) return null
    const xs = options.points.map((point) => point.left)
    const ys = options.points.map((point) => point.top)
    const left = Math.min(...xs)
    const top = Math.min(...ys)
    const width = Math.max(...xs) - left
    const height = Math.max(...ys) - top
    const dot = options.points.length === 1
    const pad = dot ? PEN_STROKE_WIDTH : 0
    return roundItem({
        object_id: crypto.randomUUID(),
        author_id: options.authorId,
        type: "LINE",
        left: left - pad,
        top: top - pad,
        width: width + pad * 2,
        height: height + pad * 2,
        angle: 0,
        points: options.points,
        color: options.color,
        stroke_width: PEN_STROKE_WIDTH,
    })
}

export function itemFromDrag(options: {
    tool: Exclude<BoardTool, "select" | "text" | "image" | "pen">
    authorId: number
    startX: number
    startY: number
    x: number
    y: number
    color: string
    fill: string
}): BoardItem | null {
    const box = bounds(options.startX, options.startY, options.x, options.y)
    const distance = Math.hypot(options.x - options.startX, options.y - options.startY)
    if (options.tool === "line" || options.tool === "arrow") {
        if (distance < 4) return null
    } else if (box.width < 4 || box.height < 4) {
        return null
    }

    const figure = options.tool === "line" ? undefined : options.tool.toUpperCase() as FigureType

    return roundItem({
        object_id: crypto.randomUUID(),
        author_id: options.authorId,
        type: options.tool === "line" ? "LINE" : "FIGURE",
        figure_type: figure,
        ...box,
        angle: 0,
        points: options.tool === "line" || options.tool === "arrow"
            ? pointsOf(options.startX, options.startY, options.x, options.y)
            : undefined,
        color: options.color,
        fill: options.tool === "line" || options.tool === "arrow" ? undefined : options.fill,
        stroke_width: 2,
    })
}

export function itemFromClick(options: {
    tool: "text" | "image"
    authorId: number
    x: number
    y: number
    color: string
    fill: string
    text: string
    imageUrl: string
}): BoardItem {
    if (options.tool === "text") {
        const text = options.text.trim() || "Текст"
        return roundItem({
            object_id: crypto.randomUUID(),
            author_id: options.authorId,
            type: "TEXT",
            left: options.x,
            top: options.y,
            width: Math.max(80, text.length * 10),
            height: 32,
            angle: 0,
            color: options.color,
            text,
            font_size: 20,
        })
    }

    return roundItem({
        object_id: crypto.randomUUID(),
        author_id: options.authorId,
        type: "IMAGE",
        left: options.x,
        top: options.y,
        width: 180,
        height: 120,
        angle: 0,
        color: options.color,
        fill: options.fill,
        image_url: options.imageUrl.trim(),
    })
}
