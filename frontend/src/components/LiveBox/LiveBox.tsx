import React from 'react';
import { LiveBoxHeader } from './LiveBoxHeader';
import { LiveBoxFooter } from './LiveBoxFooter';
import { LiveBoxContainer } from '@/components/LiveBox/LiveBoxContainer'

export type Live = API.LiveInfo;

export const LiveBox: React.FC<Live> = (live) => {


    return (
        <>
            <LiveBoxHeader />
                <LiveBoxContainer live={live} />
            <LiveBoxFooter />
        </>
    );
};
