package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/apangh/tofo/dynamodbutil"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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
		glog.Errorf("Failed to list dynamodb tables: %s\n", e)
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "Hello"

	var cb dynamodbutil.LogBackupSummary

	if e := dynamodbutil.ListBackup(ctx, client, tableName, &cb); e != nil {
		tofo.LogErr("ListBackups", e)
		glog.Errorf("Failed to list backups: %s", e)
		return
	}
}
