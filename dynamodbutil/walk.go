package dynamodbutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/golang/glog"
)

type WalkTableNameCB struct {
	i      int
	client *dynamodb.Client
}

func (w *WalkTableNameCB) Do(ctx context.Context, tableName string) error {
	var cb LogBackupSummary
	glog.Infof("Table[%d] %s", w.i, tableName)
	e := ListBackup(ctx, w.client, tableName, &cb)
	w.i++
	return e
}

func Walk(ctx context.Context, client *dynamodb.Client) error {
	cb := WalkTableNameCB{
		client: client,
	}
	return ListTables(ctx, client, &cb)
}
