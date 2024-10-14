import { LiveCard } from '@/components/LiveCard/LiveCard'
import { Live } from '@/components/LiveBox/LiveBox'

export default function Match() {
    const liveInfo: Live = {
        id: 0,
        user_id: 0,
        title: '默认标题',
        cover: '默认封面',
        status: 0,
        start_time: 0,
        end_time: 0,
        watch_count: 0,
        like_count: 0,
        comment_count: 0,
        share_count: 0,
        is_like: false,
        is_follow: false,
        is_star: false,
        is_self: false,
        author: {
            id: 0,
            username: '默认用户名',
            nickname: '默认昵称',
            avatar: 'http://qny.hallnakulaos.cn/gvb/20231119133444__avatar.jpg.png',
            gender: 0,
            role: 0,
            follower_count: 0,
            following_count: 0,
            like_count: 0,
            star_count: 0,
            self_star_count: 0,
            self_like_count: 0,
            live_count: 0,
            work_count: 0,
            friend_count: 0,
            phone: '000-0000',
            email: 'default@example.com',
            status: 0,
        },
        type: 0,
        description: '默认描述',
        player_url: 'http://localhost:8090/live/6/1726042205.live.flv',
        cover_url: 'http://qny.hallnakulaos.cn/gvb/20231115112603__OIP-C%20%281%29.jpg.png',
        is_over: false,
        category: ['默认分类'],
        tags: ['默认标签'],
        partition: '默认分区',
        room_id: 0,
        token: '默认token',
    };
    return (
        <div>
            <LiveCard live={liveInfo}></LiveCard>
        </div>
    );
}