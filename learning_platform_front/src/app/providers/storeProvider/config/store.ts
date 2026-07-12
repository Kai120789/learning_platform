import { configureStore, type Reducer, type ReducersMapObject } from "@reduxjs/toolkit";
import type { StateSchema } from "./StateSchema";
import { notificationReducer } from "@/features/notifications";
import { createReducerManager } from "./reducerManager";
import axios from "axios";

export function createReduxStore(
    initialState: StateSchema,
    asyncReducers?: ReducersMapObject<StateSchema>,
) {
    const rootReducer: ReducersMapObject<StateSchema> = {
        ...asyncReducers,
        notifications: notificationReducer,
    };

    const reducerManager = createReducerManager(rootReducer);

    const extraArg = {
        api: axios.create,
    };

    const store = configureStore({

        reducer: reducerManager.reduce as Reducer<StateSchema>,
        devTools: true,
        preloadedState: initialState,
        middleware: (getDefaultMiddleware) => getDefaultMiddleware({
            thunk: {
                extraArgument: extraArg,
            },
        }),
    });

    store.reducerManager = reducerManager;

    return store;
}

export type AppDispatch = ReturnType<typeof createReduxStore>['dispatch']
