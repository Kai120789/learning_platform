import { FiBell, FiMenu } from "react-icons/fi";
import { FaRegFontAwesomeLogoFull, FaRegUserCircle } from "react-icons/fa";
import { DropdownMenuIcons } from "@/widgets/dropdownMenu";

type TopMenuProps = {
    onClick: () => void
}

export function TopMenu({ onClick }: TopMenuProps) {
    return (
        <div className="border-b-2 border-border bg-background">
            <div className="flex flex-row p-3 lg:p-4 justify-between items-center">
                <div className="flex flex-row items-center gap-3 lg:gap-6">
                    <FiMenu onClick={onClick} className="size-5 cursor-pointer" />
                    <div className='z-50'>
                        <FaRegFontAwesomeLogoFull className="h-[32px] w-[120px]" />
                    </div>
                </div>
                <div className="flex flex-row gap-3">
                    <FiBell className="size-5" />
                    <DropdownMenuIcons trigger={
                        <FaRegUserCircle className="size-5 cursor-pointer border-none" />
                    } />
                </div>
            </div>
        </div>
    )
}