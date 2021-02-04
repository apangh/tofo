package model

import (
	"context"
	"errors"
	"time"
)

type Policy struct {
	Id                            string
	Name                          string
	Path                          string
	Arn                           string
	AttachmentCount               int32
	PermissionsBoundaryUsageCount int32
	DefaultVersionId              string
	Description                   string
	IsAttachable                  bool
	CreateDate                    time.Time
	UpdateDate                    time.Time
}

var (
	PolicyAlreadyExists = errors.New("PolicyAlreadyExists")
	PolicyNotFound      = errors.New("PolicyNotFound")
)

type PolicyModel interface {
	Insert(ctx context.Context, p *Policy) error
	Lookup(ctx context.Context, id string) (*Policy, error)
	LookupByArn(ctx context.Context, arn string) (*Policy, error)
	Dump(ctx context.Context)
}
