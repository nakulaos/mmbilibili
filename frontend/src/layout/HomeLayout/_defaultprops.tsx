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
    MenuMatchKey, MenuVideoKey,
    MessageKey,
    StarKey,
    TrendKey
} from '@/locales/locale'

const route = {
    path: '/',
    routes: [
        {
            path: '/home',
            name: MenuHomeKey,

        },
        {
            path: '/live',
            name: MenuLiveKey,

        },
        {
            path: '/match',
            name: MenuMatchKey,

        },
        {
            path: '/gamecenter',
            name: MenuGameCenterKey,

        },
        {
            path: '/video',
            name: MenuVideoKey,

        },
        {
            path: '/vipsupermarket',
            name: MenuGameCenterKey,
        },
    ],
};

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

export default route;
