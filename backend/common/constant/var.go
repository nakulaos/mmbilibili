/**
 ******************************************************************************
 * @file           : var.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/25
 ******************************************************************************
 */

package constant

var (
	UnKnowUser int = -1
)

var (
	UID       string = "UID"
	UserInfo  string = "UserInfo"
	ErrorCode string = "ErrorCode"
)

var (
	TokenBlackList string = "TokenBlackList:"
)

// 用户状态
var (
	UserStatusNormal  uint = 1
	UserNotAllowLogin uint = 2
	UserNotAllowLive  uint = 4
)

var (
	LikeAction         int = 1
	StarAction         int = 1
	UnLikeAction       int = 2
	UnStarAction       int = 2
	CommentAction      int = 1
	UnCommentAction    int = 2
	FollowUserAction   int = 1
	UnFollowUserAction int = 2
)

var (
	DanmuVideoType int = 1
	DanmuLiveType  int = 2
)

var (
	Live     string = "直播"
	LiveType string = "live"
)

var (
	IncrRoomID          string = "RoomID"
	IncrRoomIDInitValue int    = 100000
)

var (
	LikeFeedBack    string = "like"
	StarFeedBack    string = "star"
	ShareFeedBack   string = "share"
	CommentFeedBack string = "comment"
	ReadFeedBack    string = "read"
)

var (
	ListDefaultLimit int = 20
)

func HasStatus(userStatus, status uint) bool {
	return userStatus&status == status
}
