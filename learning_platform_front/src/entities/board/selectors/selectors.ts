import type { StateSchema } from "@/app/providers/storeProvider"

export const getBoards = (state: StateSchema) => state.board.data
export const getBoardsLoading = (state: StateSchema) => state.board.isLoading
