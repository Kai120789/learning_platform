import { Route, Routes, type RouteProps } from 'react-router-dom';
import { routeConfig } from './routeConfig';
import { Suspense, type JSX } from 'react';
import { Loader } from 'lucide-react';
import { AuthProvider } from '@/app/providers';
import { NotificationList } from '@/features/notifications';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


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
        <>
            <NotificationList />
            <ToastContainer
                position="top-right"
                autoClose={4500}
                hideProgressBar={false}
                newestOnTop
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable={false}
                pauseOnHover
                toastClassName="app-toast"
                progressClassName="app-toast-progress"
                closeButton
            />
            <Routes>
                <Route>
                    {Object.values(routeConfig).map((item) => render(item))}
                </Route>
            </Routes>
        </>

    )
}
