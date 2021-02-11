package model

import (
	"context"
	"errors"
)

type Trail struct {
	CWLogsGroupArn             *ARN
	CWLogsRoleArn              *ARN
	KmsKeyIdArn                *ARN
	SnsTopicArn                *ARN
	Arn                        *ARN
	HasCustomEventSelectors    bool
	HasInsightSelectors        bool
	IncludeGlobalServiceEvents bool
	IsMultiRegionTrail         bool
	IsOrganizationTrail        bool
	LogFileVallidationEnabled  bool
	HomeRegion                 string
	Name                       string
	S3BucketName               string
	S3KeyPrefix                string
}

var (
	TrailAlreadyExists = errors.New("TrailAlreadyExists")
	TrailNotFound      = errors.New("TrailNotFound")
)

type TrailModel interface {
	Insert(ctx context.Context, t *Trail) error
	Lookup(ctx context.Context, name string) (*Trail, error)
	Dump(ctx context.Context)
}
