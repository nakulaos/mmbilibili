import { useEffect, useRef } from 'react';
import Player from 'xgplayer';
import FlvPlugin from 'xgplayer-flv';
import "xgplayer/dist/index.min.css";
import Mp4Player from 'xgplayer-mp4';

interface LiveBoxContentProps {
    url: string;
    poster: string;
    pullToken: string;
    liveId: number;
}

// 测试：ffmpeg -re -i test.flv -vcodec h264 -acodec aac -f flv "rtmp://localhost:1935/live/6/1726042205?uid=6&lid=3"
export const LiveBoxContent: React.FC<LiveBoxContentProps> = ({url,poster,pullToken,liveId}) => {
    const playerRef = useRef<Player | null>(null); // 使用 useRef 存储播放器实例

    useEffect(() => {
        if(FlvPlugin.isSupported()){
            console.log('FlvPlugin is supported');
        }else{
            console.log('FlvPlugin is not supported');
        }
        playerRef.current = new Player({
            url: 'http://localhost:8090/live/6/1726042205.live.flv',
            id: 'mse',
            fluid: true,
            plugins: [FlvPlugin],
            videoInit: true,
            // poster: poster,
            isLive: true,
        });

        return () => {
            playerRef.current?.destroy();
            playerRef.current = null;
        };
    }, []);

    return (
        <>
            <div id="mse" ></div>
        </>
    );
};
