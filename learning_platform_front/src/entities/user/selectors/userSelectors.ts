import type { StateSchema } from "@/app/providers/storeProvider";

export const getUserFullData = (state: StateSchema) => state.user.data
export const getIsAuth = (state: StateSchema) => state.user.isAuth
