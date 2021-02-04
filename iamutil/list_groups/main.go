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
	if e := logToStderr.Value.Set("true"); e != nil {
		fmt.Printf("failed to setup glog: %v\n", e)
		return
	}

	ctx := context.Background()
	cfg, e := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"))
	if e != nil {
		glog.Errorf("failed to list groups: %s", e)
		return
	}
	client := iam.NewFromConfig(cfg)

	l := &iamutil.LogGroup{}

	if e := iamutil.ListGroups(ctx, client, l); e != nil {
		tofo.LogErr("listGroups", e)
		glog.Errorf("failed to list groups: %v", e)
		return
	}
}
