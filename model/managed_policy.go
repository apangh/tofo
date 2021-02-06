package model

import (
	"context"
	"errors"
	"time"
)

type ManagedPolicy struct {
	Id         string
	Name       string
	Path       string
	Arn        string
	CreateDate time.Time

	AttachmentCount               int32
	PermissionsBoundaryUsageCount int32
	DefaultVersionId              string
	Description                   string
	IsAttachable                  bool
	UpdateDate                    time.Time
}

var (
	ManagedPolicyAlreadyExists = errors.New("ManagedPolicyAlreadyExists")
	ManagedPolicyNotFound      = errors.New("ManagedPolicyNotFound")
)

type ManagedPolicyModel interface {
	Insert(ctx context.Context, p *ManagedPolicy) error
	Lookup(ctx context.Context, id string) (*ManagedPolicy, error)
	LookupByArn(ctx context.Context, arn string) (*ManagedPolicy, error)
	Dump(ctx context.Context)
}
