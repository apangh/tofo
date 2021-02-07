package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToRoleDetail(role types.RoleDetail) (*model.RoleDetail, error) {
	s, e := ToJsonPolicyDocument(aws.ToString(role.AssumeRolePolicyDocument))
	if e != nil {
		return nil, e
	}
	policies, e := ToInlinePolicyDetails(role.RolePolicyList)
	if e != nil {
		return nil, e
	}
	ip, e := ToInstanceProfiles(role.InstanceProfileList)
	if e != nil {
		return nil, e
	}
	arn, e := ToArn(role.Arn)
	if e != nil {
		return nil, e
	}
	pb, e := ToAttachedPermissionsBoundary(role.PermissionsBoundary)
	if e != nil {
		return nil, e
	}
	mp, e := ToAttachedPolicies(role.AttachedManagedPolicies)
	if e != nil {
		return nil, e
	}

	return &model.RoleDetail{
		Id:                       aws.ToString(role.RoleId),
		Name:                     aws.ToString(role.RoleName),
		Path:                     aws.ToString(role.Path),
		Arn:                      arn,
		CreateDate:               aws.ToTime(role.CreateDate),
		AssumeRolePolicyDocument: s,
		Tags:                     ToTags(role.Tags),
		PermissionsBoundary:      pb,
		LastUsed:                 ToRoleLastUsed(role.RoleLastUsed),
		ManagedPolicies:          mp,
		Policies:                 policies,
		InstanceProfiles:         ip,
	}, nil
}
