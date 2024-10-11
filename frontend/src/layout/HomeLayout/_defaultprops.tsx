import { RightBarProps } from "./RightBar/RightBar";
import {
    BulbOutlined,
    ClockCircleOutlined,
    MessageOutlined,
    NodeIndexOutlined,
    RadarChartOutlined
} from '@ant-design/icons'
import {
    CreativeCenterKey,
    HistoryKey,
    MajorMemberKey, MenuGameCenterKey,
    MenuHomeKey,
    MenuLiveKey,
    MenuMatchKey, MenuVideoKey, MenuVIPSupermarketKey,
    MessageKey,
    StarKey,
    TrendKey
} from '@/locales/locale'
import VipSuperMarket from '~/pages/VipSuperMarket'


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
      }
  ]  
}

