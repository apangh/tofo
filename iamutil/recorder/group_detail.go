package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToGroupDetail(g types.GroupDetail) (*model.GroupDetail, error) {
	policies, e := ToInlinePolicyDetails(g.GroupPolicyList)
	if e != nil {
		return nil, e
	}
	arn, e := ToArn(g.Arn)
	if e != nil {
		return nil, e
	}
	managedPolicies, e := ToAttachedPolicies(g.AttachedManagedPolicies)
	if e != nil {
		return nil, e
	}

	return &model.GroupDetail{
		Id:              aws.ToString(g.GroupId),
		Name:            aws.ToString(g.GroupName),
		Path:            aws.ToString(g.Path),
		Arn:             arn,
		CreateDate:      aws.ToTime(g.CreateDate),
		ManagedPolicies: managedPolicies,
		Policies:        policies,
	}, nil
}
