import { RightBarProps } from "./RightBar/RightBar";
import {
    BulbOutlined,
    ClockCircleOutlined,
    MessageOutlined,
    NodeIndexOutlined,
    RadarChartOutlined, SettingOutlined, TranslationOutlined
} from '@ant-design/icons'
import {
    Chinese,
    CreativeCenterKey, English,
    HistoryKey,
    MajorMemberKey, MenuGameCenterKey,
    MenuHomeKey,
    MenuLiveKey,
    MenuMatchKey, MenuSettingKey, menuThemeKey, MenuTranslationKey, MenuVideoKey, MenuVIPSupermarketKey,
    MessageKey,
    StarKey,
    TrendKey
} from '@/locales/locale'
import { useIntl } from 'react-intl'
import { Menu } from 'antd'
import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { setLanguage } from '@/store/global'


export const HomeMenuData = [
    {
        path: '/home',
        key: MenuHomeKey,

    },
    {
        path: '/live',
        key: MenuLiveKey,

    },
    {
        path: '/match',
        key: MenuMatchKey,

    },
    {
        path: '/gamecenter',
        key: MenuGameCenterKey,

    },
    {
        path: '/video',
        key: MenuVideoKey,

    },
    {
        path: '/vipsupermarket',
        key: MenuVIPSupermarketKey,
    },
]



const content = (
    <div>
        <p>Content</p>
        <p>Content</p>
    </div>
);

export const TransactionContent:React.FC = ()=>{
    const intl = useIntl();
    const dispatch = useDispatch()
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

    return (
        <div>
            <Menu items={data} onClick={handleClick}></Menu>
        </div>
    )
}



export const rightbarProps:RightBarProps = {
  items:[
    {
        title:MajorMemberKey,
        icon:<BulbOutlined />,
        content:content,
    },
    {
        title:MessageKey,
        icon:<ClockCircleOutlined />,
        content:content,
    },
    {
        title:TrendKey,
        icon:<MessageOutlined />,
        content:content,
    },
    {
        title:StarKey,
        icon:<NodeIndexOutlined />,
        content:content,
    },
    {
        title:HistoryKey,
        icon:<RadarChartOutlined />,
        content:content,
    },
      {
          title:CreativeCenterKey,
          icon:<BulbOutlined />,
          content:content,
      },
      {
          title: menuThemeKey,
          icon: <SettingOutlined />,
          content: undefined
      },
      {
          icon:<TranslationOutlined />,
          title: MenuTranslationKey,
          content: <TransactionContent />,
      },
  ]  
}

