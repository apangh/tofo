package iamutil

import (
	"context"

	"github.com/apangh/tofo/iamutil/recorder"
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type GroupDetailRecorder struct {
	orm *model.ORM
}

func (r *GroupDetailRecorder) Do(ctx context.Context, g types.GroupDetail) error {
	p, e := recorder.ToGroupDetail(g)
	if e != nil {
		return e
	}
	return r.orm.GroupDetailModel.Insert(ctx, p)
}

type RoleDetailRecorder struct {
	orm *model.ORM
}

func (r *RoleDetailRecorder) Do(ctx context.Context, rd types.RoleDetail) error {
	d, e := recorder.ToRoleDetail(rd)
	if e != nil {
		return e
	}
	return r.orm.RoleDetailModel.Insert(ctx, d)
}

type ManagedPolicyDetailRecorder struct {
	orm *model.ORM
}

func (r *ManagedPolicyDetailRecorder) Do(ctx context.Context, p types.ManagedPolicyDetail) error {
	o, e := recorder.ToManagedPolicyDetail(p)
	if e != nil {
		return e
	}
	return r.orm.ManagedPolicyDetailModel.Insert(ctx, o)
}

type UserDetailRecorder struct {
	orm *model.ORM
}

func (r *UserDetailRecorder) Do(ctx context.Context, u types.UserDetail) error {
	d, e := recorder.ToUserDetail(u)
	if e != nil {
		return e
	}
	return r.orm.UserDetailModel.Insert(ctx, d)
}

type GroupRecorder struct {
	orm *model.ORM
}

func (r *GroupRecorder) Do(ctx context.Context, group types.Group) error {
	g, e := recorder.ToGroup(group)
	if e != nil {
		return e
	}
	return r.orm.GroupModel.Insert(ctx, g)
}

type RoleRecorder struct {
	orm *model.ORM
}

func (rr *RoleRecorder) Do(ctx context.Context, role types.Role) error {
	r, e := recorder.ToRole(role)
	if e != nil {
		return e
	}
	return rr.orm.RoleModel.Insert(ctx, r)
}

type UserRecorder struct {
	orm *model.ORM
}

func (r *UserRecorder) Do(ctx context.Context, user types.User) error {
	u, e := recorder.ToUser(user)
	if e != nil {
		return e
	}
	return r.orm.UserModel.Insert(ctx, u)
}
