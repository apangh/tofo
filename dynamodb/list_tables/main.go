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

	var i int

	var listEvaluatedTableName *string

	for {
		params := &dynamodb.ListTablesInput{
			Limit:                   aws.Int32(100),
			ExclusiveStartTableName: listEvaluatedTableName,
		}

		o, err := client.ListTables(ctx, params)
		if err != nil {
			tofo.LogErr("ListTables", err)
			glog.Errorf("Failed to list tables: %s\n", err)
			return
		}
		for _, tName := range o.TableNames {
			glog.Infof("Table[%d] %s", i, tName)
			i++
		}
		if o.LastEvaluatedTableName == nil {
			break
		}
		listEvaluatedTableName = o.LastEvaluatedTableName
	}

	return
}
