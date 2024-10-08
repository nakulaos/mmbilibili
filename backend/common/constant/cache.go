package constant

var (
	AppName string = "mmbilibili:"
)

var (
	UserInfoCacheUidKey string = AppName + "UserInfoCache:Uid:"
	UserNameToUidKey    string = AppName + "GetUidFromUniqueKey:Username:"
)

var (
	LiveInfoCacheKey string = AppName + "LiveInfoCache:Lid:"
)

var (
	CategoryInfoCidKey  string = AppName + "CategoryInfoCache:Cid:"
	CategoryInfoNameKey string = AppName + "CategoryNameCache:Name:"
)

var (
	CategoryUserPortrait string = AppName + "CategoryUserPortrait:uid:" // 分类下的用户头像 eg: CategoryUserPortrait:1 FpsGame 30
	HotRoomTags          string = AppName + "HotRoomTags:"              // 热门房间标签 eg: HotRoomTags:FpsGame  room_id  30
	UserTokenCount       string = AppName + "UserTokenCount:Uid:"       // 发弹幕所用时间 eg: UserTokenCount:Uid:1  30
	CategoryRoomScore    string = AppName + "CategoryRoomScore:"        // 分类下的房间分数 eg: CategoryRoomScore:FpsGame  room_id  10
	StreamNameToLid      string = AppName + "StreamNameToLid:"          // 直播流名字到直播ID的映射
)

var (
	InitRoomScore    = 10
	StartLiveScore   = 300
	StarLiveScore    = 20
	CommentLiveScore = 30
	LikeLiveScore    = 10
	ShareLiveScore   = 40
	ReadLiveScore    = 5
)

// 统计数相关变量
// 用户相关的
var (
	AppUserFollowerCount  string = AppName + "UserFollowerCount:Uid:"  // 用户的粉丝数
	AppUserFollowingCount string = AppName + "UserFollowingCount:Uid:" // 用户的关注数
	AppUserFriendCount    string = AppName + "UserFriendCount:Uid:"    // 用户好友数
	AppUserLikeCount      string = AppName + "UserLikeCount:Uid:"      // 用户被点赞数
	AppUserStarCount      string = AppName + "UserStarCount:Uid:"      // 用户被收藏数
	AppUserWorkCount      string = AppName + "UserWorkCount:Uid:"      // 用户作品数
	AppUserSelfStarCount  string = AppName + "UserSelfStarCount:Uid:"  // 用户自己收藏作品数
	AppUserSelfLikeCount  string = AppName + "UserSelfLikeCount:Uid:"  // 用户自己点赞作品数
	AppUserLiveCount      string = AppName + "UserLiveCount:Uid:"      // 用户直播次数
)

var (
	AppLiveViewCount    string = AppName + "LiveViewCount:Lid:"    // 直播观看次数
	AppLiveLikeCount    string = AppName + "LiveLikeCount:Lid:"    // 直播点赞次数
	AppLiveCommentCount string = AppName + "LiveCommentCount:Lid:" // 直播评论次数
)
