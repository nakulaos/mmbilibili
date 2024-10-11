import { message, theme } from 'antd'
import { useIntl } from 'react-intl'
import { ProFormText } from '@ant-design/pro-components'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import {
    LoginByUsernameKey, OkKey,
    PasswordPlaceholderKey,
    UsernamePlaceholderKey
} from '@/locales/locale'
import { validatePassword, validateUsername } from '@/components/LoginModel/ValidationRules'
import { forwardRef, useImperativeHandle } from 'react'
import { loginWithUsername } from '@/api/common'
import { useDispatch, useSelector } from 'react-redux'
import { setUserInfo } from '@/store/userInfo'
import { Await } from 'react-router-dom'

interface LoginModelByUsernameProps {
    onFinish: (values: any) => void;  // 表单提交的处理函数
    onReset: () => void;  // 表单重置的处理函数
}

export const LoginModelByUsername = forwardRef((props: LoginModelByUsernameProps, ref) => {
    const intl = useIntl()
    const dispatch = useDispatch();

    const login =  async (values: any) => {
        await loginWithUsername(values).then((res) => {
            const updatedUserInfo = {
                id: res.data.userinfo.id,
                token: res.data.accessToken,
                username: values.username,
                avatar: res.data.userinfo.avatar,
                nickname: res.data.userinfo.nickname,
                email: res.data.userinfo.email,
                phone: res.data.userinfo.phone,
                followerCount: res.data.userinfo.follower_count,
                followingCount: res.data.userinfo.following_count,
                likeCount: res.data.userinfo.like_count,
                starCount: res.data.userinfo.star_count,
                selfStarCount: res.data.userinfo.self_star_count,
                selfLikeCount: res.data.userinfo.self_like_count,
                liveCount: res.data.userinfo.live_count,
                workCount: res.data.userinfo.work_count,
                friendCount: res.data.userinfo.friend_count,
                status: res.data.userinfo.status,
                gender: res.data.userinfo.gender,
                role: res.data.userinfo.role,
            };

            dispatch(setUserInfo(updatedUserInfo));
            message.success(intl.formatMessage({ id: OkKey }));
            props.onReset();
        })
    };

    useImperativeHandle(ref, () => ({
        login,
    }));

    return (
        <>
            <ProFormText
                name="username"
                fieldProps={{
                    size: 'large',
                    prefix: <UserOutlined className={'prefixIcon'} />,
                }}
                placeholder={intl.formatMessage({ id: UsernamePlaceholderKey })}
                rules={
                    validateUsername().map((item) => {
                        return {
                            ...item,
                            message: intl.formatMessage({ id: item.message })
                        }
                    })
                }
            />
            <ProFormText.Password
                name="password"
                placeholder={intl.formatMessage({ id: PasswordPlaceholderKey })}
                fieldProps={{
                    size: 'large',
                    prefix: <LockOutlined className={'prefixIcon'} />,
                }}
                rules={
                    validatePassword().map((item) => {
                        return {
                            ...item,
                            message: intl.formatMessage({ id: item.message })
                        }
                    })
                }
            />
        </>
    )
})