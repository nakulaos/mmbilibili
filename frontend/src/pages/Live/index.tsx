import BasicLayout from '@/layout/BasicLayout/BasicLayout'
import { LiveDataRoute, rightbarProps } from '@/pages/Live/data'
import { LiveBox } from '~/components/LiveBox/LiveBox';

const LiveContent = () => {
    return (
        <>
            <LiveBox />
        </>
    )
}



export default function Live() {
    return (
        <div>
            <BasicLayout route={LiveDataRoute} rightBar={rightbarProps} >
                <LiveContent />
            </BasicLayout>
        </div>
    );
}