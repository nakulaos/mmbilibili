package initializer

import (
	"backend/library/conf"
	"backend/library/metric"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	cache "github.com/mgtv-tech/jetcache-go"
	"github.com/mgtv-tech/jetcache-go/local"
	"github.com/mgtv-tech/jetcache-go/remote"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

func InitJETCache(c conf.JETCache, dalRedisClient redis.UniversalClient) cache.Cache {

	dalUUID := uuid.New().String()
	dalPubSub := dalRedisClient.Subscribe(context.Background(), c.SyncLocalCacheName)
	dalCache := cache.New(cache.WithName(c.Name),
		cache.WithLocal(local.NewTinyLFU(1000000*c.LocalLFUSize, time.Minute*time.Duration(c.LocalLFUTTL))),
		cache.WithRemote(remote.NewGoRedisV9Adapter(dalRedisClient)),
		cache.WithRefreshDuration(time.Duration(c.RefreshDuration)*time.Minute),
		cache.WithStopRefreshAfterLastAccess(time.Duration(c.StopRefreshAfterLastAccess)*time.Minute),
		cache.WithRemoteExpiry(time.Duration(c.RemoteExpiry)*time.Minute),
		cache.WithSourceId(dalUUID),
		cache.WithSyncLocal(true),
		cache.WithNotFoundExpiry(time.Duration(c.NotFoundExpiry)*time.Minute),
		cache.WithEventHandler(func(event *cache.Event) {
			bs, _ := json.Marshal(event)
			dalRedisClient.Publish(context.Background(), c.SyncLocalCacheName, string(bs))
		}),
		cache.WithStatsHandler(metric.NewJETCacheMetric(c.Name, "")),
		cache.WithErrNotFound(gorm.ErrRecordNotFound),
	)

	go func() {
		for {
			msg := <-dalPubSub.Channel()
			var event *cache.Event
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				panic(err)
			}

			fmt.Println(event.Keys)

			if event.SourceID != dalUUID {
				for _, key := range event.Keys {
					dalCache.DeleteFromLocalCache(key)
				}
			}
		}
	}()

	return dalCache
}
