import { Button } from "@/shared/ui/Button";

type ConfirmButtonProps = {
    confirmText?: string
    onClickConfirm: () => void
    cancelText?: string
    onClickCancel: () => void
}

export function ConfirmButton({
    confirmText = "Сохранить изменения",
    onClickConfirm,
    cancelText = "Сбросить",
    onClickCancel,
}: ConfirmButtonProps) {
    return (
        <div className="flex flex-col lg:flex-row gap-2 justify-end pt-2">
            <Button
                onClick={onClickConfirm}
                variant="default"
                size="lg"
            >
                {confirmText}
            </Button>
            <Button
                onClick={onClickCancel}
                variant="outline"
                size="lg"
            >
                {cancelText}
            </Button>
        </div>
    )
}