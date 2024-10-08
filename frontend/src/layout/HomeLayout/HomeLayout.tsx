import React, { useState } from 'react';
import { Outlet } from 'react-router-dom';
import { ProLayout } from '@ant-design/pro-components';
import { useSelector } from 'react-redux';
import { IntlShape, useIntl } from 'react-intl';
import route from '~/layout/HomeLayout/_defaultprops'
import { RightBar } from '~/layout/HomeLayout/RightBar/RightBar'

const HomeLayout = () => {
    const [pathname, setPathname] = useState('/home')
    // @ts-ignore
    const userInfo = useSelector((state) => state.userInfo)
    const intl = useIntl()

    return (
        <div className="home-layout">
            <ProLayout
                title="mmbilibili"
                layout="top"
                menu={{ locale: true, type: 'group' }}
                location={{ pathname }}
                route={route}
                avatarProps={{
                    src: userInfo.avatar || 'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
                    size: 'small',
                    title: userInfo.nickname || '用户'
                }}
                menuDataRender={(props) => {
                  return  props.map((item) => {
                    return {
                      ...item,
                      name: intl.formatMessage({ id: item.locale as string }),
                    }
                  })
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
                 RightBar
                }
            >
                <Outlet />
            </ProLayout>
        </div>
    )
}


export default HomeLayout;
