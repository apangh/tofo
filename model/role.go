package model

import (
	"context"
	"errors"
	"time"
)

type RoleLastUsed struct {
	Date   time.Time
	Region string
}

type Role struct {
	Id                  string
	Name                string
	Path                string
	Arn                 string
	CreateDate          time.Time
	Tags                map[string]string
	PermissionsBoundary *AttachedPermissionsBoundary

	AssumeRolePolicyDocument string
	Description              string
	MaxSessionDuration       *int32
	LastUsed                 *RoleLastUsed
}

var (
	RoleAlreadyExists = errors.New("RoleAlreadyExists")
	RoleNotFound      = errors.New("RoleNotFound")
)

type RoleModel interface {
	Insert(ctx context.Context, r *Role) error
	Lookup(ctx context.Context, Id string) (*Role, error)
	Dump(ctx context.Context)
}
