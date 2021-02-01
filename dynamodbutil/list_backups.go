package dynamodbutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BackupSummaryCB interface {
	Do(ctx context.Context, summary types.BackupSummary) error
}

func ListBackup(ctx context.Context, client *dynamodb.Client, tableName string,
	cb BackupSummaryCB) error {
	var exclusiveStartBackupArn *string
	for {
		params := &dynamodb.ListBackupsInput{
			ExclusiveStartBackupArn: exclusiveStartBackupArn,
			Limit:                   aws.Int32(100),
			TableName:               aws.String(tableName),
		}

		o, e := client.ListBackups(ctx, params)
		if e != nil {
			return e
		}

		for _, summary := range o.BackupSummaries {
			if e := cb.Do(ctx, summary); e != nil {
				return e
			}
		}

		if o.LastEvaluatedBackupArn == nil {
			return nil
		}
		exclusiveStartBackupArn = o.LastEvaluatedBackupArn
	}
}
