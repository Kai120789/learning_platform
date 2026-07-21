import { configureStore, type Reducer, type ReducersMapObject } from "@reduxjs/toolkit";
import type { StateSchema } from "./StateSchema";
import { notificationReducer } from "@/features/notifications";
import { createReducerManager } from "./reducerManager";
import { $api } from "./api";
import { userReducer } from "@/entities/user/slice/userSlice";
import { groupReducer } from "@/entities/group/slice/groupSlice";
import { subjectReducer } from "@/entities/subject/slice/slice";

export function createReduxStore(
    initialState: StateSchema,
    asyncReducers?: ReducersMapObject<StateSchema>,
) {
    const rootReducer: ReducersMapObject<StateSchema> = {
        ...asyncReducers,
        notifications: notificationReducer,
        user: userReducer,
        group: groupReducer,
        subject: subjectReducer,
    };

    const reducerManager = createReducerManager(rootReducer);

    const extraArg = {
        api: $api,
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
