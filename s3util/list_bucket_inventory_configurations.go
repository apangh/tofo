package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListBucketInventoryConfigurationsCB interface {
	Do(ctx context.Context, c types.InventoryConfiguration) error
}

func ListBucketInventoryConfigurations(ctx context.Context, client *s3.Client,
	bucketName string, cb ListBucketInventoryConfigurationsCB) error {
	var nextContinuationToken *string
	for {
		params := &s3.ListBucketInventoryConfigurationsInput{
			Bucket:            aws.String(bucketName),
			ContinuationToken: nextContinuationToken,
		}
		o, err := client.ListBucketInventoryConfigurations(ctx, params)
		if err != nil {
			return err
		}

		for _, c := range o.InventoryConfigurationList {
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
