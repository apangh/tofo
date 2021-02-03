package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/glog"
)

func Walk(ctx context.Context, client *s3.Client) error {
	params := &s3.ListBucketsInput{}

	o, e := client.ListBuckets(ctx, params)
	if e != nil {
		return e
	}
	glog.Infof("Owner: %s %s", aws.ToString(o.Owner.DisplayName), aws.ToString(o.Owner.ID))
	glog.Infof("Metadata: %v", o.ResultMetadata)
	for i, bucket := range o.Buckets {
		bucketName := aws.ToString(bucket.Name)
		glog.Infof("Bucket[%d] %s %v", i, bucketName, bucket.CreationDate)

		e := ListBucketInventoryConfigurations(ctx, client, bucketName,
			&LogInventoryConfiguration{})
		if e != nil {
			return e
		}

		e = ListBucketMetricsConfigurations(ctx, client, bucketName,
			&LogMetricsConfiguration{})
		if e != nil {
			return e
		}

		// bucket logging
		l, e := client.GetBucketLogging(ctx,
			&s3.GetBucketLoggingInput{
				Bucket: bucket.Name,
			})
		if e != nil {
			return e
		}
		if o := l.LoggingEnabled; o != nil {
			glog.Infof("Logging: TargetBucket: %s", aws.ToString(o.TargetBucket))
			glog.Infof("Logging: TargetPrefix: %s", aws.ToString(o.TargetPrefix))
			for i, t := range o.TargetGrants {
				glog.Infof("[%d]%+v", i, t)
			}
		}

		// bucket versioning
		v, e := client.GetBucketVersioning(ctx,
			&s3.GetBucketVersioningInput{
				Bucket: bucket.Name,
			})
		if e != nil {
			return e
		}
		if v.Status == "" {
			// Bucket has no versioning enabled
			e = ListObjects(ctx, client, bucketName, &LogObject{})
			if e != nil {
				return e
			}
		} else {
			glog.Infof("Bucket version status: %s", v.Status)
			e = ListObjectVersions(ctx, client, bucketName,
				&LogObjectVersion{})
			if e != nil {
				return e
			}
		}

		// multi-parts
		cb := LogPartInMultiPartUpload{
			client: client,
			bucket: bucketName,
		}
		e = ListMultiPartUploads(ctx, client, bucketName, &cb)
		if e != nil {
			return e
		}
	}

	return nil
}
