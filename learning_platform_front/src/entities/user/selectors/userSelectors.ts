import type { StateSchema } from "@/app/providers/storeProvider";

export const getUserFullData = (state: StateSchema) => state.user.data
export const getIsAuth = (state: StateSchema) => state.user.isAuth
export const getUserRole = (state: StateSchema) => state.user.data?.user.role
export const getIsInitialized = (state: StateSchema) => state.user.isInitialized
