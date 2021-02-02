package tofo

import (
	"context"

	"github.com/apangh/tofo/dynamodbutil"
	"github.com/apangh/tofo/iamutil"
	"github.com/apangh/tofo/s3util"
	"github.com/apangh/tofo/sqsutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Walk(ctx context.Context, cfg aws.Config) error {
	if e := s3util.Walk(ctx, s3.NewFromConfig(cfg)); e != nil {
		return e
	}
	cfg.Region = "us-west-2"
	if e := dynamodbutil.Walk(ctx, dynamodb.NewFromConfig(cfg)); e != nil {
		return e
	}
	if e := sqsutil.Walk(ctx, sqs.NewFromConfig(cfg)); e != nil {
		return e
	}
	if e := iamutil.Walk(ctx, iam.NewFromConfig(cfg)); e != nil {
		return e
	}
	return nil
}
