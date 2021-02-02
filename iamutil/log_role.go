package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/glog"
)

var _ ListRolesCB = (*LogRole)(nil)

type LogRole struct {
	i int
}

func (l *LogRole) Do(ctx context.Context, u types.Role) error {
	glog.Infof("%s %v %s %s %s %s %s %d %v %v %v",
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
