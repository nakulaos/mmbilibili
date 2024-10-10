declare namespace API {
    type EndLiveReq = {
        live_id: number
    }

    type FollowerListReq = {
        page?: number
        page_size?: number
        total?: number
    }

    type FollowerListResp = {
        total: number
        list: User[]
    }

    type FollowingListReq = {
        page?: number
        page_size?: number
        total?: number
    }

    type FollowingListResp = {
        total: number
        list: User[]
    }

    type FollowUserReq = {
        user_id: number
        action: number
    }

    type FriendListReq = {
        page?: number
        page_size?: number
        total?: number
    }

    type FriendListResp = {
        total: number
        list: User[]
    }

    type LiveComment = {
        comment_id: number
        user_id: number
        live_id: number
        content: string
        create_at: number
    }

    type LiveCommentListReq = {
        live_id: number
        page: number
        page_size: number
    }

    type LiveCommentListResp = {
        total: number
        list: LiveComment[]
    }

    type LiveCommentReq = {
        live_id: number
        content: string
        send_time: number
    }

    type LiveCommentResp = {
        comment_id: number
    }

    type LiveDeleteCommentReq = {
        comment_id: number
    }

    type LiveDetailReq = {
        live_id: number
    }

    type LiveDetailResp = {
        live_info: LiveInfo
    }

    type LiveInfo = {
        id: number
        user_id: number
        title: string
        cover: string
        status: number
        start_time: number
        end_time: number
        watch_count: number
        like_count: number
        comment_count: number
        share_count: number
        is_like: boolean
        is_follow: boolean
        /**  这个字段没有实际含义 */
        is_star: boolean
        /**  是否是自己的直播 */
        is_self: boolean
        author: User
        type: number
        description: string
        player_url: string
        cover_url: string
        is_over: boolean
        category: string[]
        tags: string[]
        partition: string
        room_id: number
        /**  推流token或者拉流token */
        token: string
    }

    type LiveLikeReq = {
        live_id: number
        action: number
    }

    type LiveLikeResp = {
        like_count: number
    }

    type LiveListReq = {
        page: number
        page_size: number
        type: string
        item_id: number
        category: string
        total: number
    }

    type LiveListResp = {
        total: number
        list: LiveInfo[]
    }

    type LoginResp = {
        accessToken: string
        user_id: number
        userinfo: User
    }

    type LoginWithEmailReq = {
        email: string
        password: string
    }

    type LoginWithPhoneReq = {
        phone: string
        password: string
    }

    type LoginWithUsernameReq = {
        username: string
        password: string
    }

    type LogoutReq = {}

    type RegisterReq = {
        username: string
        password: string
    }

    type StartLiveReq = {
        title: string
        cover: string
        description: string
        category: string[]
        tags: string[]
        partition: string
    }

    type UpdateUserInfoReq = {
        nickname?: string
        avatar?: string
        gender?: number
        role?: number
        phone?: string
        email?: string
    }

    type UpdateUserInfoResp = {
        userinfo: User
    }

    type User = {
        id: number
        username: string
        nickname: string
        avatar: string
        gender: number
        role: number
        follower_count: number
        following_count: number
        like_count: number
        star_count: number
        self_star_count: number
        self_like_count: number
        live_count: number
        work_count: number
        friend_count: number
        phone: string
        email: string
        status: number
    }

    type UserUploadFileResp = {
        url: string
        cover_url: string
    }
}
