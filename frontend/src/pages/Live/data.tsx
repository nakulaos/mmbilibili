import {
    ChatRoomKey,
    CreativeCenterKey,
    EntertainmentKey,
    HelpMePlayKey,
    HistoryKey,
    InteractiveGamePlayKey,
    KnowledgeKey,
    LifeKey,
    MajorMemberKey,
    MatchKey,
    MenuGameCenterKey,
    MenuHomeKey,
    MenuLiveKey,
    MenuMatchKey,
    MenuVideoKey,
    MenuVIPSupermarketKey, MessageKey, MobileGameKey, OnlineGameKey, RadioStationKey, ShoppingKey, SinglePlayerGameKey, StarKey, TrendKey,
    VirtualAnchorKey
} from '~/locales/locale'
import { RightBarProps } from '~/layout/HomeLayout/RightBar/RightBar'
import {
    BulbOutlined,
    ClockCircleOutlined,
    MessageOutlined,
    NodeIndexOutlined,
    RadarChartOutlined
} from '@ant-design/icons'
import path from 'path';


export const LiveDataRoute = {
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
            path: "OnlineGames",
            name: OnlineGameKey,
        },
        {
            path: "MobileGames",
            name: MobileGameKey,
        },
        {
            path: "SingleGame",
            name:SinglePlayerGameKey,
        },
        {
            path: "VirtualAnchor",
            name: VirtualAnchorKey,
        },
        {
            path: "Entertainment",
            name: EntertainmentKey,
        },
        {
            path: "RadioStation",
            name: RadioStationKey,
        },
        {
            path: "Match",
            name: MatchKey,
        },
        {
            path: "ChatRoom",
            name: ChatRoomKey,
        },
        {
            path: "Life",
            name: LifeKey,
        },
        {
            path: "Knowledge",
            name: KnowledgeKey,
        },
        {
            path: "HelpMePlay",
            name: HelpMePlayKey,
        },
        {
            path: "InteractiveGamePlay",
            name: InteractiveGamePlayKey,
        },
        {
            path: "Shopping",
            name: ShoppingKey,
        }

    ],
}

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