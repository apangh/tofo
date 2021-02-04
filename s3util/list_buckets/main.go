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
	if e := logToStderr.Value.Set("true"); e != nil {
		fmt.Printf("Failed to setup glog: %v\n", e)
		return
	}

	ctx := context.Background()
	cfg, e := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"))
	if e != nil {
		glog.Errorf("Failed to list buckets: %s", e)
		return
	}
	client := s3.NewFromConfig(cfg)

	params := &s3.ListBucketsInput{}

	o, e := client.ListBuckets(ctx, params)
	if e != nil {
		tofo.LogErr("ListBuckets", e)
		glog.Errorf("Failed to list buckets: %s", e)
		return
	}
	for i, bucket := range o.Buckets {
		glog.Infof("Bucket[%d] %s %v", i, aws.ToString(bucket.Name),
			bucket.CreationDate)
	}
	glog.Infof("Owner: %s %s", aws.ToString(o.Owner.DisplayName),
		aws.ToString(o.Owner.ID))
	glog.Infof("Metadata: %v", o.ResultMetadata)
}
