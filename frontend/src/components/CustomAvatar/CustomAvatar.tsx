import { Avatar, Badge, Popover, Typography } from 'antd'
import React from 'react'
import { CustomAvatarContent } from '@/components/CustomAvatar/CustomAvatarContent'

interface CustomAvatarProps {
    src?: string;
    content?: React.ReactNode;
    count?: number;
}

export const CustomAvatar: React.FC<CustomAvatarProps> = ({ src, count, content=<CustomAvatarContent ></CustomAvatarContent> }) => {
    return (
        <Popover content={content}>
            <Badge count={count} overflowCount={99}>
                <Avatar size={'large'} src={src} />
            </Badge>
        </Popover>
    )
}
