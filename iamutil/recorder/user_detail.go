package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToUserDetail(user types.UserDetail) *model.UserDetail {
	return &model.UserDetail{
		Id:         aws.ToString(user.UserId),
		Name:       aws.ToString(user.UserName),
		Path:       aws.ToString(user.Path),
		Arn:        aws.ToString(user.Arn),
		CreateDate: aws.ToTime(user.CreateDate),
		Tags:       ToTags(user.Tags),
		PermissionsBoundary: ToAttachedPermissionsBoundary(
			user.PermissionsBoundary),
		ManagedPolicies: ToAttachedPolicies(user.AttachedManagedPolicies),
		Policies:        ToInlinePolicyDetails(user.UserPolicyList),
		Groups:          ToAttachedGroups(user.GroupList),
	}
}