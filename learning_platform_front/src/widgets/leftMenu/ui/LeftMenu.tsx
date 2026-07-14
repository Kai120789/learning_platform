import { LeftMenuItemsType } from '@/shared/types/leftMenuItems'
import { LeftMenuItem } from './LeftMenuItem'
import { CgClose } from "react-icons/cg";
import { FaRegFontAwesomeLogoFull } from 'react-icons/fa';
import { ThemeSwitch } from '@/widgets/themeSwitch';
import { useTranslation } from 'react-i18next';

type LeftMenuProps = {
    isOpen: boolean
    onClick: () => void
}

export function LeftMenu({ isOpen, onClick }: LeftMenuProps) {
    const { t } = useTranslation()

    return (
        <div className={`fixed top-0 left-0 z-50 flex h-screen w-[300px] flex-col items-center 
            border-r-2 border-border bg-muted p-5 pt-0 rounded-tr-2xl rounded-br-2xl
            transition-transform duration-300 ease-in-out
            ${isOpen ? "translate-x-0" : "-translate-x-full"}`}>
            <div className='flex justify-between w-full mb-[40px]'>
                <div className='pt-4 z-50'>
                    <FaRegFontAwesomeLogoFull className="h-[40px] w-[150px]" />
                </div>
                <CgClose onClick={onClick} className='size-6 mt-[20px] cursor-pointer' />
            </div>

            <div className="flex items-start w-full h-full border-b-2 border-border">
                <nav>
                    {LeftMenuItemsType().map((item) => (
                        <LeftMenuItem onClick={onClick} key={item.path} item={item} />
                    ))}
                </nav>
            </div>

            <div className="flex flex-row items-center gap-5 pb-5 pt-5 text-lg text-primary font-medium">
                {t("theme")}
                <ThemeSwitch />
            </div>
        </div>
    )
}
