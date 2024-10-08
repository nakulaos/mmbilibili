/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/14
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string    `gorm:"type:varchar(48);not null;uniqueIndex;comment:账号名"`
	Nickname        string    `gorm:"type:varchar(48);not null"`
	Description     string    `gorm:"type:varchar(256);comment:简介"`
	Status          uint      `gorm:"type:int unsigned;default:1;not null"` // 1:正常 2:封禁
	Phone           string    `gorm:"type:varchar(48);index;null"`
	Email           string    `gorm:"type:varchar(48);index;null"`
	Password        string    `gorm:"type:varchar(256);not null"`
	Avatar          string    `gorm:"type:varchar(255);not null"`
	Role            uint      `gorm:"type:varchar(36);default:1;not null"`
	Gender          uint      `gorm:"type:int unsigned;default:1;not null"`
	Salt            string    `gorm:"type:varchar(36);not null"`
	Videos          []Video   `gorm:"foreignKey:AuthorID"`              // 一个用户可以有多个视频
	Comments        []Comment `gorm:"foreignKey:UID"`                   // 一个用户可以有多个评论
	Danmus          []Danmu   `gorm:"foreignKey:UID"`                   // 一个用户可以有多个弹幕
	FavoriteVideos  []Video   `gorm:"many2many:user_favorite_videos"`   // 一个用户可以有多个点赞视频，一个视频可以被多个用户点赞
	StarVideos      []Video   `gorm:"many2many:user_star_videos"`       // 一个用户可以有多个收藏视频，一个视频可以被多个用户收藏
	FavoriteArticle []Article `gorm:"many2many:user_favorite_articles"` // 一个用户可以有多个点赞文章，一个文章可以被多个用户点赞
	StarArticles    []Article `gorm:"many2many:user_star_articles"`     // 一个用户可以有多个收藏文章，一个文章可以被多个用户收藏
	FavoriteLives   []Live    `gorm:"many2many:user_favorite_lives"`    // 一个用户可以有多个点赞直播，一个直播可以被多个用户点赞
	Lives           []Live    `gorm:"foreignKey:UID"`                   // 一个用户可以有多场直播记录
	// 多对多自引用：关注与被关注关系
	Followers []*User `gorm:"many2many:user_follows;joinForeignKey:FollowedID;joinReferences:FollowerID"`
	Following []*User `gorm:"many2many:user_follows;joinForeignKey:FollowerID;joinReferences:FollowedID"`
	RoomID    uint    `gorm:"type:int;default:0;not null;comment:房间ID"`
	Tag       []Tag   `gorm:"many2many:user_tag;"` // 用户标签
	//FollowerCount  int     `gorm:"type:int;default:0;not null;comment:粉丝数"`
	//FollowingCount int     `gorm:"type:int;default:0;not null;comment:关注数"`
	//LikeCount      int     `gorm:"type:int;default:0;not null;comment:被点赞数"`
	//StarCount      int     `gorm:"type:int;default:0;not null;comment:被收藏数"`
	//SelfStarCount  int     `gorm:"type:int;comment:自己收藏作品数"`
	//SelfLikeCount  int     `gorm:"type:int;default:0;not null;comment:自己点赞作品数"`
	//LiveCount      int     `gorm:"type:int;default:0;not null;comment:直播次数"`
	//WorkCount      int     `gorm:"type:int;default:0;not null;comment:作品数"`
	//FriendCount    int     `gorm:"type:int;default:0;not null;comment:朋友数"`
}

/*
// Followers: 用户被哪些人关注
// Following: 用户关注了哪些人
Followers []*User `gorm:"many2many:user_follows;joinForeignKey:FollowedID;joinReferences:FollowerID"`
// joinForeignKey: 当前用户是被关注的用户，对应的是FollowedID。
// joinReferences: 关注当前用户的人，其ID是FollowerID。

Following []*User `gorm:"many2many:user_follows;joinForeignKey:FollowerID;joinReferences:FollowedID"`
// joinForeignKey: 当前用户是关注其他用户的人，对应的是FollowerID。
// joinReferences: 被当前用户关注的人，其ID是FollowedID。
*/
