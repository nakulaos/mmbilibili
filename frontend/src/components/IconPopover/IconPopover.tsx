import { Button, Flex, Popover, Typography } from "antd";

interface IconPopoverProps {
    title: string;
    icon?: React.ReactNode;
    content: React.ReactNode;
}

export const IconPopover: React.FC<IconPopoverProps> = ({ title, icon, content }) => {
    return (
        <Popover content={content}>
            <Flex gap="0" vertical justify="center" align="center">
                {icon && (
                    <Button icon={icon} style={{ border: 'none' }} />
                )}
                <Typography>
                    {title}
                </Typography>
            </Flex>
        </Popover>
    );
};
