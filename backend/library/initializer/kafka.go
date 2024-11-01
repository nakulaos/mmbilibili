package initializer

import (
	globalConf "backend/library/conf"
	"github.com/segmentio/kafka-go"
	"time"
)

func InitKafkaWriter(c *globalConf.KafkaWriter) *kafka.Writer {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(c.Brokers...),
		Topic:                  c.Topic,
		Balancer:               SwitchBalancer(c.EnableLocalOrder),
		BatchSize:              c.BatchSize,
		BatchBytes:             int64(c.BatchBytes),
		BatchTimeout:           time.Duration(c.BatchTimeout) * time.Millisecond,
		ReadTimeout:            time.Duration(c.ReadTimeout) * time.Millisecond,
		WriteTimeout:           time.Duration(c.WriteTimeout) * time.Millisecond,
		Async:                  c.EnableAsync,
		RequiredAcks:           kafka.RequiredAcks(c.RequiredAcks),
		AllowAutoTopicCreation: c.AllowAutoCreateTopic,
		MaxAttempts:            c.MaxAttempts,
	}
	return writer
}

func SwitchBalancer(enableOrder bool) kafka.Balancer {
	if enableOrder {
		return &kafka.Hash{}
	} else {
		return &kafka.LeastBytes{}
	}
}
