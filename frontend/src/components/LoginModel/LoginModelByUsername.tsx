import { theme } from 'antd'
import { useIntl } from 'react-intl'
import { ProFormText } from '@ant-design/pro-components'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import {
    LoginByUsernameKey,
    PasswordPlaceholderKey,
    UsernamePlaceholderKey
} from '@/locales/locale'
import { validatePassword, validateUsername } from '@/components/LoginModel/ValidationRules'
import { forwardRef, useImperativeHandle } from 'react'
import { loginWithUsername } from '@/api/common'

interface LoginModelByUsernameProps {
    onFinish: (values: any) => void;  // 表单提交的处理函数
}

export const LoginModelByUsername = forwardRef((props: LoginModelByUsernameProps, ref) => {
    const intl = useIntl()

    const login = (values: any) => {
        loginWithUsername(values).then((res)=>{
            console.log('Login response:', res);
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