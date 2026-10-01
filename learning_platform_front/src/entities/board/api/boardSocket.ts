import type {
    AddMessage,
    BoardClientMessage,
    BoardItem,
    BoardSocketMessage,
    DeleteMessage,
    SnapshotMessage,
    UpdateMessage,
} from "../model/types"

export function buildBoardWsUrl(boardId: number) {
    const httpBase = import.meta.env.VITE_SERVER_ENDPOINT.replace(/\/$/, "")
    const wsBase = httpBase.replace(/^http/, "ws")
    return `${wsBase}/api/board/${boardId}/ws`
}

export function applyBoardMessage(items: BoardItem[], data: BoardSocketMessage): BoardItem[] {
    switch (data.envelope.type) {
        case "snapshot":
            return (data as SnapshotMessage).items ?? []
        case "add": {
            const message = data as AddMessage
            if (items.some((item) => item.object_id === message.item.object_id)) return items
            return [...items, message.item]
        }
        case "update": {
            const message = data as UpdateMessage
            return items.map((item) => item.object_id === message.item.object_id ? message.item : item)
        }
        case "delete": {
            const message = data as DeleteMessage
            return items.filter((item) => !message.object_ids.includes(item.object_id))
        }
        default:
            return items
    }
}

export function parseBoardMessage(raw: string): BoardSocketMessage | null {
    try {
        const data = JSON.parse(raw) as BoardSocketMessage
        if (!data?.envelope?.type) return null
        return data
    } catch {
        return null
    }
}

export type { BoardClientMessage }
