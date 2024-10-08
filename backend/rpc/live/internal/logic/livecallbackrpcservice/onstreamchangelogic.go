package livecallbackrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/pkg/xerror"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhenghaoz/gorse/client"
	"strconv"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnStreamChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnStreamChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnStreamChangeLogic {
	return &OnStreamChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnStreamChangeLogic) OnStreamChange(in *live.OnStreamChangeReq) (*live.OnStreamChangeResp, error) {
	var resp = &live.OnStreamChangeResp{}
	if !in.Regist {
		// 注销流
		ret, err := l.svcCtx.Redis.GetCtx(l.ctx, fmt.Sprintf("%s%s", constant.StreamNameToLid, in.Stream))
		if err != nil {
			resp.Code = 1
			resp.Msg = "Get stream name error"
			return nil, errors.Wrapf(xerror.ServerError, "[OnStreamChange] get stream name error")
		}
		lid, err := strconv.Atoi(ret)
		if err != nil {
			resp.Code = 1
			resp.Msg = "Parse lid error"
			return nil, errors.Wrapf(xerror.ServerError, "[OnStreamChange] parse lid error: %v", err)
		}
		var liveModel = &model.Live{}
		err = l.svcCtx.Cache.Get(l.ctx, constant.LiveInfoCacheKey+fmt.Sprintf("%d", lid), liveModel)
		if err != nil {
			l.Logger.Errorf("[OnStreamChange] get live cache error: %v", err)
			liveModel, err = l.svcCtx.Dao.Live.WithContext(l.ctx).Where(l.svcCtx.Dao.Live.ID.Eq(uint(lid))).First()
			if err != nil {

				return nil, errors.Wrapf(xerror.ServerError, "[OnStreamChange] get live error: %v", err)
			}
		}

		liveModel.IsOver = 1
		err = l.svcCtx.Dao.Live.WithContext(l.ctx).Save(liveModel)
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[OnStreamChange] save live error: %v", err)
		}
		go func() {
			_, err := l.svcCtx.Redis.DelCtx(context.Background(), fmt.Sprintf("%s%s", constant.StreamNameToLid, in.Stream))
			if err != nil {
				l.Logger.Errorf("[OnStreamChange] del stream name error: %v", err)
			}
			err = modelcache.DelLiveModelCache(l.svcCtx.Cache, l.ctx, uint(lid))
			if err != nil {
				l.Logger.Errorf("[OnStreamChange] del live cache error: %v", err)
				return
			}

			var itempatch = client.ItemPatch{}
			var isHidden = true
			itempatch.IsHidden = &isHidden
			_, err = l.svcCtx.GorseClient.UpdateItem(context.Background(), fmt.Sprintf("%s:%d", constant.LiveType, liveModel.ID), itempatch)
			if err != nil {
				l.Logger.Errorf("[OnStreamChange] update item error: %v", err)
			}
		}()
	}
	return resp, nil
}
