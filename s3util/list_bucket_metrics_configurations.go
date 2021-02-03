package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListBucketMetricsConfigurationsCB interface {
	Do(ctx context.Context, m types.MetricsConfiguration) error
}

func ListBucketMetricsConfigurations(ctx context.Context, client *s3.Client,
	bucketName string, cb ListBucketMetricsConfigurationsCB) error {
	var nextContinuationToken *string
	for {
		params := &s3.ListBucketMetricsConfigurationsInput{
			Bucket:            aws.String(bucketName),
			ContinuationToken: nextContinuationToken,
		}
		o, err := client.ListBucketMetricsConfigurations(ctx, params)
		if err != nil {
			return err
		}

		for _, c := range o.MetricsConfigurationList {
			if e := cb.Do(ctx, c); e != nil {
				return e
			}
		}

		if !o.IsTruncated {
			return nil
		}

		nextContinuationToken = o.NextContinuationToken
	}
}
