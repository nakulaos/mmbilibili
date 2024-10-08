package usercommonrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/crypt"
	"backend/pkg/jwt"
	"backend/pkg/tools"
	"backend/pkg/xerror"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhenghaoz/gorse/client"
	"gorm.io/gorm"
	"strconv"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.LoginResp, error) {
	oldUser, err := l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.Username.Eq(in.Username)).First()
	if err == nil {
		return nil, errors.Wrapf(xerror.UserExistError.WithTemplateData(map[string]interface{}{
			"UID": oldUser.Username,
		}), "[Register] username exist: %v", in.Username)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerror.ServerError, "[Register] get user failed: %v", err)
	}

	userSalt, _ := tools.GenerateSalt()

	// 获取room_id
	roomID, e := GetIncrRoomID(l.ctx, l.svcCtx.Redis)
	if e != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[Register] get incr room id failed: %v", e)
	}
	newUser := &model.User{
		Username: in.Username,
		Password: crypt.PasswordEncrypt(userSalt, l.svcCtx.Config.App.Salt, in.Password),
		Salt:     userSalt,
		RoomID:   roomID,
	}

	// 数据库
	q := l.svcCtx.Dao.Begin()
	err = q.User.WithContext(l.ctx).Create(newUser)
	if err != nil {
		err := q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[Register] rollback failed: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[Register] create user failed: %v", err)
	}

	_, err = l.svcCtx.GorseClient.InsertUser(l.ctx, client.User{
		UserId: strconv.Itoa(int(newUser.ID)),
	})

	if err != nil {
		err := q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[Register] rollback failed: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[Register] insert user to gorse failed: %v", err)
	}

	err = statistics.InitializeUserRelevantCount(l.svcCtx.Redis, l.ctx, int(newUser.ID))
	if err != nil {
		err := q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[Register] rollback failed: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[Register] initialize user relevant count failed: %v", err)
	}

	err = q.Commit()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[Register] commit failed: %v", err)
	}

	l.Logger.Infof("[mysql] success to insert a user called by %s ", in.Username)

	var resp user.LoginResp
	resp.UserInfo = &user.User{
		Id:       uint32(newUser.ID),
		Username: newUser.Username,
		Nickname: newUser.Nickname,
		Avatar:   newUser.Avatar,
		Gender:   uint32(newUser.Gender),
		Role:     uint32(1),
		Phone:    newUser.Phone,
		Email:    newUser.Email,
	}
	resp.UserId = int64(newUser.ID)

	go func() {
		// 缓存用户信息
		modelcache.SetUserModelCache(l.svcCtx.Cache, newUser)

	}()

	// 获取token
	token, _ := jwt.MakeToken(int(newUser.ID), 1, l.svcCtx.Config.Jwt)
	resp.AccessToken = token

	return &resp, nil
}

func GetIncrRoomID(ctx context.Context, redisClient *redis.Redis) (uint, error) {
	ret, e := redisClient.ExistsCtx(ctx, constant.IncrRoomID)
	if e != nil {
		return 0, errors.Wrapf(xerror.ServerError, "[GetIncrRoomID] get incr room id failed: %v", e)
	}
	if !ret {
		// 不存在
		err := redisClient.SetCtx(ctx, constant.IncrRoomID, strconv.Itoa(constant.IncrRoomIDInitValue))
		if err != nil {
			return 0, errors.Wrapf(xerror.ServerError, "[GetIncrRoomID] set incr room id failed: %v", e)
		}
	}

	// 获取room_id
	roomID, e := redisClient.IncrCtx(ctx, constant.IncrRoomID)
	if e != nil {
		return 0, errors.Wrapf(xerror.ServerError, "[GetIncrRoomID] incr room id failed: %v", e)
	}
	return uint(roomID), nil
}
