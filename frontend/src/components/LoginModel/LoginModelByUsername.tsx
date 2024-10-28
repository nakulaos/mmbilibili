import { message, theme } from 'antd'
import { useIntl } from 'react-intl'
import { ProFormText } from '@ant-design/pro-components'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import {
     OkKey,
    PasswordPlaceholderKey,
    UsernamePlaceholderKey
} from '@/locales/locale'
import { validatePassword, validateUsername } from '@/components/LoginModel/ValidationRules'
import { forwardRef, useImperativeHandle } from 'react'

import { useDispatch, useSelector } from 'react-redux'

import { loginWithUsername } from '@/api/userApi'
import { setUserInfo } from '@/store/userInfo'

interface LoginModelByUsernameProps {
    onFinish: (values: any) => void;  // 表单提交的处理函数
    onReset: () => void;  // 表单重置的处理函数
}

export const LoginModelByUsername = forwardRef((props: LoginModelByUsernameProps, ref) => {
    const intl = useIntl()
    const dispatch = useDispatch();

    const login =  async (values: any) => {
        await loginWithUsername(values).then((res) => {
            const userInfo = {
                id: res.data.user_info.id,
                accessToken: res.data.access_token,
                refreshToken: res.data.refresh_token,
                username: values.username,
                avatar: res.data.user_info.avatar,
                nickname: res.data.user_info.nickname,
                email: res.data.user_info.email,
                phone: res.data.user_info.phone,
                followerCount: res.data.user_info.follower_count,
                followingCount: res.data.user_info.following_count,
                likeCount: res.data.user_info.like_count,
                starCount: res.data.user_info.star_count,
                selfStarCount: res.data.user_info.self_star_count,
                selfLikeCount: res.data.user_info.self_like_count,
                liveCount: res.data.user_info.live_count,
                workCount: res.data.user_info.work_count,
                friendCount: res.data.user_info.friend_count,
                status: res.data.user_info.status,
                gender: res.data.user_info.gender,
                role: res.data.user_info.role,
            };

            dispatch(setUserInfo(userInfo));
            props.onReset();
        }).catch((error) => {
          console.log(error)
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