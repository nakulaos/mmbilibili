import React from 'react'
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import './styles/main.sass'
import HomeLayout from '@/layout/HomeLayout/HomeLayout'
import Home from '@/pages/Home'
import BasicLayout from '@/layout/BasicLayout/BasicLayout'
import Live from '@/pages/Live'
import Match from '@/pages/Match'
import Video from '@/pages/Video'
import VipSuperMarket from '@/pages/VipSuperMarket'
import GameCenter from '@/pages/GameCenter'

const RootComponent: React.FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Navigate to="/home" replace />} />
                <Route element={<HomeLayout />}>
                    <Route path="/home" element={<Home />} />
                </Route>


                <Route path="/live" element={<Live />} />
                <Route path="/match" element={<Match />} />
                <Route path="/gamecenter" element={<GameCenter />} />
                <Route path="/video" element={<Video />} />
                <Route path="/vipsupermarket" element={<VipSuperMarket />} />

            </Routes>
        </BrowserRouter>
    )
}

export default RootComponent
