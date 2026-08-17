
declare module '*.svg' {
	import * as React from 'react';
	const SVG: React.FC<React.SVGProps<SVGSVGElement>>;
	export default SVG;
}

declare module '*.svg?url' {
	const src: string;
	export default src;
}

interface ImportMetaEnv {
	readonly VITE_SERVER_ENDPOINT: string;
	readonly VITE_S3_URL: string;
}

interface ImportMeta {
	readonly env: ImportMetaEnv;
}
