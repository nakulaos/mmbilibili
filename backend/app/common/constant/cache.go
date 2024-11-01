package constant

var (
	AppName string = "mmbilibili:"
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

var (
	UserDetailCacheFromUidKey      string = AppName + "User:Detail:Uid"
	UserDetailCacheFromUsernameKey string = AppName + "User:Detail:Username"
	UserRelevantCountFromIdKey     string = AppName + "User:RelevantCount:Uid"
	UserTokenBlackListKey          string = AppName + "User:TokenBlackList:Token"
	UserRelationshipKey            string = AppName + "User:Relationship:Uid" // 用户关系,%d 是分片值，会进行一个打散处理
)

// lock
var (
	LockUserRelation string = AppName + "User:Relation:Lock:Uid"
)

var (
	LiveInfoCacheKey string = AppName + "Live:Detail:Lid"
	LiveViewCount    string = AppName + "Live:ViewCount:Lid"    // 直播观看次数
	LiveLikeCount    string = AppName + "Live:LikeCount:Lid"    // 直播点赞次数
	LiveCommentCount string = AppName + "Live:CommentCount:Lid" // 直播评论次数
	LiveHotRoomTags  string = AppName + "Live:HotRoom:"         // 热门房间标签 eg: LiveHotRoomTags:FpsGame  room_id  30
)

var (
	CategoryInfoCidKey   string = AppName + "Category:Detail:Cid"
	CategoryInfoNameKey  string = AppName + "Category:Detail:Name"
	CategoryUserPortrait string = AppName + "CategoryUserPortrait:uid" // 分类下的用户头像 eg: CategoryUserPortrait:1 FpsGame 30
	UserTokenCount       string = AppName + "UserTokenCount:Uid"       // 发弹幕所用时间 eg: UserTokenCount:Uid:1  30
	CategoryRoomScore    string = AppName + "CategoryRoomScore"        // 分类下的房间分数 eg: CategoryRoomScore:FpsGame  room_id  10
	StreamNameToLid      string = AppName + "Live:Lid:StreamName"      // 直播流名字到直播ID的映射
)
