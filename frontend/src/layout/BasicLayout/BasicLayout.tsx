// // src/layouts/BasicLayout.jsx
// import React from 'react';
// import { Outlet } from 'react-router-dom';
// import { Layout } from 'antd'
// import BasicHeader from '~/layout/BasicLayout/BasicHeader/BasicHeader'
//
// const { Header, Footer, Sider, Content } = Layout;
//
//
//
// const BasicLayout = () => {
//     return (
//
//         <Layout >
//             <Layout >
//                 <BasicHeader></BasicHeader>
//             </Layout>
//             <Layout>
//                 <Outlet />
//             </Layout>
//             <Layout></Layout>
//         </Layout>
//         // <Layout>
//         //     <Outlet></Outlet>
//         // </Layout>
//
//
//     );
// };
//
// export default BasicLayout;

import React, {  useState } from 'react'
import {  useNavigate } from 'react-router-dom';
import { ProLayout } from '@ant-design/pro-components'

import { useIntl } from 'react-intl';
import { RightBar, RightBarProps } from '@/components/RightBar/RightBar'
import { CustomAvatar } from '@/components/CustomAvatar/CustomAvatar'
import { useSelector } from 'react-redux'


interface BasicLayoutProps {
    route:any
    rightBar:RightBarProps
    children:React.ReactNode
}

// @ts-ignore
export const BasicLayout = ({route,rightBar,children}) => {
    const [pathname, setPathname] = useState('/home')
    const userInfo = useSelector((state:any) => state.userInfo)
    const intl = useIntl()
    const navigate = useNavigate()

    return (
        <>
            <ProLayout
                id={"custom-prolayout"}
                title="mmbilibili"
                layout="top"
                menu={{ locale: true, type: 'group' }}
                location={{ pathname }}
                avatarProps={{
                    render: ()=> {
                        return (
                            <>
                                <CustomAvatar src={userInfo.avatar}></CustomAvatar>
                            </>
                        )
                    }
                }}
                menuItemRender={(item, dom) => {
                    return (
                        <div
                            onClick={() => {
                                console.log(item)
                                setPathname(item.path || '/home')
                                if(item.path?.startsWith('/')){
                                    navigate(item.path)
                                }
                            }}
                        >
                            {dom}
                        </div>
                    )
                }}
                actionsRender={
                    ()=> {
                       return <RightBar items={rightBar.items} />
                    }
                }
                menuDataRender={(props) => {
                    return  props.map((item) => {
                        return {
                            ...item,
                            name: intl.formatMessage({ id: item.name }),
                        }
                    })
                }}
                route={route}
            >
                {children}

            </ProLayout>
        </>
    )
}


export default BasicLayout;
