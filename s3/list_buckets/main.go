package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
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

	params := &s3.ListBucketsInput{}

	o, err := client.ListBuckets(ctx, params)
	if err != nil {
		tofo.LogErr("ListBuckets", err)
		glog.Errorf("Failed to list buckets: %s\n", err)
		return
	}
	for i, bucket := range o.Buckets {
		glog.Infof("Bucket[%d] %s %v", i, *bucket.Name,
			bucket.CreationDate)
	}
	glog.Infof("Owner: %s %s", *o.Owner.DisplayName, *o.Owner.ID)
	glog.Infof("Metadata: %v", o.ResultMetadata)

	return
}
