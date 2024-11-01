package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserRelevantCount struct {
	ID             int64 `gorm:"primaryKey;autoIncrement"` // 自定义ID字段
	UserID         int64 `gorm:"type:int;not null;uniqueIndex;comment:用户ID"`
	FollowerCount  int64 `gorm:"type:int;default:0;not null;comment:粉丝数"`
	FollowingCount int64 `gorm:"type:int;default:0;not null;comment:关注数"`
	LikeCount      int64 `gorm:"type:int;default:0;not null;comment:被点赞数"`
	StarCount      int64 `gorm:"type:int;default:0;not null;comment:被收藏数"`
	SelfStarCount  int64 `gorm:"type:int;comment:自己收藏作品数"`
	SelfLikeCount  int64 `gorm:"type:int;default:0;not null;comment:自己点赞作品数"`
	LiveCount      int64 `gorm:"type:int;default:0;not null;comment:直播次数"`
	WorkCount      int64 `gorm:"type:int;default:0;not null;comment:作品数"`
	FriendCount    int64 `gorm:"type:int;default:0;not null;comment:朋友数"`
	BlackCount     int64 `gorm:"type:int;default:0;not null;comment:黑名单数"`
	WhisperCount   int64 `gorm:"type:int;default:0;not null;comment:密聊数"`

	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"` // 软删除
}

const (
	TypeFollowerCount  = 1
	TypeFollowingCount = 2
	TypeLikeCount      = 3
	TypeStarCount      = 4
	TypeSelfStarCount  = 5
	TypeSelfLikeCount  = 6
	TypeLiveCount      = 7
	TypeWorkCount      = 8
	TypeFriendCount    = 9
	TypeBlackCount     = 10
	TypeWhisperCount   = 11
)

type UserRelevantCountMessage struct {
	UserID      int64          `json:"user_id"`
	CountChange map[int8]int64 `json:"count_change"`
}

func NewUserRelevantCountMessage(userID int64) *UserRelevantCountMessage {
	return &UserRelevantCountMessage{
		UserID:      userID,
		CountChange: make(map[int8]int64),
	}
}

func (m *UserRelevantCountMessage) GetUserRelevantCountMessageKey() []byte {
	key := []byte("user_relevant_count_message_")
	key = strconv.AppendInt(key, m.UserID, 10)
	return key
}
func (m *UserRelevantCountMessage) Json() []byte {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return []byte{}
	}
	return jsonData
}

func (u *UserRelevantCount) TableName() string {
	return "user_relevant_counts"
}
