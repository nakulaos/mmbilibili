package model

import (
	"github.com/vmihailenco/msgpack/v5"
	"gorm.io/gorm"
	"time"
)

const (
	RelationshipAttrNone = iota
	RelationshipAttrWhisper
	RelationshipAttrBlack     = 1 << 1
	RelationshipAttrFollowing = 1 << 2
	RelationshipAttrFriend    = 1 << 3
)

const (
	ActAddFollowing = 1
	ActDelFollowing = 2
	ActAddWhisper   = 3
	ActDelWhisper   = 4
	ActAddBlack     = 5
	ActDelBlack     = 6
	ActDelFollower  = 7
)

func RelationshipAttr(attr int64) int64 {
	if attr&RelationshipAttrBlack > 0 {
		return RelationshipAttrBlack
	}
	if attr&RelationshipAttrFriend > 0 {
		return RelationshipAttrFriend
	}
	if attr&RelationshipAttrFollowing > 0 {
		return RelationshipAttrFollowing
	}
	if attr&RelationshipAttrWhisper > 0 {
		return RelationshipAttrWhisper
	}
	return RelationshipAttrNone
}

func SetAttr(attr int64, mask int64) int64 {
	return attr | mask
}

func UnsetAttr(attr int64, mask int64) int64 {
	return attr & ^mask
}

type UserRelationship struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" msgpack:"-"` // 自定义ID字段
	CreatedAt        time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" msgpack:"-"`
	UpdatedAt        time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" msgpack:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index;comment:删除时间" msgpack:"-"`                                          // 软删除
	UserID           int64          `gorm:"not null" json:"user_id" msgpack:"-"`                                     // 用户ID
	RelatedUserID    int64          `gorm:"not null" json:"related_user_id" msgpack:"related_user_id"`               // 相关用户ID
	RelationshipAttr int64          `gorm:"not null;default:0" json:"relationship_attr" msgpack:"relationship_attr"` // 关系属性
}

func (u *UserRelationship) TableName() string {
	return "user_relationships"
}

func (u *UserRelationship) Marshal() ([]byte, error) {
	return msgpack.Marshal(u)
}

func (u *UserRelationship) Unmarshal(b []byte) error {
	return msgpack.Unmarshal(b, u)
}

func Filter(arr []*UserRelationship, attr int64) []*UserRelationship {
	var res []*UserRelationship = make([]*UserRelationship, 0)
	for _, v := range arr {
		if RelationshipAttr(v.RelationshipAttr) == attr {
			res = append(res, v)
		}
	}
	return res
}
