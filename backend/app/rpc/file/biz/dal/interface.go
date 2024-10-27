package dal

import (
	"backend/app/rpc/file/biz/model"
	"context"
)

type FileDal interface {
	CreateFileChunk(ctx context.Context, fileChunk *model.FileChunk) error
	GetFileChunkByFileHashANDUserID(ctx context.Context, hash string, userID int64) (*model.FileChunk, error)
	SaveFileChunk(ctx context.Context, fileChunk *model.FileChunk) error
}
