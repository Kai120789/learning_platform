import { FileText, Image, Link2, Video } from "lucide-react"
import type { MediaType } from "@/shared/mocks"

export const mediaIcons: Record<MediaType, typeof Video> = {
    VIDEO: Video,
    IMAGE: Image,
    DOCUMENT: FileText,
    LINK: Link2,
}
