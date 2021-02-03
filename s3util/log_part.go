package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/glog"
)

var _ ListPartsCB = (*LogPart)(nil)

type LogPart struct {
	i int
}

func (l *LogPart) Do(ctx context.Context, o types.Part) error {
	glog.Infof("[%d] Part: %s %v %d %d", l.i,
		aws.ToString(o.ETag), o.LastModified, o.PartNumber, o.Size)
	l.i++
	return nil
}
