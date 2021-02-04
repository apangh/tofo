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
	if e := logToStderr.Value.Set("true"); e != nil {
		fmt.Printf("Failed to setup glog: %v\n", e)
		return
	}

	bucketName := "test-bucket-46709394-abcd-xyz"

	ctx := context.Background()
	cfg, e := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"),
		config.WithRegion("us-west-2"))
	if e != nil {
		glog.Errorf("Failed to create bucket %s: %s", bucketName, e)
		return
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintUsWest2,
		},
	}

	o, e := client.CreateBucket(ctx, params)
	if e != nil {
		tofo.LogErr("ListBuckets", e)
		glog.Errorf("Failed to create bucket: %s: %s", bucketName, e)
		return
	}
	glog.Infof("Created bucket %s: %v", bucketName, o)
}
