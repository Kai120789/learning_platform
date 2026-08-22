type NamedPerson = {
    name: string
    surname: string
    patronymic?: string
}

export function formatTutorName(person: NamedPerson, withPatronymic = false) {
    if (withPatronymic && person.patronymic) {
        return `${person.name} ${person.patronymic} ${person.surname}`
    }
    return `${person.name} ${person.surname}`
}

export function getTutorInitials(person: Pick<NamedPerson, "name" | "surname">) {
    return `${person.name[0] ?? ""}${person.surname[0] ?? ""}`.toUpperCase()
}

export function formatTutorSubject(subject: { title?: string; type?: string }) {
    const title = subject.title?.trim()
    const type = subject.type?.trim()
    if (!title) return ""
    if (!type) return title
    return `${title} · ${type}`
}
