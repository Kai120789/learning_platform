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
        <div className={`fixed top-0 left-0 z-50 flex h-full w-[260px] flex-col items-center 
            border-r-2 border-border bg-background p-4 pt-0 rounded-tr-2xl rounded-br-2xl
            transition-transform duration-300 ease-in-out
            ${isOpen ? "translate-x-0" : "-translate-x-full"}`}>
            <div className='flex justify-between w-full mb-6'>
                <div className='pt-3 z-50'>
                    <FaRegFontAwesomeLogoFull className="h-[32px] w-[120px]" />
                </div>
                <CgClose onClick={onClick} className='size-5 mt-4 cursor-pointer' />
            </div>

            <div className="flex items-start w-full h-full border-b-2 border-border">
                <nav>
                    {LeftMenuItemsType().map((item) => (
                        <LeftMenuItem onClick={onClick} key={item.path} item={item} />
                    ))}
                </nav>
            </div>

            <div className="flex flex-row items-center gap-3 pb-4 pt-4 text-sm text-primary font-medium">
                {t("theme")}
                <ThemeSwitch />
            </div>
        </div>
    )
}
