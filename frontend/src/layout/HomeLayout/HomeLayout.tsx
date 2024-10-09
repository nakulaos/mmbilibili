import React, { useEffect, useRef, useState } from 'react'
import { NavLink, Outlet, useNavigate } from 'react-router-dom';
import { ProLayout, ProLayoutProps } from '@ant-design/pro-components'
import { useSelector } from 'react-redux';
import { IntlShape, useIntl } from 'react-intl';
import route, { rightbarProps } from '~/layout/HomeLayout/_defaultprops'
import { RightBar } from '@/layout/HomeLayout/RightBar/RightBar'
import { Image } from 'antd'
import { Logo } from '~/components/Logo/Logo'
import BasicHeader from '~/layout/BasicLayout/BasicHeader/BasicHeader'

const HomeLayout = () => {
    const [pathname, setPathname] = useState('/home')
    // @ts-ignore
    const userInfo = useSelector((state) => state.userInfo)
    const intl = useIntl()
    const navigate = useNavigate()
    const layoutRef = useRef(null)



    useEffect(() => {
        // @ts-ignore
        const header = layoutRef.current.querySelector('.ant-pro-layout-header');

        //
        // header.addEventListener("mouseenter", (e:any) => {
        //     console.log("mouseenter")
        // })
        //
        // header.addEventListener("mousemove", (e:any) => {
        //     console.log("mousemove")
        // });
        // header.addEventListener("mouseleave", (e:any)=>{
        //     console.log("mouseleave")
        // });

        // if (header) {
        //     header.style.backgroundColor = 'rgba(0, 0, 0, 0)';
        //     header.style.backgroundImage = 'url(http://qny.hallnakulaos.cn/mmbilibili.avif)'
        //     header.style.backgroundSize = 'cover';
        // }

        // const layer = document.createElement("div");
        // layer.classList.add("layer");
        // const child = document.createElement( 'img');
        // child.src = "http://qny.hallnakulaos.cn/mmbilibili.avif"
        // child.style.position = 'absolute'; // 设置为绝对定位
        // child.style.top = '0'; // 确保 child 在 layer 的顶部
        // child.style.left = '0';
        // child.style.zIndex = '-1'; // 设置为较小的 z-index 以确保其在底层
        //
        // layer.classList.add("layer");
        // layer.style.position = 'relative'; // 设置为相对定位以便使用 z-index
        // layer.appendChild(child)
        // header.appendChild(layer)


    }, [])

    return (
        <div className="home-layout"   ref={layoutRef}>
            <ProLayout
                id={"custom-prolayout"}
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

                // bgLayoutImgList={
                //     [
                //         {
                //             src: "http://qny.hallnakulaos.cn/mmbilibili.avif"
                //         }
                //     ]
                // }
                token={{
                    header: {
                        colorBgHeader: 'rgba(240, 242, 245, 0)', // 设置透明背景
                    }

                }}
                menuDataRender={(props) => {
                  return  props.map((item) => {
                    return {
                      ...item,
                      name: intl.formatMessage({ id: item.name }),
                    }
                  })
                }}

                menuItemRender={(item, dom) => {
                    return (
                        <div
                            onClick={() => {
                                console.log(item)
                                setPathname(item.path || '/home')
                                navigate(item.path || '/home')

                            }}
                        >
                            {dom}
                        </div>
                    )
                }}

                actionsRender={
                 ()=> {
                     var items = rightbarProps.items
                     items = items.map((item) => {
                         return{
                            ...item,
                            title : intl.formatMessage({ id: item.title })
                         }
                        })
                     return <RightBar items={items} />
                    }
                }
                // headerRender={(props:ProLayoutProps) => {
                //     return (
                //         <div                 style={{
                //             backgroundImage: 'url(http://qny.hallnakulaos.cn/mmbilibili.avif',
                //             backgroundSize: 'cover',
                //             height:"155px",
                //         }}>
                //             <BasicHeader></BasicHeader>
                //         </div>
                //     )
                // }}
                

            >
                <Outlet />
            </ProLayout>
        </div>
    )
}


export default HomeLayout;
