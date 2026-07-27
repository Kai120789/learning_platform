export type LessonStatus = "SCHEDULED" | "IN_PROCESS" | "COMPLETED" | "CANCELLED"

export type MediaType = "VIDEO" | "IMAGE" | "DOCUMENT" | "LINK"

export type PracticeStatus = "opened" | "completed" | "closed" | "pending"

export type MediaItemMock = {
    id: number
    lessonId: number
    s3Link: string
    s3Preview: string
    type: MediaType
    title: string
}

export type LessonMock = {
    id: number
    boardId?: number
    meetLink?: string
    startTime: string
    duration: number
    tutorId: number
    tutorName: string
    status: LessonStatus
    subjectTitle: string
    groupTitle: string
    userIds: number[]
    mediaItems: MediaItemMock[]
}

export type CourseMock = {
    id: number
    title: string
    description: string
    subjectTitle: string
    modulesCount: number
    enrolledCount: number
    progress: number
    tutorName: string
}

export type PracticeMock = {
    id: number
    title: string
    groupTitle: string
    subjectTitle: string
    type: "test" | "homework" | "video"
    status: PracticeStatus
    startTime: string
    endTime: string
    exercisesCount: number
    correctAnswersCount?: number
}

export type TutorMock = {
    id: number
    name: string
    surname: string
    city?: string
    about?: string
    subjects: string[]
    experienceYears: number
    rating: number
}

export const mockLessons: LessonMock[] = [
    {
        id: 1,
        boardId: 101,
        meetLink: "https://meet.example.com/math-a",
        startTime: "2026-07-24T10:00:00+03:00",
        duration: 60,
        tutorId: 1,
        tutorName: "Иван Петров",
        status: "SCHEDULED",
        subjectTitle: "Математика",
        groupTitle: "Группа А",
        userIds: [11, 12, 13],
        mediaItems: [
            {
                id: 1,
                lessonId: 1,
                s3Link: "#",
                s3Preview: "#",
                type: "VIDEO",
                title: "Квадратные уравнения",
            },
            {
                id: 2,
                lessonId: 1,
                s3Link: "#",
                s3Preview: "#",
                type: "DOCUMENT",
                title: "Конспект.pdf",
            },
        ],
    },
    {
        id: 2,
        boardId: 102,
        meetLink: "https://meet.example.com/physics-b",
        startTime: "2026-07-25T14:00:00+03:00",
        duration: 90,
        tutorId: 2,
        tutorName: "Анна Смирнова",
        status: "IN_PROCESS",
        subjectTitle: "Физика",
        groupTitle: "Группа Б",
        userIds: [14, 15],
        mediaItems: [
            {
                id: 3,
                lessonId: 2,
                s3Link: "#",
                s3Preview: "#",
                type: "IMAGE",
                title: "Схема опыта",
            },
        ],
    },
    {
        id: 3,
        meetLink: "https://meet.example.com/ege-math",
        startTime: "2026-07-20T09:00:00+03:00",
        duration: 120,
        tutorId: 1,
        tutorName: "Иван Петров",
        status: "COMPLETED",
        subjectTitle: "Математика ЕГЭ",
        groupTitle: "ЕГЭ-2026",
        userIds: [11, 16, 17, 18],
        mediaItems: [
            {
                id: 4,
                lessonId: 3,
                s3Link: "#",
                s3Preview: "#",
                type: "LINK",
                title: "Разбор варианта",
            },
        ],
    },
    {
        id: 4,
        startTime: "2026-07-18T16:00:00+03:00",
        duration: 60,
        tutorId: 3,
        tutorName: "Олег Козлов",
        status: "CANCELLED",
        subjectTitle: "Химия",
        groupTitle: "Группа В",
        userIds: [19],
        mediaItems: [],
    },
]

export const mockCourses: CourseMock[] = [
    {
        id: 1,
        title: "ЕГЭ Математика профиль",
        description: "Полный курс подготовки к профильному ЕГЭ",
        subjectTitle: "Математика",
        modulesCount: 12,
        enrolledCount: 24,
        progress: 45,
        tutorName: "Иван Петров",
    },
    {
        id: 2,
        title: "ОГЭ Физика",
        description: "Базовая подготовка к ОГЭ по физике",
        subjectTitle: "Физика",
        modulesCount: 8,
        enrolledCount: 16,
        progress: 20,
        tutorName: "Анна Смирнова",
    },
    {
        id: 3,
        title: "Повышение успеваемости: алгебра",
        description: "Закрепление школьной программы 8–9 классов",
        subjectTitle: "Математика",
        modulesCount: 6,
        enrolledCount: 10,
        progress: 70,
        tutorName: "Олег Козлов",
    },
]

export const mockPractices: PracticeMock[] = [
    {
        id: 1,
        title: "Тест: квадратные уравнения",
        groupTitle: "Группа А",
        subjectTitle: "Математика",
        type: "test",
        status: "opened",
        startTime: "2026-07-22T08:00:00+03:00",
        endTime: "2026-07-28T23:59:00+03:00",
        exercisesCount: 10,
    },
    {
        id: 2,
        title: "ДЗ: законы Ньютона",
        groupTitle: "Группа Б",
        subjectTitle: "Физика",
        type: "homework",
        status: "pending",
        startTime: "2026-07-20T08:00:00+03:00",
        endTime: "2026-07-25T23:59:00+03:00",
        exercisesCount: 5,
    },
    {
        id: 3,
        title: "Видео-разбор варианта",
        groupTitle: "ЕГЭ-2026",
        subjectTitle: "Математика ЕГЭ",
        type: "video",
        status: "completed",
        startTime: "2026-07-10T08:00:00+03:00",
        endTime: "2026-07-17T23:59:00+03:00",
        exercisesCount: 1,
        correctAnswersCount: 1,
    },
    {
        id: 4,
        title: "Тест: химические реакции",
        groupTitle: "Группа В",
        subjectTitle: "Химия",
        type: "test",
        status: "closed",
        startTime: "2026-07-01T08:00:00+03:00",
        endTime: "2026-07-07T23:59:00+03:00",
        exercisesCount: 15,
        correctAnswersCount: 12,
    },
]

export const mockTutors: TutorMock[] = [
    {
        id: 1,
        name: "Иван",
        surname: "Петров",
        city: "Москва",
        about: "Преподаватель математики, подготовка к ЕГЭ и ОГЭ",
        subjects: ["Математика", "Информатика"],
        experienceYears: 8,
        rating: 4.9,
    },
    {
        id: 2,
        name: "Анна",
        surname: "Смирнова",
        city: "Санкт-Петербург",
        about: "Физика для школьников и абитуриентов",
        subjects: ["Физика"],
        experienceYears: 5,
        rating: 4.7,
    },
    {
        id: 3,
        name: "Олег",
        surname: "Козлов",
        city: "Казань",
        about: "Химия и повышение успеваемости",
        subjects: ["Химия", "Биология"],
        experienceYears: 6,
        rating: 4.8,
    },
]

export const mockGroupsForSelect = [
    { id: 1, title: "Группа А" },
    { id: 2, title: "Группа Б" },
    { id: 3, title: "ЕГЭ-2026" },
    { id: 4, title: "Группа В" },
]

export type CandidateUserMock = {
    id: number
    name: string
    surname: string
    patronymic?: string
    tgUsername?: string
}

export const mockGroupCandidates: CandidateUserMock[] = [
    { id: 101, name: "Иван", surname: "Петров", tgUsername: "@ivan_petrov" },
    { id: 102, name: "Мария", surname: "Смирнова", tgUsername: "@m_smirnova" },
    { id: 103, name: "Алексей", surname: "Козлов", tgUsername: "@alex_kozlov" },
    { id: 104, name: "Екатерина", surname: "Волкова", tgUsername: "@kate_volkova" },
    { id: 105, name: "Дмитрий", surname: "Соколов", tgUsername: "@d_sokolov" },
    { id: 106, name: "Анна", surname: "Морозова", tgUsername: "@anna_moroz" },
    { id: 107, name: "Никита", surname: "Лебедев", tgUsername: "@n_lebedev" },
    { id: 108, name: "Ольга", surname: "Новикова", tgUsername: "@olga_nov" },
]

export type ScheduleSlotMock = {
    id: number
    weekday: 0 | 1 | 2 | 3 | 4 | 5 | 6
    date: string
    start: string
    end: string
    subjectTitle: string
    groupTitle: string
    status: LessonStatus
}

export const mockWeekSlots: ScheduleSlotMock[] = [
    {
        id: 1,
        weekday: 0,
        date: "2026-07-20",
        start: "09:00",
        end: "10:00",
        subjectTitle: "Математика",
        groupTitle: "Группа А",
        status: "COMPLETED",
    },
    {
        id: 2,
        weekday: 0,
        date: "2026-07-20",
        start: "14:00",
        end: "15:00",
        subjectTitle: "Информатика",
        groupTitle: "Группа А",
        status: "COMPLETED",
    },
    {
        id: 3,
        weekday: 1,
        date: "2026-07-21",
        start: "11:00",
        end: "12:30",
        subjectTitle: "Физика",
        groupTitle: "Группа Б",
        status: "COMPLETED",
    },
    {
        id: 4,
        weekday: 1,
        date: "2026-07-21",
        start: "16:00",
        end: "17:00",
        subjectTitle: "Физика",
        groupTitle: "Группа Б",
        status: "SCHEDULED",
    },
    {
        id: 5,
        weekday: 2,
        date: "2026-07-22",
        start: "09:00",
        end: "10:30",
        subjectTitle: "Математика ЕГЭ",
        groupTitle: "ЕГЭ-2026",
        status: "SCHEDULED",
    },
    {
        id: 6,
        weekday: 2,
        date: "2026-07-22",
        start: "12:00",
        end: "13:00",
        subjectTitle: "Математика ЕГЭ",
        groupTitle: "ЕГЭ-2026",
        status: "SCHEDULED",
    },
    {
        id: 7,
        weekday: 2,
        date: "2026-07-22",
        start: "18:00",
        end: "19:00",
        subjectTitle: "Разбор ДЗ",
        groupTitle: "ЕГЭ-2026",
        status: "SCHEDULED",
    },
    {
        id: 8,
        weekday: 3,
        date: "2026-07-23",
        start: "10:00",
        end: "11:00",
        subjectTitle: "Химия",
        groupTitle: "Группа В",
        status: "SCHEDULED",
    },
    {
        id: 9,
        weekday: 3,
        date: "2026-07-23",
        start: "16:00",
        end: "17:00",
        subjectTitle: "Химия",
        groupTitle: "Группа В",
        status: "SCHEDULED",
    },
    {
        id: 10,
        weekday: 4,
        date: "2026-07-24",
        start: "10:00",
        end: "11:00",
        subjectTitle: "Математика",
        groupTitle: "Группа А",
        status: "SCHEDULED",
    },
    {
        id: 11,
        weekday: 4,
        date: "2026-07-24",
        start: "13:00",
        end: "14:30",
        subjectTitle: "Математика",
        groupTitle: "Группа А",
        status: "IN_PROCESS",
    },
    {
        id: 12,
        weekday: 5,
        date: "2026-07-25",
        start: "12:00",
        end: "13:30",
        subjectTitle: "Физика",
        groupTitle: "Группа Б",
        status: "SCHEDULED",
    },
    {
        id: 13,
        weekday: 6,
        date: "2026-07-26",
        start: "11:00",
        end: "12:00",
        subjectTitle: "Консультация",
        groupTitle: "ЕГЭ-2026",
        status: "SCHEDULED",
    },
]


export const mockHomeGroups = [
    {
        id: 1,
        title: "Группа А",
        description: "Математика, базовый уровень",
        subject: {
            id: 1,
            code: "math",
            title: "Математика",
            type: "Повышение успеваемости",
        },
        users: [
            { id: 11, name: "Алексей", surname: "Иванов" },
            { id: 12, name: "Мария", surname: "Кузнецова" },
        ],
        tutorId: 1,
        tgGroupLink: "https://t.me/group_a",
    },
    {
        id: 2,
        title: "Группа Б",
        description: "Физика ОГЭ",
        subject: {
            id: 2,
            code: "physics",
            title: "Физика",
            type: "ОГЭ",
        },
        users: [
            { id: 14, name: "Пётр", surname: "Сидоров" },
        ],
        tutorId: 2,
        tgGroupLink: "https://t.me/group_b",
    },
    {
        id: 3,
        title: "ЕГЭ-2026",
        description: "Профильная математика",
        subject: {
            id: 3,
            code: "math_ege",
            title: "Математика",
            type: "ЕГЭ",
        },
        users: [
            { id: 16, name: "Ольга", surname: "Новикова" },
            { id: 17, name: "Игорь", surname: "Орлов" },
        ],
        tutorId: 1,
        tgGroupLink: "https://t.me/ege_2026",
    },
    {
        id: 4,
        title: "Группа В",
        description: "Химия",
        subject: {
            id: 4,
            code: "chem",
            title: "Химия",
            type: "Повышение успеваемости",
        },
        users: [
            { id: 19, name: "Дарья", surname: "Морозова" },
        ],
        tutorId: 3,
    },
] as const

export type MaterialFileMock = MediaItemMock & {
    subjectTitle: string
    groupTitle: string
    updatedAt: string
    sizeLabel: string
}

export type MaterialFolderMock = {
    id: string
    name: string
    kind: "subject" | "group" | "lesson"
    parentId: string | null
    itemCount: number
    lessonId?: number
}

export const mockMaterials: MaterialFileMock[] = [
    ...mockLessons.flatMap((lesson) =>
        lesson.mediaItems.map((item) => ({
            ...item,
            subjectTitle: lesson.subjectTitle,
            groupTitle: lesson.groupTitle,
            updatedAt: lesson.startTime,
            sizeLabel: item.type === "VIDEO" ? "128 MB" : item.type === "IMAGE" ? "2.4 MB" : "840 KB",
        }))
    ),
    {
        id: 10,
        lessonId: 1,
        s3Link: "#",
        s3Preview: "#",
        type: "DOCUMENT",
        title: "Домашнее задание.docx",
        subjectTitle: "Математика",
        groupTitle: "Группа А",
        updatedAt: "2026-07-22T12:00:00+03:00",
        sizeLabel: "56 KB",
    },
    {
        id: 11,
        lessonId: 1,
        s3Link: "#",
        s3Preview: "#",
        type: "IMAGE",
        title: "Формулы.png",
        subjectTitle: "Математика",
        groupTitle: "Группа А",
        updatedAt: "2026-07-21T09:30:00+03:00",
        sizeLabel: "1.1 MB",
    },
    {
        id: 12,
        lessonId: 2,
        s3Link: "#",
        s3Preview: "#",
        type: "VIDEO",
        title: "Демонстрация опыта.mp4",
        subjectTitle: "Физика",
        groupTitle: "Группа Б",
        updatedAt: "2026-07-24T18:00:00+03:00",
        sizeLabel: "210 MB",
    },
    {
        id: 13,
        lessonId: 3,
        s3Link: "#",
        s3Preview: "#",
        type: "DOCUMENT",
        title: "Вариант ЕГЭ.pdf",
        subjectTitle: "Математика ЕГЭ",
        groupTitle: "ЕГЭ-2026",
        updatedAt: "2026-07-19T11:00:00+03:00",
        sizeLabel: "3.2 MB",
    },
]

export function buildMaterialFolders(files: MaterialFileMock[]): MaterialFolderMock[] {
    const folders: MaterialFolderMock[] = []
    const subjects = new Map<string, MaterialFileMock[]>()

    for (const file of files) {
        const list = subjects.get(file.subjectTitle) ?? []
        list.push(file)
        subjects.set(file.subjectTitle, list)
    }

    for (const [subjectTitle, subjectFiles] of subjects) {
        const subjectId = `subject:${subjectTitle}`
        folders.push({
            id: subjectId,
            name: subjectTitle,
            kind: "subject",
            parentId: null,
            itemCount: 0,
        })

        const groups = new Map<string, MaterialFileMock[]>()
        for (const file of subjectFiles) {
            const list = groups.get(file.groupTitle) ?? []
            list.push(file)
            groups.set(file.groupTitle, list)
        }

        for (const [groupTitle, groupFiles] of groups) {
            const groupId = `group:${subjectTitle}:${groupTitle}`
            folders.push({
                id: groupId,
                name: groupTitle,
                kind: "group",
                parentId: subjectId,
                itemCount: 0,
            })

            const lessons = new Map<number, MaterialFileMock[]>()
            for (const file of groupFiles) {
                const list = lessons.get(file.lessonId) ?? []
                list.push(file)
                lessons.set(file.lessonId, list)
            }

            for (const [lessonId, lessonFiles] of lessons) {
                folders.push({
                    id: `lesson:${subjectTitle}:${groupTitle}:${lessonId}`,
                    name: String(lessonId),
                    kind: "lesson",
                    parentId: groupId,
                    itemCount: lessonFiles.length,
                    lessonId,
                })
            }

            const groupFolder = folders.find((f) => f.id === groupId)
            if (groupFolder) {
                groupFolder.itemCount = lessons.size
            }
        }

        const subjectFolder = folders.find((f) => f.id === subjectId)
        if (subjectFolder) {
            subjectFolder.itemCount = groups.size
        }
    }

    return folders
}
