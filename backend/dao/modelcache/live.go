package modelcache

import (
	"backend/common/constant"
	"backend/dao/model"
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
)

func SetLiveModelCache(cacheImpl *cache.Cache, newLive *model.Live) {
	_ = cacheImpl.Set(&cache.Item{
		Key:   constant.LiveInfoCacheKey + fmt.Sprintf("%d", newLive.ID),
		Value: newLive,
	})
}

func GetLiveModelCache(cacheImpl *cache.Cache, ctx context.Context, liveID uint) (liveModel *model.Live, err error) {
	liveModel = &model.Live{}
	err = cacheImpl.Get(ctx, constant.LiveInfoCacheKey+fmt.Sprintf("%d", liveID), liveModel)
	return
}

func DelLiveModelCache(cacheImpl *cache.Cache, ctx context.Context, liveID uint) (err error) {
	key := constant.LiveInfoCacheKey + fmt.Sprintf("%d", liveID)
	for i := 0; i < 10; i++ {
		err = cacheImpl.Delete(ctx, key)
		if err == nil || err == cache.ErrCacheMiss {
			break
		}
	}
	return err
}
