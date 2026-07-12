import ReactDOM from 'react-dom/client';
import './index.css'

import App from './app/App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { StoreProvider } from './app/providers/storeProvider/index.ts';

const router = createBrowserRouter([
	{ path: '/*', element: <App /> },
]);

const renderApp = async () => {
	const root = ReactDOM.createRoot(
		document.getElementById('root') as HTMLElement
	);
	root.render(
		<StoreProvider>
			<RouterProvider router={router} />
		</StoreProvider>
	);
}

renderApp()
