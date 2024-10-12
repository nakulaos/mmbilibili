import React from 'react'
import { Provider, useSelector } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react'
import RootComponent from './RootComponent'
import { persistor, store } from '@/store/store'
import { ConfigProvider } from 'antd'
import I18nComponent from '~/I18nComponent'
import { ThemeProvider } from 'antd-style';

interface AntdStyleComponentProps {
    children: React.ReactNode

}

const AntdStyleComponent:React.FC<AntdStyleComponentProps> = ({ children }) => {
    const global = useSelector((state:any) => state.global)

    return(
        <ConfigProvider >
            <ThemeProvider  themeMode={global.theme}
                // 支持传入方法，来动态响应外观
                            theme={(appearance) =>
                                appearance === 'light'
                                    ?
                                    {
                                        "components": {
                                            "Layout": {
                                                "headerBg": "rgb(255,255,255)",
                                                "headerPadding": "5 5"
                                            }
                                        }
                                    }
                                    :
                                    {
                                        "components": {
                                            "Layout": {
                                                "headerPadding": "5 5"
                                            }
                                        }
                                    }
                            }>
                {children}
            </ThemeProvider>
        </ConfigProvider>

    )

}

const App: React.FC = () => {
    // @ts-ignore

    return (
        <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
                <AntdStyleComponent >
                    <I18nComponent>
                        <RootComponent />
                    </I18nComponent>
                </AntdStyleComponent>
            </PersistGate>
        </Provider>
    )
}

export default App
