import React from 'react';
import { Flex, Input, Space } from 'antd'
import Search from 'antd/es/input/Search';
import { useIntl } from 'react-intl';
import { 
    BulbOutlined, 
    ClockCircleOutlined, 
    MessageOutlined, 
    NodeIndexOutlined, 
    RadarChartOutlined, 
    StarOutlined 
} from '@ant-design/icons';
import { IconPopover } from '~/components/IconPopover/IconPopover';
import { MenuInputSearchTextKey } from '~/locales/locale'

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
    return (
        <>
            <Space size={150}>
                <Flex gap="middle" justify="center" align="center">
                    <Search
                        placeholder={intl.formatMessage({ id: MenuInputSearchTextKey })}
                        size={'large'}
                        style={{ width: 300 }}
                    />
                </Flex>
                <Flex gap="middle" justify="center" align="center">

                    {items.map((item, index) => (
                        <IconPopover
                            key={index}
                            title={item.title} // 使用国际化方法
                            icon={item.icon}
                            content={item.content}
                        />
                    ))}
                    {/* 如果需要，直接渲染额外的图标 */}
                    {/* <RadarChartOutlined />
            <MessageOutlined />
            <StarOutlined />
            <NodeIndexOutlined />
            <ClockCircleOutlined />
            <BulbOutlined /> */}
                </Flex>
            </Space>

        </>

    );
};
