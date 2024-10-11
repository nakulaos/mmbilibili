import React from 'react'
import { Provider } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react'
import RootComponent from './RootComponent'
import { persistor, store } from '@/store/store'
import { ConfigProvider, theme } from 'antd'
import I18nComponent from '~/I18nComponent'
import { ThemeProvider } from 'antd-style';


const App: React.FC = () => {
    // @ts-ignore

    return (
        <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
                <ConfigProvider >
                    <ThemeProvider  themeMode={'auto'}
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
                        <I18nComponent>
                            <RootComponent />
                        </I18nComponent>
                    </ThemeProvider>
                </ConfigProvider>
            </PersistGate>
        </Provider>
    )
}

export default App
