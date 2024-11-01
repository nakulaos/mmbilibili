package dal

import (
	"backend/app/common/constant"
	"backend/app/rpc/user/biz/model"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	cache "github.com/mgtv-tech/jetcache-go"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
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

func userRelationshipFromMidKey(mid, shardId int64) string {
	return fmt.Sprintf("%s:%d:%d", constant.UserRelationshipKey, mid, shardId)
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

func setUserRelationshipCache(c redis.UniversalClient, ctx context.Context, mid int64, relations map[int64]*model.UserRelationship, maxCount int64, expire time.Duration) (err error) {
	// hash 分片
	pipe := c.Pipeline()
	shards := maxCount / constant.Shards
	keys := make([]string, 0, shards)
	for i := int64(0); i < shards; i++ {
		keys[i] = userRelationshipFromMidKey(mid, i)
	}

	pipe.Del(ctx, keys...)
	for _, v := range relations {
		shard := v.RelatedUserID % shards
		key := keys[shard]
		var ef []byte
		if ef, err = v.Marshal(); err != nil {
			klog.Errorf("setUserRelationshipCache.marshal user relationship failed, err: %v", err)
			return err
		}
		pipe.HSet(ctx, key, fmt.Sprintf("%d", v.RelatedUserID), ef)
	}

	for i := int64(0); i < shards; i++ {
		pipe.Expire(ctx, keys[i], expire)
	}

	if cmds, err := pipe.Exec(ctx); err != nil {
		klog.Errorf("setUserRelationshipCache.exec failed, err: %v", err)
		for i, cmd := range cmds {
			if cmd.Err() != nil {
				klog.Errorf("setUserRelationshipCache.exec failed, index: %d, err: %v", i, cmd.Err())
			}
		}
		return err
	}
	return nil
}

func addUserRelationshipCache(c redis.UniversalClient, ctx context.Context, mid int64, relations []*model.UserRelationship, maxCount int64, expire time.Duration) error {
	shardCount := maxCount / constant.Shards
	keys := make([]string, shardCount)
	for i := int64(0); i < shardCount; i++ {
		keys[i] = userRelationshipFromMidKey(mid, i)
	}

	pipe := c.Pipeline()

	for i := int64(0); i < shardCount; i++ {
		pipe.Expire(ctx, keys[i], expire)
	}

	for _, relation := range relations {
		shard := relation.RelatedUserID % shardCount
		key := keys[shard]

		ef, err := relation.Marshal()
		if err != nil {
			klog.Errorf("addUserRelationshipCache.marshal user relationship failed, err: %v", err)
			return err
		}

		pipe.HSet(ctx, key, fmt.Sprintf("%d", relation.RelatedUserID), ef)
	}

	if cmds, err := pipe.Exec(ctx); err != nil {
		klog.Errorf("setUserRelationshipCache.exec failed, err: %v", err)
		for i, cmd := range cmds {
			if cmd.Err() != nil {
				klog.Errorf("setUserRelationshipCache.exec failed, index: %d, err: %v", i, cmd.Err())
			}
		}
		return err
	}

	return nil
}

func delUserRelationshipCache(c redis.UniversalClient, ctx context.Context, mid int64, relatedUserID []int64, maxCount int64, expire time.Duration) (err error) {

	shardCount := maxCount / constant.Shards
	keys := make([]string, shardCount)
	for i := int64(0); i < shardCount; i++ {
		keys[i] = userRelationshipFromMidKey(mid, i)
	}

	pipe := c.Pipeline()

	for i := int64(0); i < shardCount; i++ {
		pipe.Expire(ctx, keys[i], expire)
	}

	for _, v := range relatedUserID {
		shard := v % shardCount
		key := keys[shard]
		pipe.HDel(ctx, key, fmt.Sprintf("%d", v))
	}

	if cmds, err := pipe.Exec(ctx); err != nil {
		for i, cmd := range cmds {
			if cmd.Err() != nil {
				klog.Errorf("delUserRelationshipCache.exec failed, index: %d, err: %v", i, cmd.Err())
			}
		}
		klog.Errorf("delUserRelationshipCache.exec failed, err: %v", err)
		return err
	}

	return nil
}

func getUserRelationshipCache(c redis.UniversalClient, ctx context.Context, mid int64, relatedUserIDs []int64, maxCount int64) (map[int64]*model.UserRelationship, error) {
	shardCount := maxCount / constant.Shards
	keys := make([]string, shardCount)
	for i := int64(0); i < shardCount; i++ {
		keys[i] = userRelationshipFromMidKey(mid, i)
	}

	pipe := c.Pipeline()
	results := make([]*redis.StringCmd, len(relatedUserIDs))

	// 按照分片哈希获取每个 relatedUserID 的关系数据
	for i, relatedUserID := range relatedUserIDs {
		shard := relatedUserID % shardCount
		key := keys[shard]
		results[i] = pipe.HGet(ctx, key, fmt.Sprintf("%d", relatedUserID))
	}

	// 执行管道命令
	_, err := pipe.Exec(ctx)
	if err != nil {
		klog.Errorf("getUserRelationshipCache.exec failed, err: %v", err)
		return nil, err
	}

	// 解析结果，返回map类型
	relationships := make(map[int64]*model.UserRelationship, len(relatedUserIDs))
	for i, cmd := range results {
		if cmd.Err() == redis.Nil {
			// relatedUserID 不存在
			klog.Infof("getUserRelationshipCache: relatedUserID %d does not exist", relatedUserIDs[i])
			continue
		} else if cmd.Err() != nil {
			// 其他错误
			klog.Errorf("getUserRelationshipCache.hget failed for relatedUserID %d, err: %v", relatedUserIDs[i], cmd.Err())
			continue
		}

		// 解析存在的关系数据
		var relation model.UserRelationship
		if err := relation.Unmarshal([]byte(cmd.Val())); err != nil {
			klog.Errorf("getUserRelationshipCache.unmarshal user relationship failed for relatedUserID %d, err: %v", relatedUserIDs[i], err)
			continue
		}

		relationships[relatedUserIDs[i]] = &relation
	}

	return relationships, nil
}

func existUserRelationshipCache(c redis.UniversalClient, ctx context.Context, mid int64, maxCount int64, expire time.Duration) (bool, error) {
	shardCount := maxCount / constant.Shards
	keys := make([]string, shardCount)
	for i := int64(0); i < shardCount; i++ {
		keys[i] = userRelationshipFromMidKey(mid, i)
	}

	for i := int64(0); i < shardCount; i++ {
		if exist, err := c.Exists(ctx, keys[i]).Result(); err != nil {
			klog.Errorf("existUserRelationshipCache.exists failed, err: %v", err)
			return false, err
		} else if exist == 0 {
			return false, nil
		}

		c.Expire(ctx, keys[i], expire)
	}

	return true, nil

}

func delAllUserRelationshipCache(c redis.UniversalClient, ctx context.Context, mid int64, maxCount int64) error {
	shardCount := maxCount / constant.Shards
	keys := make([]string, shardCount)
	for i := int64(0); i < shardCount; i++ {
		keys[i] = userRelationshipFromMidKey(mid, i)
	}

	pipe := c.Pipeline()
	for i := int64(0); i < shardCount; i++ {
		pipe.Del(ctx, keys[i])
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
