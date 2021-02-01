package dynamodbutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type TableNameCB interface {
	Do(ctx context.Context, tableName string) error
}

func ListTables(ctx context.Context, client *dynamodb.Client,
	cb TableNameCB) error {
	var listEvaluatedTableName *string

	for {
		params := &dynamodb.ListTablesInput{
			Limit:                   aws.Int32(100),
			ExclusiveStartTableName: listEvaluatedTableName,
		}

		o, e := client.ListTables(ctx, params)
		if e != nil {
			return e
		}
		for _, tName := range o.TableNames {
			if e := cb.Do(ctx, tName); e != nil {
				return e
			}
		}
		if o.LastEvaluatedTableName == nil {
			return nil
		}
		listEvaluatedTableName = o.LastEvaluatedTableName
	}
}
