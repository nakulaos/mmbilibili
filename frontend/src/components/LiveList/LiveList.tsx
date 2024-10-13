import { ProCard } from '@ant-design/pro-components'
import { useEffect, useState } from 'react'
import { HotKey, neighborKey, NewKey, RecommendKey } from '@/locales/locale'
import { useIntl } from 'react-intl'


export const LiveList=()=>{
    const [tabKey, setTabKey] = useState('')
    const intl = useIntl()
    const onClickTab=(key:string)=>{
        setTabKey(key)
    }
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
        }
    ]

    useEffect(() => {
        console.log('LiveList')
    }, [])




    return(
        <>
            <ProCard tabs={{
                tabPosition:'top',
                activeKey:tabKey,
                onChange:onClickTab,
                items: tabItems,
            } }>



            </ProCard>
        </>
    )
}