package sqsutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Walk(ctx context.Context, client *sqs.Client, orm *model.ORM) error {
	var cb LogQueue
	return ListQueues(ctx, client, &cb)
}
