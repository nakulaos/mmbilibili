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
import { Outlet, useNavigate } from 'react-router-dom';
import { ProLayout } from '@ant-design/pro-components'
import { useSelector } from 'react-redux';
import { useIntl } from 'react-intl';
import { RightBar, RightBarProps } from '@/layout/HomeLayout/RightBar/RightBar'

interface BasicLayoutProps {
    route:any
    rightBar:RightBarProps
    children:React.ReactNode
}

// @ts-ignore
export const BasicLayout = ({route,rightBar,children}) => {
    const [pathname, setPathname] = useState('/home')
    // @ts-ignore
    const userInfo = useSelector((state) => state.userInfo)
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
                    src: userInfo.avatar || 'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
                    size: 'small',
                    title: userInfo.nickname || 'user'
                }}
                menuItemRender={(item, dom) => {
                    return (
                        <div
                            onClick={() => {
                                console.log(item)
                                setPathname(item.path || '/home')
                            }}
                        >
                            {dom}
                        </div>
                    )
                }}
                actionsRender={
                    ()=> {
                        var items = rightBar.items
                        // @ts-ignore
                        items = items.map((item) => {
                            return{
                                ...item,
                                title : intl.formatMessage({ id: item.title })
                            }
                        })
                        return <RightBar items={items} />
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
