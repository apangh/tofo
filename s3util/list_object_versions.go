package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListObjectVersionsCB interface {
	Do(ctx context.Context, v types.ObjectVersion) error
}

func ListObjectVersions(ctx context.Context, client *s3.Client, bucketName string,
	cb ListObjectVersionsCB) error {

	var nextKeyMarker *string
	var nextVersionIdMarker *string

	for {
		params := &s3.ListObjectVersionsInput{
			Bucket:          aws.String(bucketName),
			MaxKeys:         1000,
			KeyMarker:       nextKeyMarker,
			VersionIdMarker: nextVersionIdMarker,
		}

		o, err := client.ListObjectVersions(ctx, params)
		if err != nil {
			return err
		}

		for _, v := range o.Versions {
			if e := cb.Do(ctx, v); e != nil {
				return e
			}
		}

		if !o.IsTruncated {
			return nil
		}
		nextKeyMarker = o.NextKeyMarker
		nextVersionIdMarker = o.NextVersionIdMarker
	}
}
