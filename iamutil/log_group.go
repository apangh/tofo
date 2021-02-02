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
