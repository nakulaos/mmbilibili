package livecallbackrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/pkg/xerror"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"strconv"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnPublicStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnPublicStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnPublicStreamLogic {
	return &OnPublicStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnPublicStreamLogic) OnPublicStream(in *live.OnPublicStreamReq) (*live.OnPublicStreamResp, error) {
	// 检查请求参数是否为空
	var resp = &live.OnPublicStreamResp{}
	if in.Params == "" {
		resp.Code = 1
		resp.Msg = "Params is empty"
		return resp, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"params": "params"}), "[OnPublicStream] Params is empty")
	}

	// 解析请求参数
	values, err := url.ParseQuery(in.Params)
	if err != nil {
		resp.Code = 1
		resp.Msg = "Params parse error"
		return resp, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"params": "params"}), "[OnPublicStream] Params parse error: %v", err)
	}

	// 获取 token 和 lid
	token := values.Get("token")
	if token == "" {
		resp.Code = 1
		resp.Msg = "Token is empty"
		return resp, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"params": "token"}), "[OnPublicStream] Token is empty")
	}

	lid := values.Get("lid")
	if lid == "" {
		resp.Code = 1
		resp.Msg = "Lid is empty"
		return resp, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"params": "lid"}), "[OnPublicStream] Lid is empty")
	}

	// 获取用户信息，首先从缓存中获取
	var liveModel = &model.Live{}
	lidi, e := strconv.Atoi(lid)
	liveModel, err = modelcache.GetLiveModelCache(l.svcCtx.Cache, l.ctx, uint(lidi))

	// 如果缓存获取失败，则从数据库中获取
	if err != nil {

		if e != nil {
			resp.Code = 1
			resp.Msg = "Lid parse error"
			return resp, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"params": "lid"}), "[OnPublicStream] Lid parse error: %v", e)
		}
		liveModel, err = l.svcCtx.Dao.Live.WithContext(l.ctx).Where(l.svcCtx.Dao.Live.ID.Eq(uint(lidi))).First()
		if err != nil {
			resp.Code = 1
			resp.Msg = "Get live error"
			return resp, errors.Wrapf(xerror.ServerError, "[OnPublicStream] get live error: %v", err)
		}
	}

	// 检查直播状态
	if liveModel.IsOver == 1 {
		resp.Code = 1
		resp.Msg = "Live is over"
		return resp, errors.Wrapf(xerror.LiveIsOverError, "[OnPublicStream] live is over")
	}

	// 验证 token
	if liveModel.PushToken != token {
		resp.Code = 1
		resp.Msg = "Token is invalid"
		return resp, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"params": "token"}), "[OnPublicStream] Token is invalid")
	}

	// 返回成功响应
	// 缓存流名字
	go func() {
		err = l.svcCtx.Redis.SetexCtx(context.Background(), fmt.Sprintf("%s%s", constant.StreamNameToLid, in.Stream), lid, 0)
		if err != nil {
			l.Logger.Errorf("[OnPublicStream] cache stream name error: %v", err)
		}
	}()
	return resp, nil
}
