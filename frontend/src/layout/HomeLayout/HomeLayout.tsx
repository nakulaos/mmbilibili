import { Button, Drawer, Flex, Layout, Menu } from 'antd'
import { Outlet, useNavigate } from 'react-router-dom'
import { Logo } from '@/components/Logo/Logo'
import {
 MenuInputSearchTextKey,
} from '@/locales/locale'
import { useIntl } from 'react-intl'
import { HomeMenuData, rightbarProps } from '@/layout/HomeLayout/_defaultprops'
import { RightBar } from '@/layout/HomeLayout/RightBar/RightBar'
import { useSelector } from 'react-redux'
import { CustomAvatar } from '@/components/CustomAvatar/CustomAvatar'
import React, { useEffect, useState } from 'react'
import Search from 'antd/es/input/Search'
import "./index.css"
import {  MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons'

const { Header, Footer, Sider, Content } = Layout;

export const HomeLayout= ()=> {
    const intl = useIntl();
    const userInfo = useSelector((state:any) => state.userInfo)
    const [isMobile, setIsMobile] = useState(window.innerWidth < 768);
    const [openDrawer, setOpenDrawer] = useState(false);
    const navigate = useNavigate();

    const homeMenuData = HomeMenuData.map((item) => {
        return({
            ...item,
            label: intl.formatMessage({ id: item.key })
        })
    })
    const rightBarItems = rightbarProps.items.map((item) => {
        return {
            ...item,
            title: intl.formatMessage({ id: item.title })
        }
    })

    const onCloseDrawer = () => {
        setOpenDrawer(false);
    };

    const onOpenDrawer = () => {
        setOpenDrawer(true);
    }

    const onClickMenu = (e:any) => {
        const item = homeMenuData.filter((item) => item.key === e.key);
        navigate(item[0]?.path || '/home');
    }

    // 使用 useEffect 监听窗口变化
    useEffect(() => {
        const handleResize = () => {
            setIsMobile(window.innerWidth < 768);
        };

        window.addEventListener('resize', handleResize);

        // 组件卸载时移除监听器
        return () => {
            window.removeEventListener('resize', handleResize);
        }
    }, []);

    return (
        <>
            <Layout>
                <Header className={"rootLayout-header"}>
                    <Flex gap={'small'} align={'center'} justify={'space-between'}>
                        <Flex gap={10} align={'center'} justify={'center'} className={'rootLayout-header-left-entry'}>
                            <Logo title={"mmbilibili"} src={"/Logo.svg"} />
                            {
                                isMobile ? (
                                    <>
                                        {
                                            openDrawer
                                                ? (<Button icon={<MenuFoldOutlined />} onClick={onCloseDrawer}></Button>)
                                                : (<Button icon={<MenuUnfoldOutlined />} onClick={onOpenDrawer}></Button>)
                                        }
                                        <Drawer open={openDrawer} onClose={onCloseDrawer} placement={'left'}>
                                            <Menu items={homeMenuData}></Menu>
                                        </Drawer>
                                    </>
                                ) : (
                                    <Menu mode="horizontal" onClick={onClickMenu} items={homeMenuData} style={{ flex: "auto", minWidth: 0 }}></Menu>
                                )
                            }
                        </Flex>
                        <div className={"rootLayout-header-middle-entry"}>
                            <Flex gap="middle" justify="center" align="center">
                                <Search
                                    placeholder={intl.formatMessage({ id: MenuInputSearchTextKey })}
                                    size={'large'}
                                />
                            </Flex>
                        </div>
                        <Flex gap={'small'} className={'rootLayout-header-right-entry'}>
                            <RightBar items={rightBarItems} />
                            <div className={"rootLayout-header-right-entry-avatar"}>
                                <CustomAvatar src={userInfo.avatar}></CustomAvatar>
                            </div>
                        </Flex>
                    </Flex>
                </Header>

                <Content>
                    <Outlet></Outlet>
                </Content>
                <Footer></Footer>
            </Layout>
        </>
    )
}
