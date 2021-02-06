package iamutil

import (
	"context"

	"github.com/apangh/tofo/iamutil/recorder"
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type GroupRecorder struct {
	orm *model.ORM
}

func (r *GroupRecorder) Do(ctx context.Context, group types.Group) error {
	g := recorder.ToGroup(group)
	return r.orm.GroupModel.Insert(ctx, g)
}

type RoleRecorder struct {
	orm    *model.ORM
	client *iam.Client
}

func (rr *RoleRecorder) Do(ctx context.Context, role types.Role) error {
	return rr.orm.RoleModel.Insert(ctx, recorder.ToRole(role))
}

type RoleRecorderForListRoles struct {
	RoleRecorder
}

func (r *RoleRecorderForListRoles) Do(ctx context.Context, role types.Role) error {
	// There is a known issue that the ListRoles does not return tags and
	// permissions boundary, need to invoke GetRole again to obtain such
	// information.
	params := iam.GetRoleInput{
		RoleName: role.RoleName,
	}
	o, e := r.client.GetRole(ctx, &params)
	if e != nil {
		return e
	}

	return r.RoleRecorder.Do(ctx, *o.Role)
}

type UserRecorder struct {
	orm    *model.ORM
	client *iam.Client
}

func (r *UserRecorder) Do(ctx context.Context, user types.User) error {
	return r.orm.UserModel.Insert(ctx, recorder.ToUser(user))
}

type UserRecorderForListUsers struct {
	UserRecorder
}

func (r *UserRecorderForListUsers) Do(ctx context.Context, user types.User) error {
	// There is a known issue that the ListUsers does not return tags and
	// permissions boundary, need to invoke GetUser again to obtain such
	// information.
	params := iam.GetUserInput{
		UserName: user.UserName,
	}
	o, e := r.client.GetUser(ctx, &params)
	if e != nil {
		return e
	}

	return r.UserRecorder.Do(ctx, *o.User)
}
