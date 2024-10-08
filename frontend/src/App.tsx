import React from 'react'
import { Provider } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react'
import RootComponent from './RootComponent'
import { persistor, store } from '@/store/store'
import { ConfigProvider } from 'antd'
import I18nComponent from '~/I18nComponent'



const App: React.FC = () => {
    // @ts-ignore

    return (
        <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
                <ConfigProvider>
                    <I18nComponent>
                        <RootComponent />
                    </I18nComponent>
                </ConfigProvider>
            </PersistGate>
        </Provider>
    )
}

export default App
