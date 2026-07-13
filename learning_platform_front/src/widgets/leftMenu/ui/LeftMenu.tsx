import { LeftMenuItemsType } from '@/shared/types/leftMenuItems'
import Logo from '@/shared/assets/Logo-Test.png'
import { LeftMenuItem } from './LeftMenuItem'
import { CgClose } from "react-icons/cg";
import { Switch } from '@/shared/ui/Switch';

type LeftMenuProps = {
    isOpen: boolean
    onClick: () => void
}

export function LeftMenu({ isOpen, onClick }: LeftMenuProps) {
    return (
        <div className={`fixed top-0 left-0 z-50 flex h-screen w-[300px] flex-col items-center 
            border-r-2 border-[#D9D9D9] bg-muted p-5 pt-0 rounded-tr-2xl rounded-br-2xl
            transition-transform duration-300 ease-in-out
            ${isOpen ? "translate-x-0" : "-translate-x-full"}`}>
            <div className='flex justify-between w-full mb-[40px]'>
                <div className='pt-4 z-50'>
                    <img className="h-[40px] w-[100px]" src={Logo} />
                </div>
                <CgClose onClick={onClick} className='size-6 mt-[20px] cursor-pointer' />
            </div>

            <div className="flex items-start w-full h-full border-b-2 border-[#D9D9D9]">
                <nav>
                    {LeftMenuItemsType().map((item) => (
                        <LeftMenuItem onClick={onClick} key={item.path} item={item} />
                    ))}
                </nav>
            </div>

            <div className="flex flex-row items-center gap-5 pb-5 pt-5 text-lg text-primary font-medium">
                Темная тема
                <Switch size="default" />
            </div>
        </div>
    )
}
