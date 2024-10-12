import { Button, Flex, Popover, Typography } from "antd";

interface IconPopoverProps {
    title: string;
    icon?: React.ReactNode;
    content: React.ReactNode;
    onClick?: () => void;
    isDropdownItem?: boolean;
}

export const IconPopover: React.FC<IconPopoverProps> = ({ title, icon, content,onClick,isDropdownItem=false }) => {
    if(isDropdownItem){
        return (
            <Button type={'text'} onClick={onClick} icon={icon}>
                {title}
            </Button>
        )
    }

    return (
        <Popover content={content} >
            <Flex gap="0" vertical justify="center" align="center">
                {icon && (
                    <Button icon={icon} style={{ border: 'none' }} onClick={onClick} />
                )}
                <Typography>
                    {title}
                </Typography>
            </Flex>
        </Popover>
    );
};
