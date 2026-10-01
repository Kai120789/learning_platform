import { useCallback, useEffect, useRef, useState } from "react"
import { applyBoardMessage, buildBoardWsUrl, parseBoardMessage } from "../api/boardSocket"
import type { BoardClientMessage, BoardItem } from "../model/types"

export type BoardSocketStatus = "connecting" | "open" | "closed"

export function useBoardSocket(boardId: number) {
    const socketRef = useRef<WebSocket | null>(null)
    const [items, setItems] = useState<BoardItem[]>([])
    const [version, setVersion] = useState(0)
    const [status, setStatus] = useState<BoardSocketStatus>("connecting")
    const [attempt, setAttempt] = useState(0)

    useEffect(() => {
        if (!Number.isInteger(boardId) || boardId <= 0) {
            setStatus("closed")
            return
        }

        setItems([])
        setVersion(0)
        setStatus("connecting")
        const socket = new WebSocket(buildBoardWsUrl(boardId))
        socketRef.current = socket

        socket.onopen = () => setStatus("open")
        socket.onclose = () => setStatus("closed")
        socket.onmessage = (event) => {
            const data = parseBoardMessage(String(event.data))
            if (!data) return
            if (data.envelope.type === "snapshot") setVersion(data.envelope.version)
            else setVersion((value) => value + 1)
            setItems((prev) => applyBoardMessage(prev, data))
        }

        return () => {
            socket.close()
            socketRef.current = null
        }
    }, [attempt, boardId])

    const send = useCallback((message: BoardClientMessage) => {
        const socket = socketRef.current
        if (!socket || socket.readyState !== WebSocket.OPEN) return false
        socket.send(JSON.stringify(message))
        setItems((prev) => applyBoardMessage(prev, message))
        setVersion((value) => value + 1)
        return true
    }, [])

    const reconnect = useCallback(() => setAttempt((value) => value + 1), [])

    return { items, version, status, send, reconnect }
}
