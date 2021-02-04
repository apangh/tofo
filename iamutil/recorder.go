package iamutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type TagsRecorder struct {
	tags map[string]string
}

func (r *TagsRecorder) Do(ctx context.Context, t types.Tag) error {
	if r.tags == nil {
		r.tags = make(map[string]string)
	}

	r.tags[*t.Key] = *t.Value
	return nil
}

type GroupRecorder struct {
	orm *model.ORM
}

func (r *GroupRecorder) ToGroup(ctx context.Context, group types.Group) (
	*model.Group, error) {
	g := &model.Group{
		Id:         aws.ToString(group.GroupId),
		Name:       aws.ToString(group.GroupName),
		Path:       aws.ToString(group.Path),
		Arn:        aws.ToString(group.Arn),
		CreateDate: aws.ToTime(group.CreateDate),
	}
	return g, nil
}

func (r *GroupRecorder) Do(ctx context.Context, group types.Group) error {
	g, e := r.ToGroup(ctx, group)
	if e != nil {
		return e
	}
	return r.orm.GroupModel.Insert(ctx, g)
}

type RoleRecorder struct {
	orm    *model.ORM
	client *iam.Client
}

func (rr *RoleRecorder) ToRole(ctx context.Context, role types.Role) (
	*model.Role, error) {
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

	// ListRoles does not return tags, ned to retreive separately
	var tr TagsRecorder
	if e := ListRoleTags(ctx, rr.client, r.Name, &tr); e != nil {
		return nil, e
	}
	r.Tags = tr.tags

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
	orm    *model.ORM
	client *iam.Client
}

func (r *UserRecorder) toUser(ctx context.Context, user types.User) (*model.User, error) {
	u := &model.User{
		Id:               aws.ToString(user.UserId),
		Name:             aws.ToString(user.UserName),
		Path:             aws.ToString(user.Path),
		Arn:              aws.ToString(user.Arn),
		CreateDate:       user.CreateDate,
		PasswordLastUsed: user.PasswordLastUsed,
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

	// ListUsers does not return tags, need to retrieve that separately
	var tr TagsRecorder
	if e := ListUserTags(ctx, r.client, u.Name, &tr); e != nil {
		return nil, e
	}
	u.Tags = tr.tags

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
