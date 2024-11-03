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

func InitKafkaReader(c *globalConf.KafkaReader) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          c.Brokers,
		GroupID:          c.GroupID,
		Topic:            c.Topic,
		StartOffset:      c.StartOffset,
		MinBytes:         c.MinBytes,
		MaxBytes:         c.MaxBytes,
		MaxWait:          time.Duration(c.MaxWaitTime) * time.Millisecond,
		CommitInterval:   time.Duration(c.CommitInterval) * time.Millisecond,
		ReadBatchTimeout: time.Duration(c.ReadBatchTimeout) * time.Millisecond,
	})
	return reader
}
