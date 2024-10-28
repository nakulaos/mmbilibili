declare namespace API {
    type CompleteMultipartReq = {
        /**  文件哈希值 */
        file_hash: string
    }

    type CompleteMultipartResp = {}

    type EndLiveReq = {
        live_id: number
    }

    type FollowerListParams = {
        /**  页码 */
        page: number
        /**  每页大小 */
        page_size: number
    }

    type FollowerListReq = {
        /**  页码 */
        page: number
        /**  每页大小 */
        page_size: number
    }

    type FollowerListResp = {
        /**  粉丝总数 */
        total: number
        /**  粉丝列表 */
        list: User[]
    }

    type FollowingListParams = {
        /**  页码 */
        page: number
        /**  每页大小 */
        page_size: number
    }

    type FollowingListReq = {
        /**  页码 */
        page: number
        /**  每页大小 */
        page_size: number
    }

    type FollowingListResp = {
        /**  关注总数 */
        total: number
        /**  关注列表 */
        list: User[]
    }

    type FollowUserReq = {
        /**  用户ID */
        user_id: number
        /**  操作，1: 关注, 2: 取消关注 */
        action: number
    }

    type FollowUserResp = {}

    type FriendListParams = {
        /**  页码 */
        page: number
        /**  每页大小 */
        page_size: number
    }

    type FriendListReq = {
        /**  页码 */
        page: number
        /**  每页大小 */
        page_size: number
    }

    type FriendListResp = {
        /**  好友总数 */
        total: number
        /**  好友列表 */
        list: User[]
    }

    type GetMultiUploadUriReq = {
        /**  文件哈希值 */
        file_hash: string
        /**  分块 ID */
        chunk_id: number
        /**  分块大小 */
        chunk_size: number
    }

    type GetMultiUploadUriResp = {
        /**  URI */
        uri: string
    }

    type GetSuccessChunksReq = {
        /**  文件哈希值 */
        file_hash: string
    }

    type GetSuccessChunksResp = {
        /**  是否已经上传 */
        is_upload: boolean
        /**  是否已经记录在数据库 */
        is_record: boolean
        /**  已经上传的分块 */
        chunks: string
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
        /**  访问令牌 */
        access_token: string
        /**  刷新令牌 */
        refresh_token: string
        /**  用户ID */
        user_id: number
        user_info: User
    }

    type LoginWithEmailReq = {
        /**  邮箱 */
        email: string
        /**  密码 */
        password: string
    }

    type LoginWithPhoneReq = {
        /**  手机号 */
        phone: string
        /**  密码 */
        password: string
    }

    type LoginWithUsernameReq = {
        /**  用户名 */
        username: string
        /**  密码 */
        password: string
    }

    type LogoutReq = {
        /**  访问令牌 */
        access_token: string
        /**  刷新令牌 */
        refresh_token: string
    }

    type LogoutResp = {}

    type NewMultiUploadReq = {
        /**  文件内容的唯一哈希值 */
        file_hash: string
        /**  分块总数 */
        chunk_total_number: number
        /**  文件名 */
        file_name: string
        /**  文件大小 */
        file_size: number
        /**  文件类型 */
        file_type: number
    }

    type NewMultiUploadResp = {}

    type OnPublicStreamReq = {
        /**  流应用名 */
        app: string
        /**  TCP 链接唯一 ID */
        id: string
        /**  推流器 IP */
        ip: string
        /**  推流 URL 参数 */
        params: string
        /**  推流器端口号 */
        port: number
        /**  推流协议, 可能是 rtsp、rtmp */
        schema: string
        /**  流 ID */
        stream: string
        /**  流虚拟主机 */
        vhost: string
        /**  服务器 ID, 通过配置文件设置 */
        mediaServerId: string
    }

    type OnPublicStreamResp = {}

    type OnStreamChangeReq = {
        /**  流应用名 */
        app: string
        /**  流注册或注销 */
        regist: boolean
        /**  rtsp 或 rtmp */
        schema: string
        /**  流 ID */
        stream: string
        /**  流虚拟主机 */
        vhost: string
        /**  服务器 ID, 通过配置文件设置 */
        mediaServerId: string
    }

    type OnStreamChangeResp = {}

    type RefreshTokenReq = {}

    type RefreshTokenResp = {
        /**  访问令牌 */
        access_token: string
        /**  刷新令牌 */
        refresh_token: string
    }

    type RegisterReq = {
        /**  用户名 */
        username: string
        /**  密码 */
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
        /**  昵称 */
        nickname: string
        /**  头像URL */
        avatar: string
        /**  性别 */
        gender: number
        /**  角色 */
        role: number
        /**  手机号 */
        phone: string
        /**  邮箱 */
        email: string
    }

    type UpdateUserInfoResp = {
        user_info: User
    }

    type User = {
        /**  用户ID */
        id: number
        /**  用户名 */
        username: string
        /**  昵称 */
        nickname: string
        /**  头像URL */
        avatar: string
        /**  性别，0: 未知, 1: 男, 2: 女 */
        gender: number
        /**  用户角色，0: 普通用户, 1: 管理员 */
        role: number
        /**  粉丝数量 */
        follower_count: number
        /**  关注数量 */
        following_count: number
        /**  点赞数 */
        like_count: number
        /**  收藏数 */
        star_count: number
        /**  自己的收藏数 */
        self_star_count: number
        /**  自己的点赞数 */
        self_like_count: number
        /**  直播次数 */
        live_count: number
        /**  作品数量 */
        work_count: number
        /**  好友数量 */
        friend_count: number
        /**  手机号 */
        phone: string
        /**  邮箱 */
        email: string
        /**  用户状态 */
        status: number
    }

    type UserUploadFileReq = {
        kind: string
    }

    type UserUploadFileResp = {
        url: string
        cover_url: string
    }
}
