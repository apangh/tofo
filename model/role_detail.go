package model

import (
	"context"
	"errors"
	"time"
)

type AttachedPermissionsBoundary struct {
	Arn *string

	// This is an optimization to cache the pointer after the first reference
	Policy *ManagedPolicyDetail
}

type AttachedPolicy struct {
	Arn  *string
	Name *string

	// This is an optimization to cache the pointer after the first reference
	Policy *ManagedPolicyDetail
}

type RoleDetail struct {
	Id                  string
	Name                string
	Path                string
	Arn                 string
	CreateDate          time.Time
	Tags                map[string]string
	PermissionsBoundary *AttachedPermissionsBoundary

	AssumeRolePolicyDocument JsonPolicyDocument
	Description              string
	MaxSessionDuration       *int32
	LastUsed                 *RoleLastUsed

	ManagedPolicies  []*AttachedPolicy
	InstanceProfiles []*InstanceProfile
	Policies         []*InlinePolicyDetail
}

var (
	RoleDetailAlreadyExists = errors.New("RoleDetailAlreadyExists")
	RoleDetailNotFound      = errors.New("RoleDetailNotFound")
)

type RoleDetailModel interface {
	Insert(ctx context.Context, r *RoleDetail) error
	Lookup(ctx context.Context, Id string) (*RoleDetail, error)
	Dump(ctx context.Context)
}
