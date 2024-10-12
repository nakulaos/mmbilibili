import React from 'react';
import { Flex, Space, Dropdown, Button, Menu } from 'antd';
import { DownOutlined } from '@ant-design/icons';



export interface RightBarProps {
    items: React.ReactNode[];
}

export const RightBar: React.FC<RightBarProps> = ({ items }) => {
    // 获取窗口宽度
    const isSmallScreen = window.innerWidth < 1400;
    const isExtraSmallScreen = window.innerWidth < 450;

    // 创建一个新的 items 数组，设置 isDropdownItem 属性
    const updatedItems = React.Children.map(items, (item: any) => {
        return React.cloneElement(item, { isDropdownItem: isExtraSmallScreen });
    });

    return (
        <>
            <Space size={isSmallScreen ? 15 : 130}>
                {isExtraSmallScreen ? (
                    <Dropdown trigger={['click']} dropdownRender={() => (
                        <Flex gap={'small'} vertical align={'flex-start'} justify={'center'} >
                            {updatedItems}
                        </Flex>
                    )}>
                        <Button icon={<DownOutlined />} />
                    </Dropdown>
                ) : (
                    <Flex gap="middle" justify="center" align="center">
                        {updatedItems}
                    </Flex>
                )}
            </Space>
        </>
    );
};