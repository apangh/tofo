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

	return &model.GroupDetail{
		Id:              aws.ToString(g.GroupId),
		Name:            aws.ToString(g.GroupName),
		Path:            aws.ToString(g.Path),
		Arn:             aws.ToString(g.Arn),
		CreateDate:      aws.ToTime(g.CreateDate),
		ManagedPolicies: ToAttachedPolicies(g.AttachedManagedPolicies),
		Policies:        policies,
	}, nil
}
