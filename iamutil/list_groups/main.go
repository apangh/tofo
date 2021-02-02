package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/apangh/tofo/iamutil"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	logToStderr := flag.Lookup("alsologtostderr")
	if err := logToStderr.Value.Set("true"); err != nil {
		fmt.Printf("failed to setup glog: %v", err)
	}

	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"))
	if err != nil {
		glog.Errorf("failed to list groups: %s", err)
		return
	}
	client := iam.NewFromConfig(config)

	l := &iamutil.LogGroup{}

	e := iamutil.ListGroups(ctx, client, l)
	if e != nil {
		tofo.LogErr("listGroups", e)
		glog.Errorf("failed to list groups: %v", e)
		return
	}
}
