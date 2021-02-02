package sqsutil

import (
	"context"

	"github.com/golang/glog"
)

var _ ListQueuesCB = (*LogQueue)(nil)

type LogQueue struct {
	i int
}

func (l *LogQueue) Do(ctx context.Context, queueURL string) error {
	glog.Infof("Queue[%d] %s", l.i, queueURL)
	l.i++
	return nil
}
