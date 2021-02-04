package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/apangh/tofo/model/mem"
	"github.com/aws/aws-sdk-go-v2/config"
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
		glog.Errorf("Failed to walk: %s", e)
		return
	}

	orm, e := mem.NewORM()
	if e != nil {
		glog.Errorf("Failed to setup orm: %s\n", e)
		return
	}

	if e := tofo.Walk(ctx, cfg, orm); e != nil {
		tofo.LogErr("Walk", e)
		glog.Errorf("Failed to walk: %s", e)
		return
	}
}
