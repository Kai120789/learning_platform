export type Board = {
    id: number
    title: string
    tutorId: number
}

export type BoardResponse = {
    id: number
    title: string
    tutor_id: number
}

export type BoardSchema = {
    data: Board[] | null
    isLoading: boolean
    error?: string
}

export type BoardAction = "add" | "update" | "delete" | "snapshot"

export type BoardItemType = "LINE" | "FIGURE" | "TEXT" | "IMAGE"

export type FigureType = "RECT" | "ELLIPSE" | "TRIANGLE" | "ARROW"

export interface BoardPoint {
    left: number
    top: number
}

export interface BoardItem {
    object_id: string
    author_id: number
    type: BoardItemType
    left: number
    top: number
    width: number
    height: number
    z_index?: number
    angle: number
    figure_type?: FigureType
    points?: BoardPoint[]
    color?: string
    fill?: string
    stroke_width?: number
    text?: string
    font_size?: number
    image_url?: string
}

export interface BoardEnvelope {
    type: BoardAction
    board_id: number
    version: number
}

export interface SnapshotMessage {
    envelope: BoardEnvelope
    items: BoardItem[] | null
}

export interface AddMessage {
    envelope: BoardEnvelope
    item: BoardItem
}

export interface UpdateMessage {
    envelope: BoardEnvelope
    item: BoardItem
}

export interface DeleteMessage {
    envelope: BoardEnvelope
    object_ids: string[]
}

export type BoardClientMessage = AddMessage | UpdateMessage | DeleteMessage

export type BoardSocketMessage = SnapshotMessage | BoardClientMessage
