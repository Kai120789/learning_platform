export function formatPrice(price: number) {
    return new Intl.NumberFormat("ru-RU").format(price)
}

export function formatRating(rating: number) {
    return Number.isInteger(rating) ? String(rating) : rating.toFixed(1)
}
