import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css'

import App from './app/App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const router = createBrowserRouter([
	{ path: '/*', element: <App /> },
]);

const renderApp = async () => {
	const root = ReactDOM.createRoot(
		document.getElementById('root') as HTMLElement
	);
	root.render(
		<React.StrictMode>
			<RouterProvider router={router} />
		</React.StrictMode>
	);
}

renderApp()
