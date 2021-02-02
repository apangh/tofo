package sqsutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type ListQueuesCB interface {
	Do(ctx context.Context, queueURL string) error
}

func ListQueues(ctx context.Context, client *sqs.Client, cb ListQueuesCB) error {
	var nextToken *string

	for {
		params := &sqs.ListQueuesInput{
			MaxResults: aws.Int32(1000),
			NextToken:  nextToken,
		}
		o, e := client.ListQueues(ctx, params)
		if e != nil {
			return e
		}

		for _, queueUrl := range o.QueueUrls {
			if e := cb.Do(ctx, queueUrl); e != nil {
				return e
			}
		}

		if o.NextToken == nil {
			return nil
		}
		nextToken = o.NextToken
	}
}
