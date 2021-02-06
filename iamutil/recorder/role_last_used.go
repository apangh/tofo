package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToRoleLastUsed(u *types.RoleLastUsed) *model.RoleLastUsed {
	if u == nil {
		return nil
	}
	return &model.RoleLastUsed{
		Date:   aws.ToTime(u.LastUsedDate),
		Region: aws.ToString(u.Region),
	}
}
