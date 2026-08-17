import { FileText, Image, Link2, Video } from "lucide-react"
import type { MaterialMediaKind } from "@/entities/material"

export const mediaIcons: Record<MaterialMediaKind, typeof Video> = {
    VIDEO: Video,
    IMAGE: Image,
    DOCUMENT: FileText,
    LINK: Link2,
}
