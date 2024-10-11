import React from 'react';
import { Flex, Space, Dropdown, Button, Menu } from 'antd';
import Search from 'antd/es/input/Search';
import { useIntl } from 'react-intl';
import { DownOutlined } from '@ant-design/icons';
import { IconPopover } from '@/components/IconPopover/IconPopover';
import { MenuInputSearchTextKey } from '@/locales/locale';
import { CustomAvatar } from '@/components/CustomAvatar/CustomAvatar'

export interface RightBarItem {
    title: string;
    icon: React.ReactNode;
    content: React.ReactNode;
}

export interface RightBarProps {
    items: RightBarItem[];
}

export const RightBar: React.FC<RightBarProps> = ({ items }) => {
    const intl = useIntl();

    // 获取窗口宽度
    const isSmallScreen = window.innerWidth < 1400;
    const isExtraSmallScreen = window.innerWidth < 450;

    const dropdownMenu = (
        <Menu>
            {items.map((item, index) => (
                <Menu.Item key={index} icon={item.icon}>
                    {item.title}
                </Menu.Item>
            ))}
        </Menu>
    );

    return (
        <>
            <Space size={isSmallScreen ? 15 : 130}>
                {isExtraSmallScreen ? (
                    <Dropdown overlay={dropdownMenu} trigger={['click']}>
                        <Button icon={<DownOutlined />}>
                        </Button>
                    </Dropdown>
                ) : (
                    <Flex gap="middle" justify="center" align="center">
                        {items.map((item, index) => (
                            <IconPopover
                                key={index}
                                title={item.title} // 使用国际化方法
                                icon={item.icon}
                                content={item.content}
                            />
                        ))}
                    </Flex>
                )}
            </Space>
        </>
    );
};
