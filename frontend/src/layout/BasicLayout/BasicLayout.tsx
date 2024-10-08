// src/layouts/BasicLayout.jsx
import React from 'react';
import { Outlet } from 'react-router-dom';


const BasicLayout = () => {
    return (
        <div className="basic-layout">
            this is basic layout
            <div className="content">
                <main>
                    <Outlet /> {/* 用于渲染嵌套路由的组件 */}
                </main>
            </div>
        </div>
    );
};

export default BasicLayout;
