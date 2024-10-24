package dal

import (
	"backend/app/common/constant"
	"context"
	"errors"
	"fmt"
	cache "github.com/mgtv-tech/jetcache-go"
	"strconv"
)

func userFromUidKey(id int64) string {
	return fmt.Sprintf("%s:%d", constant.UserDetailCacheFromUidKey, id)
}

func userFromUserNameKey(name string) string {
	return fmt.Sprintf("%s:%s", constant.UserDetailCacheFromUsernameKey, name)
}

func userRelevantCountFromIdKey(id int64) string {
	return fmt.Sprintf("%s:%d", constant.UserRelevantCountFromIdKey, id)
}

func setCacheUsernameToUid(c cache.Cache, ctx context.Context, name string, id int64) error {
	obj := fmt.Sprintf("%d", id)
	err := c.Set(ctx, userFromUserNameKey(name), cache.Value(obj))
	return err
}

func delCacheUsernameToUid(c cache.Cache, ctx context.Context, name string) error {
	return c.Delete(ctx, userFromUserNameKey(name))
}

func getCacheUsernameToUid(c cache.Cache, ctx context.Context, name string) (int64, error) {
	var id string
	if err := c.Get(ctx, userFromUserNameKey(name), &id); err != nil && !errors.Is(err, cache.ErrCacheMiss) {
		return 0, err
	}

	if id == "" {
		return -1, nil
	}

	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, err
	}

	return iid, nil
}
