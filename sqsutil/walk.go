package sqsutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Walk(ctx context.Context, client *sqs.Client) error {
	var cb LogQueue
	return ListQueues(ctx, client, &cb)
}
