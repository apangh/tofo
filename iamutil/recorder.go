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
	return r.orm.GroupDetailModel.Insert(ctx, recorder.ToGroupDetail(g))
}

type RoleDetailRecorder struct {
	orm *model.ORM
}

func (r *RoleDetailRecorder) Do(ctx context.Context, rd types.RoleDetail) error {
	return r.orm.RoleDetailModel.Insert(ctx, recorder.ToRoleDetail(rd))
}

type ManagedPolicyDetailRecorder struct {
	orm *model.ORM
}

func (r *ManagedPolicyDetailRecorder) Do(ctx context.Context, p types.ManagedPolicyDetail) error {
	return r.orm.ManagedPolicyDetailModel.Insert(ctx, recorder.ToManagedPolicyDetail(p))
}

type UserDetailRecorder struct {
	orm *model.ORM
}

func (r *UserDetailRecorder) Do(ctx context.Context, u types.UserDetail) error {
	return r.orm.UserDetailModel.Insert(ctx, recorder.ToUserDetail(u))
}

type GroupRecorder struct {
	orm *model.ORM
}

func (r *GroupRecorder) Do(ctx context.Context, group types.Group) error {
	return r.orm.GroupModel.Insert(ctx, recorder.ToGroup(group))
}

type RoleRecorder struct {
	orm *model.ORM
}

func (rr *RoleRecorder) Do(ctx context.Context, role types.Role) error {
	return rr.orm.RoleModel.Insert(ctx, recorder.ToRole(role))
}

type UserRecorder struct {
	orm *model.ORM
}

func (r *UserRecorder) Do(ctx context.Context, user types.User) error {
	return r.orm.UserModel.Insert(ctx, recorder.ToUser(user))
}
