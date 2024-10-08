package statistics

import (
	"backend/common/constant"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type LiveRelevantCount struct {
	LikeCount    int
	ViewCount    int
	CommentCount int
}

func GetLiveRelevantCount(redisClient *redis.Redis, ctx context.Context, lid int) (LiveRelevantCount, error) {
	luaScript := `
    local lid = ARGV[1]
    local likeCount = redis.call("GET", KEYS[1]..lid)
    local viewCount = redis.call("GET", KEYS[2]..lid)
    local commentCount = redis.call("GET", KEYS[3]..lid)
    return {likeCount, viewCount, commentCount}
    `

	keys := []string{
		constant.AppLiveLikeCount,    // 直播点赞次数
		constant.AppLiveViewCount,    // 直播观看次数
		constant.AppLiveCommentCount, // 直播评论次数
	}

	result, err := redisClient.EvalCtx(ctx, luaScript, keys, lid)
	if err != nil {
		return LiveRelevantCount{}, err
	}

	counts, ok := result.([]interface{})
	if !ok || len(counts) != 3 {
		return LiveRelevantCount{}, fmt.Errorf("unexpected result format")
	}

	return LiveRelevantCount{
		LikeCount:    parseInt(counts[0]),
		ViewCount:    parseInt(counts[1]),
		CommentCount: parseInt(counts[2]),
	}, nil
}

func InitializeLiveRelevantCount(redisClient *redis.Redis, ctx context.Context, lid int) error {
	luaScript := `
    local lid = ARGV[1]
    for i = 1, #KEYS do
        local key = KEYS[i] .. lid
        if redis.call("EXISTS", key) == 0 then
            redis.call("SET", key, 0)
        end
    end
    return "nil"
    `

	keys := []string{
		constant.AppLiveLikeCount,    // 直播点赞次数
		constant.AppLiveViewCount,    // 直播观看次数
		constant.AppLiveCommentCount, // 直播评论次数
	}

	_, err := redisClient.EvalCtx(ctx, luaScript, keys, lid)
	if err != nil {
		return fmt.Errorf("failed to initialize live relevant counts: %w", err)
	}
	return nil
}

func GetLivesRelevantCount(redisClient *redis.Redis, ctx context.Context, lids []uint) (map[uint]LiveRelevantCount, error) {
	luaScript := `
    local lidCount = ARGV[1]  -- 直播数量
    local results = {}
    
    for i = 1, lidCount do
        local lid = ARGV[i + 1]  -- 直播 ID
        local likeCount = redis.call("GET", KEYS[1]..lid)
        local viewCount = redis.call("GET", KEYS[2]..lid)
        local commentCount = redis.call("GET", KEYS[3]..lid)

        results[i] = {
            likeCount, viewCount, commentCount
        }
    end
    return results
    `

	keys := []string{
		constant.AppLiveLikeCount,    // 直播点赞次数
		constant.AppLiveViewCount,    // 直播观看次数
		constant.AppLiveCommentCount, // 直播评论次数
	}

	// Prepare ARGV for Lua script
	args := make([]interface{}, len(lids)+1)
	args[0] = len(lids)
	for i, lid := range lids {
		args[i+1] = lid
	}

	result, err := redisClient.EvalCtx(ctx, luaScript, keys, args...)
	if err != nil {
		return nil, err
	}

	counts, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result format")
	}

	liveCounts := make(map[uint]LiveRelevantCount)

	for i, count := range counts {
		if countValues, ok := count.([]interface{}); ok && len(countValues) == 3 {
			lid := lids[i]
			liveCounts[lid] = LiveRelevantCount{
				LikeCount:    parseInt(countValues[0]),
				ViewCount:    parseInt(countValues[1]),
				CommentCount: parseInt(countValues[2]),
			}
		}
	}

	return liveCounts, nil
}

func UpdateLiveLikeCount(redisClient *redis.Redis, ctx context.Context, lid int, action int) error {
	luaScript := `
    local likeKey = KEYS[1]  -- 直播的点赞数键

    -- 增加或减少点赞数
    if ARGV[1] == "INCR" then
        redis.call("INCR", likeKey)  -- 增加点赞数
    elseif ARGV[1] == "DECR" then
        redis.call("DECR", likeKey)  -- 减少点赞数
    end
    return "nil"
    `

	likeKey := constant.AppLiveLikeCount + fmt.Sprint(lid)

	var actionCommand string
	if action == constant.LikeAction {
		actionCommand = "INCR" // 点赞操作
	} else if action == constant.UnLikeAction {
		actionCommand = "DECR" // 取消点赞操作
	} else {
		return fmt.Errorf("invalid action: %d", action) // 无效操作
	}

	_, err := redisClient.EvalCtx(ctx, luaScript, []string{likeKey}, actionCommand)
	if err != nil {
		return fmt.Errorf("failed to update live like count: %w", err)
	}
	return nil
}

func UpdateLiveCommentCount(redisClient *redis.Redis, ctx context.Context, lid int, action int) error {
	luaScript := `
    local commentKey = KEYS[1]  -- 直播的评论数键

    -- 增加或减少评论数
    if ARGV[1] == "INCR" then
        redis.call("INCR", commentKey)  -- 增加评论数
    elseif ARGV[1] == "DECR" then
        redis.call("DECR", commentKey)  -- 减少评论数
    end
    return "nil"
    `

	commentKey := constant.AppLiveCommentCount + fmt.Sprint(lid)

	var actionCommand string
	if action == constant.CommentAction {
		actionCommand = "INCR" // 添加评论操作
	} else if action == constant.UnCommentAction {
		actionCommand = "DECR" // 删除评论操作
	} else {
		return fmt.Errorf("invalid action: %d", action) // 无效操作
	}

	_, err := redisClient.EvalCtx(ctx, luaScript, []string{commentKey}, actionCommand)
	if err != nil {
		return fmt.Errorf("failed to update live comment count: %w", err)
	}
	return nil
}
