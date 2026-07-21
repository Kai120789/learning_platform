import type { GroupSchema } from '@/entities/group/types/types'
import type { SubjectSchema } from '@/entities/subject/types/types'
import type { UserSchema } from '@/entities/user/types/types'
import type { NotificationSchema } from '@/features/notifications'
import type { AnyAction, EnhancedStore, Reducer, ReducersMapObject } from '@reduxjs/toolkit'
import type { AxiosInstance } from 'axios'
import type { NavigateOptions, To } from 'react-router'


export interface StateSchema {
    notifications: NotificationSchema
    user: UserSchema
    group: GroupSchema
    subject: SubjectSchema
}

export type StateSchemaKey = keyof StateSchema
export type MountedReducers = OptionalRecord<StateSchemaKey, boolean>

export interface ReducerManager {
    getReducerMap: () => ReducersMapObject<StateSchema>
    reduce: (state: StateSchema, action: AnyAction) => StateSchema
    add: (key: StateSchemaKey, reducer: Reducer) => void
    remove: (key: StateSchemaKey) => void
    getMountedReducers: () => MountedReducers
}

export interface ReduxStoreWithManager extends EnhancedStore<StateSchema> {
    reducerManager: ReducerManager
}

export interface ThunkExtraArg {
    api: AxiosInstance
    navigate?: (to: To, options?: NavigateOptions) => void
}

export interface ThunkConfig<T> {
    rejectValue: T
    extra: ThunkExtraArg
    state: StateSchema
}

export type OptionalRecord<K extends PropertyKey, T> = {
    [P in K]?: T;
};