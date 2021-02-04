package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/apangh/tofo/sqsutil"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
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
		glog.Errorf("Failed to list dynamodb tables: %s", e)
		return
	}
	client := sqs.NewFromConfig(cfg)

	var cb sqsutil.LogQueue

	if e := sqsutil.ListQueues(ctx, client, &cb); e != nil {
		tofo.LogErr("ListQueues", e)
		glog.Errorf("Failed to list queues: %s", e)
		return
	}
}
