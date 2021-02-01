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
	if err := logToStderr.Value.Set("true"); err != nil {
		fmt.Printf("Failed to setup glog: %v", err)
	}

	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("administrator"))
	if err != nil {
		glog.Errorf("Failed to list dynamodb tables: %s\n", err)
		return
	}
	client := dynamodb.NewFromConfig(config)
	tableName := "Hello"

	var cb dynamodbutil.LogBackupSummary

	e := dynamodbutil.ListBackup(ctx, client, tableName, &cb)
	if e != nil {
		tofo.LogErr("ListBackups", err)
		glog.Errorf("Failed to list backups: %s", err)
		return
	}
	return
}
