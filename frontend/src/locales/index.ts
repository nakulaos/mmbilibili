// src/i18n.ts
import { createIntl, createIntlCache, IntlShape } from 'react-intl';
import messages_en from '@/locales/en-US';
import messages_zh from '@/locales/zh-CN';
import { store } from '@/store/store'; // 引入你的 Redux store

const messages: Record<string, Record<string, string>> = {
    'en-US': messages_en,
    'zh-CN': messages_zh,
};

const cache = createIntlCache();
let intlInstance: IntlShape | undefined;

// 创建函数以获取当前国际化实例
export const getIntl = (): IntlShape => {
    const language = store.getState().global.language; // 从 Redux store 获取语言
    if (!intlInstance || intlInstance.locale !== language) {
        intlInstance = createIntl(
            {
                locale: language,
                messages: messages[language],
            },
            cache
        );
    }
    return intlInstance;
};

export const updateIntl = (language: string): void => {
    intlInstance = createIntl(
        {
            locale: language,
            messages: messages[language],
        },
        cache
    );
};