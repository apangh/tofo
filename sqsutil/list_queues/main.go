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
	if err := logToStderr.Value.Set("true"); err != nil {
		fmt.Printf("Failed to setup glog: %v", err)
	}

	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"))
	if err != nil {
		glog.Errorf("Failed to list dynamodb tables: %s", err)
		return
	}
	client := sqs.NewFromConfig(config)

	var cb sqsutil.LogQueue

	e := sqsutil.ListQueues(ctx, client, &cb)
	if e != nil {
		tofo.LogErr("ListQueues", err)
		glog.Errorf("Failed to list queues: %s", err)
		return
	}
	return
}
