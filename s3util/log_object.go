package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/glog"
)

var _ ListObjectsCB = (*LogObject)(nil)

type LogObject struct {
	i int
}

func (l *LogObject) Do(ctx context.Context, o types.Object) error {
	glog.Infof("[%d] Object: %s, %s, %v, %s, %s, %d, %v", l.i,
		aws.ToString(o.Key), aws.ToString(o.ETag), o.LastModified,
		aws.ToString(o.Owner.DisplayName), aws.ToString(o.Owner.ID), o.Size,
		o.StorageClass)
	l.i++
	return nil
}
