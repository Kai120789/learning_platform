import { Route, Routes, type RouteProps } from 'react-router-dom';
import { routeConfig } from './routeConfig';
import { Suspense, type JSX } from 'react';
import { Loader } from 'lucide-react';
import { AuthProvider } from '@/app/providers';
import { NotificationList } from '@/features/notifications';
import { ToastContainer } from 'react-toastify';
import { useTheme } from '@teispace/next-themes/client';


export type AppRoutesProps = RouteProps & {
    authOnly?: boolean;
    layout: JSX.Element
};

export function AppRouter() {
    const { theme, } = useTheme();

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
        <>
            <NotificationList />
            <ToastContainer
                position="top-right"
                autoClose={5000}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
                theme={theme}
            />
            <Routes>
                <Route>
                    {Object.values(routeConfig).map((item) => render(item))}
                </Route>
            </Routes>
        </>

    )
}
