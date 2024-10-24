package service

import (
	"backend/app/common/constant"
	"backend/app/common/ecode"
	"backend/app/rpc/user/biz/dal"
	"backend/app/rpc/user/biz/model"
	user "backend/app/rpc/user/kitex_gen/user"
	"backend/library/metric"
	"backend/library/tools"
	"context"
	"golang.org/x/sync/errgroup"
	"strconv"
)

type UpdateUserInfoService struct {
	ctx context.Context
} // NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	/*
		1. 判断用户是否存在
		2. 更新用户信息
		3. 并行获取用户信息和用户统计信息
	*/

	userModel := buildUserModel(req)

	f, err := dal.UserDalInstance.ExistUserByID(s.ctx, userModel.ID)
	if err != nil {
		return nil, ecode.ServerError
	}
	if !f {
		metric.IncrGauge(metric.BusinessError, constant.PromUserNotExistError)
		return nil, ecode.UserNotExistError.WithTemplateData(map[string]string{"UID": strconv.Itoa(int(req.Id))})
	}

	err = dal.UserDalInstance.UpdateUserByID(s.ctx, userModel.ID, userModel)
	if err != nil {
		return nil, ecode.ServerError
	}

	var (
		eg        errgroup.Group
		userInfo  *model.User
		userStats *model.UserRelevantCount
	)

	eg.Go(func() error {
		var e error
		userInfo, e = dal.UserDalInstance.GetUserByID(s.ctx, userModel.ID)
		return e
	})
	eg.Go(func() error {
		var e error
		userStats, e = dal.UserDalInstance.GetUserRelevantCountByID(s.ctx, userModel.ID)
		return e
	})

	err = eg.Wait()
	if err != nil {
		return nil, ecode.ServerError
	}

	resp = &user.UpdateUserInfoResp{}

	resp.UserInfo = &user.User{
		Id:             userInfo.ID,
		Username:       userInfo.Username,
		Role:           uint32(userInfo.Role),
		Avatar:         userInfo.Avatar,
		Nickname:       userInfo.Nickname,
		Gender:         uint32(userInfo.Gender),
		Phone:          tools.MaskPhone(userInfo.Phone),
		Email:          tools.MaskEmail(userInfo.Email),
		FollowerCount:  userStats.FollowerCount,
		FollowingCount: userStats.FollowingCount,
		LikeCount:      userStats.LikeCount,
		StarCount:      userStats.StarCount,
		SelfStarCount:  userStats.SelfStarCount,
		SelfLikeCount:  userStats.SelfLikeCount,
		LiveCount:      userStats.LiveCount,
		WorkCount:      userStats.WorkCount,
		FriendCount:    userStats.FriendCount,
	}

	return
}

func buildUserModel(req *user.UpdateUserInfoReq) *model.User {
	user := &model.User{}
	user.ID = int64(req.Id)

	if req.Nickname != nil {
		user.Nickname = *req.Nickname
	}

	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}

	if req.Gender != nil {
		user.Gender = uint8(*req.Gender)
	}

	if req.Role != nil {
		user.Role = uint8(*req.Role)
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	return user
}
