import { Avatar, Flex, Image, Typography  } from 'antd'
import Icon from '@ant-design/icons'
import { LogoSvg } from '~/components/LogoSvg/LogoSvg'




interface LogoProps {
    title: string;
    src: string;
}

const { Title } = Typography;

export const Logo: React.FC<LogoProps> = ({ title, src }) => {
    return (
        <>
            <Flex align={"center"} justify={"center"}>
                <Icon component={LogoSvg} style={{
                    width:"30px",
                    height:"30px",
                }}></Icon>
                <Title level={5}>{title}</Title>
            </Flex>
        </>
    );
};