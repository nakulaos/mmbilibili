import { Avatar, Flex, Image, Typography  } from 'antd'
import Icon from '@ant-design/icons'
import { LogoSvg } from '~/components/LogoSvg/LogoSvg'




interface LogoProps {
    title: string;
    src: string;
}

const { Title } = Typography;

export const Logo: React.FC<LogoProps> = ({ title, src }) => {
    const isMobile = window.innerWidth < 768;

    return (
        <>
            <Flex align={"center"} justify={"center"} gap={20}>
                <Icon component={LogoSvg} style={{
                    width:"30px",
                    height:"30px",
                }}></Icon>
                {isMobile ? null : <Title level={4}>{title}</Title>}
            </Flex>
        </>
    );
};