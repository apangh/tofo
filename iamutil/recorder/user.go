package recorder

import (
	"github.com/apangh/tofo/model"
	commonRecorder "github.com/apangh/tofo/recorder"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToUser(user types.User) (*model.User, error) {
	arn, e := commonRecorder.ToArn(user.Arn)
	if e != nil {
		return nil, e
	}
	pb, e := ToAttachedPermissionsBoundary(user.PermissionsBoundary)
	if e != nil {
		return nil, e
	}
	return &model.User{
		Id:                  aws.ToString(user.UserId),
		Name:                aws.ToString(user.UserName),
		Path:                aws.ToString(user.Path),
		Arn:                 arn,
		CreateDate:          aws.ToTime(user.CreateDate),
		PasswordLastUsed:    user.PasswordLastUsed,
		Tags:                ToTags(user.Tags),
		PermissionsBoundary: pb,
	}, nil
}
