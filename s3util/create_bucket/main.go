package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	logToStderr := flag.Lookup("alsologtostderr")
	if err := logToStderr.Value.Set("true"); err != nil {
		fmt.Printf("Failed to setup glog: %v", err)
	}

	bucketName := "test-bucket-46709394-abcd-xyz"

	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"),
		config.WithRegion("us-west-2"))
	if err != nil {
		glog.Errorf("Failed to create bucket %s: %s\n", bucketName, err)
		return
	}
	client := s3.NewFromConfig(config, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintUsWest2,
		},
	}

	o, err := client.CreateBucket(ctx, params)
	if err != nil {
		tofo.LogErr("ListBuckets", err)
		glog.Errorf("Failed to create bucket: %s: %s\n", bucketName, err)
		return
	}
	glog.Infof("Created bucket %s: %v\n", bucketName, o)

	return
}
