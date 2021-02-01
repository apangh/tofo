package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apangh/tofo"
	"github.com/aws/aws-sdk-go-v2/aws"
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

	var i int
	var exclusiveStartBackupArn *string
	for {
		params := &dynamodb.ListBackupsInput{
			ExclusiveStartBackupArn: exclusiveStartBackupArn,
			Limit:                   aws.Int32(100),
			TableName:               aws.String(tableName),
		}

		o, err := client.ListBackups(ctx, params)
		if err != nil {
			tofo.LogErr("ListBackups", err)
			glog.Errorf("Failed to list backups: %s", err)
			return
		}

		for _, summary := range o.BackupSummaries {
			glog.Infof("Backup[%d] %s %s %v %v %d %v %v", i,
				aws.ToString(summary.BackupArn),
				aws.ToString(summary.BackupName),
				summary.BackupCreationDateTime,
				summary.BackupExpiryDateTime,
				aws.ToInt64(summary.BackupSizeBytes),
				summary.BackupStatus,
				summary.BackupType)
			i++
		}

		if o.LastEvaluatedBackupArn == nil {
			break
		}
		exclusiveStartBackupArn = o.LastEvaluatedBackupArn
	}

	return
}
