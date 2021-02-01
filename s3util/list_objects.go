package s3util

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/glog"
)

func ListObjects(ctx context.Context, client *s3.Client, bucketName string) error {
	params := &s3.ListObjectsV2Input{
		Bucket:     aws.String(bucketName),
		FetchOwner: true,
		MaxKeys:    1000,
	}

	var i int
	p := s3.NewListObjectsV2Paginator(client, params)
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, obj := range page.Contents {
			glog.Infof("[%d] Object: %s, %s, %v, %s, %s, %d, %v", i,
				aws.ToString(obj.Key), aws.ToString(obj.ETag),
				obj.LastModified, aws.ToString(obj.Owner.DisplayName),
				aws.ToString(obj.Owner.ID), obj.Size,
				obj.StorageClass)
			i++
		}
	}
	return nil
}
