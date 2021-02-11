package tofo

import (
	"context"

	"github.com/apangh/tofo/cloudtrailutil"
	"github.com/apangh/tofo/dynamodbutil"
	"github.com/apangh/tofo/iamutil"
	"github.com/apangh/tofo/model"
	"github.com/apangh/tofo/s3util"
	"github.com/apangh/tofo/sqsutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Walk(ctx context.Context, cfg aws.Config, orm *model.ORM) error {
	if e := iamutil.Walk(ctx, iam.NewFromConfig(cfg), orm); e != nil {
		return e
	}
	if e := s3util.Walk(ctx, s3.NewFromConfig(cfg), orm); e != nil {
		return e
	}
	cfg.Region = "us-west-2"
	if e := cloudtrailutil.Walk(ctx, cloudtrail.NewFromConfig(cfg), orm); e != nil {
		return e
	}
	if e := dynamodbutil.Walk(ctx, dynamodb.NewFromConfig(cfg), orm); e != nil {
		return e
	}
	if e := sqsutil.Walk(ctx, sqs.NewFromConfig(cfg), orm); e != nil {
		return e
	}
	return nil
}
