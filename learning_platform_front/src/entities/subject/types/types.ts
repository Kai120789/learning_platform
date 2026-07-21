export interface SubjectSchema {
    data: SubjectData[] | null;
    isLoading: boolean;
    error?: string;
}

export type SubjectData = {
    id: number
    code: string
    title: string
    type: SubjectTypeEnum
}

export enum SubjectTypeEnum {
    EGE = "ЕГЭ",
    OGE = "ОГЭ",
    IMPROVE = "Повышение успеваемости"
}
