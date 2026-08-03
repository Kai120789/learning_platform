import { useTranslation } from "react-i18next"
import { Field, FieldGroup, FieldLabel } from "@/shared/ui/Field"
import { Input } from "@/shared/ui/Input"
import { NativeSelect } from "@/shared/ui/NativeSelect"
import type { BoardOption } from "../lib/constants"

type LessonDetailsFieldsProps = {
    meetLink: string
    onMeetLinkChange: (value: string) => void
    startTime: string
    onStartTimeChange: (value: string) => void
    duration: number
    onDurationChange: (value: number) => void
    boardId?: number | ""
    onBoardIdChange?: (value: number | "") => void
    boards?: BoardOption[]
    showBoard?: boolean
}

export function LessonDetailsFields({
    meetLink,
    onMeetLinkChange,
    startTime,
    onStartTimeChange,
    duration,
    onDurationChange,
    boardId = "",
    onBoardIdChange,
    boards = [],
    showBoard = false,
}: LessonDetailsFieldsProps) {
    const { t } = useTranslation()

    return (
        <FieldGroup className="w-full gap-3">
            <div className="grid gap-3 sm:grid-cols-2">
                <Field className="w-full sm:col-span-2">
                    <FieldLabel>{t("createLesson.meetLink")}</FieldLabel>
                    <Input
                        className="w-full"
                        type="url"
                        value={meetLink}
                        onChange={(e) => onMeetLinkChange(e.target.value)}
                        placeholder="https://"
                    />
                </Field>
                <Field className="w-full">
                    <FieldLabel>{t("createLesson.startTime")}</FieldLabel>
                    <Input
                        className="w-full"
                        required
                        type="datetime-local"
                        value={startTime}
                        onChange={(e) => onStartTimeChange(e.target.value)}
                    />
                </Field>
                <Field className="w-full">
                    <FieldLabel>{t("createLesson.duration")}</FieldLabel>
                    <Input
                        className="w-full"
                        required
                        type="number"
                        min={15}
                        step={15}
                        value={duration}
                        onChange={(e) => onDurationChange(Number(e.target.value))}
                    />
                </Field>
                {showBoard && onBoardIdChange && (
                    <Field className="w-full sm:col-span-2">
                        <FieldLabel>{t("createLesson.board")}</FieldLabel>
                        <NativeSelect
                            value={boardId}
                            disabled={boards.length === 0}
                            onChange={(e) => {
                                const value = e.target.value
                                onBoardIdChange(value === "" ? "" : Number(value))
                            }}
                        >
                            <option value="">
                                {boards.length === 0
                                    ? t("createLesson.noBoards")
                                    : t("createLesson.boardPlaceholder")}
                            </option>
                            {boards.map((board) => (
                                <option key={board.id} value={board.id}>
                                    {board.title}
                                </option>
                            ))}
                        </NativeSelect>
                    </Field>
                )}
            </div>
        </FieldGroup>
    )
}
