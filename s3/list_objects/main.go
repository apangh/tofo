package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	logToStderr := flag.Lookup("alsologtostderr")
	if err := logToStderr.Value.Set("true"); err != nil {
		fmt.Printf("Failed to setup glog: %v", err)
	}

	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"))
	if err != nil {
		glog.Errorf("Failed to list buckets: %s\n", err)
		return
	}
	client := s3.NewFromConfig(config)

	bucketName := "test-bucket-46709394-abcd-1112233"

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
			tofo.LogErr("ListObjectsV2", err)
			glog.Errorf("Failed to list objects in bucket %s: %v",
				bucketName, err)
			return
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
}
