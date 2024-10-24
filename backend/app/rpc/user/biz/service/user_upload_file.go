package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type UserUploadFileService struct {
	ctx context.Context
} // NewUserUploadFileService new UserUploadFileService
func NewUserUploadFileService(ctx context.Context) *UserUploadFileService {
	return &UserUploadFileService{ctx: ctx}
}

// Run create note info
func (s *UserUploadFileService) Run(req *user.UserUploadFileReq) (resp *user.UserUploadFileResp, err error) {
	// Finish your business logic.

	return
}
