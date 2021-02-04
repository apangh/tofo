package iamutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type UserRecorder struct {
	orm *model.ORM
}

func toUser(user types.User) *model.User {
	u := &model.User{
		Id:                 aws.ToString(user.UserId),
		Name:               aws.ToString(user.UserName),
		Path:               aws.ToString(user.Path),
		Arn:                aws.ToString(user.Arn),
		CreateDate:         user.CreateDate,
		PasswordLastUsed:   user.PasswordLastUsed,
		PermissionBoundary: "", // TODO
		Tags:               make(map[string]string),
	}
	for _, t := range user.Tags {
		u.Tags[*t.Key] = *t.Value
	}
	return u
}

func (r *UserRecorder) Do(ctx context.Context, user types.User) error {
	u := toUser(user)
	return r.orm.UserModel.Insert(ctx, u)
}
