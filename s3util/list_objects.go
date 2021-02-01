package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListObjectsCB interface {
	Do(ctx context.Context, o types.Object) error
}

func ListObjects(ctx context.Context, client *s3.Client, bucketName string,
	cb ListObjectsCB) error {
	params := &s3.ListObjectsV2Input{
		Bucket:     aws.String(bucketName),
		FetchOwner: true,
		MaxKeys:    1000,
	}

	p := s3.NewListObjectsV2Paginator(client, params)
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, obj := range page.Contents {
			if e := cb.Do(ctx, obj); e != nil {
				return e
			}
		}
	}
	return nil
}
