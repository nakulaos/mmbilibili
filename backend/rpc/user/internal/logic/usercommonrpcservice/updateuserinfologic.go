package usercommonrpcservicelogic

import (
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"time"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoReq) (*user.UpdateUserInfoResp, error) {
	// 使用errgroup处理并行操作
	var eg errgroup.Group
	uid := in.Id

	// 获取用户信息
	var userModel *model.User
	eg.Go(func() error {
		var err error
		err = modelcache.DelUserModelCache(l.svcCtx.Cache, l.ctx, uint(uid))
		if err != nil {
			return errors.Wrapf(xerror.ServerError, "[UpdateUserInfoLogic] DelUserModelCache error: %v", err)
		}

		userModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.Wrapf(xerror.UserNotExistError, "[UpdateUserInfoLogic] user not exist:%v", err)
			} else {
				return errors.Wrapf(xerror.ServerError, "[UpdateUserInfoLogic] user not exist:%v", err)
			}
		}
		// 更新用户信息
		if in.Nickname != nil {
			userModel.Nickname = *in.Nickname
		}
		if in.Avatar != nil {
			userModel.Avatar = *in.Avatar
		}
		if in.Gender != nil {
			userModel.Gender = uint(*in.Gender)
		}
		if in.Role != nil {
			userModel.Role = uint(*in.Role)
		}
		if in.Phone != nil {
			userModel.Phone = *in.Phone
		}
		if in.Email != nil {
			userModel.Email = *in.Email
		}

		// 保存更新后的用户信息
		err = l.svcCtx.Dao.User.WithContext(l.ctx).Save(userModel)
		if err != nil {
			return errors.Wrapf(xerror.ServerError, "[UpdateUserInfoLogic] update user info error:%v", err)
		}

		// 重新删除缓存（延迟500毫秒）
		go func() {
			time.Sleep(500 * time.Millisecond)
			ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
			if err := modelcache.DelUserModelCache(l.svcCtx.Cache, ctx, uint(uid)); err != nil {
				l.Logger.Errorf("[UpdateUserInfoLogic] DelUserModelCache error: %v", err)
			}
		}()
		return nil
	})

	// 获取用户统计数据
	var userStats statistics.UserRelevantCount
	eg.Go(func() error {
		var err error
		userStats, err = statistics.GetUserRelevantCount(l.svcCtx.Redis, l.ctx, int(uid))
		if err != nil {
			return errors.Wrapf(xerror.ServerError, "[UpdateUserInfoLogic] GetUserRelevantCount error:%v", err)
		}
		return nil
	})

	// 等待前面的goroutine完成
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// 填充响应
	var resp user.UpdateUserInfoResp
	resp.UserInfo = &user.User{
		Id:             uint32(userModel.ID),
		Username:       userModel.Username,
		Role:           uint32(userModel.Role),
		Avatar:         userModel.Avatar,
		Nickname:       userModel.Nickname,
		Gender:         uint32(userModel.Gender),
		Phone:          userModel.Phone,
		Email:          userModel.Email,
		FollowerCount:  int32(userStats.FollowerCount),
		FollowingCount: int32(userStats.FollowingCount),
		LikeCount:      int32(userStats.LikeCount),
		StarCount:      int32(userStats.StarCount),
		SelfStarCount:  int32(userStats.SelfStarCount),
		SelfLikeCount:  int32(userStats.SelfLikeCount),
		LiveCount:      int32(userStats.LiveCount),
		WorkCount:      int32(userStats.WorkCount),
		FriendCount:    int32(userStats.FriendCount),
	}
	return &resp, nil
}
