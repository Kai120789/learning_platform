import { Route, Routes, type RouteProps } from 'react-router-dom';
import { routeConfig } from './routeConfig';
import { AuthProvider } from '../providers/AuthProvider';
import { Suspense, type JSX } from 'react';
import { Loader } from 'lucide-react';


export type AppRoutesProps = RouteProps & {
    authOnly?: boolean;
    layout: JSX.Element
};

export function AppRouter() {
    const render = (route: AppRoutesProps) => {
        const element = (
            <Suspense fallback={<Loader />}>
                <div>
                    {route.element}
                </div>
            </Suspense>
        )

        return (
            <Route element={route.layout}>
                <Route
                    key={route.path}
                    path={route.path}
                    element={
                        route.authOnly
                            ? <AuthProvider>{element}</AuthProvider>
                            : route.element
                    }
                />
            </Route>
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
