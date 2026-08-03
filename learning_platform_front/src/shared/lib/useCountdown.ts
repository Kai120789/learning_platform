import { useEffect, useState } from "react"

function pad(value: number) {
    return String(value).padStart(2, "0")
}

export function formatCountdown(ms: number) {
    const totalSeconds = Math.max(0, Math.floor(ms / 1000))
    const days = Math.floor(totalSeconds / 86400)
    const hours = Math.floor((totalSeconds % 86400) / 3600)
    const minutes = Math.floor((totalSeconds % 3600) / 60)
    const seconds = totalSeconds % 60

    if (days > 0) {
        return `${days}д ${pad(hours)}:${pad(minutes)}:${pad(seconds)}`
    }
    return `${pad(hours)}:${pad(minutes)}:${pad(seconds)}`
}

export function useCountdown(targetIso: string | null | undefined) {
    const [now, setNow] = useState(() => Date.now())

    useEffect(() => {
        const id = window.setInterval(() => setNow(Date.now()), 1000)
        return () => window.clearInterval(id)
    }, [])

    if (!targetIso) {
        return { remainingMs: 0, isPast: true, label: "00:00:00" }
    }

    const target = new Date(targetIso).getTime()
    if (Number.isNaN(target)) {
        return { remainingMs: 0, isPast: true, label: "00:00:00" }
    }

    const remainingMs = target - now
    return {
        remainingMs,
        isPast: remainingMs <= 0,
        label: formatCountdown(remainingMs),
    }
}
