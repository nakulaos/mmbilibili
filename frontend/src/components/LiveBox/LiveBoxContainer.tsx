import React from 'react';
import { LiveBoxContent } from './LiveBoxContent';
import { Live } from '@/components/LiveBox/LiveBox';

interface LiveBoxContainerProps {
    live: Live;
}

export const LiveBoxContainer: React.FC<LiveBoxContainerProps> = ({ live }) => {

    const { player_url, id, cover_url, token } = live;

    return (
        <>
            <LiveBoxContent url={player_url} liveId={id} poster={cover_url} pullToken={token} />
        </>
    );
};
