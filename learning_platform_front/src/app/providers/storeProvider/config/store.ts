import { configureStore, type Reducer, type ReducersMapObject } from "@reduxjs/toolkit";
import type { StateSchema } from "./StateSchema";
import { notificationReducer } from "@/features/notifications";
import { createReducerManager } from "./reducerManager";
import { $api } from "./api";
import { userReducer } from "@/entities/user";
import { groupReducer } from "@/entities/group";
import { subjectReducer } from "@/entities/subject";
import { lessonReducer } from "@/entities/lesson";
import { materialReducer } from "@/entities/material";
import { tutorReducer } from "@/entities/tutor";
import { interceptor } from "@/shared/api/interceptor.ts";
import { resetStore } from "@/shared/lib/resetStore";

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
        lesson: lessonReducer,
        material: materialReducer,
        tutor: tutorReducer,
    };

    const reducerManager = createReducerManager(rootReducer);

    const rootReducerWithReset: Reducer<StateSchema> = (state, action) => (
        reducerManager.reduce(
            resetStore.match(action) ? undefined as unknown as StateSchema : state as StateSchema,
            action,
        )
    );

    const extraArg = {
        api: $api,
    };

    const store = configureStore({

        reducer: rootReducerWithReset,
        devTools: true,
        preloadedState: initialState,
        middleware: (getDefaultMiddleware) => getDefaultMiddleware({
            thunk: {
                extraArgument: extraArg,
            },
        }),
    });

    store.reducerManager = reducerManager;

    interceptor($api, store)

    return store;
}

export type AppDispatch = ReturnType<typeof createReduxStore>['dispatch']
export type AppStore = ReturnType<typeof createReduxStore>
