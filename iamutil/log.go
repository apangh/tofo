package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/glog"
)

var _ ListGroupsCB = (*LogGroup)(nil)

type LogGroup struct {
	i int
}

func (l *LogGroup) Do(ctx context.Context, g types.Group) error {
	glog.Infof("Group[%d] %s %v %s %s %s", l.i, aws.ToString(g.Arn), g.CreateDate,
		aws.ToString(g.GroupId), aws.ToString(g.GroupName), aws.ToString(g.Path))
	l.i++
	return nil
}

var _ ListRolesCB = (*LogRole)(nil)

type LogRole struct {
	i int
}

func (l *LogRole) Do(ctx context.Context, u types.Role) error {
	glog.Infof("Role[%d] %s %v %s %s %s %s %s %d %v %v %v", l.i,
		aws.ToString(u.Arn), u.CreateDate,
		aws.ToString(u.Path), aws.ToString(u.RoleId),
		aws.ToString(u.RoleName),
		aws.ToString(u.AssumeRolePolicyDocument),
		aws.ToString(u.Description),
		aws.ToInt32(u.MaxSessionDuration),
		u.PermissionsBoundary,
		u.RoleLastUsed,
		u.Tags)
	l.i++
	return nil
}

var _ ListUsersCB = (*LogUser)(nil)

type LogUser struct {
	i int
}

func (l *LogUser) Do(ctx context.Context, u types.User) error {
	glog.Infof("User[%d] %s %v %s %s %s %v %v %v", l.i,
		aws.ToString(u.Arn), u.CreateDate, aws.ToString(u.Path),
		aws.ToString(u.UserId), aws.ToString(u.UserName), u.PasswordLastUsed,
		u.PermissionsBoundary, u.Tags)
	l.i++
	return nil
}
