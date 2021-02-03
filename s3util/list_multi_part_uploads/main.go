package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/apangh/tofo/s3util"
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
		glog.Errorf("Failed to list multi-part uploads: %s\n", err)
		return
	}
	client := s3.NewFromConfig(config)

	bucketName := "test-bucket-46709394-abcd-1112233"

	e := s3util.ListMultiPartUploads(ctx, client, bucketName, &s3util.LogMultiPartUpload{})
	if e != nil {
		tofo.LogErr("ListMultiPartUploads", e)
		glog.Errorf("Failed to list multi-part uploads in bucket %s: %v",
			bucketName, e)
		return
	}
}
