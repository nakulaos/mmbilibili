package model

import (
	"gorm.io/gorm"
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

	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"` // 软删除
}

func (u *UserRelevantCount) TableName() string {
	return "user_relevant_counts"
}
