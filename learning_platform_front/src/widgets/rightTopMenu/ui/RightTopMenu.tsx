import { FiBell, FiMenu } from "react-icons/fi";
import { FaRegUserCircle } from "react-icons/fa";
import { DropdownMenuIcons } from "@/widgets/dropdownMenu/ui/DropdownMenuIcons";


type RightTopMenuProps = {
    onClick: () => void
}

export function RightTopMenu({ onClick }: RightTopMenuProps) {
    return (
        <div className="border-b-2 border-[#D9D9D9]">
            <div className="flex flex-row p-[20px] justify-between">
                <FiMenu onClick={onClick} className="size-6 cursor-pointer" />
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