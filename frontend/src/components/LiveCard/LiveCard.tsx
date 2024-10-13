import { Avatar, Image } from 'antd'
import './LiveCard.scss'
import { Live } from '@/components/LiveBox/LiveBox'
import { ThunderboltOutlined } from '@ant-design/icons'
import { useEffect, useRef, useState } from 'react'
import Player from 'xgplayer'
import FlvPlugin from 'xgplayer-flv'
import Mp4Player from 'xgplayer-mp4'

interface LiveCardProps {
    live: Live;
}

export const LiveCard: React.FC<LiveCardProps> = ({ live }) => {
    const { author } = live;
    const [isHovered, setIsHovered] = useState(false);
    const [showPlayer, setShowPlayer] = useState(false);
    const player = useRef<Player | null>(null)
    let timeoutId: NodeJS.Timeout | null = null;

    const handleMouseEnter = () => {
        setIsHovered(true);
        timeoutId = setTimeout(() => {
            setShowPlayer(true);
        }, 2000);
    };

    const handleMouseLeave = () => {
        setIsHovered(false);
        if (timeoutId) {
            clearTimeout(timeoutId);
        }
        setShowPlayer(false);
    };
    useEffect(() => {
        return () => {
            if (timeoutId) {
                clearTimeout(timeoutId);
            }
        };
    }, []);

    useEffect(() => {
        if(showPlayer){
            if(!player.current){
                player.current = new Player({
                    id: 'xgplayer',
                    url: live.player_url,
                    fluid: true,
                    plugins: [FlvPlugin],
                    videoInit: true,
                })

        }}

        return () => {
            player.current?.destroy()
            player.current = null
        }
    }, [showPlayer])


    return (
        <div className="live-card-container">
            <div className="live-card">
                <div className="card-banner">
                    <div className="banner-header"></div>
                    <div className="banner-content">
                        <div
                            className="image-container"
                            onMouseEnter={handleMouseEnter} // 添加鼠标悬停事件
                            onMouseLeave={handleMouseLeave} // 添加鼠标离开事件
                        >
                            {!showPlayer ? (
                                <div className="image" draggable={false}>
                                    <Image
                                        src={live.cover_url}
                                        preview={false}
                                        onDragStart={(e) => e.preventDefault()}
                                    />
                                    <div className="top-right-corner-overlay-element" draggable={false}>
                                        <ThunderboltOutlined />
                                    </div>
                                    <div className="left-lower-corner-overlay-element">
                                        王者荣耀
                                    </div>
                                </div>
                            ) : (
                                <div className="video-container">
                                    <div id="xgplayer"/>
                                </div>
                            )}
                        </div>
                    </div>
                    <div className="banner-footer"></div>
                </div>

                {/* 卡片底部，显示头像和用户名 */}
                <div className="card-footer">
                    <Avatar src={author.avatar} />
                    <span>{author.nickname}</span>
                </div>
            </div>
        </div>
    );
};
