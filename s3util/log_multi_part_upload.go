package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/glog"
)

var _ ListMultiPartUploadsCB = (*LogMultiPartUpload)(nil)

type LogMultiPartUpload struct {
	i int
}

func (l *LogMultiPartUpload) Do(ctx context.Context, o types.MultipartUpload) error {
	glog.Infof("[%d] MultiPartUpload: %v %s %s %s %s %s %v %s", l.i,
		o.Initiated, aws.ToString(o.Initiator.DisplayName),
		aws.ToString(o.Initiator.ID), aws.ToString(o.Key),
		aws.ToString(o.Owner.DisplayName), aws.ToString(o.Owner.ID),
		o.StorageClass, aws.ToString(o.UploadId))
	l.i++
	return nil
}
