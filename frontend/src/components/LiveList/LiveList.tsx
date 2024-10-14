import { useEffect, useState } from 'react';
import { HotKey, neighborKey, NewKey, RecommendKey } from '@/locales/locale';
import { useIntl } from 'react-intl';
import { liveInfos } from '@/components/LiveList/data';
import { LiveCard } from '@/components/LiveCard/LiveCard';
import './LiveList.scss';

export const LiveList = () => {
    const [tabKey, setTabKey] = useState('');
    const intl = useIntl();

    const onClickTab = (key: string) => {
        setTabKey(key);
    };

    const tabItems = [
        {
            key: NewKey,
            label: intl.formatMessage({ id: NewKey }),
        },
        {
            key: HotKey,
            label: intl.formatMessage({ id: HotKey }),
        },
        {
            key: RecommendKey,
            label: intl.formatMessage({ id: RecommendKey }),
        },
        {
            key: neighborKey,
            label: intl.formatMessage({ id: neighborKey }),
        },
    ];

    const data = liveInfos;

    useEffect(() => {
        console.log('LiveList');
    }, []);

    return (
        <div className="live-list-container">
            <div className="tabs">
                {tabItems.map((item) => (
                    <div
                        key={item.key}
                        className={`tab ${tabKey === item.key ? 'active' : ''}`}
                        onClick={() => onClickTab(item.key)}
                    >
                        {item.label}
                    </div>
                ))}
            </div>
            <div className="live-cards-container">
                {data.map((live, index) => (
                    <div key={index} className="live-card-wrapper">
                        <LiveCard live={live} />
                    </div>
                ))}
            </div>
        </div>
    );
};
