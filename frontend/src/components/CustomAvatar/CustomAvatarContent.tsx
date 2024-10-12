import React, { useState } from 'react'
import { useIntl } from 'react-intl'
import {
    AuthorizationErrorKey,
    CustomAvatarContentAfterLoginYouCanKey,
    CustomAvatarContentMultiTerminalSynchronousPlaybackRecordKey,
    customAvatarContentNoAccountKey,
    customAvatarContentPostBulletCommentsKey,
    CustomAvatarContentWatchHighDefinitionVideosForFreeKey,
    customAvatarContentWatchingPopularTVDramasAndMoviesKey,
    followersKey,
    followingsKey,
    loginKey,
    logoutKey,
    OkKey,
    personalCenterKey,
    recommendationServiceKey,
    registerKey,
    submissionManagementKey,
    worksKey
} from '@/locales/locale'
import { Button, Col, Divider, Flex, Menu, message, Modal, Progress, Row, Space, Tag, Typography } from 'antd'
import {
    ClockCircleOutlined,
    FireOutlined,
    NodeIndexOutlined,
    PlayCircleOutlined,
    PlaySquareOutlined,
    SelectOutlined,
    SwitcherOutlined,
    UserOutlined
} from '@ant-design/icons'
import { LoginModel } from '@/components/LoginModel/LoginModel'
import { useDispatch, useSelector } from 'react-redux'
import { StatisticCard } from '@ant-design/pro-components'
import { logout } from '@/api/follow'
import { clearToken, clearUserInfo } from '@/store/userInfo'

const { Title } = Typography


export const CustomAvatarContent: React.FC = () => {
    const intl = useIntl()
    const userInfo = useSelector((state: any) => state.userInfo)
    const [current, setCurrent] = useState('') // 默认选中的菜单项
    const dispatch = useDispatch()
    const handleClick = async (e: any) => {
        console.log('Clicked menu item:', e.key) // 获取被点击的菜单项的 key
        if (e.key === logoutKey) {
            // 退出登录
            await logout({}).then((res) => {
                dispatch(clearUserInfo())
                dispatch(clearToken())
                message.success(intl.formatMessage({ id: OkKey }))
            },(error)=>{
                dispatch(clearUserInfo())
                dispatch(clearToken())
                message.error(intl.formatMessage({id:AuthorizationErrorKey}))
            })
        }
    }

    const [visibilityForLoginModal, setVisibilityForLoginModal] = useState(false)
    return (
        <div>
            {userInfo.id ? (
                <>
                    <Flex gap={'10'} vertical style={{ width: '300px' }}>
                        <Flex gap={0} align={'center'} justify={'center'}>
                            <Title level={5} style={{ color: 'magenta' }}>
                                {userInfo.nickname || userInfo.username}
                            </Title>
                        </Flex>
                        <Flex align={'center'} justify={'center'}>
                            <Tag color="volcano">大会员</Tag>
                        </Flex>
                        <Flex align={'center'} justify={'center'}>
                            <Typography style={{ fontSize: '12px' }}>
                                硬币: 755.6 B币: 0
                            </Typography>
                        </Flex>
                        <Row>
                            <Col span={2}><Typography style={{ fontSize: '12px', color: 'orange', fontWeight: 'bold' }}>
                                Lv4</Typography>
                            </Col>
                            <Col span={1}></Col>
                            <Col span={18}>
                                <Progress percent={40} showInfo={false} status={'active'} />
                            </Col>
                            <Col span={1}></Col>
                            <Col span={2}>
                                <Typography style={{ fontSize: '12px', color: 'gray', fontWeight: 'bold' }}>
                                    Lv5
                                </Typography>
                            </Col>
                        </Row>
                        <StatisticCard.Group>
                            <StatisticCard
                                onClick={() => {
                                    handleClick({ key: followingsKey })
                                }}
                                statistic={{
                                    title: intl.formatMessage({ id: followingsKey }),
                                    value: userInfo.followingCount
                                }}
                            />
                            <StatisticCard
                                onClick={() => {
                                    handleClick({ key: followersKey })
                                }}
                                statistic={{
                                    title: intl.formatMessage({ id: followersKey }),
                                    value: userInfo.followerCount
                                }}
                            />
                            <StatisticCard
                                onClick={() => {
                                    handleClick({ key: worksKey })
                                }}
                                statistic={{
                                    title: intl.formatMessage({ id: worksKey }),
                                    value: userInfo.workCount
                                }}
                            />
                        </StatisticCard.Group>
                        <Menu onClick={handleClick} selectedKeys={[current]}>
                            <Menu.Item key={personalCenterKey} icon={<UserOutlined />}>
                                {intl.formatMessage({ id: personalCenterKey })}
                            </Menu.Item>
                            <Menu.Item key={submissionManagementKey} icon={<NodeIndexOutlined />}>
                                {intl.formatMessage({ id: submissionManagementKey })}
                            </Menu.Item>
                            <Menu.Item key={recommendationServiceKey} icon={<FireOutlined />}>
                                {intl.formatMessage({ id: recommendationServiceKey })}
                            </Menu.Item>
                            <Divider />
                            <Menu.Item key={logoutKey} icon={<SelectOutlined />}>
                                {intl.formatMessage({ id: logoutKey })}
                            </Menu.Item>
                        </Menu>
                    </Flex>

                </>
            ) : (
                <>
                    <Space direction="vertical" size={15}>
                        <div>
                            {intl.formatMessage({ id: CustomAvatarContentAfterLoginYouCanKey })}
                        </div>
                        <Flex gap={'small'} align={'center'} justify={'space-between'}>
                            <Flex gap={'small'}>
                                <Space size={10}>
                                    <PlayCircleOutlined style={{ color: 'skyblue' }} />
                                    {intl.formatMessage({ id: CustomAvatarContentWatchHighDefinitionVideosForFreeKey })}
                                </Space>
                                <Space size={5}>
                                    <ClockCircleOutlined style={{ color: 'skyblue' }} />
                                    {intl.formatMessage({ id: CustomAvatarContentMultiTerminalSynchronousPlaybackRecordKey })}
                                </Space>
                            </Flex>
                        </Flex>
                        <Flex gap={'small'} align={'center'} justify={'space-between'}>
                            <Flex gap={'small'}>
                                <Space size={10}>
                                    <SwitcherOutlined style={{ color: 'skyblue' }} />
                                    {intl.formatMessage({ id: customAvatarContentPostBulletCommentsKey })}
                                </Space>
                                <Space size={5}>
                                    <PlaySquareOutlined style={{ color: 'skyblue' }} />
                                    {intl.formatMessage({ id: customAvatarContentWatchingPopularTVDramasAndMoviesKey })}
                                </Space>
                            </Flex>
                        </Flex>
                        <Button size={'large'} type={'primary'} block onClick={() => {
                            setVisibilityForLoginModal(!visibilityForLoginModal)
                        }}>{intl.formatMessage({ id: loginKey })}</Button>
                        <Flex gap={'small'} align={'center'} justify={'center'}>
                            <div>
                                {intl.formatMessage({ id: customAvatarContentNoAccountKey })}
                                <Button type={'link'}> {intl.formatMessage({ id: registerKey })}</Button>
                            </div>
                        </Flex>
                    </Space>
                    <Modal open={visibilityForLoginModal}
                           onOk={() => {
                               setVisibilityForLoginModal(false)
                           }}
                           onCancel={() => {
                               setVisibilityForLoginModal(false)
                           }}
                           onClose={() => {
                               setVisibilityForLoginModal(false)
                           }}
                           footer={null}
                    >
                        <LoginModel></LoginModel>
                    </Modal>
                </>


            )}
        </div>
    )
}
