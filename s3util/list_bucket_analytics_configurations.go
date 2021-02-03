package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListBucketAnalyticsConfigurationsCB interface {
	Do(ctx context.Context, c types.AnalyticsConfiguration) error
}

func ListBucketAnalyticsConfigurations(ctx context.Context, client *s3.Client,
	bucketName string, cb ListBucketAnalyticsConfigurationsCB) error {
	var nextContinuationToken *string
	for {
		params := &s3.ListBucketAnalyticsConfigurationsInput{
			Bucket:            aws.String(bucketName),
			ContinuationToken: nextContinuationToken,
		}
		o, err := client.ListBucketAnalyticsConfigurations(ctx, params)
		if err != nil {
			return err
		}

		for _, c := range o.AnalyticsConfigurationList {
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
