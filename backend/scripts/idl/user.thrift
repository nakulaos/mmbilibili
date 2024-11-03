// 用户基本信息结构体
struct User {
    1: i64 Id (api.body = "id")                          // 用户ID
    2: string Username (api.body = "username")           // 用户名
    3: string Nickname (api.body = "nickname")           // 昵称
    4: string Avatar (api.body = "avatar")               // 头像URL
    5: i64 Gender (api.body = "gender")                  // 性别，0: 未知, 1: 男, 2: 女
    6: i64 Role (api.body = "role")                       // 用户角色，0: 普通用户, 1: 管理员
    7: i64 FollowerCount (api.body = "follower_count")   // 粉丝数量
    8: i64 FollowingCount (api.body = "following_count") // 关注数量
    9: i64 LikeCount (api.body = "like_count")           // 点赞数
    10: i64 StarCount (api.body = "star_count")          // 收藏数
    11: i64 SelfStarCount (api.body = "self_star_count") // 自己的收藏数
    12: i64 SelfLikeCount (api.body = "self_like_count") // 自己的点赞数
    13: i64 LiveCount (api.body = "live_count")          // 直播次数
    14: i64 WorkCount (api.body = "work_count")          // 作品数量
    15: i64 FriendCount (api.body = "friend_count")      // 好友数量
    16: string Phone (api.body = "phone")                // 手机号
    17: string Email (api.body = "email")                // 邮箱
    18: i64 Status (api.body = "status")                  // 用户状态
}

// 用户名登录请求
struct LoginWithUsernameReq {
    1: string Username (api.body = "username,required" api.vd="len($)>=3 && len($)<=12 && username($)")   // 用户名
    2: string Password (api.body = "password,required" api.vd="len($)>=8 && len($)<=20 && password($)")   // 密码
}

// 邮箱登录请求
struct LoginWithEmailReq {
    1: string Email (api.body = "email,required" api.vd="email($)")         // 邮箱
    2: string Password (api.body = "password,required" api.vd="len($)>=8 && len($)<=20 && password($)")   // 密码
}

// 手机号登录请求
struct LoginWithPhoneReq {
    1: string Phone (api.body = "phone,required" api.vd="phone($)")         // 手机号
    2: string Password (api.body = "password,required" api.vd="len($)>=8 && len($)<=20 && password($)")   // 密码
}

// 用户注册请求
struct RegisterReq {
    1: string Username (api.body = "username,required" api.vd="len($)>=3 && len($)<=12 && username($)")   // 用户名
    2: string Password (api.body = "password,required" api.vd="len($)>=8 && len($)<=20 && password($)")   // 密码
}

// 登录响应
struct LoginResp {
    1: string AccessToken (api.body = "access_token") // 访问令牌
    2: string RefreshToken (api.body = "refresh_token") // 刷新令牌
    3: i64 UserID (api.body = "user_id")               // 用户ID
    4: User UserInfo (api.body = "user_info")          // 用户信息
}

// 更新用户信息请求
struct UpdateUserInfoReq {
    1: string Nickname (api.body = "nickname" api.vd="len($)==0 || (len($)<=12 && len($)>=3)")   // 昵称
    2: string Avatar (api.body = "avatar" api.vd="len($)>=0 " )       // 头像URL
    3: i64 Gender (api.body = "gender" api.vd="$>=0")           // 性别
    4: i64 Role (api.body = "role"  api.vd="$>=0")               // 角色
    5: string Phone (api.body = "phone" api.vd="len($)==0 || phone($)")         // 手机号
    6: string Email (api.body = "email" api.vd="len($)==0 || email($)")         // 邮箱
}

// 关注/取关用户请求
struct FollowUserReq {
    1: i64 UserID (api.body = "user_id,required" api.vd="$>=1")         // 用户ID
    2: i64 Action (api.body = "action,required" api.vd="$>=1")           // 操作，1: 关注, 2: 取消关注
}

// 粉丝列表请求
struct FollowerListReq {
    1: i64 ActionID (api.body = "action_id,required" api.vd="$>=1")       // 动作ID
    2: i64 Page (api.body = "page,required" api.vd="$>=1")                // 页码
    3: i64 PageSize (api.body = "page_size,required" api.vd="$>=1")       // 每页大小
    4: i64 Total (api.body = "total")
}

// 粉丝列表响应
struct FollowerListResp {
    1: i64 Total (api.body = "total")             // 粉丝总数
    2: list<User> List (api.body = "list")        // 粉丝列表
}

// 关注列表请求
struct FollowingListReq {
    1: i64 ActionID (api.body = "action_id,required" api.vd="$>=1")       // 动作ID
    2: i64 Page (api.body = "page,required" api.vd="$>=1")                // 页码
    3: i64 PageSize (api.body = "page_size,required" api.vd="$>=1")       // 每页大小
    4: i64 Total (api.body = "total")
}

// 关注列表响应
struct FollowingListResp {
    1: i64 Total (api.body = "total")              // 关注总数
    2: list<User> List (api.body = "list")         // 关注列表
}

// 好友列表请求
struct FriendListReq {
    1: i64 ActionID (api.body = "action_id,required" api.vd="$>=1")       // 动作ID
    2: i64 Page (api.body = "page,required" api.vd="$>=1")                // 页码
    3: i64 PageSize (api.body = "page_size,required" api.vd="$>=1")       // 每页大小
    4: i64 Total (api.body = "total")
}

// 好友列表响应
struct FriendListResp {
    1: i64 Total (api.body = "total")              // 好友总数
    2: list<User> List (api.body = "list")         // 好友列表
}

// 用户上传文件请求
struct UserUploadFileReq {
    1: string Kind (api.body = "kind,required")             // 文件类型，avatar: 头像, cover: 封面, video: 视频
}

// 用户上传文件响应
struct UserUploadFileResp {
    1: string FileURL (api.body = "file_url")      // 文件URL
    2: string CoverURL (api.body = "cover_url")    // 封面URL
}

// 更新用户信息响应
struct UpdateUserInfoResp {
    1: User UserInfo (api.body = "user_info")      // 用户信息
}

// 用户登出响应
struct LogoutResp {
}

// 关注用户响应
struct FollowUserResp {
}

// logout
struct LogoutReq {
    1: string AccessToken (api.body="access_token,required") // 访问令牌
    2: string RefreshToken (api.body="refresh_token,required") // 刷新令牌
}

struct RefreshTokenReq {

}

struct RefreshTokenResp {
    1: string AccessToken (api.body="access_token") // 访问令牌
    2: string RefreshToken (api.body="refresh_token") // 刷新令牌
}

// 用户服务接口定义
service UserAPI {
    LoginResp LoginWithUsername(1: LoginWithUsernameReq req) (api.post = "/v1/user/login/username") // 用户名登录
    LoginResp LoginWithEmail(1: LoginWithEmailReq req) (api.post = "/v1/user/login/email")           // 邮箱登录
    LoginResp LoginWithPhone(1: LoginWithPhoneReq req) (api.post = "/v1/user/login/phone")           // 手机号登录
    LoginResp Register(1: RegisterReq req) (api.post = "/v1/user/register")                          // 用户注册
    UpdateUserInfoResp UpdateUserInfo(1: UpdateUserInfoReq req) (api.post = "/v1/auth/user/info")    // 更新用户信息
    LogoutResp Logout(1: LogoutReq req) (api.post = "/v1/auth/user/logout")                                                // 用户登出
    FollowUserResp FollowUser(1: FollowUserReq req) (api.post = "/v1/auth/user/follow")              // 关注或取消关注用户
    FollowerListResp FollowerList(1: FollowerListReq req) (api.get = "/v1/user/followers")      // 获取粉丝列表
    FollowingListResp FollowingList(1: FollowingListReq req) (api.get = "/v1/user/following")   // 获取关注列表
    FriendListResp FriendList(1: FriendListReq req) (api.get = "/v1/user/friends")             // 获取好友列表
    UserUploadFileResp UserUploadFile(1: UserUploadFileReq req) (api.post = "/v1/auth/user/upload")   // 上传文件
    RefreshTokenResp RefreshToken(1: RefreshTokenReq req) (api.post = "/v1/auth/user/refresh")   // 刷新令牌
}
