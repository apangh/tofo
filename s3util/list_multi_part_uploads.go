package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListMultiPartUploadsCB interface {
	Do(ctx context.Context, u types.MultipartUpload) error
}

func ListMultiPartUploads(ctx context.Context, client *s3.Client, bucketName string,
	cb ListMultiPartUploadsCB) error {
	var nextKeyMarker *string
	var nextUploadIdMarker *string
	for {
		params := &s3.ListMultipartUploadsInput{
			Bucket:         aws.String(bucketName),
			MaxUploads:     1000,
			KeyMarker:      nextKeyMarker,
			UploadIdMarker: nextUploadIdMarker,
		}
		o, err := client.ListMultipartUploads(ctx, params)
		if err != nil {
			return err
		}

		for _, u := range o.Uploads {
			if e := cb.Do(ctx, u); e != nil {
				return e
			}
		}

		if !o.IsTruncated {
			return nil
		}
		nextKeyMarker = o.NextKeyMarker
		nextUploadIdMarker = o.NextUploadIdMarker
	}
}
