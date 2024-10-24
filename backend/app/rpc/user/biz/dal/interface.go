package dal

import (
	"backend/app/rpc/user/biz/model"
	"context"
)

type UserDal interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	ExistUserByUserName(ctx context.Context, name string) (bool, error)
	ExistUserByID(ctx context.Context, id int64) (bool, error)
	UpdateUserByID(ctx context.Context, id int64, user *model.User) error
	DeleteUserByID(ctx context.Context, id int64) error
	CreateUser(ctx context.Context, user *model.User) error
	GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]*model.User, error)
	GetFollowersByUserID(ctx context.Context, userID int64, total, offset int) (map[int64]*model.User, error)
	GetFollowingsByUserID(ctx context.Context, userID int64, total, offset int) (map[int64]*model.User, error)
	GetFriendsByUserID(ctx context.Context, userID int64, total, offset int) (map[int64]*model.User, error)
	GetUserByUserName(ctx context.Context, userName string) (*model.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserRelevantCountByID(ctx context.Context, id int64) (*model.UserRelevantCount, error)
	AddTokenToBlackList(ctx context.Context, token string) error
}
