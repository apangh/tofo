package iamutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type RoleRecorder struct {
	orm *model.ORM
}

func (rr *RoleRecorder) ToRole(ctx context.Context, role types.Role) (*model.Role, error) {
	r := &model.Role{
		Id:                       aws.ToString(role.RoleId),
		Name:                     aws.ToString(role.RoleName),
		Path:                     aws.ToString(role.Path),
		Arn:                      aws.ToString(role.Arn),
		CreateDate:               role.CreateDate,
		AssumeRolePolicyDocument: aws.ToString(role.AssumeRolePolicyDocument),
		Description:              aws.ToString(role.Description),
		MaxSessionDuration:       role.MaxSessionDuration,
	}
	if l := role.RoleLastUsed; l != nil {
		r.LastUsed = &model.LastUsed{
			Date:   l.LastUsedDate,
			Region: aws.ToString(l.Region),
		}
	}
	// TODO - there is a known issue that ListRole does not return PermissionsBoundary
	if pb := role.PermissionsBoundary; pb != nil {
		p, e := rr.orm.PolicyModel.LookupByArn(ctx,
			aws.ToString(pb.PermissionsBoundaryArn))
		if e != nil {
			return nil, e
		}
		r.PermissionBoundary = p
	}
	for _, t := range role.Tags {
		r.Tags[*t.Key] = *t.Value
	}
	return r, nil
}

func (rr *RoleRecorder) Do(ctx context.Context, role types.Role) error {
	r, e := rr.ToRole(ctx, role)
	if e != nil {
		return e
	}
	return rr.orm.RoleModel.Insert(ctx, r)
}

type UserRecorder struct {
	orm *model.ORM
}

func (r *UserRecorder) toUser(ctx context.Context, user types.User) (*model.User, error) {
	u := &model.User{
		Id:               aws.ToString(user.UserId),
		Name:             aws.ToString(user.UserName),
		Path:             aws.ToString(user.Path),
		Arn:              aws.ToString(user.Arn),
		CreateDate:       user.CreateDate,
		PasswordLastUsed: user.PasswordLastUsed,
		Tags:             make(map[string]string),
	}
	// TODO - there is a known issue that ListUser does not return PermissionsBoundary
	if pb := user.PermissionsBoundary; pb != nil {
		p, e := r.orm.PolicyModel.LookupByArn(ctx,
			aws.ToString(pb.PermissionsBoundaryArn))
		if e != nil {
			return nil, e
		}
		u.PermissionBoundary = p
	}
	for _, t := range user.Tags {
		u.Tags[*t.Key] = *t.Value
	}
	return u, nil
}

func (r *UserRecorder) Do(ctx context.Context, user types.User) error {
	u, e := r.toUser(ctx, user)
	if e != nil {
		return e
	}
	return r.orm.UserModel.Insert(ctx, u)
}

type PolicyRecorder struct {
	orm *model.ORM
}

func toPolicy(p types.Policy) *model.Policy {
	return &model.Policy{
		Id:                            aws.ToString(p.PolicyId),
		Name:                          aws.ToString(p.PolicyName),
		Path:                          aws.ToString(p.Path),
		Arn:                           aws.ToString(p.Arn),
		AttachmentCount:               aws.ToInt32(p.AttachmentCount),
		PermissionsBoundaryUsageCount: aws.ToInt32(p.PermissionsBoundaryUsageCount),
		DefaultVersionId:              aws.ToString(p.DefaultVersionId),
		Description:                   aws.ToString(p.Description),
		IsAttachable:                  p.IsAttachable,
		CreateDate:                    aws.ToTime(p.CreateDate),
		UpdateDate:                    aws.ToTime(p.UpdateDate),
	}
}

func (r *PolicyRecorder) Do(ctx context.Context, policy types.Policy) error {
	p := toPolicy(policy)
	return r.orm.PolicyModel.Insert(ctx, p)
}
