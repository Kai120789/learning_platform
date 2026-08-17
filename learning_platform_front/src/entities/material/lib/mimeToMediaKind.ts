import type { MaterialMediaKind } from "../types/types"

export function mimeToMediaKind(mimeType: string): MaterialMediaKind {
    const mime = mimeType.toLowerCase()

    if (mime.startsWith("video/")) return "VIDEO"
    if (mime.startsWith("image/")) return "IMAGE"
    if (mime.startsWith("audio/")) return "DOCUMENT"
    if (
        mime.includes("pdf")
        || mime.includes("document")
        || mime.includes("msword")
        || mime.includes("sheet")
        || mime.includes("presentation")
        || mime.startsWith("text/")
    ) {
        return "DOCUMENT"
    }

    return "DOCUMENT"
}
