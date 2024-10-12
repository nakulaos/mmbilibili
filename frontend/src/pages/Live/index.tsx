import BasicLayout from '@/layout/BasicLayout/BasicLayout'
import { LiveBox } from '~/components/LiveBox/LiveBox';
import {
    ChatRoomKey,
    EntertainmentKey, HelpMePlayKey, InteractiveGamePlayKey, KnowledgeKey, LifeKey, MatchKey,
    MenuHomeKey,
    MenuLiveKey,
    MobileGameKey,
    OnlineGameKey, RadioStationKey, ShoppingKey,
    SinglePlayerGameKey,
    VirtualAnchorKey
} from '@/locales/locale'
import { RightBarProps } from '@/components/RightBar/RightBar'
import {
    CreativeCenterIconPopover,
    HistoryIconPopover, LanguageIconPopover,
    MajorMemberIconPopover,
    MessageIconPopover, StarIconPopover, ThemeIconPopover,
    TrendIconPopover
} from '@/components/IconPopover/IconPopoverInstance'
import React from 'react'



const LiveContent = () => {

    return (
        <>
            <LiveBox />
        </>
    )
}



export default function Live() {
    const LiveDataRoute = {
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
    const rightBarProps: RightBarProps = {
        items: [
            <MajorMemberIconPopover />,
            <MessageIconPopover />,
            <TrendIconPopover />,
            <StarIconPopover />,
            <HistoryIconPopover />,
            <CreativeCenterIconPopover />,
            <ThemeIconPopover />,
            <LanguageIconPopover />
        ]
    }

    return (
        <div>
            <BasicLayout route={LiveDataRoute} rightBar={rightBarProps} >
                <LiveContent />
            </BasicLayout>
        </div>
    );
}