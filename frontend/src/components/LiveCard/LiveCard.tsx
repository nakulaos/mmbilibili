import { Avatar, Image } from 'antd';
import './LiveCard.scss';
import { Live } from '@/components/LiveBox/LiveBox';
import { ThunderboltOutlined } from '@ant-design/icons';
import { useEffect, useRef, useState } from 'react';
import Player from 'xgplayer';
import FlvPlugin from 'xgplayer-flv.js';

interface LiveCardProps {
    live: Live;
}

export const LiveCard: React.FC<LiveCardProps> = ({ live }) => {
    const { author } = live;
    const [showPlayer, setShowPlayer] = useState(false);
    const player = useRef<Player | null>(null);
    let timeoutId: NodeJS.Timeout | null = null;

    const handleMouseEnter = () => {
        const delay = 2000;

        timeoutId = setTimeout(() => {
            setShowPlayer(true);
        }, delay);
    };

    const handleMouseLeave = () => {
        setShowPlayer(false);
        if (timeoutId) {
            clearTimeout(timeoutId);
            timeoutId = null;
        }
    };

    useEffect(() => {
        return () => {
            if (timeoutId) {
                clearTimeout(timeoutId);
            }
        };
    }, []);

    useEffect(() => {
        if (showPlayer) {
            if (!player.current) {
                player.current = new Player({
                    id: `xgplayer-${live.id}`,
                    url: live.player_url,
                    fluid: true,
                    plugins: [FlvPlugin],
                    autoplay: true,
                    disableProgress: true,
                    volume: 0,
                    closeInactive: true,
                    closeVideoClick: true,
                    closeVideoDblclick: true,
                    controls: false,
                });
            }
        }

        return () => {
            player.current?.destroy();
            player.current = null;
        };
    }, [showPlayer]);

    return (
        <div className="live-card-container">
            <div
                className="live-card"
                onMouseEnter={handleMouseEnter}
                onMouseLeave={handleMouseLeave}
            >
                <div className="card-banner">
                    <div className="banner-header"></div>
                    <div className="banner-content">
                        <div className="banner-container">
                            {!showPlayer ? (
                                <>
                                    <div className={'image-content'} draggable={false}>
                                        <img
                                            src={live.cover_url}
                                            // preview={false}
                                            draggable={false}
                                            onDragStart={(e) => e.preventDefault()}
                                            className="image"
                                        />
                                    </div>
                                    <div className="top-right-corner-overlay-element" draggable={false}>
                                        <ThunderboltOutlined />
                                    </div>
                                    <div className="left-lower-corner-overlay-element">
                                        王者荣耀
                                    </div>
                                </>

                            ) : (
                                <div className={'video-container'}>
                                    <div id={`xgplayer-${live.id}`} />
                                </div>
                            )}
                        </div>
                    </div>
                    <div className="banner-footer"></div>
                </div>

                <div className="card-footer">
                    <div className={'avatar-container'}>
                        <Avatar src={author.avatar} />
                    </div>
                    <div className={"action-container"}></div>
                    <span>{author.nickname}</span>
                </div>
            </div>
        </div>
    );
};
