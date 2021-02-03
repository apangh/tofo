package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/glog"
)

var _ ListObjectVersionsCB = (*LogObjectVersion)(nil)

type LogObjectVersion struct {
	i int
}

func (l *LogObjectVersion) Do(ctx context.Context, v types.ObjectVersion) error {
	glog.Infof("[%d] ObjectVersion: %s %s %v %v %s %s %d %v %s", l.i,
		aws.ToString(v.Key), aws.ToString(v.ETag), v.IsLatest, v.LastModified,
		aws.ToString(v.Owner.DisplayName), aws.ToString(v.Owner.ID), v.Size,
		v.StorageClass, aws.ToString(v.VersionId))
	l.i++
	return nil
}
