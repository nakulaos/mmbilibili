package modelcache

import (
	"backend/common/constant"
	"backend/dao/model"
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	"time"
)

func SetCategoryModelCache(cacheImpl *cache.Cache, partition, category string, newCategory *model.Category) {
	_ = cacheImpl.Set(&cache.Item{
		Key:   fmt.Sprintf("%s%s:%s", constant.CategoryInfoNameKey, partition, category),
		Value: newCategory,
		TTL:   time.Minute * 30,
	})
	_ = cacheImpl.Set(&cache.Item{
		Key:   fmt.Sprintf("%s%d", constant.CategoryInfoCidKey, newCategory.ID),
		Value: newCategory,
		TTL:   time.Minute * 30,
	})
}

func GetCategoryFromNameAndPartition(cacheImpl *cache.Cache, ctx context.Context, partition, category string) (categoryModel *model.Category, err error) {
	categoryModel = &model.Category{}
	err = cacheImpl.Get(ctx, fmt.Sprintf("%s%s:%s", constant.CategoryInfoNameKey, partition, category), categoryModel)
	return
}

func DelCategoryModelCacheFromPartitionAndCategory(cacheImpl *cache.Cache, ctx context.Context, partition, category string) (err error) {
	key := fmt.Sprintf("%s%s:%s", constant.CategoryInfoNameKey, partition, category)
	for i := 0; i < 10; i++ {
		err = cacheImpl.Delete(ctx, key)
		if err == nil || err == cache.ErrCacheMiss {
			break
		}
	}
	return err
}

func DelCategoryModelCacheFromId(cacheImpl *cache.Cache, ctx context.Context, cid uint) (err error) {
	key := fmt.Sprintf("%s%d", constant.CategoryInfoCidKey, cid)
	for i := 0; i < 10; i++ {
		err = cacheImpl.Delete(ctx, key)
		if err == nil || err == cache.ErrCacheMiss {
			break
		}
	}
	return err
}
