import BasicLayout from '@/layout/BasicLayout/BasicLayout'
import { LiveDataRoute, rightbarProps } from '@/pages/Live/data'

const LiveContent = () => {
    return (
        <div>
            LiveContent
        </div>
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