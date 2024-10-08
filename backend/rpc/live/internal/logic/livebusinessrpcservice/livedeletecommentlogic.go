package livebusinessrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveDeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveDeleteCommentLogic {
	return &LiveDeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveDeleteCommentLogic) LiveDeleteComment(in *live.LiveDeleteCommentReq) (*live.LiveCommentResp, error) {
	// 获取评论信息
	commentModel, err := l.svcCtx.Dao.Danmu.WithContext(l.ctx).Where(l.svcCtx.Dao.Danmu.ID.Eq(uint(in.CommentId))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerror.CommentNotExistError, "[LiveComment] get comment error: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] get comment error: %v", err)
	}

	// 获取直播信息
	liveModel, err := l.svcCtx.Dao.Live.WithContext(l.ctx).Preload(l.svcCtx.Dao.Live.Categories).Where(l.svcCtx.Dao.Live.ID.Eq(commentModel.OwnerID)).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.LiveNotExistError, "[LiveComment] get live error: %v", err)
	}

	// 获取用户ID
	uid := in.UserId
	userModel, _ := modelcache.GetUserModelCacheFromId(l.svcCtx.Cache, l.ctx, uint(uid))
	if userModel.ID == 0 {
		userModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrapf(xerror.UserNotExistError, "[LiveComment] get user failed: %v", err)
			} else {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] get user failed: %v", err)
			}
		}
	}
	if constant.HasRole(int(userModel.Role), constant.AdminRole) == false && constant.HasRole(int(userModel.Role), constant.RootRole) == false && userModel.ID != commentModel.UID {
		return nil, errors.Wrapf(xerror.UserRoleNotAllowError, "[LiveComment] comment delete error")
	}

	// 删除评论
	q := l.svcCtx.Dao.Begin()
	_, err = q.Danmu.WithContext(l.ctx).Where(l.svcCtx.Dao.Danmu.ID.Eq(uint(in.CommentId))).Delete()
	if err != nil {
		l.Logger.Errorf("[LiveComment] delete comment error: %v", err)
		err = q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] rollback error: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] delete comment error: %v", err)
	}

	//liveModel.CommentCount -= 1
	//err = q.Live.WithContext(l.ctx).Save(liveModel)
	//if err != nil {
	//	l.Logger.Errorf("[LiveComment] save live error: %v", err)
	//	err := q.Rollback()
	//	if err != nil {
	//		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] rollback error: %v", err)
	//	}
	//	return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] save live error: %v", err)
	//}

	err = statistics.UpdateLiveCommentCount(l.svcCtx.Redis, l.ctx, int(liveModel.ID), constant.UnCommentAction)
	if err != nil {
		l.Logger.Errorf("[LiveComment] update comment count error: %v", err)
		err = q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] rollback error: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] update comment count error: %v", err)
	}

	err = q.Commit()
	if err != nil {
		l.Logger.Errorf("[LiveComment] commit error: %v", err)
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] commit error: %v", err)
	}

	go func() {
		//err := l.svcCtx.Cache.Delete(l.ctx, constant.LiveInfoCacheKey+string(liveModel.ID))
		//if err != nil {
		//	l.Logger.Errorf("[LiveComment] delete cache error: %v", err)
		//}
		RecommendUpdateWithCommentLive(l.svcCtx.Redis, l.svcCtx.GorseClient, l.Logger, liveModel, constant.UnCommentAction)
	}()
	var resp = &live.LiveCommentResp{}
	resp.CommentId = uint32(in.CommentId)

	return resp, nil
}
