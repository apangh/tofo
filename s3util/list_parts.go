package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListPartsCB interface {
	Do(ctx context.Context, p types.Part) error
}

func ListParts(ctx context.Context, client *s3.Client, bucketName, key, UploadId string,
	cb ListPartsCB) error {
	params := &s3.ListPartsInput{
		Bucket:   aws.String(bucketName),
		Key:      aws.String(key),
		UploadId: aws.String(UploadId),
		MaxParts: 1000,
	}

	p := s3.NewListPartsPaginator(client, params)
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, p := range page.Parts {
			if e := cb.Do(ctx, p); e != nil {
				return e
			}
		}
	}
	return nil
}
