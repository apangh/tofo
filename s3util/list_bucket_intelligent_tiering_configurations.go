package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListBucketIntelligentTieringConfigurationsCB interface {
	Do(ctx context.Context, c types.IntelligentTieringConfiguration) error
}

func ListBucketIntelligentTieringConfigurations(ctx context.Context, client *s3.Client,
	bucketName string, cb ListBucketIntelligentTieringConfigurationsCB) error {
	var nextContinuationToken *string
	for {
		params := &s3.ListBucketIntelligentTieringConfigurationsInput{
			Bucket:            aws.String(bucketName),
			ContinuationToken: nextContinuationToken,
		}
		o, err := client.ListBucketIntelligentTieringConfigurations(ctx, params)
		if err != nil {
			return err
		}

		for _, c := range o.IntelligentTieringConfigurationList {
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
