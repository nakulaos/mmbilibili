import { AlipayCircleOutlined, TaobaoCircleOutlined, WeiboCircleOutlined } from "@ant-design/icons";
import { LoginForm, ProFormCheckbox, ProFormInstance } from '@ant-design/pro-components';
import { Flex, Tabs, Typography } from 'antd';
import { useRef, useState } from 'react';
import { useIntl } from 'react-intl';
import {
    AutomaticLoginKey,
    LoginByEmailKey,
    LoginByOtherKey,
    LoginByPhoneKey,
    LoginByUsernameKey
} from '@/locales/locale';
import { LoginModelByUsername } from '@/components/LoginModel/LoginModelByUsername';

type LoginType = 'phone' | 'username' | 'email';

export const LoginModel: React.FC = () => {
    const [loginType, setLoginType] = useState<LoginType>('username');

    const usernameLoginRef = useRef<{ login: (values: any) => void } | null>(null);
    const formRef = useRef<ProFormInstance>();

    const handleReset = () => {
        formRef.current?.resetFields();
    }

    const handleFinish = async (values: any) => {
        // 调用子组件的自定义方法
        if (loginType === 'username' && usernameLoginRef.current) {
            await usernameLoginRef.current.login(values);  // 调用子组件的 login 方法
        }
    };

    const loginMap = [
        {
            type: 'username',
            component: <LoginModelByUsername onReset={handleReset} onFinish={handleFinish} ref={usernameLoginRef} />
        },
        {
            type: 'phone',
            component: <></>
        },
        {
            type: 'email',
            component: <></>
        }
    ];

    const intl = useIntl();

    // Prepare the items for Tabs
    const tabItems = [
        {
            key: 'username',
            label: intl.formatMessage({ id: LoginByUsernameKey }),
        },
        {
            key: 'phone',
            label: intl.formatMessage({ id: LoginByPhoneKey }),
        },
        {
            key: 'email',
            label: intl.formatMessage({ id: LoginByEmailKey }),
        },
    ];

    return (
        <>
            <LoginForm
                title="mmbilibili"
                subTitle="mmbilibili is a platform for watching videos and live broadcasts"
                onFinish={handleFinish}
                formRef={formRef}
                actions={
                    <Flex align={"center"} justify={"space-between"}>
                        <Typography>
                            {intl.formatMessage({ id: LoginByOtherKey })}
                        </Typography>
                        <div>
                            <AlipayCircleOutlined />
                            <TaobaoCircleOutlined />
                            <WeiboCircleOutlined />
                        </div>
                    </Flex>
                }
            >
                <Tabs
                    centered
                    activeKey={loginType}
                    items={tabItems}
                    onChange={(activeKey) => setLoginType(activeKey as LoginType)}
                />
                {
                    loginMap.map((item) => {
                        if (item.type === loginType) {
                            return (
                                <div key={item.type}>{item.component}</div> // Add a unique key prop here
                            );
                        }
                        return null; // Ensure that you return null if the condition is not met
                    })
                }
                <div style={{ marginBlockEnd: 24 }}>
                    <ProFormCheckbox noStyle name="autoLogin">
                        {intl.formatMessage({ id: AutomaticLoginKey })}
                    </ProFormCheckbox>
                    <a style={{ float: 'right' }}>
                        忘记密码
                    </a>
                </div>
            </LoginForm>
        </>
    );
}
