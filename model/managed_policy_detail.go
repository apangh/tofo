package model

import (
	"context"
	"errors"
	"time"
)

type ManagedPolicyDetail struct {
	Id         string
	Name       string
	Path       string
	Arn        *ARN
	CreateDate time.Time

	AttachmentCount               int32
	PermissionsBoundaryUsageCount int32
	DefaultVersionId              string
	Description                   string
	IsAttachable                  bool
	UpdateDate                    time.Time

	Versions map[string]*PolicyVersion
}

type PolicyVersion struct {
	IsDefaultVersion bool
	CreateDate       time.Time
	Document         *JsonPolicyDocument
	VersionId        string
}

var (
	ManagedPolicyDetailAlreadyExists = errors.New("ManagedPolicyDetailAlreadyExists")
	ManagedPolicyDetailNotFound      = errors.New("ManagedPolicyDetailNotFound")
)

type ManagedPolicyDetailModel interface {
	Insert(ctx context.Context, p *ManagedPolicyDetail) error
	Lookup(ctx context.Context, Id string) (*ManagedPolicyDetail, error)
	LookupByArn(ctx context.Context, arn string) (*ManagedPolicyDetail, error)
	Dump(ctx context.Context)
}
