import { Button } from "@/shared/ui/Button";
import { useTranslation } from "react-i18next";

type ConfirmButtonProps = {
    confirmText?: string
    onClickConfirm: () => void
    cancelText?: string
    onClickCancel: () => void
}

export function ConfirmButton({
    confirmText,
    onClickConfirm,
    cancelText,
    onClickCancel,
}: ConfirmButtonProps) {
    const { t } = useTranslation()

    return (
        <div className="flex flex-col sm:flex-row gap-2 justify-end pt-1">
            <Button
                onClick={onClickConfirm}
                variant="default"
                size="sm"
            >
                {confirmText ?? t("common.save")}
            </Button>
            <Button
                onClick={onClickCancel}
                variant="outline"
                size="sm"
            >
                {cancelText ?? t("common.cancel")}
            </Button>
        </div>
    )
}
