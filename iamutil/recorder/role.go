package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToRole(role types.Role) *model.Role {
	return &model.Role{
		Id:                       aws.ToString(role.RoleId),
		Name:                     aws.ToString(role.RoleName),
		Path:                     aws.ToString(role.Path),
		Arn:                      aws.ToString(role.Arn),
		CreateDate:               aws.ToTime(role.CreateDate),
		AssumeRolePolicyDocument: aws.ToString(role.AssumeRolePolicyDocument),
		Description:              aws.ToString(role.Description),
		MaxSessionDuration:       role.MaxSessionDuration,
		Tags:                     ToTags(role.Tags),
		PermissionsBoundary: ToAttachedPermissionsBoundary(
			role.PermissionsBoundary),
		LastUsed: ToRoleLastUsed(role.RoleLastUsed),
	}
}

func ToRoles(r []types.Role) []*model.Role {
	res := make([]*model.Role, 0, len(r))
	for _, a := range r {
		res = append(res, ToRole(a))
	}
	return res
}
