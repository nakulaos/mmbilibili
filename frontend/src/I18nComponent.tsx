import React, { useEffect } from 'react'
import messages_en from '@/locales/en-US';
import messages_zh from '@/locales/zh-CN';
import { useSelector } from 'react-redux';
import { IntlProvider } from 'react-intl';

const messages = {
    'en-US': messages_en,
    'zh-CN': messages_zh,
};

interface I18nComponentProps {
    children: React.ReactNode; // 定义 children 的类型
}

const I18nComponent: React.FC<I18nComponentProps> = ({ children }) => {
    // @ts-ignore
    const language = useSelector((state) => state.global.language);

    useEffect(() => {
        console.log('Current language:', language);
    }, [language]);

    // @ts-ignore
    return (
        <IntlProvider locale={language} messages={messages[language] }>
            {children}
        </IntlProvider>
    );
};

export default I18nComponent;
