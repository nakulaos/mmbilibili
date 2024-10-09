import { Flex, Image } from 'antd'
import { Logo } from '@/components/Logo/Logo'


export default function BasicHeader () {

    return(
        <>
            <Flex gap={"small"} justify={"space-between"} align={"center"} style={{
                backgroundImage: 'url(http://qny.hallnakulaos.cn/mmbilibili.avif',
                backgroundSize: 'cover',
                height: '155px',
            }}
            >
                <Flex className={"left-entry"}>
                    <Logo title={"mmbilibili"} src={"/Logo.svg"} />
                </Flex>



            </Flex>

        </>
    )
}