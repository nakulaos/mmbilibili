import React from 'react'
import { useIntl } from 'react-intl'
import { useDispatch, useSelector } from 'react-redux'
import {
    Chinese, CreativeCenterKey,
    English, HistoryKey,
    MajorMemberKey,
    menuThemeKey,
    MenuTranslationKey,
    MessageKey,
    StarKey, TrendKey
} from '@/locales/locale'
import { setLanguage, setTheme } from '@/store/global'
import { IconPopover } from '@/components/IconPopover/IconPopover'
import { Menu } from 'antd'
import {
    BulbOutlined,
    ClockCircleOutlined,
    MessageOutlined,
    MoonOutlined,
    RadarChartOutlined,
    SunOutlined, TranslationOutlined
} from '@ant-design/icons'


export interface LanguageIconPopoverProps {
    isDropdownItem?: boolean;
}

export const LanguageIconPopover:React.FC<LanguageIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const dispatch = useDispatch()
    const icon = <TranslationOutlined/>
    const data = [
        {
            key: Chinese,
            label: intl.formatMessage({ id: Chinese })
        },
        {
            key: English,
            label: intl.formatMessage({ id: English })
        }
    ]

    const handleClick = (e:any) => {
        dispatch(setLanguage(e.key))
    }

    const title = intl.formatMessage({ id: MenuTranslationKey })

    return (
        <>
            <IconPopover icon={icon} isDropdownItem={isDropdownItem} title={title} content={<Menu items={data} onClick={handleClick}></Menu>}></IconPopover>
        </>
    )
}

export interface MajorMemberIconPopoverProps {
    isDropdownItem?: boolean;
}

export const MajorMemberIconPopover:React.FC<MajorMemberIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const title = intl.formatMessage({ id: MajorMemberKey });
    const icon = <BulbOutlined />;

    const Content = () => {
        return (
            <>
                <p>Content</p>
                <p>Content</p>
            </>
        );
    };

    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} icon={icon} content={<Content />} />
        </>
    );
};


export interface MessageIconPopoverProps {
    isDropdownItem?: boolean;
}

export const MessageIconPopover:React.FC<MessageIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const title = intl.formatMessage({ id: MessageKey });
    const icon = <ClockCircleOutlined />;
    const Content = () => {
        return (
            <>
                <p>Content</p>
                <p>Content</p>
            </>
        );
    };
    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} icon={icon} content={<Content />} />
        </>
    );
}


export interface StarIconPopoverProps {
    isDropdownItem?: boolean;
}

export const StarIconPopover:React.FC<StarIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const title = intl.formatMessage({ id: StarKey });
    const icon = <MessageOutlined />;

    const Content = () => {
        return (
            <>
                <p>Content</p>
                <p>Content</p>
            </>
        );
    };

    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} icon={icon} content={<Content />} />
        </>
    );
}


export interface ThemeIconPopoverProps {
    isDropdownItem?: boolean;

}


export const ThemeIconPopover:React.FC<ThemeIconPopoverProps>= ({isDropdownItem})=>{

    const intl = useIntl()
    const title = intl.formatMessage({ id: menuThemeKey })
    const global = useSelector((state:any) => state.global)
    const dispatch = useDispatch()
    const icon = global.theme === 'dark' ? <SunOutlined /> : <MoonOutlined />
    const handleThemeChange = () => {
        if(global.theme === 'dark') {
            dispatch(setTheme('light'));
        }else{
            dispatch(setTheme('dark'));
        }}

    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} content={undefined} icon={icon} onClick={handleThemeChange} />
        </>
    )
}


export interface TrendIconPopoverProps {
    isDropdownItem?: boolean;
}

export const TrendIconPopover:React.FC<TrendIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const title = intl.formatMessage({ id: TrendKey });
    const icon = <MessageOutlined />;
    const Content = () => {
        return (
            <>
                <p>Content</p>
                <p>Content</p>
            </>
        );
    };
    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} icon={icon} content={<Content />} />
        </>
    );
}

export interface HistoryIconPopoverProps {
    isDropdownItem?: boolean;
}

export const HistoryIconPopover:React.FC<HistoryIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const title = intl.formatMessage({ id: HistoryKey });
    const icon = <RadarChartOutlined />;
    const Content = () => {
        return (
            <>
                <p>Content</p>
                <p>Content</p>
            </>
        );
    };
    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} icon={icon} content={<Content />} />
        </>
    );
}

export interface CreativeCenterIconPopoverProps{
    isDropdownItem?: boolean;
}

export const CreativeCenterIconPopover:React.FC<CreativeCenterIconPopoverProps> = ({isDropdownItem}) => {
    const intl = useIntl();
    const title = intl.formatMessage({ id: CreativeCenterKey });
    const icon = <BulbOutlined />;
    const Content = () => {
        return (
            <>
                <p>Content</p>
                <p>Content</p>
            </>
        );
    };
    return (
        <>
            <IconPopover isDropdownItem={isDropdownItem} title={title} icon={icon} content={<Content />} />
        </>
    );
}