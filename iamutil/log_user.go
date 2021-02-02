package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/glog"
)

var _ ListUsersCB = (*LogUser)(nil)

type LogUser struct {
	i int
}

func (l *LogUser) Do(ctx context.Context, u types.User) error {
	glog.Infof("%s %v %s %s %s %v %v %v", aws.ToString(u.Arn), u.CreateDate,
		aws.ToString(u.Path), aws.ToString(u.UserId), aws.ToString(u.UserName),
		u.PasswordLastUsed, u.PermissionsBoundary, u.Tags)
	l.i++
	return nil
}
