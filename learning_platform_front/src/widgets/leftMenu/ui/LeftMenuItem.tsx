import type { LeftMenuItemTab } from "@/shared/types/leftMenuItems";
import { Button } from "@/shared/ui/Button";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { FaCaretDown, FaCaretUp } from "react-icons/fa";

type LeftMenuItemProps = {
    item: LeftMenuItemTab;
    onClick: () => void
};

const LeftMenuItem = ({ item, onClick }: LeftMenuItemProps) => {
    const navigate = useNavigate();
    const [isOpen, setIsOpen] = useState(false);

    const hasChildren = !!item.childrens?.length;

    const onClickItem = () => {
        onClick()
        navigate(item.path);
    };

    return (
        <div className="flex flex-col items-start">
            <div className="flex flex-row items-center mb-2">
                <Button
                    variant="link"
                    onClick={onClickItem}
                    className="flex items-center gap-2 text-lg cursor-pointer"
                >
                    <item.icon className="size-6" />

                    {item.text}


                </Button>
                {hasChildren &&
                    (isOpen ? (
                        <FaCaretUp onClick={() => setIsOpen((prev) => !prev)} className="size-5 cursor-pointer" />
                    ) : (
                        <FaCaretDown onClick={() => setIsOpen((prev) => !prev)} className="size-5 cursor-pointer" />
                    ))
                }
            </div>

            {isOpen &&
                item.childrens?.map((child) => (
                    <div key={child.path} className="ml-10">
                        <LeftMenuItem onClick={onClick} item={child} />
                    </div>
                ))
            }
        </div>
    );
};

export default LeftMenuItem;