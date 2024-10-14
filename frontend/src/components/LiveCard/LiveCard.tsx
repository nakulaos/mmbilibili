import { Avatar, Image } from 'antd'
import './LiveCard.scss'
import { Live } from '@/components/LiveBox/LiveBox'
import { ThunderboltOutlined } from '@ant-design/icons'
import { useEffect, useRef, useState } from 'react'
import Player from 'xgplayer'
import FlvPlugin from 'xgplayer-flv.js'
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
    const [isFirstHover, setIsFirstHover] = useState(true);

    const handleMouseEnter = () => {
        setIsHovered(true);
        const delay = isFirstHover ? 2000 : 800;

        timeoutId = setTimeout(() => {
            setShowPlayer(true);
            setIsFirstHover(false);
        }, delay);
    };

    const handleMouseLeave = () => {
        setIsHovered(false);
        if (timeoutId) {
            clearTimeout(timeoutId);
            timeoutId = null;
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
                    autoplay: true,
                    disableProgress: true,
                    volume: 0,
                    closeInactive: true,
                    closeVideoClick: true,
                    closeVideoDblclick: true,
                    controls: false,
                })
            }else{
                player.current?.play()
            }
        }else{
            if(player.current){
                player.current?.pause()
            }

        }

        return () => {
            player.current?.destroy()
            player.current = null
        }
    }, [showPlayer])


    return (
        <div className="live-card-container">
            <div className="live-card"
                 onMouseEnter={handleMouseEnter}
                 onMouseLeave={handleMouseLeave}>
                <div className="card-banner">
                    <div className="banner-header"></div>
                    <div className="banner-content" >
                        <div
                            className={'image-container'}
                        >
                            {!showPlayer ? (
                                <div className={showPlayer?'image-content hidden':'image-content visible'} draggable={false}>
                                    <Image
                                        src={live.cover_url}
                                        preview={false}
                                        onDragStart={(e) => e.preventDefault()}
                                        height={168.75}
                                        width={300}
                                    />
                                    <div className="top-right-corner-overlay-element" draggable={false}>
                                        <ThunderboltOutlined />
                                    </div>
                                    <div className="left-lower-corner-overlay-element">
                                        王者荣耀
                                    </div>
                                </div>
                            ) : (
                                <div className={showPlayer ? 'video-container visible' : 'video-container hidden'}>
                                    <div id="xgplayer"/>
                                </div>
                            )}
                        </div>
                    </div>
                    <div className="banner-footer"></div>
                </div>

                <div className="card-footer">
                    <div className={"avatar-container"}>
                        <Avatar src={author.avatar} />
                    </div>
                    <div className={"action-container"}>

                    </div>

                    <span>{author.nickname}</span>
                </div>
            </div>
        </div>
    );
};
