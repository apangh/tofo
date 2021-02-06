package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToUser(user types.User) *model.User {
	return &model.User{
		Id:               aws.ToString(user.UserId),
		Name:             aws.ToString(user.UserName),
		Path:             aws.ToString(user.Path),
		Arn:              aws.ToString(user.Arn),
		CreateDate:       aws.ToTime(user.CreateDate),
		PasswordLastUsed: user.PasswordLastUsed,
		Tags:             ToTags(user.Tags),
		PermissionsBoundary: ToAttachedPermissionsBoundary(
			user.PermissionsBoundary),
	}
}
