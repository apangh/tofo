package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToRoleDetail(role types.RoleDetail) *model.RoleDetail {
	return &model.RoleDetail{
		Id:                       aws.ToString(role.RoleId),
		Name:                     aws.ToString(role.RoleName),
		Path:                     aws.ToString(role.Path),
		Arn:                      aws.ToString(role.Arn),
		CreateDate:               aws.ToTime(role.CreateDate),
		AssumeRolePolicyDocument: aws.ToString(role.AssumeRolePolicyDocument),
		Tags:                     ToTags(role.Tags),
		PermissionsBoundary: ToAttachedPermissionsBoundary(
			role.PermissionsBoundary),
		LastUsed:         ToRoleLastUsed(role.RoleLastUsed),
		ManagedPolicies:  ToAttachedPolicies(role.AttachedManagedPolicies),
		Policies:         ToInlinePolicyDetails(role.RolePolicyList),
		InstanceProfiles: ToInstanceProfiles(role.InstanceProfileList),
	}

}
