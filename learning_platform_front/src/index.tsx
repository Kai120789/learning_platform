import ReactDOM from 'react-dom/client';
import './index.css'

import App from './app/App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { StoreProvider } from './app/providers/storeProvider/index.ts';
import { ThemeProvider } from '@teispace/next-themes/client';
import "./app/providers/i18n/config";
import { TooltipProvider } from './shared/ui/Tooltip.tsx';

const router = createBrowserRouter([
	{ path: '/*', element: <App /> },
]);

const renderApp = async () => {
	const root = ReactDOM.createRoot(
		document.getElementById('root') as HTMLElement
	);
	root.render(
		<StoreProvider>
			<ThemeProvider attribute="class">
				<TooltipProvider>
					<RouterProvider router={router} />
				</TooltipProvider>
			</ThemeProvider>
		</StoreProvider>
	);
}

renderApp()
