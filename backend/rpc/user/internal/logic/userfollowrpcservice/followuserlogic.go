package userfollowrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowUserLogic) FollowUser(in *user.FollowUserReq) (*user.FollowUserResp, error) {
	var userModel = &model.User{}
	var followModel = &model.User{}
	var err error
	var userid = in.UserId
	var actionid = in.ActionId

	userModel, _ = modelcache.GetUserModelCacheFromId(l.svcCtx.Cache, l.ctx, uint(userid))
	if userModel.ID == 0 {
		userModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(userid))).First()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] get user error: %v", err)
		}
	}
	followModel, _ = modelcache.GetUserModelCacheFromId(l.svcCtx.Cache, l.ctx, uint(actionid))
	if followModel.ID == 0 {
		followModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(actionid))).First()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] get follower user error: %v", err)
		}
	}

	q := l.svcCtx.Dao.Begin()
	defer q.Commit()
	if int(in.Action) == constant.FollowUserAction {
		ret, err := q.User.Following.Where(q.User.ID.Eq(uint(actionid))).Model(userModel).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] get following error: %v", err)
		}
		if len(ret) != 0 {
			return nil, errors.Wrapf(xerror.FollowingIsExistError, "[FollowUser] following is exist")
		}

		// 插入 following
		//
		//userModel.FollowingCount++
		//followModel.FollowerCount++

		err = q.User.Following.Model(userModel).Append(followModel)
		if err != nil {
			q.Rollback()
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] following error: %v", err)
		}

		err = statistics.UpdateUserFollowCounts(l.svcCtx.Redis, l.ctx, int(userid), int(actionid), constant.FollowUserAction)
		if err != nil {
			q.Rollback()
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] update user follow counts error: %v", err)
		}

		//err = q.User.WithContext(l.ctx).Save(userModel, followModel)
		//if err != nil {
		//	q.Rollback()
		//	return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] save user error: %v", err)
		//}
	} else if int(in.Action) == constant.UnFollowUserAction {
		ret, err := q.User.Following.Where(q.User.ID.Eq(uint(actionid))).Model(userModel).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] get following error: %v", err)
		}
		if len(ret) == 0 {
			return nil, errors.Wrapf(xerror.FollowingIsExistError, "[FollowUser] following is not exist")
		}

		// 插入 following
		//userModel.FollowingCount--
		//followModel.FollowerCount--

		err = q.User.Following.Model(userModel).Delete(followModel)
		if err != nil {
			q.Rollback()
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] following error: %v", err)
		}

		err = statistics.UpdateUserFollowCounts(l.svcCtx.Redis, l.ctx, int(userid), int(actionid), constant.UnFollowUserAction)
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] update user follow counts error: %v", err)
		}

		//
		//err = q.User.WithContext(l.ctx).Save(userModel, followModel)
		//if err != nil {
		//	q.Rollback()
		//	return nil, errors.Wrapf(xerror.ServerError, "[FollowUser] save user error: %v", err)
		//}

	}

	//go func() {
	//	err = modelcache.DelUserModelCache(l.svcCtx.Cache, l.ctx, uint(userid))
	//	if err != nil {
	//		l.Logger.Errorf("[FollowUser] delete user cache error: %v", err)
	//	}
	//	err = modelcache.DelUserModelCache(l.svcCtx.Cache, l.ctx, uint(actionid))
	//	if err != nil {
	//		l.Logger.Errorf("[FollowUser] delete user cache error: %v", err)
	//	}
	//}()

	return nil, nil

}
