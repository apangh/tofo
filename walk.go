package tofo

import (
	"context"

	"github.com/apangh/tofo/dynamodbutil"
	"github.com/apangh/tofo/s3util"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Walk(ctx context.Context, cfg aws.Config) error {
	if e := dynamodbutil.Walk(ctx, dynamodb.NewFromConfig(cfg)); e != nil {
		return e
	}
	if e := s3util.Walk(ctx, s3.NewFromConfig(cfg)); e != nil {
		return e
	}
	return nil
}
