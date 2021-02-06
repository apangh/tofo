package model

import (
	"context"
	"errors"
	"time"
)

type AttachedGroup struct {
	Name string
}

type GroupDetail struct {
	Id         string
	Name       string
	Path       string
	Arn        string
	CreateDate time.Time

	ManagedPolicies []*AttachedPolicy
	Policies        []*InlinePolicyDetail
}

var (
	GroupDetailAlreadyExists = errors.New("GroupDetailAlreadyExists")
	GroupDetailNotFound      = errors.New("GroupDetailNotFound")
)

type GroupDetailModel interface {
	Insert(ctx context.Context, u *GroupDetail) error
	Lookup(ctx context.Context, Id string) (*GroupDetail, error)
	Dump(ctx context.Context)
}
