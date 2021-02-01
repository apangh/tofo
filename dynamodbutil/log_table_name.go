package dynamodbutil

import (
	"context"

	"github.com/golang/glog"
)

type LogTableName struct {
	i int
}

func (l *LogTableName) Do(ctx context.Context, tableName string) error {
	glog.Infof("Table[%d] %s", l.i, tableName)
	l.i++
	return nil
}

var _ TableNameCB = (*LogTableName)(nil)
