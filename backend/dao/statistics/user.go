package statistics

import (
	"backend/common/constant"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
)

type UserRelevantCount struct {
	FollowerCount  int
	FollowingCount int
	FriendCount    int
	LikeCount      int
	StarCount      int
	WorkCount      int
	SelfStarCount  int
	SelfLikeCount  int
	LiveCount      int
}

func GetUserRelevantCount(redisClient *redis.Redis, ctx context.Context, uid int) (UserRelevantCount, error) {
	luaScript := `
    local uid = ARGV[1]
    local followerCount = redis.call("GET", KEYS[1]..uid)
    local followingCount = redis.call("GET", KEYS[2]..uid)
    local friendCount = redis.call("GET", KEYS[3]..uid)
    local likeCount = redis.call("GET", KEYS[4]..uid)
    local starCount = redis.call("GET", KEYS[5]..uid)
    local workCount = redis.call("GET", KEYS[6]..uid)
    local selfStarCount = redis.call("GET", KEYS[7]..uid)
    local selfLikeCount = redis.call("GET", KEYS[8]..uid)
    local liveCount = redis.call("GET", KEYS[9]..uid)
    return {followerCount, followingCount, friendCount, likeCount, starCount, workCount, selfStarCount, selfLikeCount, liveCount}
    `

	keys := []string{
		constant.AppUserFollowerCount,
		constant.AppUserFollowingCount,
		constant.AppUserFriendCount,
		constant.AppUserLikeCount,
		constant.AppUserStarCount,
		constant.AppUserWorkCount,
		constant.AppUserSelfStarCount,
		constant.AppUserSelfLikeCount,
		constant.AppUserLiveCount,
	}

	result, err := redisClient.EvalCtx(ctx, luaScript, keys, uid)
	if err != nil {
		return UserRelevantCount{}, err
	}

	counts, ok := result.([]interface{})
	if !ok || len(counts) != 9 {
		return UserRelevantCount{}, fmt.Errorf("unexpected result format")
	}

	return UserRelevantCount{
		FollowerCount:  parseInt(counts[0]),
		FollowingCount: parseInt(counts[1]),
		FriendCount:    parseInt(counts[2]),
		LikeCount:      parseInt(counts[3]),
		StarCount:      parseInt(counts[4]),
		WorkCount:      parseInt(counts[5]),
		SelfStarCount:  parseInt(counts[6]),
		SelfLikeCount:  parseInt(counts[7]),
		LiveCount:      parseInt(counts[8]),
	}, nil
}

func parseInt(value interface{}) int {
	if value == nil {
		return 0
	}
	strVal, ok := value.(string)
	if !ok {
		return 0
	}
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0
	}
	return intVal
}

func InitializeUserRelevantCount(redisClient *redis.Redis, ctx context.Context, uid int) error {
	luaScript := `
    local uid = ARGV[1]
    for i = 1, #KEYS do
        local key = KEYS[i] .. uid
        if redis.call("EXISTS", key) == 0 then
            redis.call("SET", key, 0)
        end
    end
    return "nil"
    `
	keys := []string{
		constant.AppUserFollowerCount,
		constant.AppUserFollowingCount,
		constant.AppUserFriendCount,
		constant.AppUserLikeCount,
		constant.AppUserStarCount,
		constant.AppUserWorkCount,
		constant.AppUserSelfStarCount,
		constant.AppUserSelfLikeCount,
		constant.AppUserLiveCount,
	}

	_, err := redisClient.EvalCtx(ctx, luaScript, keys, uid)
	if err != nil {
		return fmt.Errorf("failed to initialize user relevant counts: %w", err)
	}
	return nil
}

func IncrUserRelevantCount(redisClient *redis.Redis, ctx context.Context, uid int, key string) error {
	luaScript := `
	local uid = ARGV[1]
	local key = ARGV[2]
	local count = redis.call("INCR", key..uid)
	return count
	`
	_, err := redisClient.EvalCtx(ctx, luaScript, []string{key}, uid)
	if err != nil {
		return fmt.Errorf("failed to increment user relevant count: %w", err)
	}
	return nil
}

func DecrUserRelevantCount(redisClient *redis.Redis, ctx context.Context, uid int, key string) error {
	luaScript := `
	local uid = ARGV[1]
	local key = ARGV[2]
	local count = redis.call("DECR", key..uid)
	return count
	`
	_, err := redisClient.EvalCtx(ctx, luaScript, []string{key}, uid)
	if err != nil {
		return fmt.Errorf("failed to decrement user relevant count: %w", err)
	}
	return nil
}

func UpdateUserFollowCounts(redisClient *redis.Redis, ctx context.Context, followerID, followedID int, action int) error {
	luaScript := `
	local followerKey = KEYS[1]  -- 粉丝数键
	local followedKey = KEYS[2]    -- 关注数键

	-- 增加或减少粉丝数和关注数
	if ARGV[1] == "INCR" then
		redis.call("INCR", followedKey)  -- 增加被关注者的粉丝数
		redis.call("INCR", followerKey)   -- 增加关注者的关注数
	elseif ARGV[1] == "DECR" then
		redis.call("DECR", followedKey)  -- 减少被关注者的粉丝数
		redis.call("DECR", followerKey)   -- 减少关注者的关注数
	end
    return "nil"
	`

	followerKey := constant.AppUserFollowerCount + fmt.Sprint(followedID)  // 粉丝数
	followedKey := constant.AppUserFollowingCount + fmt.Sprint(followerID) // 关注数

	var actionCommand string
	if action == constant.FollowUserAction {
		actionCommand = "INCR" // 关注操作
	} else if action == constant.UnFollowUserAction {
		actionCommand = "DECR" // 取消关注操作
	} else {
		return fmt.Errorf("invalid action: %d", action) // 无效操作
	}

	_, err := redisClient.EvalCtx(ctx, luaScript, []string{followerKey, followedKey}, actionCommand)
	if err != nil {
		return fmt.Errorf("failed to update user follow counts: %w", err)
	}
	return nil
}

func GetUsersRelevantCount(redisClient *redis.Redis, ctx context.Context, uids []uint) (map[uint]UserRelevantCount, error) {
	luaScript := `
    local uidCount = ARGV[1]  -- 用户数量
    local results = {}
    
    for i = 1, uidCount do
        local uid = ARGV[i + 1]  -- 用户 ID
        local followerCount = redis.call("GET", KEYS[1]..uid)
        local followingCount = redis.call("GET", KEYS[2]..uid)
        local friendCount = redis.call("GET", KEYS[3]..uid)
        local likeCount = redis.call("GET", KEYS[4]..uid)
        local starCount = redis.call("GET", KEYS[5]..uid)
        local workCount = redis.call("GET", KEYS[6]..uid)
        local selfStarCount = redis.call("GET", KEYS[7]..uid)
        local selfLikeCount = redis.call("GET", KEYS[8]..uid)
        local liveCount = redis.call("GET", KEYS[9]..uid)

        results[i] = {
            followerCount, followingCount, friendCount,
            likeCount, starCount, workCount,
            selfStarCount, selfLikeCount, liveCount
        }
    end
    return results
    `

	keys := []string{
		constant.AppUserFollowerCount,
		constant.AppUserFollowingCount,
		constant.AppUserFriendCount,
		constant.AppUserLikeCount,
		constant.AppUserStarCount,
		constant.AppUserWorkCount,
		constant.AppUserSelfStarCount,
		constant.AppUserSelfLikeCount,
		constant.AppUserLiveCount,
	}

	// Prepare ARGV for Lua script
	args := make([]interface{}, len(uids)+1)
	args[0] = len(uids)
	for i, uid := range uids {
		args[i+1] = uid
	}

	result, err := redisClient.EvalCtx(ctx, luaScript, keys, args...)
	if err != nil {
		return nil, err
	}

	counts, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result format")
	}

	userCounts := make(map[uint]UserRelevantCount)

	for i, count := range counts {
		if countValues, ok := count.([]interface{}); ok && len(countValues) == 9 {
			uid := uids[i]
			userCounts[uid] = UserRelevantCount{
				FollowerCount:  parseInt(countValues[0]),
				FollowingCount: parseInt(countValues[1]),
				FriendCount:    parseInt(countValues[2]),
				LikeCount:      parseInt(countValues[3]),
				StarCount:      parseInt(countValues[4]),
				WorkCount:      parseInt(countValues[5]),
				SelfStarCount:  parseInt(countValues[6]),
				SelfLikeCount:  parseInt(countValues[7]),
				LiveCount:      parseInt(countValues[8]),
			}
		}
	}

	return userCounts, nil
}
