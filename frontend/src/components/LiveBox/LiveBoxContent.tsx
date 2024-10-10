import { useEffect, useRef } from 'react';
import Player from 'xgplayer';
import FlvPlugin from 'xgplayer-flv';
import "xgplayer/dist/index.min.css";
import Mp4Player from 'xgplayer-mp4';

export const LiveBoxContent: React.FC = () => {
    const playerRef = useRef<Player | null>(null); // 使用 useRef 存储播放器实例

    useEffect(() => {
        playerRef.current = new Player({
            url: 'http://qny.hallnakulaos.cn/7_%E6%8A%93%E4%BD%8F%E4%BB%96.mp4',
            id: 'mse',
            fluid: true,
            plugins: [FlvPlugin,Mp4Player],
            videoInit: true,
            poster: 'http://qny.hallnakulaos.cn/eightqueen.png',
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
