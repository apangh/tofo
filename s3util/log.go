package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

var _ ListMultiPartUploadsCB = (*LogPartInMultiPartUpload)(nil)

type LogPartInMultiPartUpload struct {
	LogMultiPartUpload
	client *s3.Client
	bucket string
}

func (l *LogPartInMultiPartUpload) Do(ctx context.Context, o types.MultipartUpload) error {
	if e := l.LogMultiPartUpload.Do(ctx, o); e != nil {
		return e
	}
	return ListParts(ctx, l.client, l.bucket, aws.ToString(o.Key),
		aws.ToString(o.UploadId), &LogPart{})
}

type LogInventoryConfiguration struct {
	i int
}

func (l *LogInventoryConfiguration) Do(ctx context.Context,
	c types.InventoryConfiguration) error {
	glog.Infof("IC[%d] %s %v %s %v %s %s %v %v %v %v %v", l.i,
		aws.ToString(c.Destination.S3BucketDestination.Bucket),
		c.Destination.S3BucketDestination.Format,
		aws.ToString(c.Destination.S3BucketDestination.AccountId),
		c.Destination.S3BucketDestination.Encryption,
		aws.ToString(c.Destination.S3BucketDestination.Prefix),
		aws.ToString(c.Id), c.IncludedObjectVersions, c.IsEnabled, c.Schedule,
		c.Filter, c.OptionalFields)
	l.i++
	return nil
}

type LogMetricsConfiguration struct {
	i int
}

func (l *LogMetricsConfiguration) Do(ctx context.Context,
	c types.MetricsConfiguration) error {
	glog.Infof("MC[%d] %+v", l.i, c)
	l.i++
	return nil
}
