package model

import (
	"context"
	"errors"
	"time"
)

type UserDetail struct {
	Id                  string
	Name                string
	Path                string
	Arn                 string
	CreateDate          time.Time
	Tags                map[string]string
	PermissionsBoundary *AttachedPermissionsBoundary

	ManagedPolicies []*AttachedPolicy
	Groups          []*AttachedGroup
	Policies        []*InlinePolicyDetail
}

var (
	UserDetailAlreadyExists = errors.New("UserDetailAlreadyExists")
	UserDetailNotFound      = errors.New("UserDetailNotFound")
)

type UserDetailModel interface {
	Insert(ctx context.Context, u *UserDetail) error
	Lookup(ctx context.Context, Id string) (*UserDetail, error)
	Dump(ctx context.Context)
}
