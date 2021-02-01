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
	for i, bucket := range o.Buckets {
		glog.Infof("Bucket[%d] %s %v", i, aws.ToString(bucket.Name),
			bucket.CreationDate)

		l := &LogObject{}

		e := ListObjects(ctx, client, aws.ToString(bucket.Name), l)
		if e != nil {
			return e
		}
	}
	glog.Infof("Owner: %s %s", aws.ToString(o.Owner.DisplayName),
		aws.ToString(o.Owner.ID))
	glog.Infof("Metadata: %v", o.ResultMetadata)

	return nil
}
