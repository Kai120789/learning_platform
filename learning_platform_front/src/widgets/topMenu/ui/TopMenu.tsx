import { FiBell, FiMenu } from "react-icons/fi";
import { FaRegFontAwesomeLogoFull, FaRegUserCircle } from "react-icons/fa";
import { DropdownMenuIcons } from "@/widgets/dropdownMenu/ui/DropdownMenuIcons";

type TopMenuProps = {
    onClick: () => void
}

export function TopMenu({ onClick }: TopMenuProps) {
    return (
        <div className="border-b-2 border-border">
            <div className="flex flex-row p-[20px] justify-between items-center">
                <div className="flex flex-row items-center gap-10">
                    <FiMenu onClick={onClick} className="size-6 cursor-pointer" />
                    <div className='z-50'>
                        <FaRegFontAwesomeLogoFull className="h-[40px] w-[150px]" />
                    </div>
                </div>
                <div className="flex flex-row  gap-4">
                    <FiBell className="size-6" />
                    <DropdownMenuIcons trigger={
                        <FaRegUserCircle className="size-6 cursor-pointer border-none" />
                    } />
                </div>

            </div>
        </div>
    )
}