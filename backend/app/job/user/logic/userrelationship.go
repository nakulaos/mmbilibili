package logic

import (
	"backend/app/job/user/conf"
	"backend/app/rpc/user/biz/model"
	"backend/library/initializer"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/executors"
	"gorm.io/gorm"
)

type UserRelationshipService struct {
	consumer *kafka.Reader
	e        *executors.BulkExecutor
	db       *gorm.DB
	c        *conf.UserRelevantCountService
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewUserRelationshipService(c *conf.UserRelevantCountService) *UserRelationshipService {
	ctx, cancel := context.WithCancel(context.Background())

	consumer := initializer.InitKafkaReader(&c.UserRelevantCountConsumer)
	db := initializer.InitGormDBFromMysql(c.UserRelevantCountMysql)

	u := &UserRelationshipService{
		db:       db,
		consumer: consumer,
		c:        c,
		ctx:      ctx,
		cancel:   cancel,
	}

	u.e = executors.NewBulkExecutor(u.Execute(),
		executors.WithBulkInterval(time.Millisecond*time.Duration(c.UserRelevantCountBulkExecutor.FlushInterval)),
		executors.WithBulkTasks(c.UserRelevantCountBulkExecutor.TaskSize),
	)

	return u
}

func (s *UserRelationshipService) Execute() func(tasks []any) {
	return func(tasks []any) {
		// 创建一个 map，用于存储每个用户的更新信息
		m := make(map[int64]*model.UserRelevantCountMessage)

		// 合并相同用户的 CountChange 增量
		for _, task := range tasks {
			msg := task.(kafka.Message)
			value := msg.Value
			key := string(msg.Key)
			parts := strings.Split(key, ":")
			if len(parts) != 2 {
				fmt.Println("Invalid key format")
				return
			}
			uid, err := strconv.Atoi(parts[1])
			if err != nil {
				klog.Errorf("UserRelationshipService.Atoi(%v) err: %v", key, err)
				continue
			}

			i64uid := int64(uid)
			userRelevantCountMsg := &model.UserRelevantCountMessage{}
			if err := userRelevantCountMsg.Unmarshal(value); err != nil {
				klog.Errorf("UserRelationshipService.Unmarshal(%v) err: %v", value, err)
				continue
			}

			if _, ok := m[i64uid]; !ok {
				m[i64uid] = userRelevantCountMsg
			} else {
				for k, v := range userRelevantCountMsg.CountChange {
					m[i64uid].CountChange[k] += v
				}
			}
		}

		for userID, userCountMsg := range m {
			um := make(map[string]interface{})
			for countType, change := range userCountMsg.CountChange {
				field := model.SwitchCountType(countType)
				if change < 0 {
					um[field] = gorm.Expr(field+" - ?", -change)
					continue
				} else if change == 0 {
					continue
				} else {
					um[field] = gorm.Expr(field+" + ?", change)
				}
			}

			maxRetries := 3
			var err error

			for attempt := 1; attempt <= maxRetries; attempt++ {
				err = s.db.Model(&model.UserRelevantCount{}).Where("user_id = ?", userID).Updates(um).Error
				if err == nil {
					break
				}

				klog.Errorf("Attempt %d to update user_id %d failed: %v", attempt, userID, err)
				time.Sleep(100 * time.Millisecond) // 等待 100ms 后重试
			}

			if err != nil {
				klog.Errorf("Failed to update user_id %d after %d attempts: %v", userID, maxRetries, err)
			}
		}
	}
}

func (s *UserRelationshipService) Start() {

	for {
		select {
		case <-s.ctx.Done():
			klog.Infof("UserRelationshipService consumer stopped by context cancellation")
			return
		default:
			m, err := s.consumer.ReadMessage(s.ctx)
			if err != nil {
				if err == context.Canceled {
					klog.Infof("UserRelationshipService consumer read canceled")
					return
				}
				klog.Errorf("UserRelationshipService.consumer.ReadMessage err: %v", err)
				break
			}

			err = s.e.Add(m)
			if err != nil {
				klog.Errorf("UserRelationshipService.bulkexecutor.add err: %v", err)
			}
		}
	}

}

func (s *UserRelationshipService) Stop() {
	s.cancel()

	if err := s.consumer.Close(); err != nil {
		klog.Errorf("Failed to close Kafka consumer: %v", err)
	}

	s.e.Flush()
	s.e.Wait()
}
