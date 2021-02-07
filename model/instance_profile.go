package model

import (
	"context"
	"errors"
	"time"
)

type InstanceProfile struct {
	Arn        *ARN
	CreateDate time.Time
	Id         string
	Name       string
	Path       string
	Roles      []*Role
}

var (
	InstanceProfileAlreadyExists = errors.New("InstanceProfileAlreadyExists")
	InstanceProfileNotFound      = errors.New("InstanceProfileNotFound")
)

type InstanceProfileModel interface {
	Insert(ctx context.Context, i *InstanceProfile) error
	Lookup(ctx context.Context, Id string) (*InstanceProfile, error)
	Dump(ctx context.Context)
}
