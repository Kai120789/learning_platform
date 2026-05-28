import React from 'react'
import { Route, Routes, type RouteProps } from 'react-router-dom';
import { routeConfig } from './routeConfig';


export type AppRoutesProps = RouteProps & {
    authOnly?: boolean;
};

const AppRouter = () => {
    const render = (route: AppRoutesProps) => {
        return (
            <Route
                key={route.path}
                path={route.path}
                element={
                    route.element
                }
            />
        )
    }

    return (
        <Routes>
            <Route>
                {Object.values(routeConfig).map((item) => render(item))}
            </Route>
        </Routes>
    )
}

export default AppRouter