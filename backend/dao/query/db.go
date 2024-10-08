package query

import (
	"context"
	"gorm.io/gorm"
)

func (q *Query) DB() *gorm.DB {
	return q.db.WithContext(context.Background())
}
