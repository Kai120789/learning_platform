import { IoMoonOutline } from "react-icons/io5";
import { FiSun } from "react-icons/fi";
import { useTheme } from '@teispace/next-themes/client';


import { Switch } from "@/shared/ui/Switch";

export function ThemeSwitch() {
    const { theme, setTheme } = useTheme();

    const toggleTheme = () => {
        setTheme(theme === 'dark' ? 'light' : 'dark')
    };

    return (
        <div className="flex gap-3">
            <FiSun />
            <Switch
                className="border border-ring"
                checked={theme == 'dark'}
                onCheckedChange={toggleTheme}
            />
            <IoMoonOutline size={17} />
        </div>
    )
}
