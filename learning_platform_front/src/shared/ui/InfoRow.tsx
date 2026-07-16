type InfoRowProps = {
    label: string
    value: string
}

export function InfoRow({ label, value }: InfoRowProps) {
    return (
        <div className="flex items-center justify-between border-b pb-2 last:border-0 last:pb-0">
            <span className="text-sm text-muted-foreground">
                {label}
            </span>

            <span className="font-medium text-right">
                {value}
            </span>
        </div>
    )
}