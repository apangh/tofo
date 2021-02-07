package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToRole(role types.Role) (*model.Role, error) {
	arn, e := ToArn(role.Arn)
	if e != nil {
		return nil, e
	}
	pb, e := ToAttachedPermissionsBoundary(role.PermissionsBoundary)
	if e != nil {
		return nil, e
	}
	return &model.Role{
		Id:                       aws.ToString(role.RoleId),
		Name:                     aws.ToString(role.RoleName),
		Path:                     aws.ToString(role.Path),
		Arn:                      arn,
		CreateDate:               aws.ToTime(role.CreateDate),
		AssumeRolePolicyDocument: aws.ToString(role.AssumeRolePolicyDocument),
		Description:              aws.ToString(role.Description),
		MaxSessionDuration:       role.MaxSessionDuration,
		Tags:                     ToTags(role.Tags),
		PermissionsBoundary:      pb,
		LastUsed:                 ToRoleLastUsed(role.RoleLastUsed),
	}, nil
}

func ToRoles(r []types.Role) ([]*model.Role, error) {
	res := make([]*model.Role, 0, len(r))
	for _, a := range r {
		r, e := ToRole(a)
		if e != nil {
			return nil, e
		}
		res = append(res, r)
	}
	return res, nil
}
